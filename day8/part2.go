package main

import (
	"fmt"
)

type tree struct {
	children []*tree
	data []int
}

func treeSize(t tree) int {
	total := 0;
	if len(t.children) == 0 {
		for i := 0; i < len(t.data); i++ {
			total += t.data[i]
		}
	} else {
		for i := 0; i < len(t.data); i++ {
			if (t.data[i] > 0 && t.data[i] <= len(t.children)) {
				total += treeSize(*t.children[t.data[i] - 1])
			}

		}
	}
	return total;
}

func loadTree() *tree {
	var childrenCount, dataCount int
	fmt.Scanf("%d",&childrenCount)
	fmt.Scanf("%d",&dataCount)
	var newTree tree
	newTree.children = make([]*tree, childrenCount)
	newTree.data = make([]int, dataCount)
	for i := 0; i < childrenCount; i++ {
		newTree.children[i] = loadTree()
	}

	for j := 0; j < dataCount; j++ {
		fmt.Scanf("%d",&newTree.data[j])
	}
	return &newTree
}

func main(){
	rootTree := loadTree()
	fmt.Printf("%d\n",treeSize(*rootTree))
}
