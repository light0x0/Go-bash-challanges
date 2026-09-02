package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 || l == nil {
		return nil
	}
	index := 0
	cp := l
	for cp != nil && index < pos {
		if index != pos {
			cp = cp.Next
			index++
		}
	}

	return cp
}
