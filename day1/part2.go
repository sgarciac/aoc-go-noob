package main

import (
	"fmt"
)

func main () {
	presence := make(map[int]bool);
	numbers := ReadInts();
	current := 0
	presence[current] = true;
	current = numbers[0];
	index := 1;

	for !presence[current]  {
		presence[current] = true;
		current += numbers[index];
		index++;
		if (index == len(numbers)) {
			index = 0;
		}
	}

	fmt.Printf("%d\n", current);
}
