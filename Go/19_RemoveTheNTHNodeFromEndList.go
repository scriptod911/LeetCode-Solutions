func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	q := head
	for i := 0; i < n; i++ {
		p = p.Next
	}
	if p == nil {
		return head.Next
	}
	for p.Next != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return head
}