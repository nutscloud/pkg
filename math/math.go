package math

// MaxUint returns the maximum values in a and b
func MaxUint(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

// MinUint returns the minimum values in a and b
func MinUint(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

// IsRangeOverlapUint judge whether [aStart, aEnd] and [bStart, bEnd] overlap
func IsRangeOverlapUint(aStart, aEnd, bStart, bEnd uint) bool {
	return int(MinUint(aEnd, bEnd))-int(MaxUint(aStart, bStart)) >= 0
}
