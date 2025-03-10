package algo

func Search(nums []int, target int) int {
	// [l,r]
	l := 0
	r := len(nums) - 1
	for l <= r {
		middle := (l + r) / 2
		if nums[middle] < target {
			l = middle + 1
		} else if nums[middle] > target {
			r = middle - 1
		} else {
			return middle
		}
	}
	return -1
}
