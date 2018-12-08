package main

import (
	"fmt"
)

var total = 0

func readTree() {
	var childrenCount, dataCount int
	fmt.Scanf("%d",&childrenCount)
	fmt.Scanf("%d",&dataCount)
	for i := 0; i < childrenCount; i++ {
		readTree()
	}
	for j := 0; j < dataCount; j++ {
		var entry int
		fmt.Scanf("%d",&entry)
		total += entry
	}
}

func main(){
	readTree()
	fmt.Printf("%d", total)
}
