// Package ch15 - Writing fast Go.
package ch15

import (
	"io"
)

// ==== EX 15.1 | fmt.Fscan is slow; buffer it at least | Core
// PROMPT: Write `func SumFscan(r io.Reader) int` that reads N then N numbers with `fmt.Fscan` on a `bufio.Reader`, and compare it with the byte-level reader from 1.2 on 10^6 numbers.
// EXAMPLE: `SumFscan(strings.NewReader("3\n10 20 30"))` -> `60`
// BIGO: O(input) either way, but `fmt.Fscan` straight on `os.Stdin` (no `bufio`) makes a system call per read and is far slower still.
func SumFscan(r io.Reader) int {
	panic("todo")
}

// ==== END

// ==== EX 15.2 | One buffered writer for all output | Core
// PROMPT: Write `func SquaresOutput(n int) string` producing lines "i i*i" for 1..n through one `bufio.Writer` using `strconv.AppendInt` on a scratch buffer, never `fmt.Println` per line.
// EXAMPLE: `SquaresOutput(3)` -> `"1 1\n2 4\n3 9\n"`
// BIGO: O(output).
func SquaresOutput(n int) string {
	panic("todo")
}

// ==== END

// ==== EX 15.3 | Preallocate with make | Core
// PROMPT: Write `func Evens(n int) []int` that collects 0, 2, 4, ... below n into a slice made with `make([]int, 0, n/2+1)`, so `append` never reallocates.
// EXAMPLE: `Evens(9)` -> `[0 2 4 6 8]`, and `cap` never changes while appending
// BIGO: O(N). Without the capacity hint, append grows by doubling: still O(N) amortized, but with about log N copies and garbage.
func Evens(n int) []int {
	panic("todo")
}

// ==== END

// ==== EX 15.4 | A counting slice beats a map | Core
// PROMPT: Write `func CountsSlice(a []int, maxValue int) []int` using a `[]int` of buckets, and compare its speed with a `map[int]int` version on 10^6 values.
// EXAMPLE: `CountsSlice([]int{2, 0, 2, 3}, 3)` -> `[1 0 2 1]`
// BIGO: Both O(N), but a slice index is one memory access while a map lookup hashes the key and probes buckets: often 10 times faster.
func CountsSlice(a []int, maxValue int) []int {
	panic("todo")
}

// ==== END

// ==== EX 15.5 | Flatten a 2D grid | Core
// PROMPT: Write `func SumFlat(flat []int, rows, cols int) int` that sums row by row over a 1D slice indexed `r*cols + c`, and explain why row-major traversal is faster than column-major.
// EXAMPLE: rows 2, cols 3, flat `[1 2 3 4 5 6]` -> `21`
// BIGO: O(R * C). Walking memory in order keeps the CPU cache full; jumping a whole row per step misses it.
func SumFlat(flat []int, rows, cols int) int {
	panic("todo")
}

// ==== END

// ==== EX 15.6 | strings.Builder with Grow | Core
// PROMPT: Write `func Repeat(s string, n int) string` with a `strings.Builder`, calling `sb.Grow(len(s) * n)` first, and compare it with `+=` in a loop.
// EXAMPLE: `Repeat("ab", 3)` -> `"ababab"`
// BIGO: O(N) with one allocation; `+=` is O(N^2) because every step copies the whole string so far.
func Repeat(s string, n int) string {
	panic("todo")
}

// ==== END

// ==== EX 15.7 | Graphs in flat slices: forward-star | Stretch
// PROMPT: Write `func CSR(n int, edges [][2]int) (start, to []int)` building directed adjacency: `start` has n + 1 offsets and the neighbours of u are `to[start[u]:start[u+1]]`.
// EXAMPLE: n 3, edges `[[0 1] [0 2] [2 1]]` -> start `[0 2 2 3]`, to `[1 2 1]`
// BIGO: O(N + E) with two slices and no per-node allocation, versus a `[][]int` with one slice per node.
func CSR(n int, edges [][2]int) (start, to []int) {
	panic("todo")
}

// ==== END

// ==== EX 15.8 | The operations budget | Core
// PROMPT: Write `func Verdict(n int, complexity string) string` estimating operations for "n", "nlogn", "n2" or "n3" and returning "fast" (<= 10^8), "tight" (<= 10^9) or "too slow".
// EXAMPLE: `Verdict(200_000, "nlogn")` -> `"fast"`; `Verdict(20_000, "n2")` -> `"tight"`; `Verdict(100_000, "n2")` -> `"too slow"`
// BIGO: O(1). Compiled Go does roughly 10^8 to 10^9 simple operations a second, about the same as Java and C++ without the warm-up.
func Verdict(n int, complexity string) string {
	panic("todo")
}

// ==== END
