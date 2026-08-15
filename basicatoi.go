package piscine

func BasicAtoi(s string) int {
    num := 0
    for _,value := range s{
        num = num * 10 + int(value - '0')
    }
    return num
}