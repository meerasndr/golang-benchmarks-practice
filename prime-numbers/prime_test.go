package prime

import (
	"testing"
)

var primeTests = []struct{
	n uint64
	expected bool
}{
	{1, false},
	{2, true},
	{3, true},
	{4, false},
	{5, true},
	{6, false},
	{7, true},
	{17, true},
	{24, false},
	{53, true},
	{113, true},
}

func TestFullTraverse(t *testing.T){
	for _, tt := range primeTests{
		actual := FullTraverse(tt.n)
		if actual != tt.expected{
			t.Errorf("FullTraverse(%d): Expected Result: %t, Actual Result:%t", tt.n, tt.expected, actual)
		}
	}
}

func TestMidway(t *testing.T){
	for _, tt := range primeTests {
		actual := Midway(tt.n)
		if actual != tt.expected{
			t.Errorf("Midway(%d): Expected Result: %t, Actual Result: %t", tt.n, tt.expected, actual)
		}
	}
}

func TestTrialDivision(t *testing.T){
	for _, tt := range primeTests {
		actual := TrialDivision(tt.n)
		if actual != tt.expected{
			t.Errorf("Midway(%d): Expected Result: %t, Actual Result: %t", tt.n, tt.expected, actual)
		}
	}
}

var result bool

func benchmarkFullTraverse(n uint64, b *testing.B ){
	var r bool
	for i := 0; i < b.N; i++ {
		r = FullTraverse(n)
	}
	result = r
}

func benchmarkMidway(n uint64, b *testing.B){
	var r bool
	for i := 0; i < b.N; i++ {
		r = Midway(n)
	}
	result = r
}

func benchmarkTrialDivision(n uint64, b *testing.B){
	var r bool
	for i := 0; i < b.N; i++ {
		r = TrialDivision(n)
	}
	result = r
}

func BenchmarkFullTraverse36(b *testing.B){benchmarkFullTraverse(36, b)}
func BenchmarkMidway36(b *testing.B){benchmarkMidway(36, b)}
func BenchmarkTrialDivision36(b *testing.B){benchmarkTrialDivision(36, b)}

func BenchmarkFullTraverse15dig(b *testing.B){benchmarkFullTraverse(999999999999999, b)}
func BenchmarkMidway15dig(b *testing.B){benchmarkMidway(999999999999999, b)}
func BenchmarkTrialDivision15dig(b *testing.B){benchmarkTrialDivision(999999999999999, b)}

func BenchmarkFullTraverse19dig(b *testing.B){benchmarkFullTraverse(9999999999999991111, b)}
func BenchmarkMidway19dig(b *testing.B){benchmarkMidway(9999999999999991111, b)}
func BenchmarkTrialDivision19dig(b *testing.B){benchmarkTrialDivision(9999999999999991111, b)}