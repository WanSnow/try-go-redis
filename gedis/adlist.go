package gedis

const (
	AL_START_HEAD = 0
	AL_START_TAIL = 1
)

type ListNode struct {
	Prev *ListNode
	Next *ListNode
	Value interface{}
}

type ListIter struct {
	Next *ListNode
	Direction int64
}

type List struct {
	Head *ListNode
	Tail *ListNode
	Dup func(ptr interface{}) interface{}
	Free func(ptr interface{})
	Match func(ptr interface{}, key interface{})int64
	Len uint64
}

func ListCreate() *List{
	return  new(List)
}

func ListRelease(list *List) {
	var current, next *ListNode

	current = list.Head
	for i := list.Len;i != 0; i-- {
		next = current.Next
		if list.Free != nil {
			list.Free(current.Value)
		}
		current = nil
		current = next
	}
	list = nil
}

func ListAddNodeHead(list *List, value interface{}) *List{
	node := new(ListNode)
	node.Value = value
	if list.Len == 0 {
		list.Head = node
		list.Tail = node
	} else {
		node.Next = list.Head
		list.Head.Prev = node
		list.Head = node
	}
	list.Len++
	return list
}

func ListAddNodeTail(list *List, value interface{}) *List {
	node := new(ListNode)
	node.Value = value
	if list.Len == 0 {
		list.Head = node
		list.Tail = node
	} else {
		node.Prev = list.Tail
		list.Tail.Next = node
		list.Tail = node
	}
	list.Len++
	return list
}

func ListInsertNode(list *List, oldNode *ListNode, value interface{}, after int) *List {
	node := new(ListNode)
	node.Value = value
	if after != 0 {
		node.Prev = oldNode
		node.Next = oldNode.Next
		if list.Tail == oldNode {
			list.Tail = node
		}
	} else {
		node.Next = oldNode
		node.Prev = oldNode.Prev
		if list.Head == oldNode {
			list.Head = node
		}
	}
	if node.Prev != nil{
		node.Prev.Next = node
	}
	if node.Next != nil{
		node.Next.Prev = node
	}
	list.Len++
	return list
}

func ListDelNode(list *List, node *ListNode) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		list.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		list.Tail = node.Prev
	}
	if list.Free != nil {
		list.Free(node.Value)
	}
	node = nil
	list.Len--
}

func ListGetIterator(list *List, direction int64) *ListIter{
	iter := new(ListIter)

	if direction == AL_START_HEAD{
		iter.Next = list.Head
	} else {
		iter.Next = list.Tail
	}
	iter.Direction = direction
	return iter
}

func ListReleaseIterator(iter *ListIter) {
	iter = nil
}

func ListRewind(list *List, li *ListIter) {
	li.Next = list.Head
	li.Direction = AL_START_HEAD
}

func ListRewindTail(list *List, li *ListIter) {
	li.Next = list.Tail
	li.Direction = AL_START_TAIL
}

func ListNext(iter *ListIter) *ListNode {
	current := iter.Next

	if current != nil{
		if iter.Direction == AL_START_HEAD{
			iter.Next = current.Next
		} else {
			iter.Next = current.Prev
		}
	}
	return current
}

func ListDup(orig *List) *List{
	cp := ListCreate()
	iter := new(ListIter)
	if cp == nil {
		return nil
	}
	cp.Dup = orig.Dup
	cp.Free = orig.Free
	cp.Match = orig.Match
	iter = ListGetIterator(orig, AL_START_HEAD)
	for node := ListNext(iter); node != nil;node = ListNext(iter){
		var value interface{}
		if cp.Dup != nil{
			value = cp.Dup(node.Value)
			if value == nil {
				ListRelease(cp)
				ListReleaseIterator(iter)
				return nil
			}
		} else {
			value = node.Value
		}
		if ListAddNodeTail(cp, value) == nil {
			ListRelease(cp)
			ListReleaseIterator(iter)
			return nil
		}
	}
	ListReleaseIterator(iter)
	return nil
}

func ListSearchKey(list *List, key interface{}) *ListNode {
	iter := ListGetIterator(list, AL_START_HEAD)
	for node := ListNext(iter); node != nil;node = ListNext(iter) {
		if list.Match != nil{
			if list.Match(node.Value, key) != 0{
				ListReleaseIterator(iter)
				return node
			}
		} else {
			if key == node.Value {
				ListReleaseIterator(iter)
				return node
			}
		}
	}
	ListReleaseIterator(iter)
	return nil
}

func ListIndex(list *List, index int64) *ListNode{
	n := new(ListNode)

	if index < 0{
		index = (-index)-1
		n = list.Tail
		for ;index != 0 && n != nil;index--{
			n = n.Prev
		}
	} else {
		n = list.Head
		for ;index != 0 && n != nil;index--{
			n = n.Next
		}
	}
	return n
}
