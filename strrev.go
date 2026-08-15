package main

func StrRev(s string) string {
	reversed := ""

	sum := len(s) - 1

	for sum >= 0 {
		reversed += string(s[sum])
		sum--
	}

	return reversed
}
