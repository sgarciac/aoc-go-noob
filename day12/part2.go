package main

import (
	"fmt"
	avl "github.com/emirpasic/gods/trees/avltree"
	"os"
	"bufio"
	"strings"
)

var rules = make(map[[5]bool]bool)

type segment struct {
	L2,L1,C,R1,R2 bool
}


// global state. keep two arrays
func lineToGarden(line string) *avl.Tree {
	tree := avl.NewWithIntComparator()
	for pos, char := range line {
		if char == '#' {
			tree.Put(pos, true)
		}
	}
	return tree
}

func nextGarden(tree *avl.Tree) *avl.Tree {
	newTree := avl.NewWithIntComparator()
	max := tree.Right().Key.(int)
	min := tree.Left().Key.(int)
	for i := min - 5; i < max + 5; i++ {
		_,L2 := tree.Get(i-2)
		_,L1 := tree.Get(i-1)
		_,C := tree.Get(i)
		_,R1 := tree.Get(i+1)
		_,R2 := tree.Get(i+2)

		if rules[[5]bool{L2,L1,C,R1,R2}] {
			newTree.Put(i,true)
		}
	}
	return newTree
}



func cycleSize(garden *avl.Tree) (int,int) {
	loops := 0
	for ng := garden; true; ng = nextGarden(ng){
		fmt.Printf("%d %d\n",loops,len(ng.Keys()))
		fmt.Println(ng.Keys())
		loops++
	}
	return -1,-1
}


func runeToPlant(r rune) bool{
	if r == '.' {
		return false
	} else {
		return true
	}
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	initialStateLine := scanner.Text()[15:]
	garden := lineToGarden(initialStateLine)
	scanner.Scan()
	scanner.Text()
	for scanner.Scan() {
		str := scanner.Text()
		if(len(str) > 0){
			if strings.HasSuffix(str, "=> #") {
				var rule [5]bool
				for pos,c := range str[0:5] {
					rule[pos] = runeToPlant(c)
				}
				rules[rule] = true
			}
		}
	}
	// finally it was very silly
	for i := 0; i < 1000; i++ {
		garden = nextGarden(garden)
	}
	total := 0
	for _, val := range garden.Keys() {
		total += val.(int)
	}
	fmt.Println(total)
	fmt.Println(len(garden.Keys()))
	fmt.Println(((50000000000 - 1000) * 72) + 74022)
}
