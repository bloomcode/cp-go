package ch18

import (
	"cp/internal/check"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEx18_1(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "n.txt")
	os.WriteFile(path, []byte("3\n\n-1\n10\n"), 0o644)
	total, err := SumFile(path)
	check.True(t, total == 12 && err == nil, "sum")
	_, err = SumFile(filepath.Join(dir, "missing.txt"))
	check.True(t, errors.Is(err, os.ErrNotExist), "missing file reported")
	os.WriteFile(path, []byte("1\nx\n"), 0o644)
	_, err = SumFile(path)
	check.True(t, err != nil, "bad number reported")
}

func TestEx18_2(t *testing.T) {
	path := filepath.Join(t.TempDir(), "w.txt")
	os.WriteFile(path, []byte("The cat and the hat.\nThe end, and THE start!\n"), 0o644)
	top, err := TopWords(path, 2)
	check.True(t, err == nil, "read")
	check.Eq(t, top, []WordCount{{"the", 4}, {"and", 2}})
	top, _ = TopWords(path, 4)
	check.Eq(t, top, []WordCount{{"the", 4}, {"and", 2}, {"cat", 1}, {"end", 1}})
	top, _ = TopWords(path, 99)
	check.Eq(t, len(top), 6)
}

func TestEx18_3(t *testing.T) {
	path := filepath.Join(t.TempDir(), "report.csv")
	os.WriteFile(path, []byte("old contents\n"), 0o644)
	n, err := WriteReport(path, map[string]int{"ana": 7, "bo": 9, "cy": 7})
	check.True(t, n == 3 && err == nil, "three lines")
	data, _ := os.ReadFile(path)
	check.Eq(t, string(data), "bo,9\nana,7\ncy,7\n")
	n, _ = WriteReport(path, map[string]int{})
	data, _ = os.ReadFile(path)
	check.True(t, n == 0 && len(data) == 0, "empty report")
}

func TestEx18_4(t *testing.T) {
	dir := t.TempDir()
	small := filepath.Join(dir, "a.txt")
	os.WriteFile(small, []byte("ab\nabcd\n\n"), 0o644)
	var b strings.Builder
	for i := range 100_000 {
		b.WriteString(strings.Repeat("x", i%97) + "\n")
	}
	b.WriteString(strings.Repeat("y", 70_000))
	big := filepath.Join(dir, "big.txt")
	os.WriteFile(big, []byte(b.String()), 0o644)
	n, err := LongestLine(small)
	check.True(t, n == 4 && err == nil, "small file")
	n, err = LongestLine(big)
	check.True(t, n == 70_000 && err == nil, "long last line")
}

func TestEx18_5(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "log.txt")
	os.WriteFile(log, []byte("error: disk\nok\nerror: net"), 0o644)
	n, err := CopyMatching(log, filepath.Join(dir, "out.txt"), "error")
	check.True(t, n == 2 && err == nil, "two lines copied")
	data, _ := os.ReadFile(filepath.Join(dir, "out.txt"))
	check.Eq(t, string(data), "error: disk\nerror: net\n")
	n, _ = CopyMatching(log, filepath.Join(dir, "none.txt"), "panic")
	check.Eq(t, n, 0)
	_, err = CopyMatching(filepath.Join(dir, "missing.txt"), filepath.Join(dir, "x.txt"), "a")
	check.True(t, errors.Is(err, os.ErrNotExist), "missing source")
}

func TestEx18_6(t *testing.T) {
	dir := t.TempDir()
	write := func(name, text string) string {
		path := filepath.Join(dir, name)
		os.WriteFile(path, []byte(text), 0o644)
		return path
	}
	avg, err := AverageByTeam(write("games.csv", "team,points\nred,10\nblue,4\nred,7\n"))
	check.True(t, err == nil, "read")
	check.Eq(t, avg, map[string]float64{"red": 8.5, "blue": 4})
	avg, _ = AverageByTeam(write("none.csv", "team,points\n"))
	check.Eq(t, avg, map[string]float64{})
	avg, _ = AverageByTeam(write("quoted.csv", "team,points\n\"red, the best\",3\n"))
	check.Eq(t, avg, map[string]float64{"red, the best": 3})
}
