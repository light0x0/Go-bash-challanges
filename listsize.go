// Empty file
package piscine

func ListSize(l *List) int {
  count := 0
  now := l.Head
  for now != nil{
    count++
    now = now.Next
  }
  
  return count
}