package main

import (
	"fmt"

	"github.com/menggggggg/algorithm/list"
)

func main() {
	n5 := list.New("5", nil)
	n4 := list.New("4", n5)
	n3 := list.New("3", n4)
	n2 := list.New("2", n3)
	n1 := list.New("1", n2)

	new := list.Reverse(n1)
	for new != nil {
		fmt.Print(new.Data(), " ")
		new = new.Next()
	}
	fmt.Println()
}
