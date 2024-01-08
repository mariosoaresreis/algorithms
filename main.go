package main

import (
	"algorithms/com/marioreis/domain"
	"fmt"
)

func main() {
	var tree5 domain.Tree = domain.NewTree(nil, nil, 5)
	var tree8 domain.Tree = domain.NewTree(nil, nil, 8)
	var tree15 domain.Tree = domain.NewTree(nil, nil, 15)
	var tree13 domain.Tree = domain.NewTree(nil, nil, 13)
	var tree10 domain.Tree = domain.NewTree(&tree5, &tree8, 10)
	var tree20 domain.Tree = domain.NewTree(&tree15, &tree13, 20)
	var tree1 domain.Tree = domain.NewTree(&tree10, &tree20, 1)

	fmt.Println(domain.FindTree(2, &tree1))
	fmt.Println(domain.FindTree(20, &tree1))
	fmt.Println(domain.FindTree(15, &tree1))
	fmt.Println(domain.FindTree(16, &tree1))
}
