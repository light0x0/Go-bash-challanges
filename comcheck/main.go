package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, value := range args {
		if value == "galaxy" || value == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}
