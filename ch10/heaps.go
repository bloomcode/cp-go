// Package ch10 - Stacks, queues and container/heap.
package ch10

import (
	"container/list"
)

// ==== EX 10.1 | A slice as a stack | Core
// PROMPT: Write `func Balanced(s string) bool` for brackets `()[]{}` using a `[]byte` stack: push with `append`, pop with `st = st[:len(st)-1]`.
// EXAMPLE: `Balanced("{[()]}")` -> `true`; `Balanced("([)]")` -> `false`
// BIGO: O(N).
func Balanced(s string) bool {
	panic("todo")
}

// ==== END

// ==== EX 10.2 | A slice as a queue: BFS on a grid | Core
// PROMPT: Write `func ShortestPath(grid []string) int` returning the fewest moves from the top-left to the bottom-right cell through '.' cells (4 directions), or -1. Use a slice queue with a `head` index instead of re-slicing.
// EXAMPLE: `[]string{"..#", "#..", "#.."}` -> `4`
// BIGO: O(R * C).
func ShortestPath(grid []string) int {
	panic("todo")
}

// ==== END

// ==== EX 10.3 | container/heap: a min-heap of ints | Core
// PROMPT: Write `type IntHeap []int` implementing `heap.Interface` (`Len`, `Less`, `Swap`, `Push`, `Pop`), then `func DrainAscending(a []int) []int` that pushes everything with `heap.Push` and pops with `heap.Pop`.
// EXAMPLE: `DrainAscending([]int{5, 1, 4})` -> `[1 4 5]`
// BIGO: O(N log N).
// NOTE: Call the package functions `heap.Push(h, x)` and `heap.Pop(h)`, never your own `h.Push` directly: only the package versions restore heap order.
type IntHeap []int

func (h IntHeap) Len() int           { panic("todo") }
func (h IntHeap) Less(i, j int) bool { panic("todo") }
func (h IntHeap) Swap(i, j int)      { panic("todo") }
func (h *IntHeap) Push(x any)        { panic("todo") }
func (h *IntHeap) Pop() any {
	panic("todo")
}

func DrainAscending(a []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 10.4 | A max-heap by flipping Less | Core
// PROMPT: Write `type MaxHeap struct{ IntHeap }` that overrides only `Less`, and `func DrainDescending(a []int) []int`. Use `heap.Init` to heapify the whole slice in O(N) instead of N pushes.
// EXAMPLE: `DrainDescending([]int{5, 1, 4})` -> `[5 4 1]`
// BIGO: O(N) to `Init`, O(log N) per pop.
type MaxHeap struct{ IntHeap }

func (h MaxHeap) Less(i, j int) bool { panic("todo") }

func DrainDescending(a []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 10.5 | K largest with a size-k heap | Core
// PROMPT: Write `func KLargest(a []int, k int) []int` keeping an `IntHeap` of size k, returned in descending order.
// EXAMPLE: `KLargest([]int{7, 2, 9, 4, 9, 1}, 3)` -> `[9 9 7]`
// BIGO: O(N log k).
func KLargest(a []int, k int) []int {
	panic("todo")
}

// ==== END

// ==== EX 10.6 | A priority queue of structs | Core
// PROMPT: Write `type Job struct{ Cost, ID int }` and `type JobHeap []Job` ordered by cost then ID, and `func ProcessByCost(jobs []Job) []int` returning IDs in processing order.
// EXAMPLE: `[{5 1} {2 7} {5 0}]` -> `[7 0 1]`
// BIGO: O(N log N). Dijkstra's queue (16.7) is this same heap.
type Job struct{ Cost, ID int }

type JobHeap []Job

func (h JobHeap) Len() int { panic("todo") }
func (h JobHeap) Less(i, j int) bool {
	panic("todo")
}
func (h JobHeap) Swap(i, j int) { panic("todo") }
func (h *JobHeap) Push(x any)   { panic("todo") }
func (h *JobHeap) Pop() any {
	panic("todo")
}

func ProcessByCost(jobs []Job) []int {
	panic("todo")
}

// ==== END

// ==== EX 10.7 | Running median with two heaps | Practice
// PROMPT: Write `func RunningMedian(a []int) []float64`: after each value, the median so far. Keep a `MaxHeap` for the low half and an `IntHeap` for the high half.
// EXAMPLE: `RunningMedian([]int{5, 15, 1, 3})` -> `[5 10 5 4]`
// BIGO: O(N log N).
func RunningMedian(a []int) []float64 {
	panic("todo")
}

// ==== END

// ==== EX 10.8 | Sliding window maximum | Practice
// PROMPT: Write `func WindowMax(a []int, k int) []int` keeping a slice of indices whose values decrease from front to back (a monotonic deque), dropping from the front when an index leaves the window.
// EXAMPLE: `WindowMax([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)` -> `[3 3 5 5 6 7]`
// BIGO: O(N): each index enters and leaves once.
func WindowMax(a []int, k int) []int {
	panic("todo")
}

// ==== END

// ==== EX 10.9 | An LRU cache with container/list | Practice
// PROMPT: Write `type LRU struct` combining a `*list.List` (most recent at the front) with a `map[int]*list.Element`, and methods `Get(key int) (int, bool)` and `Put(key, value int)` that evict from the back when over capacity.
// EXAMPLE: capacity 2: put 1, put 2, get 1, put 3 -> key 2 is evicted
// BIGO: O(1) per operation. `list.Element.Value` is an `any`, so store a small struct and type-assert it back.
type LRU struct {
	cap   int
	order *list.List
	items map[int]*list.Element
}

type entry struct{ key, value int }

func NewLRU(capacity int) *LRU {
	panic("todo")
}

func (c *LRU) Get(key int) (int, bool) {
	panic("todo")
}

func (c *LRU) Put(key, value int) {
	panic("todo")
}

// ==== END
