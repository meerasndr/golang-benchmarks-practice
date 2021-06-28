package fib

import "testing"

var fibTests = []struct {
	n int // input 
	expected int //expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func TestFib(t *testing.T){
	for _, tt := range fibTests{
		actual := Fib(tt.n)
		if actual != tt.expected{
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

var result int
func benchmarkFib(i int, b *testing.B){
	var r int 
	for n := 0; n < b.N; n++ {
		r = Fib(i)
	}
	result = r
}

func BenchmarkFib1(b *testing.B) {benchmarkFib(1, b)}
func BenchmarkFib2(b *testing.B) {benchmarkFib(2, b)}
func BenchmarkFib3(b *testing.B) {benchmarkFib(3, b)}
func BenchmarkFib10(b *testing.B) {benchmarkFib(10, b)}
func BenchmarkFib20(b *testing.B) {benchmarkFib(20, b)}
func BenchmarkFib50(b * testing.B) {benchmarkFib(50, b)}
