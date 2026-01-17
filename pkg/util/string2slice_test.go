package util

import "testing"

func BenchmarkString2UintSlice(b *testing.B) {
	for b.Loop() {
		String2UintSlice("1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20")
	}
}
