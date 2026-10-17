// Package ch02 - Numbers and overflow.
package ch02


// ==== EX 2.1 | Sized integers wrap silently | Core
// PROMPT: Write `func Wrapped() int32` returning `math.MaxInt32 + 1` computed through a variable, and `func Product(a, b int32) int64` that converts before multiplying.
// EXAMPLE: `Wrapped()` -> `-2147483648`; `Product(100000, 100000)` -> `10000000000` (the int32 product would be `1410065408`)
// BIGO: O(1). Go reports overflow only for constants (`int32(math.MaxInt32 + 1)` will not compile); arithmetic on variables wraps without warning.
// NOTE: Convert one operand first: `int64(a) * int64(b)`. `int64(a * b)` converts the already-wrapped result.
func Wrapped() int32 {
	panic("todo")
}

func Product(a, b int32) int64 {
	panic("todo")
}

// ==== END

// ==== EX 2.2 | Detect overflow before it happens | Core
// PROMPT: Write `func MulFits(a, b int64) bool` that reports whether `a * b` fits in an int64, using `bits.Mul64` on the absolute values to get the full 128-bit product.
// EXAMPLE: `MulFits(3_000_000_000, 3)` -> `true`; `MulFits(4_000_000_000, 4_000_000_000)` -> `false`; `MulFits(math.MinInt64, 1)` -> `true`
// BIGO: O(1). `bits.Mul64` compiles to a single machine instruction.
func MulFits(a, b int64) bool {
	panic("todo")
}

func absU(x int64) uint64 {
	panic("todo")
}

// ==== END

// ==== EX 2.3 | Truncating division versus floor division | Core
// PROMPT: Go's `/` truncates toward zero and `%` takes the sign of the dividend. Write `func TruncDivMod(a, b int) (int, int)` with the operators and `func FloorDivMod(a, b int) (int, int)` that rounds the quotient down, as Python's `//` does.
// EXAMPLE: `TruncDivMod(-7, 2)` -> `-3, -1`; `FloorDivMod(-7, 2)` -> `-4, 1`
// BIGO: O(1).
// NOTE: `x % m` can be negative, so `arr[x % m]` panics for negative x. Write `((x % m) + m) % m` for a non-negative remainder.
func TruncDivMod(a, b int) (int, int) {
	panic("todo")
}

func FloorDivMod(a, b int) (int, int) {
	panic("todo")
}

// ==== END

// ==== EX 2.4 | Ceiling division | Core
// PROMPT: Write `func CeilDiv(a, b int) int` for b > 0 that works for negative a too, using your floor division.
// EXAMPLE: `CeilDiv(7, 2)` -> `4`; `CeilDiv(-7, 2)` -> `-3`; `CeilDiv(6, 3)` -> `2`
// BIGO: O(1). The common shortcut `(a + b - 1) / b` is only right for a >= 0.
func CeilDiv(a, b int) int {
	panic("todo")
}

// ==== END

// ==== EX 2.5 | Modular arithmetic | Core
// PROMPT: Write `const MOD = 1_000_000_007`, `func ModPow(base, exp, mod int) int` by repeated squaring, and `func ModInv(a int) int` using Fermat's little theorem.
// EXAMPLE: `ModPow(2, 10, MOD)` -> `1024`; `ModPow(2, 100, MOD)` -> `976371285`; `123456 * ModInv(123456) % MOD` -> `1`
// BIGO: O(log exp). Every product of two values below MOD is below 10^18 and fits in int.
const MOD = 1_000_000_007

func ModPow(base, exp, mod int) int {
	panic("todo")
}

func ModInv(a int) int {
	panic("todo")
}

// ==== END

// ==== EX 2.6 | math/big for huge numbers | Core
// PROMPT: Write `func FactorialDigits(n int) int` returning the number of digits of n! using `math/big`.
// EXAMPLE: `FactorialDigits(20)` -> `19` (still fits in an int64); `FactorialDigits(100)` -> `158`
// BIGO: O(N) multiplications of growing numbers. `big.Int` methods write into a receiver (`f.Mul(f, x)`) to avoid allocating a new number each time.
func FactorialDigits(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 2.7 | Never compare floats with == | Core
// PROMPT: Write `func Close(a, b, eps float64) bool`.
// EXAMPLE: `Close(0.1 + 0.2, 0.3, 1e-9)` -> `true`, while `0.1 + 0.2 == 0.3` (with variables) is `false`
// BIGO: O(1).
func Close(a, b, eps float64) bool {
	panic("todo")
}

// ==== END

// ==== EX 2.8 | An exact integer square root | Core
// PROMPT: Write `func Isqrt(n int) int` that starts from `int(math.Sqrt(float64(n)))` and corrects it up or down, because the float can be off by one for large n.
// EXAMPLE: `Isqrt(99)` -> `9`; `Isqrt(999_999_999_999_999_999)` -> `999999999`; `Isqrt(math.MaxInt64)` -> `3037000499`
// BIGO: O(1).
// NOTE: Compare with division (`r > n / r`) rather than `r * r > n`: near `math.MaxInt64` the square overflows and turns negative.
func Isqrt(n int) int {
	panic("todo")
}

// ==== END

// ==== EX 2.9 | Digits of an integer | Core
// PROMPT: Write `func DigitSum(n int) int` and `func ReverseDigits(n int) int` with `% 10` and `/ 10`, keeping the sign for the reverse.
// EXAMPLE: `DigitSum(9045)` -> `18`; `ReverseDigits(-1200)` -> `-21`
// BIGO: O(D) for D digits.
func DigitSum(n int) int {
	panic("todo")
}

func ReverseDigits(n int) int {
	panic("todo")
}

// ==== END
