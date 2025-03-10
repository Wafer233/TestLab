package algo

//209. Minimum Size Subarray Sum

func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	cur_sum := 0
	window_len := 100000 + 1

	for right < len(nums) {
		cur_sum += nums[right]
		for cur_sum >= target {
			if right-left+1 < window_len {
				window_len = right - left + 1
			}
			cur_sum -= nums[left]
			left++
		}
		right++
	}

	if window_len == 100000+1 {
		return 0
	} else {
		return window_len
	}

}
