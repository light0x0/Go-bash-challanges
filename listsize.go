// Empty file
package piscine


type NodeLx struct {
	Data interface{}
	Next *NodeLx
}

type Listx struct {
	Head *NodeLx
	Tail *NodeLx
}

func ListSize(l *Listx) int {
  count := 0
  now := l.Head
  for now != nil{
    count++
    now = now.Next
  }
  
  return count
}