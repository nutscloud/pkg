package math

import (
	"testing"
)

// TestMaxUint test MaxUint
func TestMaxUint(t *testing.T) {
	max := MaxUint(1, 2)
	if max != 2 {
		t.Error("[TestMaxUint error] the maximum value of 1 and 2 should be 2")
	}
}

// TestMinUint test MinUint
func TestMinUint(t *testing.T) {
	min := MinUint(1, 2)
	if min != 1 {
		t.Error("[TestMinUint error] the minimum value of 1 and 2 should be 1")
	}
}

// TestIsRangeOverlapUint test IsRangeOverlapUint
func TestIsRangeOverlapUint(t *testing.T) {
	if IsRangeOverlapUint(11, 20, 1, 10) {
		t.Error("[TestMinUint error] There is no overlap between [11, 20] and [1, 10]")
	}

	if !IsRangeOverlapUint(11, 20, 1, 15) {
		t.Error("[TestMinUint error] There is overlap between [11, 20] and [1, 15]")
	}

	if !IsRangeOverlapUint(11, 20, 1, 11) {
		t.Error("[TestMinUint error] There is overlap between [11, 20] and [1, 11]")
	}
}
