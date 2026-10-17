package ch03

import (
	"testing"
	"unicode/utf8"

	"cp/internal/check"
)

func TestEx3_1(t *testing.T) {
	check.Eq(t, []string{Reverse("golang"), Reverse("héllo")}, []string{"gnalog", "olléh"})
	check.True(t, utf8.ValidString(Reverse("héllo")), "valid utf8")
}

func TestEx3_2(t *testing.T) {
	nb, nr := Lengths("héllo")
	check.Eq(t, []int{nb, nr}, []int{6, 5})
	check.Eq(t, Runes("héllo")[1], 'é')
}

func TestEx3_3(t *testing.T) {
	c := LetterCounts("banana")
	check.Eq(t, []int{c[0], c[1], c[13], c[25]}, []int{3, 1, 2, 0})
}

func TestEx3_4(t *testing.T) {
	check.Eq(t, []string{Caesar("xyz abc", 3), Caesar("abc", -1), Caesar("abc", -27)}, []string{"abc def", "zab", "zab"})
}

func TestEx3_5(t *testing.T) {
	check.True(t, IsPalindrome("A man, a plan, a canal: Panama") && !IsPalindrome("race a car"), "palindrome")
}

func TestEx3_6(t *testing.T) {
	check.Eq(t, ReverseWords("  the sky   is blue "), "blue is sky the")
}

func TestEx3_7(t *testing.T) {
	check.Eq(t, NumbersCSV(5), "1,2,3,4,5")
}

func TestEx3_8(t *testing.T) {
	check.True(t, IsAnagram("listen", "silent") && !IsAnagram("rat", "car"), "anagram")
}

func TestEx3_9(t *testing.T) {
	check.Eq(t, []string{RLE("aaabccdddd"), RLE("aaaaaaaaaaaa")}, []string{"a3b1c2d4", "a12"})
}
