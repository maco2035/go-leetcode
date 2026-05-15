func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1  {
		merged := mergeTwoLists(lists[0],lists[1])
		lists = append(lists[2:] , merged)
	}
	return lists[0]
}


func mergeTwoLists(l1, l2 *ListNode) (*ListNode) {
	Head :=  &ListNode{}
	Current := Head

	for (l1 != nil || l2!= nil ){

		if l1 != nil && (l2 == nil || l1.Val <= l2.Val) {
			Current.Next = l1
			l1 = l1.Next
		} else {
			Current.Next = l2
			l2 = l2.Next
		}
		Current = Current.Next
	}

	return Head.Next
}
