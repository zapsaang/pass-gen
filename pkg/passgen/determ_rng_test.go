package passgen

import (
	"testing"
)

func TestNewDetermRNG(t *testing.T) {
	rng := newDetermRNG("testseed")

	if rng == nil {
		t.Fatal("newDetermRNG() returned nil")
	}
	if string(rng.seed) != "testseed" {
		t.Errorf("seed = %q, want 'testseed'", string(rng.seed))
	}
	if rng.counter != 0 {
		t.Errorf("counter = %d, want 0", rng.counter)
	}
	if rng.buffer != nil {
		t.Error("buffer should be nil initially (lazy initialization)")
	}
	if rng.ptr != 0 {
		t.Errorf("ptr = %d, want 0", rng.ptr)
	}
}

func TestDetermRNG_Deterministic(t *testing.T) {
	rng1 := newDetermRNG("seed123")
	rng2 := newDetermRNG("seed123")

	for i := 0; i < 100; i++ {
		v1 := rng1.Intn(100)
		v2 := rng2.Intn(100)
		if v1 != v2 {
			t.Errorf("Iteration %d: rng1.Intn(100) = %d, rng2.Intn(100) = %d", i, v1, v2)
		}
	}
}

func TestDetermRNG_DifferentSeeds(t *testing.T) {
	rng1 := newDetermRNG("seed1")
	rng2 := newDetermRNG("seed2")

	vals1 := make([]int, 10)
	vals2 := make([]int, 10)

	for i := 0; i < 10; i++ {
		vals1[i] = rng1.Intn(256)
		vals2[i] = rng2.Intn(256)
	}

	allSame := true
	for i := 0; i < 10; i++ {
		if vals1[i] != vals2[i] {
			allSame = false
			break
		}
	}
	if allSame {
		t.Error("Different seeds should produce different sequences")
	}
}

func TestDetermRNG_NextByte(t *testing.T) {
	rng := newDetermRNG("byteseed")

	bytes := make([]byte, 100)
	for i := 0; i < 100; i++ {
		bytes[i] = rng.nextByte()
	}

	rng2 := newDetermRNG("byteseed")
	for i := 0; i < 100; i++ {
		b := rng2.nextByte()
		if b != bytes[i] {
			t.Errorf("Byte %d: got %d, want %d", i, b, bytes[i])
		}
	}
}

func TestDetermRNG_Refill(t *testing.T) {
	rng := newDetermRNG("refillseed")

	if rng.buffer != nil {
		t.Error("Buffer should be nil before first use")
	}

	_ = rng.nextByte()
	if rng.buffer == nil {
		t.Error("Buffer should not be nil after first nextByte()")
	}
	if len(rng.buffer) != 32 {
		t.Errorf("Buffer length = %d, want 32", len(rng.buffer))
	}

	initialCounter := rng.counter
	for i := rng.ptr; i < len(rng.buffer); i++ {
		_ = rng.nextByte()
	}
	_ = rng.nextByte()
	if rng.counter != initialCounter+1 {
		t.Errorf("Counter = %d, want %d", rng.counter, initialCounter+1)
	}
}

func TestDetermRNG_Intn_Zero(t *testing.T) {
	rng := newDetermRNG("zeroseed")

	result := rng.Intn(0)
	if result != 0 {
		t.Errorf("Intn(0) = %d, want 0", result)
	}
}

func TestDetermRNG_Intn_Negative(t *testing.T) {
	rng := newDetermRNG("negseed")

	result := rng.Intn(-5)
	if result != 0 {
		t.Errorf("Intn(-5) = %d, want 0", result)
	}
}

func TestDetermRNG_Intn_One(t *testing.T) {
	rng := newDetermRNG("oneseed")

	for i := 0; i < 100; i++ {
		result := rng.Intn(1)
		if result != 0 {
			t.Errorf("Intn(1) = %d, want 0", result)
		}
	}
}

func TestDetermRNG_Intn_Range(t *testing.T) {
	rng := newDetermRNG("rangeseed")

	maxVal := 50
	for i := 0; i < 1000; i++ {
		result := rng.Intn(maxVal)
		if result < 0 || result >= maxVal {
			t.Errorf("Intn(%d) = %d, out of range [0, %d)", maxVal, result, maxVal)
		}
	}
}

func TestDetermRNG_Intn_Distribution(t *testing.T) {
	rng := newDetermRNG("distseed")

	maxVal := 10
	counts := make([]int, maxVal)
	iterations := 10000

	for i := 0; i < iterations; i++ {
		result := rng.Intn(maxVal)
		counts[result]++
	}

	expected := iterations / maxVal
	tolerance := expected / 2

	for i, count := range counts {
		if count < expected-tolerance || count > expected+tolerance {
			t.Logf("Warning: Bucket %d has %d values (expected ~%d)", i, count, expected)
		}
	}

	for i, count := range counts {
		if count == 0 {
			t.Errorf("Bucket %d has no values, distribution may be broken", i)
		}
	}
}

func TestDetermRNG_Intn_ModuloBias(t *testing.T) {
	rng := newDetermRNG("biasseed")

	maxVal := 100
	counts := make([]int, maxVal)
	iterations := 100000

	for i := 0; i < iterations; i++ {
		result := rng.Intn(maxVal)
		counts[result]++
	}

	expected := float64(iterations) / float64(maxVal)
	var variance float64
	for _, count := range counts {
		diff := float64(count) - expected
		variance += diff * diff
	}
	variance /= float64(maxVal)

	maxVariance := expected * 0.5
	if variance > maxVariance {
		t.Logf("Variance %f may indicate modulo bias (threshold: %f)", variance, maxVariance)
	}
}

func TestDetermRNG_LargeMax(t *testing.T) {
	rng := newDetermRNG("largeseed")

	for i := 0; i < 100; i++ {
		result := rng.Intn(256)
		if result < 0 || result >= 256 {
			t.Errorf("Intn(256) = %d, out of range", result)
		}
	}
}

func TestDetermRNG_EmptySeed(t *testing.T) {
	rng := newDetermRNG("")

	result := rng.Intn(100)
	if result < 0 || result >= 100 {
		t.Errorf("Intn(100) with empty seed = %d, out of range", result)
	}

	rng2 := newDetermRNG("")
	result2 := rng2.Intn(100)
	if result != result2 {
		t.Error("Empty seed should still be deterministic")
	}
}

func TestDetermRNG_LongSeed(t *testing.T) {
	longSeed := ""
	for i := 0; i < 10000; i++ {
		longSeed += "a"
	}

	rng := newDetermRNG(longSeed)
	result := rng.Intn(100)

	if result < 0 || result >= 100 {
		t.Errorf("Intn(100) with long seed = %d, out of range", result)
	}

	rng2 := newDetermRNG(longSeed)
	result2 := rng2.Intn(100)
	if result != result2 {
		t.Error("Long seed should still be deterministic")
	}
}

func TestDetermRNG_SequenceConsistency(t *testing.T) {
	rng := newDetermRNG("consistency_test_seed")

	expected := []int{}
	for i := 0; i < 10; i++ {
		expected = append(expected, rng.Intn(256))
	}

	rng2 := newDetermRNG("consistency_test_seed")
	for i := 0; i < 10; i++ {
		result := rng2.Intn(256)
		if result != expected[i] {
			t.Errorf("Value %d: got %d, want %d", i, result, expected[i])
		}
	}
}
