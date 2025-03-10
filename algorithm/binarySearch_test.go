package algorithm

import "testing"

func TestSearch(t *testing.T) {
	ret := Search([]int{1, 2, 9}, 9)
	print(ret)
}
