package util

func SumU64(nums ...uint64) uint64 {
	total := uint64(0)
	for _, n := range nums {
		total += n
	}
	return total
}
