package main

import "fmt"

func main(){
	limit := 10551358
	total := 0
	for i := 1; i <= limit; i++ {
		if limit % i == 0 {
			total += i
		}
	}
	fmt.Println(total)
}
