package main

import (
	"bufio"
	"os"
	"fmt"
)

func countOf(carte [][]rune, typ rune, col int, row int) int {
	rows := len(carte)
	cols := len(carte[0])
	total := 0
	for rowDelta := -1; rowDelta <= 1; rowDelta++ {
		for colDelta := -1; colDelta <= 1; colDelta++ {
			neighborRow := row + rowDelta
			neighborCol := col + colDelta
			if neighborRow >= 0 &&
				neighborRow < rows &&
				neighborCol >= 0 &&
				neighborCol < cols &&
				!(neighborCol == col && neighborRow == row) &&
				carte[neighborRow][neighborCol] == typ {
				total++
			}
		}
	}
	return total
}

func readCarte() [][]rune {
	var carte [][]rune
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			var row []rune
			for _, rune := range line {
				row = append(row, rune)
			}
			carte = append(carte, row)
		}
	}
	return carte
}

func printCarte(carte [][]rune){
	for _,row := range carte {
		for _,rune := range row {
			fmt.Printf("%c", rune)
		}
		fmt.Println()
	}
}


func scoreCarte(carte [][]rune) int{
	lumber := 0
	wood := 0
	for _,row := range carte {
		for _,rune := range row {
			switch rune {
			case '#':
				lumber++
			case '|':
				wood++
			}
		}
	}
	return wood * lumber

}


func nextCarte(carte [][]rune) [][]rune {
	newCarte := make([][]rune, len(carte))
	for y := 0; y < len(carte); y++ {
		newCarte[y] = make([]rune,len(carte[0]))
		for x := 0; x < len(carte[0]); x++ {
			newCarte[y][x] = carte[y][x]
			switch(carte[y][x]) {
			case '.':
				if countOf(carte, '|', x, y) >= 3 {
					newCarte[y][x] = '|'
				}
			case '|':
				if countOf(carte, '#', x, y) >= 3{
					newCarte[y][x] = '#'
				}
			case '#':
				if countOf(carte, '#', x, y) == 0 ||
					countOf(carte, '|', x, y) == 0 {
					newCarte[y][x] = '.'
				}
			}
		}
	}
	return newCarte
}

func main(){
	originalCarte := readCarte()
	var newCarte = originalCarte
	for i := 0; i < 10; i++ {
		newCarte = nextCarte(newCarte)
	}
	printCarte(newCarte)
	fmt.Println(scoreCarte(newCarte))
}
