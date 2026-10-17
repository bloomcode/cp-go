// Package ch01 - Fast input and output.
// Judges pipe text through standard input. Every exercise takes the input as a string
// (or an io.Reader) so it can be tested; in a contest you pass os.Stdin instead.
package ch01

import (
	"bufio"
	"io"
)

// ==== EX 1.1 | Read every integer with bufio.Scanner | Core
// PROMPT: Write `func ReadInts(r io.Reader) []int` using a `bufio.Scanner` split with `bufio.ScanWords`, converting each token with `strconv.Atoi`.
// EXAMPLE: `ReadInts(strings.NewReader("3 -4\n10"))` -> `[3 -4 10]`
// BIGO: O(input). With the default `ScanLines`, a line longer than 64 KB stops the scanner; call `sc.Buffer(make([]byte, 1<<20), 1<<26)` before reading huge lines.
// NOTE: `ScanWords` ignores line breaks, which protects you from judges that wrap input unexpectedly.
func ReadInts(r io.Reader) []int {
	panic("todo")
}

// ==== END

// ==== EX 1.2 | A byte-level FastReader | Core
// PROMPT: Write `type FastReader struct` wrapping a `*bufio.Reader`, with `NewFastReader(r io.Reader)` and `func (f *FastReader) Int() int` that skips whitespace and parses digits (and a leading minus) byte by byte, with no allocation per number.
// EXAMPLE: `f := NewFastReader(strings.NewReader("3\n10 -20\n30"))` -> `f.Int()` returns `3`, then `10`, `-20`, `30`
// BIGO: O(1) amortized per digit. Several times faster than `fmt.Fscan`, which uses reflection for every value.
type FastReader struct {
	r *bufio.Reader
}

func NewFastReader(r io.Reader) *FastReader {
	panic("todo")
}

func (f *FastReader) Int() int {
	panic("todo")
}

// ==== END

// ==== EX 1.3 | N, then N numbers | Core
// PROMPT: Write `func ReadArray(input string) []int`: the first number is N, followed by N integers on any number of lines. Use your `FastReader` and allocate the slice once with `make([]int, n)`.
// EXAMPLE: `ReadArray("4\n5 1 9 2")` -> `[5 1 9 2]`
// BIGO: O(N).
func ReadArray(input string) []int {
	panic("todo")
}

// ==== END

// ==== EX 1.4 | Read a character grid | Core
// PROMPT: Write `func ReadGrid(input string) [][]byte` for input "R C" followed by R rows of C characters, using a `bufio.Scanner` with `ScanWords`. Keep rows as `[]byte` so cells can be changed.
// EXAMPLE: `ReadGrid("2 3\n#.#\n..#")` -> rows `"#.#"` and `"..#"` as byte slices
// BIGO: O(R * C).
// NOTE: `sc.Bytes()` points into the scanner's buffer and is overwritten on the next `Scan`; copy it (`[]byte(sc.Text())` does) before keeping it.
func ReadGrid(input string) [][]byte {
	panic("todo")
}

// ==== END

// ==== EX 1.5 | T test cases into one bufio.Writer | Core
// PROMPT: Write `func SumCases(input string) string`: first T, then T lines of `a b` (each up to 10^18). Print each `a + b` on its own line through one `bufio.Writer` wrapped around a `strings.Builder`, and flush once at the end.
// EXAMPLE: `SumCases("2\n1 2\n1000000000000 1")` -> `"3\n1000000000001\n"`
// BIGO: O(T). Writing to `os.Stdout` directly makes a system call per print, the most common reason a correct Go solution times out.
// NOTE: `int` is 64 bits on every judge that runs amd64 or arm64, so 10^18 fits; `int32` would not.
func SumCases(input string) string {
	panic("todo")
}

// ==== END

// ==== EX 1.6 | Format numbers exactly | Core
// PROMPT: Write `func Fmt(x float64) string` with exactly 6 decimals and `func FmtRow(name string, score int) string` with the name left-aligned in 10 characters and the score right-aligned in 5.
// EXAMPLE: `Fmt(2.0 / 3)` -> `"0.666667"`; `FmtRow("ann", 42)` -> `"ann          42"`
// BIGO: O(1). `%v` on a float prints the shortest representation (`0.6666666666666666`), which judges expecting 6 places reject.
func Fmt(x float64) string {
	panic("todo")
}

func FmtRow(name string, score int) string {
	panic("todo")
}

// ==== END

// ==== EX 1.7 | One line of output without fmt | Core
// PROMPT: Write `func JoinLine(values []int) string` that builds "3 1 2" with `strconv.AppendInt` into a single `[]byte`, adding spaces between values only.
// EXAMPLE: `JoinLine([]int{3, 1, 2})` -> `"3 1 2"`
// BIGO: O(N), with one growing buffer and no per-number strings.
func JoinLine(values []int) string {
	panic("todo")
}

// ==== END

// ==== EX 1.8 | The contest template | Core
// PROMPT: Write `func Solve(r io.Reader, w io.Writer)` for "print the maximum of N numbers", using `FastReader` and a `bufio.Writer` with `defer out.Flush()`, plus `func Run(input string) string` for testing. In a contest `main` is just `Solve(os.Stdin, os.Stdout)`. Keep this as your template.
// EXAMPLE: `Run("5\n3 9 -1 4 7")` -> `"9\n"`
// BIGO: O(N). Forgetting `Flush` prints nothing at all; `defer` makes it impossible to forget.
func Solve(r io.Reader, w io.Writer) {
	panic("todo")
}

func Run(input string) string {
	panic("todo")
}

// ==== END
