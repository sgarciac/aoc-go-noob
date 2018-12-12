package main

import (
	"fmt"
	"os"
	"bufio"
)

// global state. keep two arrays
var rules = make(map[int]bool)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan();
	initialStateLine := scanner.Text()[15:]
	fmt.Println(initialStateLine);
}
