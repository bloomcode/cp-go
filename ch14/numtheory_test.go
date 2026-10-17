package ch14

import (
	"math/big"
	"math/rand"
	"testing"

	"cp/internal/check"
)

func TestEx14_1(t *testing.T) {
	check.Eq(t, []int{GCD(84, 36), GCD(-84, 36), LCM(4, 6), LCM(1_000_000_000, 999_999_999)}, []int{12, 12, 12, 999999999000000000})
}

func TestEx14_2(t *testing.T) {
	check.Eq(t, PrimesUpTo(30), []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29})
	check.Eq(t, len(PrimesUpTo(1_000_000)), 78498)
}

func TestEx14_3(t *testing.T) {
	check.Eq(t, Factorize(360), [][2]int{{2, 3}, {3, 2}, {5, 1}})
	check.Eq(t, Factorize(2*1_000_000_007), [][2]int{{2, 1}, {1_000_000_007, 1}})
}

func TestEx14_4(t *testing.T) {
	check.Eq(t, MulMod(999_999_999_999_999_989, 2, 1_000_000_000_000_000_000), uint64(999999999999999978))
	r := rand.New(rand.NewSource(1))
	for i := 0; i < 500; i++ {
		a, b, m := r.Uint64(), r.Uint64(), r.Uint64()|1
		want := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).SetUint64(a), new(big.Int).SetUint64(b)), new(big.Int).SetUint64(m))
		check.Eq(t, MulMod(a, b, m), want.Uint64())
		e := r.Uint64() >> 1
		wantPow := new(big.Int).Exp(new(big.Int).SetUint64(a), new(big.Int).SetUint64(e), new(big.Int).SetUint64(m))
		check.Eq(t, PowMod(a, e, m), wantPow.Uint64())
	}
}

func TestEx14_5(t *testing.T) {
	check.Eq(t, []int{NewBinom(10).C(10, 3), NewBinom(100000).C(100000, 50000), NewBinom(5).C(3, 5)}, []int{120, 149033233, 0})
}

func TestEx14_6(t *testing.T) {
	g, x, y := ExtGCD(240, 46)
	check.Eq(t, []int{g, 240*x + 46*y}, []int{2, 2})
	inv, ok := ModInverse(3, 10)
	check.Eq(t, []any{inv, ok}, []any{7, true})
	_, ok = ModInverse(4, 10)
	check.Eq(t, ok, false)
}

func TestEx14_7(t *testing.T) {
	check.Eq(t, Divisors(36), []int{1, 2, 3, 4, 6, 9, 12, 18, 36})
}

func TestEx14_8(t *testing.T) {
	check.Eq(t, []string{FibBig(100), FibBig(1)}, []string{"354224848179261915075", "1"})
	check.True(t, ProbablyPrime("1000000000000000003") && !ProbablyPrime("1000000000000000001"), "primality")
}
