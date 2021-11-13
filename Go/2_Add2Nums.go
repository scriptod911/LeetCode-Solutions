func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode { // l1, l2 are the two linked list
	var carry int // carry is the carry of the addition
	var head *ListNode // head is the head of the result linked list
	var tail *ListNode // tail is the tail of the result linked list
	
	for l1 != nil || l2 != nil || carry != 0 { // while l1, l2, and carry are not nil
		var sum int // sum is the sum of the two nodes
		if l1 != nil { // if l1 is not nil
			sum += l1.Val // sum is the sum of the two nodes
			l1 = l1.Next // l1 is the next node of l1
		}
		if l2 != nil { // if l2 is not nil
			sum += l2.Val // sum is the sum of the two nodes
			l2 = l2.Next // l2 is the next node of l2
		} 
		sum += carry // sum is the sum of the two nodes and carry
		carry = sum / 10 // carry is the carry of the addition
		sum = sum % 10 // sum is the sum of the two nodes and carry
		
		if head == nil { // if head is nil
			head = &ListNode{sum, nil} // head is the head of the result linked list
			tail = head // tail is the tail of the result linked list
		} else { // if head is not nil
			tail.Next = &ListNode{sum, nil} // tail is the next node of tail
			tail = tail.Next // tail is the tail of the result linked list
		}
	}
	
	return head // return the head of the result linked list
}