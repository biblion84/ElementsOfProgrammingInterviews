package main

import "fmt"

func main() {
	fmt.Println(isPrime(18))
}

// We could use different optimizations strategies, like erastostene sieve

func isPrime(x int) []int {

	sieve := make([]int, x)
	result := []int{}

	for i := 2; i < len(sieve); i++ {
		if sieve[i] == -1 {
			continue
		}
		result = append(result, i)
		for j := i + i; j < len(sieve); j += i {
			sieve[j] = -1
		}
	}

	return result
}
