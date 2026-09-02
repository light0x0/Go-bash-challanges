package piscine

func ListSize(l *List) int {
	current := l.Head
	count := 0

	for current != nil {
		count += 1
		current = current.Next

	}
    //testtest
	return count
}
