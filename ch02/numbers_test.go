package ch02

import (
	"math"
	"testing"

	"cp/internal/check"
)

func TestEx2_1(t *testing.T) {
	check.Eq(t, Wrapped(), int32(-2147483648))
	check.Eq(t, Product(100000, 100000), int64(10000000000))
	a, b := int32(100000), int32(100000)
	check.Eq(t, a*b, int32(1410065408)) // what int32 arithmetic gives
}

func TestEx2_2(t *testing.T) {
	check.True(t, MulFits(3_000_000_000, 3), "3e9*3")
	check.True(t, !MulFits(4_000_000_000, 4_000_000_000), "4e9*4e9")
	check.True(t, MulFits(math.MinInt64, 1), "min*1")
	check.True(t, !MulFits(math.MinInt64, -1), "min*-1")
	check.True(t, MulFits(-(1<<62), 2), "-2^63")
	check.True(t, !MulFits(1<<62, 2), "2^63")
	check.True(t, MulFits(0, math.MinInt64), "0")
}

func TestEx2_3(t *testing.T) {
	q, r := TruncDivMod(-7, 2)
	check.Eq(t, []int{q, r}, []int{-3, -1})
	q, r = FloorDivMod(-7, 2)
	check.Eq(t, []int{q, r}, []int{-4, 1})
	q, r = FloorDivMod(7, -2)
	check.Eq(t, []int{q, r}, []int{-4, -1})
	q, r = FloorDivMod(6, 3)
	check.Eq(t, []int{q, r}, []int{2, 0})
}

func TestEx2_4(t *testing.T) {
	check.Eq(t, []int{CeilDiv(7, 2), CeilDiv(-7, 2), CeilDiv(6, 3)}, []int{4, -3, 2})
}

func TestEx2_5(t *testing.T) {
	check.Eq(t, []int{ModPow(2, 10, MOD), ModPow(2, 100, MOD)}, []int{1024, 976371285})
	check.Eq(t, 123456*ModInv(123456)%MOD, 1)
}

func TestEx2_6(t *testing.T) {
	check.Eq(t, []int{FactorialDigits(20), FactorialDigits(100)}, []int{19, 158})
}

func TestEx2_7(t *testing.T) {
	x, y := 0.1, 0.2
	check.True(t, Close(x+y, 0.3, 1e-9) && !Close(0.1, 0.2, 1e-9), "close")
	check.True(t, x+y != 0.3, "float inexact")
}

func TestEx2_8(t *testing.T) {
	check.Eq(t, []int{Isqrt(99), Isqrt(100), Isqrt(999_999_999_999_999_999), Isqrt(math.MaxInt64)}, []int{9, 10, 999999999, 3037000499})
}

func TestEx2_9(t *testing.T) {
	check.Eq(t, []int{DigitSum(9045), ReverseDigits(-1200)}, []int{18, -21})
}
