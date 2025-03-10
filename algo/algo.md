
# Algorithm in LeetCode


Summarized by Wafer

## Array
### Binary Search

    704. Binary Search
    
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

### Two Pointer
#### Case 1
    Template

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
####
    977. Squares of a Sorted Array

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
#### Case 2
    Template

	slow = 0
	fast = 0
	while 没有遍历完：
		if 满足要求的特殊条件:
			slow += 1
		fast += 1
	return 合适的值
####
    27. Remove Element

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

### Sliding Window
    Template

    left = 0
	right = 0
	while right < len(nums):
		window.append(nums[right])
		while 窗口需要缩小:
			# ... 可维护答案
			window.popleft()
			left += 1

		# 向右侧增大窗口
		right += 1
####

    209. Minimum Size Subarray Sum

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

## Linked List
### Dummy Head
#### Case 1
    203. Remove Linked List Elements

    type ListNode struct {
        Val  int
        Next *ListNode
    }

    func removeElements(head *ListNode, val int) *ListNode {
        dummyHead := &ListNode{
            Next: head,
        }
        curr := dummyHead
        for curr.Next != nil {
            if curr.Next.Val == val {
                curr.Next = curr.Next.Next
            } else {
                curr = curr.Next
            }
        }
        return dummyHead.Next
    }

#### Case 2

    type SingleNode struct {
        Val  int         // 节点的值
        Next *SingleNode // 下一个节点的指针
    }

    type MyLinkedList struct {
        dummyHead *SingleNode // 虚拟头节点
        Size      int         // 链表大小
    }

    /** Initialize your data structure here. */
    func Constructor() MyLinkedList {
        newNode := &SingleNode{ // 创建新节点
            0,
            nil,
        }
        return MyLinkedList{ // 返回链表
            dummyHead: newNode,
            Size:      0,
        }

    }

    func (this *MyLinkedList) Get(index int) int {
        if index < 0 || index >= this.Size { // 如果索引无效则返回-1
            return -1
        }
        // 让cur等于真正头节点
        cur := this.dummyHead // 设置当前节点为真实头节点
        for index > 0 {
            cur = cur.Next
            index--
        }
        return cur.Next.Val // 返回节点值
    }

    func (this *MyLinkedList) AddAtHead(val int) {
        newNode := &SingleNode{Val: val}   // 创建新节点
        newNode.Next = this.dummyHead.Next // 新节点指向当前头节点
        this.dummyHead.Next = newNode      // 新节点变为头节点
        this.Size++                        // 链表大小增加1
    }

    func (this *MyLinkedList) AddAtTail(val int) {
        newNode := &SingleNode{Val: val} // 创建新节点
        cur := this.dummyHead            // 设置当前节点为虚拟头节点
        for cur.Next != nil {            // 遍历到最后一个节点
            cur = cur.Next
        }
        cur.Next = newNode // 在尾部添加新节点
        this.Size++        // 链表大小增加1
    }

    func (this *MyLinkedList) AddAtIndex(index int, val int) {
        if index < 0 || index > this.Size { // 如果索引大于链表长度，直接返回
            return
        }
        newNode := &SingleNode{Val: val} // 创建新节点
        cur := this.dummyHead            // 设置当前节点为虚拟头节点
        for i := 0; i < index; i++ {     // 遍历到指定索引的前一个节点
            cur = cur.Next
        }
        newNode.Next = cur.Next // 新节点指向原索引节点
        cur.Next = newNode      // 原索引的前一个节点指向新节点
        this.Size++             // 链表大小增加1
    }

    /** Delete the index-th node in the linked list, if the index is valid. */
    func (this *MyLinkedList) DeleteAtIndex(index int) {
        if index < 0 || index >= this.Size { // 如果索引无效则直接返回
            return
        }
        cur := this.dummyHead        // 设置当前节点为虚拟头节点
        for i := 0; i < index; i++ { // 遍历到要删除节点的前一个节点
            cur = cur.Next
        }
        cur.Next = cur.Next.Next // 当前节点直接指向下下个节点，即删除了下一个节点
        this.Size--              // 注意删除节点后应将链表大小减一
    }

