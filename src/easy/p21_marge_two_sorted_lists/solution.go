package main

// Runtime beats 100%, while memory beats 66.51%.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) (list *ListNode) {
	list = list1
	if list2 == nil {
		return
	}
	if list1 == nil {
		list = list2
		return
	}

	node1, node2 := list1, list2
	if node1.Val > node2.Val {
		node1, node2 = node2, node1
	}

	list = node1
	for node2 != nil {
		next := node1.Next
		if next == nil {
			node1.Next = node2
			break
		}
		if next.Val <= node2.Val {
			node1 = next
			continue
		}
		node1.Next = node2
		node1 = node2
		node2 = node2.Next
		node1.Next = next
	}
	return
}
