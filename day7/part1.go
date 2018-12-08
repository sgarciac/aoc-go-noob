package main

import (
	"fmt"
)

func main() {
	readInput();
	makeRootsAvailable();
	for len(available) > 0 {
		head := popAvailable()
		done[head] = true
		fmt.Printf("%c",head + offset)
		for i := 0; i <= maxNode; i++ {
			if !done[i] && goesTo(head,i) && allParentsDone(i) {
				makeAvailable(i)
			}
		}
	}
	fmt.Printf("\n")
}
