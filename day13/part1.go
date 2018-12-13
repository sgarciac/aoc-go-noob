package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"sort"
)

type rotation int

const (
	acw rotation = iota // anticlockwise
	none
	cw // clockwise
)

func nextRotation(r rotation) rotation{
	return (r + 1) % 3
}

type position struct {
	x,y int
}

type direction struct {
	x,y int
}

type cart struct {
	id int
	pos position
	dir direction
	r rotation
}

func (this cart) String() string {
	return fmt.Sprintf("%d) [%d,%d] (%d,%d) %d",this.id, this.pos.x, this.pos.y, this.dir.x, this.dir.y, this.r)
}

var roads [1000][1000]rune
var carts []cart

func createCart(id, x, y int,rune rune) cart {
	dir := direction{0,0}
	switch rune {
	case 'v':
		dir = direction{0,1}
	case '^':
		dir = direction{0,-1}
	case '<':
		dir = direction{-1,0}
	case '>':
		dir = direction{1,0}
	}
	return cart{id, position{x, y}, dir, acw}
}

func sortCarts(){
	sort.Slice(carts,func(i,j int) bool {
		return carts[i].pos.y < carts[j].pos.y || (carts[i].pos.y == carts[j].pos.y && carts[i].pos.x < carts[j].pos.x)
	})
}

func readRoads(){
	scanner := bufio.NewScanner(os.Stdin)
	row := 0
	cartsCount := 0
	for scanner.Scan() {
		str := scanner.Text()
		for col, char := range str {
			if strings.ContainsRune("|-+/\\",char) {
				roads[row][col] = char
			} else if strings.ContainsRune("^v<>",char) {
				carts = append(carts, createCart(cartsCount, col, row,char))
				cartsCount++
			}
		}
		row++
	}
	sortCarts()
}

func advanceCart(c cart) cart {
	newx := c.pos.x + c.dir.x
	newy := c.pos.y + c.dir.y
	var newdx = c.dir.x
	var newdy = c.dir.y
	var newr = c.r
	switch roads[newx][newy] {
	case '/':
		newdx = c.dir.y * -1
		newdy = c.dir.x * -1
	case '\\':
		newdx = c.dir.y
		newdy = c.dir.x
	case '+':
		switch c.r {
		case acw:
			newdx = c.dir.y
			newdy = c.dir.x * -1
		case cw:
			newdx = c.dir.y
			newdy = c.dir.x * -1
		}

	}
	return cart{c.id,position{newx,newy},direction{newdx,newdy},newr}
}

func main(){
	readRoads()
}
