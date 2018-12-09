package main

import (
	"fmt"
)

type node struct {
	next *node
	previous *node
	value int
}

func createCircularList(value int) *node {
	head := node{nil,nil, value}
	head.next = &head
	head.previous = &head
	return &head
}

// insert after and return inserted
func insertAfter(value int, n *node) *node {
	newNode := node{n.next, n, value}
	n.next.previous = &newNode
	n.next = &newNode
	return &newNode
}

// removes and returns next
func remove(n *node) *node {
	n.previous.next = n.next
	n.next.previous = n.previous
	return n.next
}

func rotateForward(count int, n *node) *node{
	for i := 0; i < count; i++ {
		n = n.next
	}
	return n
}

func rotateBackward(count int, n *node) *node{
	for i := 0; i < count; i++ {
		n = n.previous
	}
	return n
}

// asumes value in list
func skipTo(value int, n *node) *node{
	for ;n.value != value; n = n.next {

	}
	return n
}

func printList(n *node) {
	fmt.Printf("(%d) ",n.value)
	for current := n.next; current != n; current = current.next {
		fmt.Printf("(%d) ",current.value)
	}
	fmt.Println();
}

func main(){
	var score [500]int
	target := 7143100
	players := 476
	currentPlayer := 0
	current := createCircularList(0)
	for marble := 1; marble <= target; marble++ {
		if (marble % 23 != 0){
			current = rotateForward(1,current)
			current = insertAfter(marble, current)
		} else {
			score[currentPlayer] += marble
			current = rotateBackward(7, current)
			score[currentPlayer] += current.value
			current = remove(current)
		}

		currentPlayer = (currentPlayer + 1) % players

	}

	maxScore := -1
	for i := 0 ; i < players; i++ {
		if score[i] > maxScore {
			maxScore = score[i]
		}
	}
	fmt.Println(maxScore)
}
