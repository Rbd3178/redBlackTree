package main

import (
	"fmt"
	"github.com/Rbd3178/redBlackTree/tree"
)

func main() {
	fmt.Println("Test")
	myTree := tree.New()
	v, found := myTree.Search(2)
	fmt.Println(v, found)
}
