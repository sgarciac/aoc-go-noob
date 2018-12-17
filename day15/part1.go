package main

import (
	"github.com/yourbasic/graph"
	"bufio"
	"os"
	"fmt"
	"sort"
)

var verybig = 1000000
var verylow = -1000000

var elf = 0
var goblin = 1

var width int
var height int

type position struct {
	x int
	y int
}

type neighbors struct {
	up position
	right position
	down position
	left position
}

type character struct {
	p position
	chartype int
	hitpoints int
	attackpoints int
}

var paths *graph.Mutable
var openPositions = make(map[position]bool)
var cinPosition map[position]*character
var characters []*character

func getNeighbors (p position) neighbors{
	up := position{p.x,p.y - 1}
	down := position{p.x,p.y + 1}
	left := position{p.x - 1,p.y}
	right := position{p.x + 1, p.y}
	return neighbors{up,right,down,left}
}

func posToId(p position) int {
	return (p.y * width) + p.x
}

func idToPos(i int) position {
	rows := i / width
	col := i % width
	return position{col, rows}
}

func removeEntriesTo(p position){
	ns := getNeighbors(p)
	for _,dir := range []position{ns.up,ns.down,ns.left,ns.right} {
		if openPositions[dir] && paths.Edge(posToId(dir),posToId(p)){
			paths.Delete(posToId(dir),posToId(p))
		}
	}
}

func addEntriesTo(p position){
	ns := getNeighbors(p)
	for _,dir := range []position{ns.up,ns.down,ns.left,ns.right} {
		if openPositions[dir] {
			paths.AddCost(posToId(dir), posToId(p), 1)
		}
	}
}

func printState() {
	for _, char := range characters {
		fmt.Printf("%d (%d,%d)\n", char.chartype, char.p.x, char.p.y)
		paths.Visit(posToId(char.p), func (w int, c int64) bool {
			dir := idToPos(w)
			fmt.Printf("  (%d,%d)",dir.x, dir.y)
			return false
		})
		fmt.Println()
	}
}

