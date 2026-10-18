package ch19

import (
	"cp/internal/check"
	"sync/atomic"
	"testing"
	"time"
)

func TestEx19_1(t *testing.T) {
	check.Eq(t, ParallelSquares([]int{1, 2, 3}), []int{1, 4, 9})
	check.Eq(t, ParallelSquares([]int{}), []int{})
	many, want := make([]int, 50), make([]int, 50)
	for i := range many {
		many[i], want[i] = i, i*i
	}
	check.Eq(t, ParallelSquares(many), want)
}

func TestEx19_2(t *testing.T) {
	many := make([]int, 10_000)
	for i := range many {
		many[i] = i
	}
	check.Eq(t, []int{PipelineTotal([]int{4, 5, 6}), PipelineTotal(nil), PipelineTotal(many)}, []int{15, 0, 49_995_000})
}

func TestEx19_3(t *testing.T) {
	check.Eq(t, PoolMap(func(x int) int { return x * 10 }, []int{1, 2, 3}, 2), []int{10, 20, 30})
	check.Eq(t, PoolMap(func(x int) int { return x }, []int{}, 3), []int{})
	var active, peak atomic.Int32
	items, want := make([]int, 12), make([]int, 12)
	for i := range items {
		items[i], want[i] = i, -i
	}
	got := PoolMap(func(x int) int {
		n := active.Add(1)
		for p := peak.Load(); n > p && !peak.CompareAndSwap(p, n); p = peak.Load() {
		}
		time.Sleep(10 * time.Millisecond)
		active.Add(-1)
		return -x
	}, items, 3)
	check.Eq(t, got, want)
	check.True(t, peak.Load() >= 1 && peak.Load() <= 3, "never more than 3 at once")
}

func TestEx19_4(t *testing.T) {
	check.Eq(t, []int{LockedCount(4, 1000), LockedCount(1, 0), LockedCount(8, 25_000)}, []int{4000, 0, 200_000})
}

func TestEx19_5(t *testing.T) {
	v, ok := ResultWithin(func() int { return 42 }, time.Second)
	check.True(t, v == 42 && ok, "fast task")
	start := time.Now()
	v, ok = ResultWithin(func() int { time.Sleep(300 * time.Millisecond); return 1 }, 50*time.Millisecond)
	check.True(t, v == 0 && !ok, "slow task times out")
	check.True(t, time.Since(start) < 250*time.Millisecond, "did not wait for the slow task")
}

func TestEx19_6(t *testing.T) {
	hundred, thousand := make([]int, 100), make([]int, 1000)
	for i := range hundred {
		hundred[i] = i + 1
	}
	for i := range thousand {
		thousand[i] = i
	}
	check.Eq(t, []int{ParallelSum(hundred, 4), ParallelSum(nil, 3), ParallelSum([]int{5}, 8), ParallelSum(thousand, 7)}, []int{5050, 0, 5, 499_500})
}
