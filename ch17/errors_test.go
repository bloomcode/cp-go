package ch17

import (
	"cp/internal/check"
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestEx17_1(t *testing.T) {
	port, err := ParsePort(" 8080 ")
	check.Eq(t, port, 8080)
	check.True(t, err == nil, "valid port")
	for text, want := range map[string]string{"http": `not a number: "http"`, "0": "out of range: 0", "70000": "out of range: 70000"} {
		_, err := ParsePort(text)
		check.True(t, err != nil && err.Error() == want, want)
	}
}

func TestEx17_2(t *testing.T) {
	left, err := Withdraw(100, 30)
	check.True(t, left == 70 && err == nil, "enough money")
	_, err = Withdraw(20, 50)
	var funds *InsufficientFundsError
	check.True(t, errors.As(err, &funds), "errors.As finds the type")
	check.Eq(t, []int{funds.Needed, funds.Available}, []int{50, 20})
	check.Eq(t, err.Error(), "needed 50, available 20")
	check.True(t, errors.As(fmt.Errorf("withdraw: %w", err), &funds), "found through wrapping too")
}

func TestEx17_3(t *testing.T) {
	n, err := NumberSetting(map[string]string{"port": "80"}, "port")
	check.True(t, n == 80 && err == nil, "valid setting")
	_, err = NumberSetting(map[string]string{}, "port")
	check.True(t, errors.Is(err, ErrMissing), "missing key")
	check.Eq(t, err.Error(), `setting "port": missing`)
	_, err = NumberSetting(map[string]string{"port": "eighty"}, "port")
	check.True(t, errors.Is(err, strconv.ErrSyntax) && !errors.Is(err, ErrMissing), "bad number keeps the strconv cause")
}

func TestEx17_4(t *testing.T) {
	var log []string
	step := func(name string) func() error { return func() error { log = append(log, name); return nil } }
	check.True(t, RunSteps([]func() error{step("a"), step("b")}, &log) == nil, "no error")
	check.Eq(t, log, []string{"start", "a", "b", "done"})
	log = nil
	boom := errors.New("boom")
	err := RunSteps([]func() error{step("a"), func() error { return boom }, step("never")}, &log)
	check.True(t, err == boom, "first error returned")
	check.Eq(t, log, []string{"start", "a", "done"})
	log = nil
	func() {
		defer func() { recover() }()
		RunSteps([]func() error{func() error { panic("crash") }}, &log)
	}()
	check.Eq(t, log, []string{"start", "done"})
}

func TestEx17_5(t *testing.T) {
	want := []string{"open a", "open b", "work", "close b", "close a"}
	var log []string
	check.True(t, UseTwo(&log, false) == nil, "no error")
	check.Eq(t, log, want)
	log = nil
	check.True(t, UseTwo(&log, true) != nil, "error returned")
	check.Eq(t, log, want)
}

func TestEx17_6(t *testing.T) {
	values, err := ParseAges([]string{"34", "x", "200", " 7 "})
	check.Eq(t, values, []int{34, 7})
	check.Eq(t, err.Error(), "row 2: not a number: x\nrow 3: out of range: 200")
	values, err = ParseAges(nil)
	check.True(t, len(values) == 0 && err == nil, "no rows")
	values, err = ParseAges([]string{"0", "150"})
	check.True(t, err == nil, "all good")
	check.Eq(t, values, []int{0, 150})
	_, err = ParseAges([]string{"-1"})
	check.Eq(t, err.Error(), "row 1: out of range: -1")
}
