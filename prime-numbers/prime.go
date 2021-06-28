// There are a few simple approaches to check if a number is prime
// Benchmark them and see what happens when n increases

package prime 

func FullTraverse(n int) bool {
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func Midway(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func TrialDivision(n int) bool{
	
	return true
}