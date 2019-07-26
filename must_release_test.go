// +build !debug

package must_test

import (
	"bytes"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/openacid/must"
	"github.com/stretchr/testify/require"
)

func TestOK(t *testing.T) {

	ta := require.New(t)

	val := 0

	must.Be(func() {
		val = 1
	})

	ta.Equal(0, val)

	ta.NotPanics(func() {
		must.Be(func() {
			must.Be.True(false)
		})
	})

}

func TestNotPanicEvenIfMustBeNotSatisfied(t *testing.T) {

	ta := require.New(t)

	ta.NotPanics(func() {
		must.Be.Condition(func() bool {
			return false
		})
	})

	ta.NotPanics(func() {
		must.Be.Contains("Hello word", "foo")
	})

	ta.NotPanics(func() {
		must.Be.DirExists("./no_such_dir")
	})

	ta.NotPanics(func() {
		must.Be.ElementsMatch([]int{1, 2, 2}, []int{1, 2, 3})
	})
	ta.NotPanics(func() {
		must.Be.Empty([]int{1})
	})
	ta.NotPanics(func() {
		must.Be.Equal(1, 2)
	})
	ta.NotPanics(func() {
		must.Be.EqualError(errors.New("foo"), "bar")
	})
	ta.NotPanics(func() {
		must.Be.EqualValues(uint32(1), int32(-1))
	})
	ta.NotPanics(func() {
		must.Be.Error(nil)
	})
	ta.NotPanics(func() {
		must.Be.Exactly(int32(1), int64(1))
	})
	ta.NotPanics(func() {
		must.Be.Fail("failit")
	})
	ta.NotPanics(func() {
		must.Be.FailNow("failit")
	})
	ta.NotPanics(func() {
		must.Be.False(true)
	})
	ta.NotPanics(func() {
		must.Be.FileExists("foo")
	})
	ta.NotPanics(func() {
		must.Be.Implements((*io.Reader)(nil), 1)
	})
	ta.NotPanics(func() {
		must.Be.InDelta(1, 2, 0.1)
	})
	ta.NotPanics(func() {
		must.Be.InDeltaMapValues(map[int]int{1: 10, 2: 20}, map[int]int{1: 20, 2: 30}, 1)
	})
	ta.NotPanics(func() {
		must.Be.InDeltaSlice([]int{1, 2}, []int{5, 10}, 1)
	})
	ta.NotPanics(func() {
		must.Be.InEpsilon(10, 20, 0.1)
	})
	ta.NotPanics(func() {
		must.Be.InEpsilonSlice([]int{1, 2}, []int{5, 10}, 0.1)
	})
	ta.NotPanics(func() {
		must.Be.IsType(int(0), uint64(0))
	})
	ta.NotPanics(func() {
		must.Be.JSONEq(`{"hello": "hello", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	})
	ta.NotPanics(func() {
		must.Be.Len("123", 2)
	})
	ta.NotPanics(func() {
		must.Be.Nil(1)
	})
	ta.NotPanics(func() {
		must.Be.NoError(errors.New("foo"))
	})
	ta.NotPanics(func() {
		must.Be.NotContains("hello word", "hello")
	})
	ta.NotPanics(func() {
		must.Be.NotEmpty("")
	})
	ta.NotPanics(func() {
		must.Be.NotEqual(1, 1)
	})
	ta.NotPanics(func() {
		must.Be.NotNil(nil)
	})
	ta.NotPanics(func() {
		must.Be.NotPanics(func() {
			panic("foo")
		})
	})
	ta.NotPanics(func() {
		must.Be.NotRegexp("^abc", "abc")
	})
	ta.NotPanics(func() {
		must.Be.NotSubset([]int{1, 2, 3}, []int{1, 2})
	})
	ta.NotPanics(func() {
		must.Be.NotZero(0)
	})
	ta.NotPanics(func() {
		must.Be.NotPanics(func() {})
	})
	ta.NotPanics(func() {
		must.Be.PanicsWithValue("panic-foo", func() { panic("panic-value") })
	})
	ta.NotPanics(func() {
		must.Be.Regexp("^start", "st")
	})
	ta.NotPanics(func() {
		must.Be.Subset([]int{1, 2, 3}, []int{1, 2, 4})
	})
	ta.NotPanics(func() {
		must.Be.True(false)
	})
	ta.NotPanics(func() {
		must.Be.WithinDuration(time.Time{}, time.Now(), 10*time.Second)
	})
	ta.NotPanics(func() {
		must.Be.Zero(1)
	})
}

func TestNotPanicIfMustBeSatisfied(t *testing.T) {

	ta := require.New(t)

	ta.NotPanics(func() {
		must.Be.Condition(func() bool {
			return true
		})
	})

	ta.NotPanics(func() {
		must.Be.Contains("Hello word", "Hello")
	})

	ta.NotPanics(func() {
		must.Be.DirExists("./examples")
	})

	ta.NotPanics(func() {
		must.Be.ElementsMatch([]int{1, 2, 2}, []int{1, 2, 2})
	})
	ta.NotPanics(func() {
		must.Be.Empty([]int{})
	})
	ta.NotPanics(func() {
		must.Be.Equal(1, 1)
	})
	ta.NotPanics(func() {
		must.Be.EqualError(errors.New("foo"), "foo")
	})
	ta.NotPanics(func() {
		must.Be.EqualValues(uint32(1), int32(1))
	})
	ta.NotPanics(func() {
		must.Be.Error(errors.New("foo"))
	})
	ta.NotPanics(func() {
		must.Be.Exactly(int32(1), int32(1))
	})
	ta.NotPanics(func() {
		// must.Be.Fail("failit")
	})
	ta.NotPanics(func() {
		// must.Be.FailNow("failit")
	})
	ta.NotPanics(func() {
		must.Be.False(false)
	})
	ta.NotPanics(func() {
		must.Be.FileExists("README.md")
	})
	ta.NotPanics(func() {
		must.Be.Implements((*io.Reader)(nil), bytes.NewBuffer([]byte{}))
	})
	ta.NotPanics(func() {
		must.Be.InDelta(1, 2, 2)
	})
	ta.NotPanics(func() {
		must.Be.InDeltaMapValues(map[int]int{1: 10, 2: 20}, map[int]int{1: 20, 2: 30}, 50)
	})
	ta.NotPanics(func() {
		must.Be.InDeltaSlice([]int{1, 2}, []int{5, 10}, 50)
	})
	ta.NotPanics(func() {
		must.Be.InEpsilon(10, 20, 2)
	})
	ta.NotPanics(func() {
		must.Be.InEpsilonSlice([]int{1, 2}, []int{5, 10}, 100)
	})
	ta.NotPanics(func() {
		must.Be.IsType(int(0), int(0))
	})
	ta.NotPanics(func() {
		must.Be.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	})
	ta.NotPanics(func() {
		must.Be.Len("123", 3)
	})
	ta.NotPanics(func() {
		must.Be.Nil(nil)
	})
	ta.NotPanics(func() {
		must.Be.NoError(nil)
	})
	ta.NotPanics(func() {
		must.Be.NotContains("hello word", "foo")
	})
	ta.NotPanics(func() {
		must.Be.NotEmpty("123")
	})
	ta.NotPanics(func() {
		must.Be.NotEqual(1, 2)
	})
	ta.NotPanics(func() {
		must.Be.NotNil(1)
	})
	ta.NotPanics(func() {
		must.Be.NotPanics(func() {
		})
	})
	ta.NotPanics(func() {
		must.Be.NotRegexp("^abc", "foo")
	})
	ta.NotPanics(func() {
		must.Be.NotSubset([]int{1, 2, 3}, []int{1, 2, 4})
	})
	ta.NotPanics(func() {
		must.Be.NotZero(1)
	})
	ta.NotPanics(func() {
		must.Be.Panics(func() { panic("foo") })
	})
	ta.NotPanics(func() {
		must.Be.PanicsWithValue("panic-value", func() { panic("panic-value") })
	})
	ta.NotPanics(func() {
		must.Be.Regexp("^start", "start")
	})
	ta.NotPanics(func() {
		must.Be.Subset([]int{1, 2, 3}, []int{1, 2})
	})
	ta.NotPanics(func() {
		must.Be.True(true)
	})
	ta.NotPanics(func() {
		must.Be.WithinDuration(time.Now(), time.Now(), 10*time.Second)
	})
	ta.NotPanics(func() {
		must.Be.Zero(0)
	})
}
