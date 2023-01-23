package main

import "fmt"

// https://leetcode.com/problems/merge-two-sorted-lists/
// Runtime beats 100%, while memory beats 66.51%.
func main() {
	list1 := toList([]int{2})
	list2 := toList([]int{1, 3})
	fmt.Println(toSlice(mergeTwoLists(list1, list2)))
}

func toList(values []int) *ListNode {
	list := &ListNode{0, nil}
	node := list
	for _, value := range values {
		node.Next = &ListNode{value, nil}
		node = node.Next
	}
	return list.Next
}

func toSlice(list *ListNode) []int {
	values := []int{}
	for list != nil {
		values = append(values, list.Val)
		list = list.Next
	}
	return values
}

type ListNode struct {
	Val  int
	Next *ListNode
}
