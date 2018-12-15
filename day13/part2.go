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
	crashed bool
}

func (this cart) String() string {
	return fmt.Sprintf("%d) [%d,%d] (%d,%d) %d",this.id, this.pos.x, this.pos.y, this.dir.x, this.dir.y, this.r)
}

var roads [1000][1000]rune

func printRoads() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 50; j++ {
			if roads[i][j] != 0 {
				fmt.Printf("%s",string(roads[i][j]))
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Println();
	}
}

func printCarts(carts []cart){
	for _, cart := range carts {
		fmt.Println(cart)
	}
}


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
	return cart{id, position{x, y}, dir, acw, false}
}

func sortCarts(carts []cart){
	sort.Slice(carts,func(i,j int) bool {
		return carts[i].pos.y < carts[j].pos.y || (carts[i].pos.y == carts[j].pos.y && carts[i].pos.x < carts[j].pos.x)
	})
}

func readRoads() []cart{
	var carts []cart
	scanner := bufio.NewScanner(os.Stdin)
	row := 0
	cartsCount := 0
	for scanner.Scan() {
		str := scanner.Text()
		for col, char := range str {
			if strings.ContainsRune("+/\\",char) {
				roads[row][col] = char
			} else if strings.ContainsRune("^v<>",char) {
				carts = append(carts, createCart(cartsCount, col, row,char))
				cartsCount++
			}
		}
		row++
	}
	return carts
}

func advanceCart(c cart) cart {
	newx := c.pos.x + c.dir.x
	newy := c.pos.y + c.dir.y
	var newdx = c.dir.x
	var newdy = c.dir.y
	var newr = c.r
	switch roads[newy][newx] {
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
			newdy = c.dir.x
			if (c.dir.y == 0){
				newdx *= -1
				newdy *= -1
			}
		case cw:
			newdx = c.dir.y
			newdy = c.dir.x
			if (c.dir.x == 0){
				newdx *= -1
				newdy *= -1
			}
		}
		newr = nextRotation(c.r)
	}
	return cart{c.id,position{newx,newy},direction{newdx,newdy},newr,c.crashed}
}

func removeCart(id int, carts []cart) []cart {
	i := 0
	for i < len(carts) {
		if carts[i].id == id {
			break
		}
		i++;
	}
	if i < len(carts){
		return append(carts[:i],carts[i+1:]...)
	} else {
		return carts
	}
}

func indexesOfUncollisionedCartsInPos(x,y int, carts []cart) []int{
	var is []int
	for pos, cart := range carts {
		if (!cart.crashed && cart.pos.x == x && cart.pos.y == y){
			is = append(is, pos)
		}
	}
	return is
}

func tic(carts []cart) []cart {
	sortCarts(carts)
	for i := 0; i < len(carts); i++ {
		carts[i] = advanceCart(carts[i])
		collisions := indexesOfUncollisionedCartsInPos(carts[i].pos.x, carts[i].pos.y, carts)
		if len(collisions) > 1 { // actual collisions
			fmt.Println("boom!")
			for _,j := range collisions {
				carts[j].crashed = true
			}
		}
	}
	var newCarts []cart
	for _,cart := range carts {
		if !cart.crashed {
			newCarts = append(newCarts, cart)
		}
	}
	return newCarts
}

func main(){
	carts := readRoads()
	loops := 0
	for true {
		carts := tic(carts)
		if len(carts) == 1 {
			fmt.Printf("%d,%d\n",carts[0].pos.x,carts[0].pos.y)
			break
		}
		loops++
	}
	fmt.Printf("%d cycles\n",loops)

}
