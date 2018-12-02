package main

import (
	"bufio"
	"os"
)

func ReadLines() ([]string){
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string;
	for scanner.Scan() {
		str := scanner.Text();
		if(len(str) > 0){
			lines = append(lines, str);
		}
	}
	return lines;
}

