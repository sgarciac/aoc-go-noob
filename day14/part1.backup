package main

import (
	"fmt"
)


var recipes = []int{3,7}
var elf1 = 0
var elf2 = 1

func create(){
	sum := recipes[elf1] + recipes[elf2]
	if sum < 10 {
		recipes = append(recipes, sum)
	} else {
		recipes = append(recipes, 1)
		recipes = append(recipes, sum - 10)
	}
	elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
	elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
}

func main() {
	input := 760221
	for len(recipes) < (input + 10) {
		create()
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d",recipes[i + input])
	}
}
