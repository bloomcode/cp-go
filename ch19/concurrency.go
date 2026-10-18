// Package ch19: Concurrency.
package ch19

import (
	"time"
)

// ==== EX 19.1 | Start workers and wait for all | Core
// PROMPT: Write `ParallelSquares(nums []int) []int` that starts one goroutine per number, each storing its number's square in its own slot of the result slice, and waits for all of them with a `sync.WaitGroup` before returning.
// EXAMPLE: `[1 2 3]` -> `[1 4 9]`
// BIGO: O(N) work spread over N threads; each writes only its own slot, so no lock is needed.
func ParallelSquares(nums []int) []int {
	panic("todo")
}

// ==== END

// ==== EX 19.2 | A producer and a consumer | Core
// PROMPT: Write `PipelineTotal(nums []int) int` with a producer goroutine that sends each number on a channel and then closes it, while the caller ranges over the channel adding values up. Closing the channel is how the producer says it is finished.
// EXAMPLE: `[4 5 6]` -> `15`
// BIGO: O(N): every value passes through the queue once; the queue does the locking and the waiting.
func PipelineTotal(nums []int) int {
	panic("todo")
}

// ==== END

// ==== EX 19.3 | A pool with a fixed number of workers | Practice
// PROMPT: Write `PoolMap[T, R any](fn func(T) R, items []T, workers int) []R` that applies `fn` to every item using exactly `workers` goroutines, which read item indexes from a channel and write each result into its slot, and returns the results in input order once a `sync.WaitGroup` says all workers are done.
// EXAMPLE: `func(x int) int { return x * 10 }, [1 2 3], 2` -> `[10 20 30]`
// BIGO: About ceil(N / W) rounds of work for W workers; the results keep the input order however the work finishes.
func PoolMap[T, R any](fn func(T) R, items []T, workers int) []R {
	panic("todo")
}

// ==== END

// ==== EX 19.4 | A counter protected by a lock | Core
// PROMPT: Write `LockedCount(workers, increments int) int` that starts `workers` goroutines, each adding 1 to a shared counter `increments` times, and returns the final count. Guard every update with one `sync.Mutex`, because `count++` is a read, an add and a write that two goroutines can interleave (`go test -race` reports it if you forget).
// EXAMPLE: `4, 1000` -> `4000`
// BIGO: O(T * K) increments; the lock makes each read-modify-write happen as one step, so none is lost.
func LockedCount(workers, increments int) int {
	panic("todo")
}

// ==== END

// ==== EX 19.5 | Stop waiting after a timeout | Practice
// PROMPT: Write `ResultWithin[T any](task func() T, timeout time.Duration) (T, bool)` that runs `task` in a goroutine and returns its result and `true` if it finishes within `timeout`, or the zero value and `false` otherwise. The goroutine sends on a channel with a buffer of one (so a late task can still finish and exit); `select` between that channel and `time.After(timeout)`.
// EXAMPLE: `func() int { return 42 }, time.Second` -> `42, true`; a task that sleeps 300 ms with `50 * time.Millisecond` -> `0, false`
// BIGO: Waits at most the timeout, however long the task takes.
func ResultWithin[T any](task func() T, timeout time.Duration) (T, bool) {
	panic("todo")
}

// ==== END

// ==== EX 19.6 | Split a sum across workers | Practice
// PROMPT: Write `ParallelSum(nums []int, workers int) int` that cuts `nums` into `workers` contiguous chunks of near-equal size, sums each chunk in its own goroutine into its own slot of a `partial` slice, waits for all with a `sync.WaitGroup`, and adds the partial sums.
// EXAMPLE: `[1..100], 4` -> `5050`
// BIGO: O(N / W) per worker for W workers, then O(W) to combine; each worker owns one slot of the partial results.
// NOTE: In CPython, threads take turns running Python code (the GIL), so this shows the pattern rather than a speed-up; use processes for CPU-heavy work.
func ParallelSum(nums []int, workers int) int {
	panic("todo")
}

// ==== END
