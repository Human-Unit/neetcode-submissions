func topKFrequent(nums []int, k int) []int {
	group := make(map[int]int)
	for _, ch := range nums {
		group[ch]++
	}
	var sorted []int
	for a := range group {
		inserted := false
		for i, b := range sorted {
			if group[a] > group[b] {
				sorted = append(sorted, 0)
				copy(sorted[i+1:], sorted[i:])
				sorted[i] = a
				inserted = true
				break
			}
		}
		if !inserted {
			sorted = append(sorted, a)
		}
	}
	return sorted[:k]
}
