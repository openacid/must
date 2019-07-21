package main

import (
	"fmt"
	"math/bits"

	"github.com/openacid/must"
)

var mustbe = must.Be

func rshift(a, b int) int {

	// "go build" emits a single No-op instruction.
	// "go build -tags debug" will call the function and to the checking.
	mustbe.OK(func() {
		mustbe.NotZero(b)
		mustbe.True(bits.TrailingZeros(uint(a)) > 2, "a must be multiple of 8")
	})

	return a >> uint(b)
}

func main() {
	fmt.Println(rshift(0xf, 1)) // panic at line 19 with "go run -tags debug"
}
