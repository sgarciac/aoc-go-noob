package main

import (
	"os"
	"regexp"
	"bufio"
	"fmt"
)

var maxNode = -1
var offset = int('A')

var validEntry = regexp.MustCompile(`^Step (.) must be finished before step (.) can begin\.$`)

var graph [26][26]bool // graph[x][y] = true if arrow from node x to node y

var available []int // a sorted list of available nodes
var done [26]bool // the done nodes

func setFromTo(from, to int) {
	graph[from][to] = true
}

func goesTo(from, to int) bool{
	return graph[from][to]
}

func printAvailable(){
	for _,i := range available {
		fmt.Printf("%c", i + offset)
	}
	fmt.Printf("\n")
}

func allParentsDone(node int) bool{
	for i := 0; i <= maxNode; i++ {
		if goesTo(i,node) && !done[i] {
			return false
		}
	}
	return true
}


func makeAvailable(newNode int){
	for i, node := range available {
		if newNode < node {
			// insert the newNode
			available = append(available, -1)
			copy(available[i+1:], available[i:])
			available[i] = newNode
			return
		} else if newNode == node {
			// if already there, dont do anything
			return
		}
	}
	available = append(available, newNode)
}

func popAvailable() int {
	head := available[0]
	available = available[1:]
	return head
}

func makeRootsAvailable() {
	for i := 0; i <= maxNode; i++ {
		parents := 0
		for j := 0; j <= maxNode; j++ {
			if (goesTo(j,i)) {
				parents++
			}
		}
		if(parents == 0){
			makeAvailable(i)
		}
	}
}

func readInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text();
		if(len(str) > 0){
			match := validEntry.FindStringSubmatch(str)
			from := int(match[1][0]) - offset
			to := int(match[2][0]) - offset
			setFromTo(from,to)
			if (to > maxNode){
				maxNode = to
			}
		}
	}
}
