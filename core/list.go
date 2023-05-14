package core

type ListNode struct {
	Content CacheElem
	Prev    *ListNode
	Next    *ListNode
}

type List struct {
	dummyHead *ListNode
	dummyTail *ListNode
	Length    int64
}

func (lst *List) Init() {
	lst.dummyHead = &ListNode{
		Content: NewDummyElem(),
		Prev:    nil,
		Next:    nil,
	}
	lst.dummyTail = &ListNode{
		Content: NewDummyElem(),
		Prev:    nil,
		Next:    nil,
	}
	lst.Length = 0
}

func (lst *List) Size() int64 {
	return lst.Length
}

func (lst *List) Find(value interface{}) (interface{}, bool) {
	for iterator := lst.dummyHead.Next; iterator != lst.dummyTail; iterator = iterator.Next {
		// move node to list head
		if iterator.Content.Value() == value {
			lst.remove(iterator)
			lst.lPushNode(iterator)
			return iterator.Content.Value(), true
		}
	}
	return nil, false
}

func (lst *List) lPush(elem string, expiration int64) {
	node := ListNode{
		Content: NewElem(elem, expiration),
		Prev:    nil,
		Next:    nil,
	}
	lst.lPushNode(&node)
}

func (lst *List) lPushNode(node *ListNode) {
	node.Next = lst.dummyHead.Next
	lst.dummyHead.Next = node
	node.Next.Prev = node
	node.Prev = lst.dummyHead
}

func (lst *List) lPop() interface{} {
	node := lst.dummyHead.Next
	lst.dummyHead.Next = node.Next
	node.Next.Prev = lst.dummyHead
	return node.Content.Value()
}

func (lst *List) rPush(elem interface{}, expiration int64) {
	node := ListNode{
		Content: NewElem(elem, expiration),
		Prev:    nil,
		Next:    nil,
	}
	lst.rPushNode(&node)
}

func (lst *List) rPushNode(node *ListNode) {
	node.Next = lst.dummyTail
	node.Prev = lst.dummyTail.Prev
	node.Prev.Next = node
	lst.dummyTail.Prev = node
}

func (lst *List) rPop() interface{} {
	node := lst.dummyTail.Prev
	lst.dummyTail.Prev = node.Prev
	node.Prev.Next = lst.dummyTail
	return node.Content.Value()
}

func (lst *List) remove(node *ListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
