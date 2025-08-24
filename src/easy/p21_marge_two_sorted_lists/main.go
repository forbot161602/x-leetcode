package main

import "fmt"

func main() {
	solution := mergeTwoLists
	list1 := toList([]int{2})
	list2 := toList([]int{1, 3})
	fmt.Println(toSlice(solution(list1, list2)))
}

type ListNode struct {
	Val  int
	Next *ListNode
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
