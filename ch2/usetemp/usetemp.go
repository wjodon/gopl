package main

import (
	"fmt"

	tc "github.com/wjodon/gopl/ch2/tempconv.1"
)

func main() {
	var temp1F tc.Fahrenheit = 33
	var temp1C tc.Celsius = 45
	var temp1K tc.Kelvin = 22

	fmt.Println(temp1F)
	fmt.Println(temp1F, "=", tc.FToC(temp1F))
	fmt.Println(temp1F, "=", tc.FToK(temp1F))

	fmt.Println(temp1C)
	fmt.Println(temp1C, "=", tc.CToF(temp1C))
	fmt.Println(temp1C, "=", tc.CToK(temp1C))

	fmt.Println(temp1K)
	fmt.Println(temp1K, "=", tc.KToF(temp1K))
	fmt.Println(temp1K, "=", tc.KToC(temp1K))
}
