// Package ch17: Errors.
package ch17

import (
	"errors"
)

// ==== EX 17.1 | Parse or report an error | Core
// PROMPT: Write `ParsePort(text string) (int, error)` that returns the port number in `text` (spaces around it allowed). Return an error reading `not a number: "text"` when it is not an integer, and `out of range: N` when it is outside 1 to 65535. Errors are values: return them, don't panic.
// EXAMPLE: `ParsePort(" 8080 ")` -> `8080, nil`; `ParsePort("http")` -> `0, not a number: "http"`
// BIGO: O(L) for a string of length L.
func ParsePort(text string) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 17.2 | A custom error type | Core
// PROMPT: Create `type InsufficientFundsError struct{ Needed, Available int }` with an `Error() string` method (on the pointer) reading `needed N, available M`. Write `Withdraw(balance, amount int) (int, error)` that returns the new balance, or an `*InsufficientFundsError` when `amount` is more than `balance`. Callers find it with `errors.As`.
// EXAMPLE: `Withdraw(100, 30)` -> `70, nil`; `Withdraw(20, 50)` -> an error with `Needed == 50`, `Available == 20`
// BIGO: O(1).
type InsufficientFundsError struct{ Needed, Available int }

func (e *InsufficientFundsError) Error() string {
	panic("todo")
}

func Withdraw(balance, amount int) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 17.3 | Wrap an error with context | Practice
// PROMPT: Declare `var ErrMissing = errors.New("missing")`. Write `NumberSetting(settings map[string]string, key string) (int, error)` that parses `settings[key]`, but wraps the failure with `fmt.Errorf("setting %q: %w", key, err)`, using `ErrMissing` when the key is absent and the `strconv` error when the number is bad. `errors.Is` then sees the cause through the context.
// EXAMPLE: `NumberSetting(map[string]string{"port": "80"}, "port")` -> `80, nil`; with no `port` -> `setting "port": missing`, and `errors.Is(err, ErrMissing)`
// BIGO: O(L) to parse the value.
var ErrMissing = errors.New("missing")

func NumberSetting(settings map[string]string, key string) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 17.4 | Cleanup that always runs | Core
// PROMPT: Write `RunSteps(steps []func() error, log *[]string) error` that appends `"start"` to `*log`, runs the steps until one returns an error (return it), and appends `"done"` however that ends, even if a step panics. A `defer` registered at the start runs on every way out.
// EXAMPLE: two steps that log `a` and `b` -> `*log` is `[start a b done]`
// BIGO: O(S) for S steps.
func RunSteps(steps []func() error, log *[]string) error {
	panic("todo")
}

// ==== END

// ==== EX 17.5 | Close resources in reverse order | Practice
// PROMPT: Create `type Tracked struct` with `Open(name string, log *[]string) *Tracked`, which appends `open name`, and a `Close()` method that appends `close name`. Write `UseTwo(log *[]string, fail bool) error` that opens `a` then `b`, deferring each `Close` right after opening it, appends `work`, and returns an error when `fail` is true. Deferred calls run last-in first-out, so `b` closes before `a`.
// EXAMPLE: `UseTwo(&log, false)` -> `log` is `[open a open b work close b close a]`
// BIGO: O(1).
type Tracked struct {
	name string
	log  *[]string
}

func Open(name string, log *[]string) *Tracked {
	panic("todo")
}

func (t *Tracked) Close() { panic("todo") }

func UseTwo(log *[]string, fail bool) error {
	panic("todo")
}

// ==== END

// ==== EX 17.6 | Collect every error instead of stopping | Practice
// PROMPT: Write `ParseAges(rows []string) ([]int, error)` that parses each row as an age from 0 to 150, returning the good ages in order (a non-nil slice) and, if any row was bad, one error built with `errors.Join` from `row N: not a number: text` or `row N: out of range: A` errors (rows count from 1, spaces trimmed); `nil` when all rows are good.
// EXAMPLE: `["34" "x" "200" " 7 "]` -> `[34 7]` and an error printing `row 2: not a number: x` and `row 3: out of range: 200` on two lines
// BIGO: O(total length of the rows).
func ParseAges(rows []string) ([]int, error) {
	panic("todo")
}

// ==== END
