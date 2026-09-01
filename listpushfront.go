package piscine

func ListPushFront(l *List, data interface{}) {
	new_node := &NodeL{data, l.Head}
	l.Head = new_node

	if l.Tail == nil {
		l.Tail = new_node
	}
}