func printMap(){
	for row := 0; row < height;row++ {
		for col := 0; col < width;col++ {
			if ch, ok := cinPosition[position{col,row}]; ok {
				c := 'G'
				if ch.chartype == 0 {
					c = 'E'
				}
				if ch.hitpoints <= 0 {
					c = '.'
				}
				fmt.Printf("%s",string(c))
			} else if (openPositions[position{col,row}]) {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		for _, char := range characters {
			if char.p.y == row {
				fmt.Printf(" %s:%d",char.p, char.hitpoints)
			}
		}
		fmt.Println()
	}
}

func prepareCharacters(){
	sort.Slice(characters,func(i,j int) bool {
		return characters[i].p.y < characters[j].p.y ||
			(characters[i].p.y == characters[j].p.y && characters[i].p.x < characters[j].p.x)
	})
	cinPosition = make(map[position]*character)
	for _, c := range characters {
		cinPosition[c.p] = c
	}
}

func pickAttack(c *character) (bool, *character){
	ct := (c.chartype + 1) % 2 // pick enemy type
	ns := getNeighbors(c.p)
	var minCharacter *character
	var minHitpoint = verybig
	for _, p := range []position{ns.up,ns.left,ns.right,ns.down} {
		if target, ok := cinPosition[p]; ok {
			if target.chartype == ct {
				if target.hitpoints > 0 && target.hitpoints < minHitpoint {
					minCharacter = target
					minHitpoint = target.hitpoints
				}
			}
		}
	}
	if minHitpoint != verybig {
		return true, minCharacter
	} else {
		return false, nil
	}
}

func (this character) String() string {
	return fmt.Sprintf("(%d) [%d,%d] <%d>",this.chartype, this.p.x, this.p.y,this.hitpoints)
}

func (this position) String() string {
	return fmt.Sprintf("(%d,%d)",this.x, this.y)
}


func pickMove(c *character) (bool, position){
	ct := (c.chartype + 1) % 2 // pick enemy type
	// find the squares from which we can attack an enemy
	//fmt.Printf("picking move for %s\n", c)
	hasEnemies := false
	enemyRanges := make(map[position]bool)
	for _, enemy := range characters {
		if enemy.hitpoints > 0 && enemy.chartype == ct {
			paths.Visit(posToId(enemy.p), func (w int, c int64) bool {
				enemyRanges[idToPos(w)] = true
				hasEnemies = true
				return false
			})
		}
	}
	//fmt.Printf("enemy ranges: %s\n",enemyRanges)
	//
	if !hasEnemies {
		return false, position{-1,-1}
	}
	var minDirection position
	var minDistance = int64(verybig)
	ns := getNeighbors(c.p)
	for _, p := range []position{ns.up,ns.left,ns.right,ns.down} {
		if paths.Edge(posToId(c.p),posToId(p)) {
			_, distances := graph.ShortestPaths(paths, posToId(p))
			for i := 0; i < len(distances); i++ {
				if enemyRanges[idToPos(i)] {
					if distances[i] != -1 && (distances[i] < minDistance || (distances[1] == minDistance && (isBefore(p, minDirection)))){
						minDistance = distances[i]
						minDirection = p
					}
				}
			}
		}

	}
	if minDistance != int64(verybig) {
		return true, minDirection
	} else {
		return false, position{-1,-1}
	}
}

func isBefore(p, p2 position) bool {
	return p.y < p2.y || (p.y == p2.y && p.x < p2.x)
}


func act(c *character) {
	canAttack, attackChar := pickAttack(c)
	if canAttack {
		attack(c, attackChar)
	} else {
		canMove, movePos := pickMove(c)
		if canMove {
			move(c,movePos)
			canAttack, attackChar := pickAttack(c)
				if canAttack {
					attack(c, attackChar)
				}
		}
	}
}

func move(c *character, np position) {
	fmt.Printf("   MOVE %d from (%d,%d) to (%d,%d)\n",c.chartype, c.p.x, c.p.y, np.x, np.y)
	addEntriesTo(c.p)
	delete(cinPosition, c.p)
	c.p = np
	removeEntriesTo(np)
	cinPosition[np] = c
}

func attack(c, t *character){
	fmt.Printf("   %s ATTACKS %s\n", c, t)
	t.hitpoints -= c.attackpoints
	// target died! eeek!
	if t.hitpoints <= 0 {
		addEntriesTo(t.p)
	}
}

func finished() (bool,int) {
	elves := 0
	goblins := 0
	for _, c := range characters {
		if c.hitpoints >= 0 {
			if c.chartype == 0 {
				elves += c.hitpoints
			} else {
				goblins += c.hitpoints
			}
		}
	}
	if elves == 0 {
		return true, goblins
	}
	if goblins == 0 {
		return true, elves
	}
	return false, -1
}

func aliveChars() int {
	total := 0
	for _, ch := range characters {
		if (ch.hitpoints > 0) {
			total++
		}
	}
	return total;
}

func turn() (bool, int) {
	prepareCharacters()
	acs := aliveChars()
	acted := 0
	for _, c := range characters {
		if c.hitpoints >= 0 {
			act(c)
			acted++
			if f,ps := finished(); acted != acs && f{
				fmt.Printf("FINISHED BEFORE THE END %d %d\n", acted, acs)
				return true,ps
			}
		}
	}
	return false,-1
}

func readMap () {
	scanner := bufio.NewScanner(os.Stdin)
	row := 0
	// load all open positions and characters
	for scanner.Scan() {
		str := scanner.Text()
		if (len(str) > 0) {
			width = len(str)
			for col, char := range str {
				p := position{col,row}
				if char == '.' || char == 'E' || char == 'G' {
					openPositions[p] = true
					switch char {
					case 'E':
						characters = append(characters,&character{p,elf,200,3})
					case 'G':
						characters = append(characters,&character{p,goblin,200,3})
					}
				}
			}
			row++
			height++
		}
	}
	// create all paths among open positions
	paths = graph.New(width * height)
	for p, _ := range openPositions {
		addEntriesTo(p)
	}
	// make all positions with characters unnaccesibles
	for _,char := range characters {
		removeEntriesTo(char.p)
	}
}

func main(){
	readMap()
	printMap()
	prepareCharacters()
	turns := 0
	points := 0
	f := false
	limit := 1000000000
	for true {
		fmt.Printf("Turn %d\n",turns)
		f,points = turn()
		printMap()
		if f || turns == limit {
			break
		} else {
			turns++
		}
	}
	fmt.Println(" *************************************** ")
	printMap()

	fmt.Println(turns)
	fmt.Println(points)
	fmt.Println(turns * points)
//	fmt.Printf("%d turns, %d resting hitpoints", turns, points)
//	printState()
//	mydebug = true
//	prepareCharacters()
//	_, p := pickMove(characters[3])
//	fmt.Printf("%d,%d\n",p.x, p.y)


}
