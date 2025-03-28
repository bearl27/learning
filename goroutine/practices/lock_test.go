package practices

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func worker(lock *sync.Mutex, lockCount int, criticalSectionTime time.Duration, timer time.Duration) int {
	count := 0
	begin := time.Now()
	for time.Since(begin) <= timer {
		for i := 0; i < lockCount; i++ {
			lock.Lock()
			time.Sleep(criticalSectionTime)
			lock.Unlock()
		}
		count++
	}
	return count
}

func BenchmarkLockBalance(b *testing.B) {
	testCases := []struct {
		name                string
		lockCount           int
		criticalSectionTime time.Duration
		goroutineCount      int
		timer               time.Duration
	}{
		{
			name:                "greedy_1_3ns_1",
			lockCount:           1,
			criticalSectionTime: 3 * time.Nanosecond,
			goroutineCount:      1,
			timer:               1 * time.Second,
		},
		{
			name:                "polite_3_1ns_1",
			lockCount:           3,
			criticalSectionTime: 1 * time.Nanosecond,
			goroutineCount:      1,
			timer:               1 * time.Second,
		},
		{
			name:                "greedy_1_3ns_10",
			lockCount:           1,
			criticalSectionTime: 3 * time.Nanosecond,
			goroutineCount:      10,
			timer:               1 * time.Second,
		},
		{
			name:                "polite_3_1ns_10",
			lockCount:           3,
			criticalSectionTime: 1 * time.Nanosecond,
			goroutineCount:      10,
			timer:               1 * time.Second,
		},
		// 他のテストケースを追加
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				var wg sync.WaitGroup
				var lock sync.Mutex
				results := make(chan int, tc.goroutineCount)

				for i := 0; i < tc.goroutineCount; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						results <- worker(&lock, tc.lockCount, tc.criticalSectionTime, tc.timer)
					}()
				}

				wg.Wait()
				close(results)

				totalCount := 0
				for count := range results {
					totalCount += count
				}
				fmt.Println(tc.name, "total count:", totalCount)
			}
		})
	}
}
