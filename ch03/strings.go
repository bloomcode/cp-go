// Package ch03 - Strings, bytes and runes.
package ch03


// ==== EX 3.1 | Reverse by runes, not bytes | Core
// PROMPT: Write `func Reverse(s string) string` that converts to `[]rune`, swaps from both ends and converts back, so multi-byte characters survive.
// EXAMPLE: `Reverse("golang")` -> `"gnalog"`; `Reverse("héllo")` -> `"olléh"`
// BIGO: O(N). Reversing the bytes instead would split the two bytes of `é` and produce invalid UTF-8.
func Reverse(s string) string {
	panic("todo")
}

// ==== END

// ==== EX 3.2 | len counts bytes | Core
// PROMPT: Write `func Lengths(s string) (bytes, runes int)` using `len` and `utf8.RuneCountInString`, and `func Runes(s string) []rune` collected with `for _, r := range s`.
// EXAMPLE: `Lengths("héllo")` -> `6, 5`; indexing `s[1]` gives a byte (195), while `range` gives the rune `'é'`
// BIGO: `len` is O(1); counting runes is O(N). Contest input is almost always ASCII, where bytes and runes coincide.
func Lengths(s string) (bytes, runes int) {
	panic("todo")
}

func Runes(s string) []rune {
	panic("todo")
}

// ==== END

// ==== EX 3.3 | Count letters with a fixed array | Core
// PROMPT: Write `func LetterCounts(s string) [26]int` counting 'a'..'z' by indexing with `s[i] - 'a'`.
// EXAMPLE: `LetterCounts("banana")[0]` -> `3` (a); index 1 (b) -> `1`; index 13 (n) -> `2`
// BIGO: O(N), and the array is a value: returning it copies 26 ints, no allocation.
func LetterCounts(s string) [26]int {
	panic("todo")
}

// ==== END

// ==== EX 3.4 | Caesar shift on bytes | Core
// PROMPT: Write `func Caesar(s string, k int) string` shifting lowercase letters by k (which may be negative) and leaving other bytes alone, building the result in a `[]byte`.
// EXAMPLE: `Caesar("xyz abc", 3)` -> `"abc def"`; `Caesar("abc", -1)` -> `"zab"`
// BIGO: O(N). Wrap with `((x % 26) + 26) % 26` because `%` can return a negative value.
func Caesar(s string, k int) string {
	panic("todo")
}

// ==== END

// ==== EX 3.5 | Palindrome with two pointers | Core
// PROMPT: Write `func IsPalindrome(s string) bool` that ignores case and anything that is not a letter or digit, using two indices over a `[]rune` and the `unicode` package.
// EXAMPLE: `IsPalindrome("A man, a plan, a canal: Panama")` -> `true`
// BIGO: O(N) time.
func IsPalindrome(s string) bool {
	panic("todo")
}

// ==== END

// ==== EX 3.6 | Fields and Join | Core
// PROMPT: Write `func ReverseWords(s string) string` with `strings.Fields` (which splits on any run of whitespace) and `strings.Join`.
// EXAMPLE: `ReverseWords("  the sky   is blue ")` -> `"blue is sky the"`
// BIGO: O(N). `strings.Split(s, " ")` would produce empty strings for repeated spaces.
func ReverseWords(s string) string {
	panic("todo")
}

// ==== END

// ==== EX 3.7 | Build strings with strings.Builder | Core
// PROMPT: Write `func NumbersCSV(n int) string` returning "1,2,...,n" with a `strings.Builder` and `strconv.Itoa`.
// EXAMPLE: `NumbersCSV(5)` -> `"1,2,3,4,5"`
// BIGO: O(N) amortized. `s += x` in a loop copies the whole string every time: O(N^2).
func NumbersCSV(n int) string {
	panic("todo")
}

// ==== END

// ==== EX 3.8 | Anagram check with counts | Core
// PROMPT: Write `func IsAnagram(a, b string) bool` for lowercase words with one `[26]int`: add for a, subtract for b, then check every count is zero.
// EXAMPLE: `IsAnagram("listen", "silent")` -> `true`; `IsAnagram("rat", "car")` -> `false`
// BIGO: O(N), versus O(N log N) for sorting both.
func IsAnagram(a, b string) bool {
	panic("todo")
}

// ==== END

// ==== EX 3.9 | Run-length encoding | Core
// PROMPT: Write `func RLE(s string) string` turning runs of the same byte into byte-then-count.
// EXAMPLE: `RLE("aaabccdddd")` -> `"a3b1c2d4"`
// BIGO: O(N).
// NOTE: `string(rune(3))` is the character with code 3, not "3". Use `strconv.Itoa` for numbers.
func RLE(s string) string {
	panic("todo")
}

// ==== END
