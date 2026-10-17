// Package ch08 - Structs, methods and interfaces.
package ch08


// ==== EX 8.1 | A struct with a pointer receiver | Core
// PROMPT: Write `type Counter struct{ n int }` with `func (c *Counter) Add(k int) error` (rejecting negative k) and `func (c *Counter) Value() int`.
// EXAMPLE: `var c Counter; c.Add(2); c.Add(5)` -> `c.Value()` is `7`; `c.Add(-1)` returns an error
// BIGO: O(1). The zero value `Counter{}` is ready to use, so no constructor is needed.
type Counter struct{ n int }

func (c *Counter) Add(k int) error {
	panic("todo")
}

func (c *Counter) Value() int { panic("todo") }

// ==== END

// ==== EX 8.2 | The value-receiver trap | Core
// PROMPT: Write `type Box struct{ N int }` with `func (b Box) IncValue()` and `func (b *Box) IncPointer()`. Call each on the same box and report what changed.
// EXAMPLE: `b := Box{}; b.IncValue(); b.IncPointer()` -> `b.N` is `1`: the value receiver changed a copy
// BIGO: O(1). A value receiver copies the whole struct on every call; use pointer receivers for methods that modify, and for large structs.
type Box struct{ N int }

func (b Box) IncValue() { panic("todo") }

func (b *Box) IncPointer() { panic("todo") }

// ==== END

// ==== EX 8.3 | Embedding promotes methods | Core
// PROMPT: Write `type Animal struct{ Name string }` with `func (a Animal) Describe() string`, and `type Dog struct{ Animal; Breed string }`. Call `Describe` on a Dog directly.
// EXAMPLE: `Dog{Animal{"Rex"}, "lab"}.Describe()` -> `"I am Rex"`
// BIGO: O(1). Embedding is composition: a Dog has an Animal and borrows its methods; it is not a subtype.
type Animal struct{ Name string }

func (a Animal) Describe() string { panic("todo") }

type Dog struct {
	Animal
	Breed string
}

// ==== END

// ==== EX 8.4 | Interfaces are satisfied implicitly | Core
// PROMPT: Write `type Shape interface{ Area() float64 }`, types `Rect{W, H float64}` and `Circle{R float64}` with `Area` methods, and `func TotalArea(shapes []Shape) float64`.
// EXAMPLE: `TotalArea([]Shape{Rect{2, 3}, Circle{1}})` -> `6 + pi`, about `9.1416`
// BIGO: O(N). No `implements` keyword: having the method is enough.
type Shape interface{ Area() float64 }

type Rect struct{ W, H float64 }

func (r Rect) Area() float64 { panic("todo") }

type Circle struct{ R float64 }

func (c Circle) Area() float64 { panic("todo") }

func TotalArea(shapes []Shape) float64 {
	panic("todo")
}

// ==== END

// ==== EX 8.5 | fmt.Stringer | Core
// PROMPT: Give `type Vec struct{ X, Y int }` a `String() string` method returning "(x, y)", and check that `fmt.Sprint` and `%v` use it automatically.
// EXAMPLE: `fmt.Sprint(Vec{1, 2})` -> `"(1, 2)"`
// BIGO: O(1). Do not call `fmt.Sprint(v)` inside `v.String()`: it calls `String` again forever.
type Vec struct{ X, Y int }

func (v Vec) String() string { panic("todo") }

// ==== END

// ==== EX 8.6 | Type switches | Core
// PROMPT: Write `func Kind(v any) string` with a type switch returning "int", "string", "[]int", "Shape" (for anything with an `Area` method) or "other".
// EXAMPLE: `Kind(3)` -> `"int"`; `Kind(Rect{1, 1})` -> `"Shape"`; `Kind(2.5)` -> `"other"`
// BIGO: O(1). Cases are tried in order, so put concrete types before interfaces.
func Kind(v any) string {
	panic("todo")
}

// ==== END

// ==== EX 8.7 | Custom error types | Core
// PROMPT: Write `type NotFoundError struct{ Key string }` with an `Error()` method, `func Find(m map[string]int, key string) (int, error)`, and show `errors.As` recovering the key from a wrapped error.
// EXAMPLE: `_, err := Find(m, "zz")`; `var nf *NotFoundError; errors.As(fmt.Errorf("lookup: %w", err), &nf)` -> `true`, `nf.Key` is `"zz"`
// BIGO: O(1).
type NotFoundError struct{ Key string }

func (e *NotFoundError) Error() string { panic("todo") }

func Find(m map[string]int, key string) (int, error) {
	panic("todo")
}

func KeyOf(err error) (string, bool) {
	panic("todo")
}

// ==== END

// ==== EX 8.8 | A validated value type | Practice
// PROMPT: Write `type Interval struct{ Lo, Hi int }` with `func NewInterval(lo, hi int) (Interval, error)` that rejects lo > hi, and methods `Len() int` and `Overlaps(o Interval) bool`.
// EXAMPLE: `a, _ := NewInterval(1, 5); b, _ := NewInterval(5, 9)` -> `a.Overlaps(b)` is `true`; `NewInterval(3, 1)` returns an error
// BIGO: O(1). Structs of comparable fields compare with `==` and work as map keys for free.
type Interval struct{ Lo, Hi int }

func NewInterval(lo, hi int) (Interval, error) {
	panic("todo")
}

func (iv Interval) Len() int { panic("todo") }

func (iv Interval) Overlaps(o Interval) bool { panic("todo") }

// ==== END
