package main

import (
	"fmt"
	"math/bits"

	"github.com/openacid/must"
)

func rshift(a, b int) int {

	// "go build" emits a single No-op instruction.
	// "go build -tags debug" will call the function and to the checking.
	must.Be(func() {
		t := must.Be
		t.NotZero(b)
		t.True(bits.TrailingZeros(uint(a)) > 2,
			"a must be multiple of 8")
	})

	return a >> uint(b)
}

func main() {
	// panic at line 19 with "go run -tags debug"
	fmt.Println(rshift(0xf, 1))
}
