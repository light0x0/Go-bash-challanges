// Empty file
package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for index, value := range s {
		if value >= 'A' && value <= 'Z' {
			runes[index] = (value-'A'+14)%26 + 'A'
		} else if value >= 'a' && value <= 'z' {
			runes[index] = (value-'a'+14)%26 + 'a'
		}
	}

	return string(runes)
}
