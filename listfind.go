package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	cp := l.Head

	for cp != nil {
		if comp(cp.Data, ref) {
			return &cp.Data
		}
		cp = cp.Next
	}

	return nil
}
