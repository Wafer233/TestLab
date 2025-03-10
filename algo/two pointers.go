package algo

//	977. Squares of a Sorted Array
//	27. Remove Element

//	case 1: Opposite Pointers
//	template
/*
	[left,right]

	left, right = 0, len(nums) - 1
	while left <= right:
		if 满足要求的特殊条件:
			return 符合条件的值
		elif 一定条件 1:
			left += 1
		elif 一定条件 2:
			right -= 1
		return 没找到 或 找到对应值
*/

func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	res := make([]int, len(nums))
	i := len(nums) - 1
	for left <= right {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			res[i] = nums[left] * nums[left]
			left++
		} else {
			res[i] = nums[right] * nums[right]
			right--
		}
		i--
	}
	return res
}

//	--------------------------------------------------

//	case 2: Fast and Slow Pointers
//	template
/*
	slow = 0
	fast = 0
	while 没有遍历完：
		if 满足要求的特殊条件:
			slow += 1
		fast += 1
	return 合适的值
*/
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
