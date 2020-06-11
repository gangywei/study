package test

import "container/heap"

//https://leetcode-cn.com/problems/linked-list-cycle/
//使用快慢指针，注意命名规范
func hasCycle(head *ListNode) bool {
	if head == nil{
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow{
			return true
		}
	}
	return false
}

//https://leetcode-cn.com/problems/merge-k-sorted-lists/
//合并 K 个排序的链表
type PQ []*ListNode

func (p PQ) Len() int {
	return len(p)
}

func (p PQ) Swap (i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PQ) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func (p *PQ) Push(x interface{}) {
	node := x.(*ListNode)
	*p = append(*p, node)
}

func (p *PQ) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListNode {
		Val: -1,
		Next: nil,
	}
	t := h
	if len(lists) == 0 {
		return h.Next
	}

	pq := make(PQ, 10)
	for i, _ := range lists {
		if lists[i] != nil {
			pq = append(pq, lists[i])
		}
	}
	//初始化堆
	heap.Init(&pq)

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*ListNode)
		next := item.Next

		item.Next = t.Next
		t.Next = item
		t = item

		if next != nil {
			heap.Push(&pq, next)
		}
	}

	return h.Next
}