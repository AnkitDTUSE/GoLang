package main

import (
	"fmt"
	"math"
)

type ErrNegativeValue float64

func (e ErrNegativeValue) Error() string {
	return fmt.Sprint("Cannot give sqrt of -ve number:", float64(e))
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeValue(x)
	}

	z := 1.0

	for {
		prev := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev) < 1e-10 {
			break
		}
	}

	return z, nil

}

func excersice() {
	fmt.Println(sqrt(4))
	fmt.Println(sqrt(-2))
}
