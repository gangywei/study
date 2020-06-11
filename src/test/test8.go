package test

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	var listSize, midIndex int
	tempNode := head
	for tempNode.Next != nil {
		listSize++
		tempNode = tempNode.Next
	}
	midIndex = listSize/2 + listSize%2
	for midIndex > 0 {
		head = head.Next
		midIndex--
	}

	return head
}

/**
另一个解法快慢指针
func middleNode(head *ListNode) *ListNode {
    fast, slow := head, head
    for {
        if fast.Next != nil && fast.Next.Next != nil {
            fast = fast.Next.Next
            slow = slow.Next
        } else if fast.Next != nil {
            return slow.Next
        } else {
            return slow
        }
    }
}
 */
