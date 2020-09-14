package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
)

func main() {
	const PRECISION = 200 // bits of precision
	argument, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalln("The argument is incorrect.")
	}

	steps := int(math.Log2(PRECISION))

	// Initial values
	a := new(big.Float).SetPrec(PRECISION).SetInt64(argument)
	half := new(big.Float).SetPrec(PRECISION).SetFloat64(.5)

	x := new(big.Float).SetPrec(PRECISION).SetInt64(1) // initial estimate

	t := new(big.Float) // temporal variable

	for i := 0; i <= steps; i++ {
		t.Quo(a, x)    // t = a / x_n
		t.Add(x, t)    // t = x_n + (2.0 / x_n)
		x.Mul(half, t) // x_{n+1} = 0.5 * t
	}

	fmt.Printf("sqrt(2) = %.50f\n", x)

	// The error between 2 and x*x.
	t.Mul(x, x) // t = x*x
	fmt.Printf("error = %e\n", t.Sub(a, t))
}
