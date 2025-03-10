package algo

//	707. Design Linked List
//	203. Remove Linked List Elements
//	add dummyHead

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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

//	--------------------------------------------------

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
