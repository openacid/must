package main

import (
	"fmt"
	"math/bits"

	"github.com/openacid/must"
)

var mustbe = must.Be

func rshift(a, b int) int {

	// Simple check:
	//   "go build" emits a No-op instruction.
	//   "go build -tags debug" will check "b" not to be zero.
	mustbe.NotZero(b)

	// Inappropriate: argument expressions are still evaluated.
	mustbe.True(bits.TrailingZeros(uint(a)) > 2)

	return a >> uint(b)
}

func main() {
	fmt.Println(rshift(0xf, 1)) // panic at line 20 with "go run -tags debug"
}
