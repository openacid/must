// Package disabled implements disabled checking functions.
//
// Since 0.1.0
package disabled

import (
	"time"

	"github.com/stretchr/testify/assert"
)

type beTyp func(f func())

func be(f func()) {}

// Be calls a checking function.
// If f() fails Be() would panic.
//
// Since 0.1.2
var Be beTyp = be

func (fl *beTyp) Condition(comp assert.Comparison, msgAndArgs ...interface{})            {}
func (fl *beTyp) Contains(s, contains interface{}, msgAndArgs ...interface{})            {}
func (fl *beTyp) DirExists(path string, msgAndArgs ...interface{})                       {}
func (fl *beTyp) ElementsMatch(listA, listB interface{}, msgAndArgs ...interface{})      {}
func (fl *beTyp) Empty(object interface{}, msgAndArgs ...interface{})                    {}
func (fl *beTyp) Equal(expected, actual interface{}, msgAndArgs ...interface{})          {}
func (fl *beTyp) EqualError(theError error, errString string, msgAndArgs ...interface{}) {}
func (fl *beTyp) EqualValues(expected, actual interface{}, msgAndArgs ...interface{})    {}
func (fl *beTyp) Error(err error, msgAndArgs ...interface{})                             {}
func (fl *beTyp) Exactly(expected, actual interface{}, msgAndArgs ...interface{})        {}
func (fl *beTyp) Fail(failureMessage string, msgAndArgs ...interface{})                  {}
func (fl *beTyp) FailNow(failureMessage string, msgAndArgs ...interface{})               {}
func (fl *beTyp) False(value bool, msgAndArgs ...interface{})                            {}
func (fl *beTyp) FileExists(path string, msgAndArgs ...interface{})                      {}
func (fl *beTyp) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
}
func (fl *beTyp) InDelta(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {}
func (fl *beTyp) InDeltaMapValues(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {
}
func (fl *beTyp) InDeltaSlice(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {
}
func (fl *beTyp) InEpsilon(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
}
func (fl *beTyp) InEpsilonSlice(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
}
func (fl *beTyp) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {}
func (fl *beTyp) JSONEq(expected string, actual string, msgAndArgs ...interface{})               {}
func (fl *beTyp) Len(object interface{}, length int, msgAndArgs ...interface{})                  {}
func (fl *beTyp) Nil(object interface{}, msgAndArgs ...interface{})                              {}
func (fl *beTyp) NoError(err error, msgAndArgs ...interface{})                                   {}
func (fl *beTyp) NotContains(s, contains interface{}, msgAndArgs ...interface{})                 {}
func (fl *beTyp) NotEmpty(object interface{}, msgAndArgs ...interface{})                         {}
func (fl *beTyp) NotEqual(expected, actual interface{}, msgAndArgs ...interface{})               {}
func (fl *beTyp) NotNil(object interface{}, msgAndArgs ...interface{})                           {}
func (fl *beTyp) NotPanics(f assert.PanicTestFunc, msgAndArgs ...interface{})                    {}
func (fl *beTyp) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{})           {}
func (fl *beTyp) NotSubset(list, subset interface{}, msgAndArgs ...interface{})                  {}
func (fl *beTyp) NotZero(i interface{}, msgAndArgs ...interface{})                               {}
func (fl *beTyp) Panics(f assert.PanicTestFunc, msgAndArgs ...interface{})                       {}
func (fl *beTyp) PanicsWithValue(expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
}
func (fl *beTyp) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {}
func (fl *beTyp) Subset(list, subset interface{}, msgAndArgs ...interface{})        {}
func (fl *beTyp) True(value bool, msgAndArgs ...interface{})                        {}
func (fl *beTyp) WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
}
func (fl *beTyp) Zero(i interface{}, msgAndArgs ...interface{}) {}
