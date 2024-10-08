package twosums

// Given an array of integers nums and an integer target, return indices of the
// two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func TwoSums(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if remainderIndex, ok := m[target-v]; ok {
			return []int{i, remainderIndex}
		}

		m[v] = i
	}

	return nil
}

// Time Complexity O(n)
// Space complexity O(n)
