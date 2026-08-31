package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	new_node := &NodeL{data, nil}

	if l.Head == nil {
		l.Head = new_node
	} else {
		l.Tail.Next = new_node
	}
	l.Tail = new_node
}
