package main

import (
	"fmt"
)

type ErrNegativeSquare float64

func (e ErrNegativeSquare) Error() string {
	return "cannot Square negative number: -2"
}

func Square(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSquare(x)
	}
	return x*x, nil
}

func main() {
	fmt.Println(Square(2))
	fmt.Println(Square(-2))
}
