package piscine

func ListReverse(l *List) {
	var prev *NodeL

	cp := l.Head
	old := l.Head

	for cp != nil {
		next := cp.Next
		cp.Next = prev
		prev = cp
		cp = next
	}
	l.Head = prev
	l.Tail = old
}
