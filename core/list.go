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

func (lst *List) Find(key interface{}) (interface{}, RespCode) {
	for iterator := lst.dummyHead.Next; iterator != lst.dummyTail; iterator = iterator.Next {
		// move node to list head
		if iterator.Content.Key() == key {
			lst.remove(iterator)
			lst.lPushNode(iterator)
			if iterator.Content.IsExpired() {
				return iterator.Content.Value(), Stale
			} else {
				return iterator.Content.Value(), HIT
			}
		}
	}
	return nil, MISS
}

func (lst *List) Add(elem CacheElem) {
	lst.rPush(elem)
}

// Delete returns true if some elem is actually deleted, else returns false
func (lst *List) Delete(key interface{}) bool {
	for iterator := lst.dummyHead.Next; iterator != lst.dummyTail; iterator = iterator.Next {
		// move node to list head
		if iterator.Content.Key() == key {
			lst.remove(iterator)
			lst.lPushNode(iterator)
			return true
		}
	}
	return false
}

func (lst *List) lPush(elem CacheElem) {
	node := ListNode{
		Content: elem,
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

func (lst *List) rPush(elem CacheElem) {
	node := ListNode{
		Content: elem,
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
