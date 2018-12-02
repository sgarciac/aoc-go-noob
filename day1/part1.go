package main

import (
	"fmt"
)

func main () {
	numbers := ReadInts();
	total := 0
	for _, value := range numbers {
		total += value;
	}
	fmt.Printf("%d\n", total);
}
