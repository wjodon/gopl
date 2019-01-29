// cf converts its numeric argument to Celsius, Fahrenheit, and Kelvin.
package main

import (
	"fmt"
	"os"
	"strconv"

	tc "github.com/wjodon/gopl/ch2/tempconv.1"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tc.Fahrenheit(t)
		c := tc.Celsius(t)
		k := tc.Kelvin(t)
		fmt.Printf("\n%s = %s = %s\n%s = %s = %s\n%s = %s = %s\n",
			f, tc.FToC(f), tc.FToK(f),
			c, tc.CToF(c), tc.CToK(c),
			k, tc.KToC(k), tc.KToF(k))
	}

}
