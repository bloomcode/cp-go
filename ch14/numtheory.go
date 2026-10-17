// Package ch14 - Math and number theory.
package ch14


// ==== EX 14.1 | gcd and lcm without overflow | Core
// PROMPT: Write `func GCD(a, b int) int` with Euclid's algorithm and `func LCM(a, b int) int` that divides before multiplying.
// EXAMPLE: `GCD(84, 36)` -> `12`; `LCM(4, 6)` -> `12`; `LCM(1_000_000_000, 999_999_999)` -> `999999999000000000`
// BIGO: O(log min(a, b)). `a / GCD(a, b) * b` keeps the intermediate value small.
func GCD(a, b int) int {
	panic("todo")
}

func LCM(a, b int) int {
	panic("todo")
}

// ==== END

// ==== EX 14.2 | Sieve of Eratosthenes | Core
// PROMPT: Write `func Sieve(n int) []bool` where `composite[i]` is false exactly for primes, plus `func PrimesUpTo(n int) []int`.
// EXAMPLE: `PrimesUpTo(30)` -> `[2 3 5 7 11 13 17 19 23 29]`
// BIGO: O(N log log N).
func Sieve(n int) []bool {
	panic("todo")
}

func PrimesUpTo(n int) []int {
	panic("todo")
}

// ==== END

// ==== EX 14.3 | Prime factorization | Core
// PROMPT: Write `func Factorize(n int) [][2]int` returning `{prime, exponent}` pairs by trial division up to sqrt(n).
// EXAMPLE: `Factorize(360)` -> `[[2 3] [3 2] [5 1]]`
// BIGO: O(sqrt N). Loop with `p <= n/p` instead of `p*p <= n` to stay inside int.
func Factorize(n int) [][2]int {
	panic("todo")
}

// ==== END

// ==== EX 14.4 | Exact 128-bit modular multiplication | Practice
// PROMPT: For a modulus near 10^18, `a * b % m` overflows. Write `func MulMod(a, b, m uint64) uint64` with `bits.Mul64` (the full 128-bit product) and `bits.Rem64`, and `func PowMod(b, e, m uint64) uint64` on top of it.
// EXAMPLE: `MulMod(999_999_999_999_999_989, 2, 1_000_000_000_000_000_000)` -> `999999999999999978`
// BIGO: O(1) per multiply, O(log e) multiplies for PowMod. Go is one of the few languages with this in the standard library.
func MulMod(a, b, m uint64) uint64 {
	panic("todo")
}

func PowMod(b, e, m uint64) uint64 {
	panic("todo")
}

// ==== END

// ==== EX 14.5 | nCr modulo a prime | Core
// PROMPT: Write `type Binom struct{ fact, inv []int }` with `NewBinom(n int) *Binom` that precomputes factorials and inverse factorials mod 1_000_000_007, and `func (b *Binom) C(n, r int) int`.
// EXAMPLE: `NewBinom(10).C(10, 3)` -> `120`; `NewBinom(100000).C(100000, 50000)` -> `149033233`
// BIGO: O(N) precompute, O(1) per query.
const MOD = 1_000_000_007

type Binom struct{ fact, inv []int }

func NewBinom(n int) *Binom {
	panic("todo")
}

func (b *Binom) C(n, r int) int {
	panic("todo")
}

// ==== END

// ==== EX 14.6 | Extended Euclid | Practice
// PROMPT: Write `func ExtGCD(a, b int) (g, x, y int)` with a*x + b*y = g, and use it for `func ModInverse(a, m int) (int, bool)` when m is not prime.
// EXAMPLE: `ExtGCD(240, 46)` -> g `2` with 240*x + 46*y = 2; `ModInverse(3, 10)` -> `7, true`; `ModInverse(4, 10)` -> `0, false`
// BIGO: O(log min(a, b)).
func ExtGCD(a, b int) (g, x, y int) {
	panic("todo")
}

func ModInverse(a, m int) (int, bool) {
	panic("todo")
}

// ==== END

// ==== EX 14.7 | Divisors in sqrt time | Core
// PROMPT: Write `func Divisors(n int) []int` returning all divisors in ascending order by pairing d with n / d.
// EXAMPLE: `Divisors(36)` -> `[1 2 3 4 6 9 12 18 36]`
// BIGO: O(sqrt N).
func Divisors(n int) []int {
	panic("todo")
}

// ==== END

// ==== EX 14.8 | math/big when int is not enough | Core
// PROMPT: Write `func FibBig(n int) string` returning the nth Fibonacci number with `big.Int`, and `func ProbablyPrime(n string) bool` with `ProbablyPrime(20)`.
// EXAMPLE: `FibBig(100)` -> `"354224848179261915075"`; `ProbablyPrime("1000000000000000003")` -> `true`
// BIGO: O(N) additions of growing numbers. Reuse receivers (`a.Add(a, b)`) to avoid an allocation per step.
func FibBig(n int) string {
	panic("todo")
}

func ProbablyPrime(n string) bool {
	panic("todo")
}

// ==== END
