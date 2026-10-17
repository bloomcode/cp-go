// Package ch13 - Bit manipulation.
package ch13


// ==== EX 13.1 | Test, set, clear, toggle | Core
// PROMPT: Write `func IsSet(x uint, i int) bool`, `func Set(x uint, i int) uint`, `func Clear(x uint, i int) uint` and `func Toggle(x uint, i int) uint`. Go also has `&^` (AND NOT), made for clearing bits.
// EXAMPLE: x = 0b1010: `IsSet(x, 1)` -> `true`; `Set(x, 0)` -> `11`; `Clear(x, 3)` -> `2`; `Toggle(x, 2)` -> `14`
// BIGO: O(1).
func IsSet(x uint, i int) bool { panic("todo") }

func Set(x uint, i int) uint { panic("todo") }

func Clear(x uint, i int) uint { panic("todo") }

func Toggle(x uint, i int) uint { panic("todo") }

// ==== END

// ==== EX 13.2 | Shifts past the width give zero | Core
// PROMPT: Write `func Shift32(x int32, n uint) int32` and `func Shift64(x int64, n uint) int64` returning `x << n`. Compare `Shift32(1, 40)` with Java, where the count is reduced mod 32.
// EXAMPLE: `Shift32(1, 40)` -> `0`; `Shift64(1, 40)` -> `1099511627776`; plain `1 << 40` as an untyped constant is fine in an `int`
// BIGO: O(1). Go does not mask the shift count: shifting by the width or more gives 0 (or -1 for a negative value shifted right).
func Shift32(x int32, n uint) int32 { panic("todo") }

func Shift64(x int64, n uint) int64 { panic("todo") }

// ==== END

// ==== EX 13.3 | Signed versus unsigned right shift | Core
// PROMPT: Go has one `>>`: it copies the sign bit for signed types and fills with zeros for unsigned ones. Write `func Midpoint(lo, hi int) int` as `int(uint(lo+hi) >> 1)` for non-negative bounds, and `func ZeroFill32(x int32, n uint) int32` that shifts the unsigned bits.
// EXAMPLE: `Midpoint(math.MaxInt64 - 2, math.MaxInt64)` -> `9223372036854775806`; `-16 >> 2` -> `-4`; `ZeroFill32(-16, 28)` -> `15`
// BIGO: O(1). `sort.Search` computes its midpoint exactly this way.
func Midpoint(lo, hi int) int { panic("todo") }

func ZeroFill32(x int32, n uint) int32 { panic("todo") }

// ==== END

// ==== EX 13.4 | math/bits | Core
// PROMPT: Write `func BitFacts(x uint64) (ones, trailing, highest int)` with `bits.OnesCount64`, `bits.TrailingZeros64` and `bits.Len64(x) - 1`.
// EXAMPLE: `BitFacts(40)` -> `2, 3, 5`
// BIGO: O(1); these compile to single CPU instructions.
func BitFacts(x uint64) (ones, trailing, highest int) {
	panic("todo")
}

// ==== END

// ==== EX 13.5 | Power of two and lowest set bit | Core
// PROMPT: Write `func IsPowerOfTwo(x int) bool` with `x & (x - 1)` and `func LowestBit(x int) int` with `x & -x`.
// EXAMPLE: `IsPowerOfTwo(64)` -> `true`; `IsPowerOfTwo(0)` -> `false`; `LowestBit(40)` -> `8`
// BIGO: O(1). `x & -x` drives Fenwick trees.
func IsPowerOfTwo(x int) bool { panic("todo") }

func LowestBit(x int) int { panic("todo") }

// ==== END

// ==== EX 13.6 | Subsets as bitmasks | Core
// PROMPT: Write `func SubsetsByMask(a []int) [][]int` looping `mask` from 0 to `1<<n - 1`.
// EXAMPLE: `SubsetsByMask([]int{1, 2})` -> `[[] [1] [2] [1 2]]`
// BIGO: O(N * 2^N); fine up to n ~ 20.
func SubsetsByMask(a []int) [][]int {
	panic("todo")
}

// ==== END

// ==== EX 13.7 | XOR tricks | Core
// PROMPT: Write `func LoneValue(a []int) int` where every value appears twice except one, and `func Missing(a []int) int` for 0..n with one value missing.
// EXAMPLE: `LoneValue([]int{4, 1, 2, 1, 2})` -> `4`; `Missing([]int{3, 0, 1})` -> `2`
// BIGO: O(N) time, O(1) space.
func LoneValue(a []int) int {
	panic("todo")
}

func Missing(a []int) int {
	panic("todo")
}

// ==== END

// ==== EX 13.8 | Iterate the submasks of a mask | Stretch
// PROMPT: Write `func Submasks(mask int) []int` listing every submask from largest to 0 with `s = (s - 1) & mask`.
// EXAMPLE: `Submasks(0b101)` -> `[5 4 1 0]`
// BIGO: O(2^popcount). Summed over all masks of n bits this is O(3^n).
func Submasks(mask int) []int {
	panic("todo")
}

// ==== END
