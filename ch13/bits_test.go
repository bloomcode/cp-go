package ch13

import (
	"math"
	"testing"

	"cp/internal/check"
)

func TestEx13_1(t *testing.T) {
	var x uint = 0b1010
	check.True(t, IsSet(x, 1) && !IsSet(x, 0), "isset")
	check.Eq(t, []uint{Set(x, 0), Clear(x, 3), Toggle(x, 2)}, []uint{11, 2, 14})
}

func TestEx13_2(t *testing.T) {
	check.Eq(t, []int32{Shift32(1, 40), Shift32(1, 31)}, []int32{0, math.MinInt32})
	check.Eq(t, Shift64(1, 40), int64(1099511627776))
}

func TestEx13_3(t *testing.T) {
	check.Eq(t, Midpoint(math.MaxInt64-2, math.MaxInt64), 9223372036854775806)
	check.Eq(t, ZeroFill32(-16, 28), int32(15))
}

func TestEx13_4(t *testing.T) {
	o, tz, h := BitFacts(40)
	check.Eq(t, []int{o, tz, h}, []int{2, 3, 5})
}

func TestEx13_5(t *testing.T) {
	check.True(t, IsPowerOfTwo(64) && !IsPowerOfTwo(0) && !IsPowerOfTwo(12), "pow2")
	check.Eq(t, LowestBit(40), 8)
}

func TestEx13_6(t *testing.T) {
	check.Eq(t, SubsetsByMask([]int{1, 2}), [][]int{{}, {1}, {2}, {1, 2}})
}

func TestEx13_7(t *testing.T) {
	check.Eq(t, []int{LoneValue([]int{4, 1, 2, 1, 2}), Missing([]int{3, 0, 1})}, []int{4, 2})
}

func TestEx13_8(t *testing.T) {
	check.Eq(t, Submasks(0b101), []int{5, 4, 1, 0})
}
