package main

import (
	"fmt"
	"math"
	"time"
)

// 素数判定関数
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 素数をリストとして返す関数
func findPrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	limit := 1000000

	// 時間計測開始
	start := time.Now()

	// 素数を計算
	primes := findPrimes(limit)

	// 結果と経過時間の表示
	fmt.Println("Found primes:", len(primes))
	fmt.Println("Execution time:", time.Since(start))
}
