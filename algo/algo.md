
# Algorithm in LeetCode


Summarized by Wafer

## Reference
[代码随想录](https://programmercarl.com/)

[算法通关手册](https://algo.itcharge.cn/00.Introduction/01.Data-Structures-Algorithms/)

[AC算法笔记](https://github.com/AlanChaw/interview/blob/master/%E7%AE%97%E6%B3%95%E7%AC%94%E8%AE%B0.md)

## Binary Search
#### 二分查找
搜索无重复、有序的数组


#### 例题
[704. Binary Search](https://leetcode.com/problems/binary-search/description/)





## Two Pointer
#### 快慢指针
删除数组特定值的元素，等价于创建不等于目标值的慢数组

#### 对撞指针
满足无重复、有序的数组又同时满足回文条件（以0对称）


#### 例题
[27. Remove Element](https://leetcode.com/problems/remove-element/description/)

[977. Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/description/)


## Sliding Window
#### 滑动窗口
数组的子数组问题，换算成求窗口大小

[209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/description/)


## Matrix
#### 矩阵
主要是学会go怎么创建二维切片

    mat := make([][]int,n)
    for i := range mat {
        mat[i] = make([]int,n)
    }

#### 例题
[59. Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/description/)


## Linked List
#### 链表
链表的插入删除等操作最好是遍历到前一个Node，因此引入 dummy head 方便遍历
#### 例题
[203. Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/description/)

[707. Design Linked List](https://leetcode.com/problems/design-linked-list/description/)


## Linked List Two Ptr
#### 链表指针
快慢指针，快的先走n步，用于删除节点。

快慢指针，反转节点，一定要记住从dummyHead开始走，会方便很多；还有一个思路是temp来存可能会丢掉的值
#### 例题
[206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/description/)

[24. Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/description/)

[19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/)

[160. Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/description/)

## Hash Table
#### slices包
这个包巨好用，点击 [此链接](https://github.com/Wafer233/testLab/blob/master/golang/slices/slices.go)
#### 切片与映射
哈希表的在golang中的表达无非两种，slice和map，那么golang在遍历他们会发生什么？

    for index, num := range nums {
        //遍历切片
    }

    for key, value := range mapping {
        //遍历映射
    }

    if value, exist := map[key]; !exist {
        //不存在时会怎么样
    }

#### 哈希表
哈希表可以用于计数，比如说在求两数之和，可以把遍历过的值存进去。

两数之和变种，为了方便起见，我们可以把他分为两两一组。
#### 例题
[242. Valid Anagram](https://leetcode.com/problems/valid-anagram/description/)

[349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/description/)

[1. Two Sum](https://leetcode.com/problems/two-sum/description/)

[454. 4Sum II](https://leetcode.com/problems/4sum-ii/description/)

[15. 3Sum](https://leetcode.com/problems/3sum/description/)

[18. 4Sum](https://leetcode.com/problems/4sum/description/)

[347. Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/description/)


## String
#### 字符串
字符串类型也是力扣上的常考的东西了，golang的字符串有一个特点，就是他不能改，然后它还能遍历，通常来说要有这些操作。

    for index, ch := range str {
        ch == '1'
    }

    b := []byte(str)
    for index, value := range b {
        b[index] = 'a'
    }

    r := []rune(str)
    for index, value := range b {
        r[index] = 'a'
    }



#### strings包
这个包巨好用，点击 [此链接](https://github.com/Wafer233/testLab/blob/master/golang/strings/strings.go)

#### 反转与重复子串
看到反转都要有条件反射了！对撞指针！golang交换很简单

重复子串，这个很新，移动匹配，增加一段掐头去尾能找到自己
    
    s[left], s[right] =  s[right],s[left]   

#### 例题
[344. Reverse String](https://leetcode.com/problems/reverse-string/description/)

[541. Reverse String II](https://leetcode.com/problems/reverse-string-ii/description/)

[151. Reverse Words in a String](https://leetcode.com/problems/reverse-words-in-a-string/description/)

[459. Repeated Substring Pattern](https://leetcode.com/problems/repeated-substring-pattern/description/)

## Stack
#### 栈和队列
其实这个东西在go中很好理解，一个先进先出，一个先进后出，因此想要模拟的话直接操作切片

    //stack
    push -> append(stack,val)
    pop -> stack[:len(stack)-1]
    peek -> stack[len(stack)-1]

    //queue
    inQueue -> append(stack,val)
    deQueue -> stack[1:]
    peek -> stack[0]

#### 实现功能
栈实现括号有效，逆波兰表达式

*双向队列实现滑窗最大值

#### 例题
[232. Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/)

[225. Implement Stack using Queues](https://leetcode.com/problems/implement-stack-using-queues/description/)

[20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/description/)

[1047. Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/description/)

[150. Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/description/)

[239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/description/)



## Binary Tree
#### 二叉树
无非就是操作在夹在递归函数中的位置，表示了前中后遍历,注意层序遍历要按照前序的方式进行递归

    递归遍历三部曲
    1.递归函数携带参数/返回值（全局，局部）
    2.中止条件
    3.单层递归逻辑
    4.全局和局部变量要啰嗦一嘴，看看他的生命周期和作用域（例如二叉树的深度，回溯算法的startIndex）


#### 例题
[144. Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/description/)

[145. Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/)

[94. Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/description/)

[102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/description/)

## Binary Tree Ops



[226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/description/)

[106. Construct Binary Tree from Inorder and Postorder](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/)

[654. Maximum Binary Tree](https://leetcode.com/problems/maximum-binary-tree/description/)

[617. Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/description/)





## Binary Tree Attribute

[101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/description/)

[104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/description/)

[111. Minimum Depth of Binary Tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/description/)

[222. Count Complete Tree Nodes](https://leetcode.com/problems/count-complete-tree-nodes/description/)

[110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/description/)

[257. Binary Tree Paths](https://leetcode.com/problems/binary-tree-paths/description/)

[404. Sum of Left Leaves](https://leetcode.com/problems/sum-of-left-leaves/description/)

[513. Find Bottom Left Tree Value](https://leetcode.com/problems/find-bottom-left-tree-value/description/)

[112. Path Sum](https://leetcode.com/problems/path-sum/description/)


## BST

[700. Search in a Binary Search Tree](https://leetcode.com/problems/search-in-a-binary-search-tree/description/)

[98. Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/description/)

[530. Minimum Absolute Difference in BST](https://leetcode.com/problems/minimum-absolute-difference-in-bst/description/)

[501. Find Mode in Binary Search Tree](https://leetcode.com/problems/find-mode-in-binary-search-tree/description/)

[236. Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/)

[701. Insert into a Binary Search Tree](https://leetcode.com/problems/insert-into-a-binary-search-tree/description/)

[538. Convert BST to Greater Tree](https://leetcode.com/problems/convert-bst-to-greater-tree/description/)


## Backtracking Combination

[77. Combinations](https://leetcode.com/problems/combinations/description/)

[216. Combination Sum III](https://leetcode.com/problems/combination-sum-iii/description/)

[17. Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/)

[39. Combination Sum](https://leetcode.com/problems/combination-sum/description/)

[40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii/description/)

## Backtracking Segmentation

[131. Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/description/)

[93. Restore IP Addresses](https://leetcode.com/problems/restore-ip-addresses/description/)

## Backtracking Subset

[78. Subsets](https://leetcode.com/problems/subsets/description/)

[90. Subsets II](https://leetcode.com/problems/subsets-ii/description/)

[491. Non-decreasing Subsequences](https://leetcode.com/problems/non-decreasing-subsequences/description/)

## Backtracking Permutation

[46. Permutations](https://leetcode.com/problems/permutations/description/)

[47. Permutations II](https://leetcode.com/problems/permutations-ii/description/)

## Backtracking Chessboard

[51. N-Queens](https://leetcode.com/problems/n-queens/description/)

## DP

[509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/description/)

[70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/description/)

[746. Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/description/)

[62. Unique Paths](https://leetcode.com/problems/unique-paths/description/)

[63. Unique Paths II](https://leetcode.com/problems/unique-paths-ii/description/)

[343. Integer Break](https://leetcode.com/problems/integer-break/description/)

[96. Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/description/)

## DP 0-1 Knapsack

[416. Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/description/)

[1049. Last Stone Weight II](https://leetcode.com/problems/last-stone-weight-ii/description/)

[494. Target Sum](https://leetcode.com/problems/target-sum/description/)

[474. Ones and Zeroes](https://leetcode.com/problems/ones-and-zeroes/description/)

## DP Complete Knapsack

[518. Coin Change II](https://leetcode.com/problems/coin-change-ii/description/)

[377. Combination Sum IV](https://leetcode.com/problems/combination-sum-iv/description/)

[322. Coin Change](https://leetcode.com/problems/coin-change/description/)

[279. Perfect Squares](https://leetcode.com/problems/perfect-squares/description/)

[139. Word Break](https://leetcode.com/problems/word-break/description/)

## DP House Robber

[198. House Robber](https://leetcode.com/problems/house-robber/description/)

[213. House Robber II](https://leetcode.com/problems/house-robber-ii/description/)

[337. House Robber III](https://leetcode.com/problems/house-robber-iii/description/)

## DP Stock

[121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)

[122. Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/)

[123. Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/description/)

[188. Best Time to Buy and Sell Stock IV](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/description/)

[309. Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/)

[714. Best Time to Buy and Sell Stock with Transa](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/)

## DP Discontinuous Subsequence

[300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)

[1143. Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/description/)

[1035. Uncrossed Lines](https://leetcode.com/problems/uncrossed-lines/description/)

[392. Is Subsequence](https://leetcode.com/problems/is-subsequence/description/)

## DP Continuous Subsequence/ Subarray

[674. Longest Continuous Increasing Subsequence](https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/)

[718. Maximum Length of Repeated Subarray](https://leetcode.com/problems/maximum-length-of-repeated-subarray/description/)

[53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/description/)

## DP Edit Distance

[115. Distinct Subsequences](https://leetcode.com/problems/distinct-subsequences/description/)

[583. Delete Operation for Two Strings](https://leetcode.com/problems/delete-operation-for-two-strings/description/)

[72. Edit Distance](https://leetcode.com/problems/edit-distance/description/)

## DP Palindrome Array

[647. Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/description/)

[516. Longest Palindromic Subsequence](https://leetcode.com/problems/longest-palindromic-subsequence/description/)


