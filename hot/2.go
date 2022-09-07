package hot

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cursor1 := l1
	cursor2 := l2
	for cursor1.Next != nil {
		add(cursor1, cursor2)
		if cursor2.Next != nil {
			cursor2 = cursor2.Next
		} else {
			/*len(l1)>len(l2)*/
			check(cursor1)
			return l1
		}
		cursor1 = cursor1.Next
	}
	/*len(l1)<l2 */
	if cursor2.Next != nil {
		add(cursor2, cursor1)
		check(cursor2.Next)
		cursor1.Val = cursor2.Val
		cursor1.Next = cursor2.Next
		return l1
	} else { /* same len*/
		add(cursor1, cursor2)
		return l1
	}

}
func add(cursor1 *ListNode, cursor2 *ListNode) {
	if cursor1.Val+cursor2.Val < 10 {
		cursor1.Val += cursor2.Val
	} else {
		if cursor1.Next == nil {
			cursor1.Next = new(ListNode)
		}
		cursor1.Val += cursor2.Val - 10
		cursor1.Next.Val += 1
	}
}
func check(cursor *ListNode) {
	for cursor.Next != nil {
		if cursor.Val >= 10 {
			cursor.Val -= 10
			cursor.Next.Val += 1
		}
		cursor = cursor.Next
	}
	if cursor.Val >= 10 {
		cursor.Next = new(ListNode)
		cursor.Val -= 10
		cursor.Next.Val = 1
	}
}
