package main

import (
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count
// (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// LoopPopCount uses loop
func LoopPopCount(x uint64) int {

	var sum byte
	for i := uint64(0); i < 8; i++ {
		seg := byte(x >> (i * 8))
		b := pc[seg]
		sum += b
		fmt.Printf("%d) %08b = %d: += %d = %d\n", i, seg, seg, b, sum)
	}
	fmt.Println()
	return int(sum)
}

// ShiftPopCount returns population count using shift based algorithm
func ShiftPopCount(x uint64) int {
	workx := x
	var sum int
	for ; workx > 0; workx = (workx >> 1) {
		if workx&1 == 1 {
			sum++
		}
		fmt.Printf("%064b - %d\n", workx, sum)
	}
	fmt.Println()
	return sum
}

// ClearPopCount returns population count by clearing successive
// rightmost bits
func ClearPopCount(x uint64) int {
	workx := x
	var sum int
	fmt.Printf("%064b - %d\n", workx, sum)
	for workx > 0 {
		workx = workx & (workx - 1)
		sum++
		fmt.Printf("%064b - %d\n", workx, sum)
	}
	fmt.Println()
	return sum
}

func printIndex(x int) {
	fmt.Printf("%d, %b, %d, %d, %d, %d\n", x, x, PopCount(uint64(x)),
		LoopPopCount(uint64(x)),
		ShiftPopCount(uint64(x)),
		ClearPopCount(uint64(x)))
	fmt.Println()
}

func main() {
	/*for i := range pc {
		fmt.Printf("i: %d i(bin):%b - val: %b val(int) %d\n", i, i, pc[i], int(pc[i]))
	}

	for i := 0; i <= 255; i++ {
		fmt.Println(i, "PopCount:", PopCount(uint64(i)))
	}
	*/
	printIndex(0)
	printIndex(1)
	printIndex(45)
	printIndex(100)
	printIndex(453847584)
}
