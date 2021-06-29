// There are a few simple approaches to check if a number is prime
// Benchmark them and see what happens when n increases

package prime

import "math"

func FullTraverse(n uint64) bool {
	if n < 2 {return false}
	var i uint64
	for i = 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func Midway(n uint64) bool {
	if n < 2 {return false}
	var i uint64
	for i = 2; i <= n/2; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func TrialDivision(n uint64) bool{
	if n < 2 {return false}
	var i uint64
	sqroot := uint64(math.Sqrt(float64(n)))
	for i = 2; i <= sqroot ; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}