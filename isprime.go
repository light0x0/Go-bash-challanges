package piscine

func IsPrime(nb int) bool {
	from := 2

	if nb < 2 {
		return false
	}
	for from < nb {
		if nb%from == 0 {
			return false
		}
		from++
	}
	return true
}
