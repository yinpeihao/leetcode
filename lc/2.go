package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	current := head
	var prev *ListNode

	for l1 != nil && l2 != nil {
		current.Val += l1.Val + l2.Val
		current.Next = &ListNode{
			Val:  current.Val / 10,
			Next: nil,
		}
		current.Val = current.Val % 10
		prev = current
		current = current.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		current.Val += l1.Val
		current.Next = &ListNode{
			Val:  current.Val / 10,
			Next: nil,
		}
		current.Val = current.Val % 10
		prev = current
		current = current.Next
		l1 = l1.Next
	}
	for l2 != nil {
		current.Val += l2.Val
		current.Next = &ListNode{
			Val:  current.Val / 10,
			Next: nil,
		}
		current.Val = current.Val % 10
		prev = current
		current = current.Next
		l2 = l2.Next
	}
	if current.Val == 0 && prev != nil {
		prev.Next = nil
	}
	return head
}
