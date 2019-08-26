package must_test

import (
	"math/bits"

	"github.com/openacid/must"
)

func CheckedRShift(a, b int) int {

	must.Be.OK(func() {
		must.Be.NotZero(b)
		must.Be.True(bits.TrailingZeros(uint(a)) > 2,
			"a must be multiple of 8")
	})

	return a >> uint(b)
}
