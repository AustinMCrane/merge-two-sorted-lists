package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}
	return nil
}

func listToArray(list *ListNode) []int {
	a := []int{}
	p := list
	for {
		if p == nil {
			break
		}

		a = append(a, p.Val)
		p = p.Next
	}

	return a
}

func main() {
}
