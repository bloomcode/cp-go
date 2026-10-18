// Package ch18: Files and Streams.
package ch18


// ==== EX 18.1 | Sum the numbers in a file | Core
// PROMPT: Write `SumFile(path string) (int, error)` that returns the sum of the integers in a text file, one per line, skipping blank lines. Open it with `os.Open`, `defer f.Close()`, and read lines with a `bufio.Scanner`; return any open, parse or scan error.
// EXAMPLE: a file holding `3\n\n-1\n10\n` -> `12, nil`
// BIGO: O(size of the file) time, O(1) extra space when read line by line.
func SumFile(path string) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 18.2 | Most common words in a file | Core
// PROMPT: A word is a run of letters, compared in lower case. Create `type WordCount struct{ Word string; Count int }` and write `TopWords(path string, k int) ([]WordCount, error)` that returns the `k` most common words in a file, most frequent first and ties in alphabetical order. Count in a map while scanning lines (`strings.FieldsFunc` splits on non-letters), then sort.
// EXAMPLE: `The cat and the hat.\nThe end, and THE start!\n`, `2` -> `[{the 4} {and 2}]`
// BIGO: O(size + W log W) for W distinct words: one pass to count, one sort.
type WordCount struct {
	Word  string
	Count int
}

func TopWords(path string, k int) ([]WordCount, error) {
	panic("todo")
}

// ==== END

// ==== EX 18.3 | Write a report | Core
// PROMPT: Write `WriteReport(path string, scores map[string]int) (int, error)` that writes one line `name,score` per entry, highest score first and ties by name, replacing any existing file, and returns how many lines it wrote. Create the file with `os.Create`, write through a `bufio.Writer`, and remember that `Flush` and `Close` can fail too.
// EXAMPLE: `{ana: 7, bo: 9, cy: 7}` -> writes `bo,9\nana,7\ncy,7\n` and returns `3, nil`
// BIGO: O(N log N) for the sort, O(N) to write.
func WriteReport(path string, scores map[string]int) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 18.4 | Stream a large file | Practice
// PROMPT: Write `LongestLine(path string) (int, error)` that returns the length of the longest line in a file (not counting the newline), holding only one line in memory at a time with a `bufio.Reader`'s `ReadString('\n')` (a `bufio.Scanner` would fail on lines over 64 KB), never `os.ReadFile`.
// EXAMPLE: a file with lines `ab`, `abcd`, “ -> `4, nil`
// BIGO: O(size) time and O(longest line) memory: only one line is ever held, however big the file is.
func LongestLine(path string) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 18.5 | Copy matching lines to another file | Practice
// PROMPT: Write `CopyMatching(source, target, needle string) (int, error)` that writes every line of `source` containing `needle` to `target`, each followed by a newline, and returns how many it wrote. Scan the source and write through a `bufio.Writer`, closing both files and reporting the first error.
// EXAMPLE: lines `error: disk`, `ok`, `error: net` with needle `error` -> writes the two error lines, returns `2, nil`
// BIGO: O(size of the input) time, O(longest line) memory: both files are streamed.
func CopyMatching(source, target, needle string) (int, error) {
	panic("todo")
}

// ==== END

// ==== EX 18.6 | Averages from a CSV file | Practice
// PROMPT: A CSV file has the header `team,points` and one row per game. Write `AverageByTeam(path string) (map[string]float64, error)` that returns each team's average points. Read the records with `encoding/csv` (which handles quoted fields), skip the header, and keep a total and a count per team.
// EXAMPLE: rows `red,10`, `blue,4`, `red,7` -> `map[blue:4 red:8.5]`
// BIGO: O(size of the file) time, O(T) space for T teams: keep a running total and count per team.
func AverageByTeam(path string) (map[string]float64, error) {
	panic("todo")
}

// ==== END
