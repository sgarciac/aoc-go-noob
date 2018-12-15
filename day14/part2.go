package main

import (
	"fmt"
)

var recipes = []int{3,7,1,0,1,0}
var elf1 = 4
var elf2 = 3

func create() int{
	sum := recipes[elf1] + recipes[elf2]
	numEntries := 1
	if sum < 10 {
		recipes = append(recipes, sum)
	} else {
		recipes = append(recipes, 1)
		recipes = append(recipes, sum - 10)
		numEntries = 2
	}
	elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
	elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	return numEntries
}

func equalSlices(slice1, slice2 []int) bool{
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}


func main() {
	input := []int{7,6,0,2,2,1}
	firstIndex := 0
	for true {
		numEntries := create()
		if equalSlices(recipes[len(recipes) - len(input):],input) {
			firstIndex = len(recipes) - len(input)
			break
		}
		if numEntries == 2 && equalSlices(recipes[len(recipes) - (len(input) + 1):len(recipes)-1],input){
			fmt.Println("got it!")
			firstIndex = len(recipes) - (len(input) + 1)
			break
		}
	}
	fmt.Println(firstIndex)
}
