package main

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInts() ([]int){
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []int;
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, i);
	}
	return numbers;
}

