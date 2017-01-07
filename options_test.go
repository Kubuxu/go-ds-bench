package main

import "testing"

func TestOptionsSimpleRange(t *testing.T) {
	start := BenchOptions{1, 100}
	end := BenchOptions{1 << 10, 100}

	opts := OptionsRange2pow(start, end, 11)
	if len(opts) != 11 {
		t.Fatalf("length is %d, should be %d", len(opts), 11)
	}

	for k, v := range opts {
		if 1<<uint(k) != v.PrePrimeCount {
			t.Errorf("expected PrePrimeCount=%d, got %d @%d", 1<<uint(k), v.PrePrimeCount, k)
		}
	}
}
func TestOptionsBoth(t *testing.T) {
	start := BenchOptions{1, 1}
	end := BenchOptions{1 << 10, 1 << 10}

	opts := OptionsRange2pow(start, end, 11)
	if len(opts) != 11*11 {
		t.Fatalf("length is %d, should be %d", len(opts), 11*11)
	}
}
