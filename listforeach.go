package piscine

func ListForEach(l *List, f func(*NodeL)) {
	cp := l.Head
	for cp != nil {
		f(cp)
		cp = cp.Next
	}
}
