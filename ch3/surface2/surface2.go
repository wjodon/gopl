// Surface1 computes an SVG rendering of a 3-D surface function.
package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	width, height = 600, 600            // canvas size in pixels
	cells         = 100                 // number of cells
	xyrange       = 10.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.0025     // pixels per z unit
	zcolor        = 100
	angle         = math.Pi / 6 // angle of x, y axes (=30 deg)
)

var s = flag.Bool("s", false, "display min/max scaled values of x and y")
var v = flag.Bool("v", false, "display min/max values of x, y, and z")

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30 deg), cos(30 deg)

func main() {
	var maxX, maxY, maxZ, minX, minY, minZ, maxSx, maxSy, minSx, minSy float64

	flag.Parse()
	if !(*s || *v) {
		fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			asx, asy := isoshift(ax, ay, az)
			bsx, bsy := isoshift(bx, by, bz)
			csx, csy := isoshift(cx, cy, cz)
			dsx, dsy := isoshift(dx, dy, dz)
			if !(math.IsInf(asx, 0) ||
				math.IsInf(asy, 0) ||
				math.IsInf(bsx, 0) ||
				math.IsInf(bsy, 0) ||
				math.IsInf(csx, 0) ||
				math.IsInf(csy, 0) ||
				math.IsInf(dsx, 0) ||
				math.IsInf(dsy, 0)) {

				if *s {
					maxSx = maxArr5([5]float64{maxSx, asx, bsx, csx, dsx})
					maxSy = maxArr5([5]float64{maxSy, asy, bsy, csy, dsy})
					minSx = minArr5([5]float64{minSx, asx, bsx, csx, dsx})
					minSy = minArr5([5]float64{minSy, asy, bsy, csy, dsy})
				} else if *v {
					maxX = maxArr5([5]float64{maxX, ax, bx, cx, dx})
					maxY = maxArr5([5]float64{maxY, ay, by, cy, dy})
					maxZ = maxArr5([5]float64{maxZ, az, bz, cz, dz})
					minX = minArr5([5]float64{minX, ax, bx, cx, dx})
					minY = minArr5([5]float64{minY, ay, by, cy, dy})
					minZ = minArr5([5]float64{minZ, az, bz, cz, dz})
				} else {
					lowZ := minArr4([4]float64{az, bz, cz, dz})
					var color string
					if int(lowZ) < zcolor {
						color = "#0000ff"
					} else {
						color = "#ff0000"
					}
					fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='%s'/><!-- %g -->\n",
						asx, asy, bsx, bsy, csx, csy, dsx, dsy, color, lowZ)
				}
			}
		}
	}
	if *s {
		fmt.Println("SCALED:: Min x:", minSx, "y:", minSy, "- Max x:", maxSx, "y:", maxSy)
	} else if *v {
		fmt.Println("Max x:", maxX, "y:", maxY, "z:", maxZ)
	} else {
		fmt.Println("</svg>")
	}

}

func minArr4(a [4]float64) float64 {
	min := math.MaxFloat64
	for i := 0; i < len(a); i++ {
		min = math.Min(min, a[i])
	}
	return min
}

func minArr5(a [5]float64) float64 {
	min := math.MaxFloat64
	for i := 0; i < len(a); i++ {
		min = math.Min(min, a[i])
	}
	return min
}

func maxArr5(a [5]float64) float64 {
	var max float64
	for i := 0; i < len(a); i++ {
		max = math.Max(max, a[i])
	}
	return max
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	return x, y, z
}

func isoshift(x, y, z float64) (float64, float64) {
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,xy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

/*
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
*/

func f(x, y float64) float64 {
	return math.Pow(x, 2) + math.Pow(y, 2) +
		25*(math.Pow(math.Sin(x), 2)+
			math.Pow(math.Sin(y), 2))
}
