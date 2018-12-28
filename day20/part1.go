package main

import (
	"github.com/yourbasic/graph"
	"fmt"
	"bufio"
	"os"
)

// DATA STRUCTURES
type Point struct {
	x int
	y int
}

type Regex []*Branch

type Branch []*Step

type Step struct { //a direction OR a regex. no idea how to to this in go
	direction rune
	regex *Regex
}

type Rooms struct {
	rooms map[Point]int
	total int
}

// PRINT FUNCTIONS FOR DSs
func (this Step) String() string {
	if this.regex != nil {
		return fmt.Sprintf("(%v)",this.regex)
	} else {
		return fmt.Sprintf("'%c'",this.direction)
	}
}

func (this *Branch) String() string {
	return fmt.Sprintf("<%v>",*this)
}

func (this *Regex) String() string {
	return fmt.Sprintf("(%v)",*this)
}


// READING THE INPUT
func readBranch(r *bufio.Reader) *Branch {
	var b Branch
	for {
		if c, _, err := r.ReadRune(); err != nil { //end of file
			return &b
		} else {
			switch c {
			case ')','|','\n':
				r.UnreadRune()
				return &b
			case '(':
				b = append(b,&Step{0,readRegex(r)})
			default:
				b = append(b,&Step{c,nil})
			}
		}
	}
}


func (r *Regex) directionsCount() int {
	total := 0
	for _,b := range *r {
		total += b.directionsCount()
	}
	return total
}

func (b *Branch) directionsCount() int {
	total := 0
	for _,s := range *b {
		if s.regex != nil {
			total += s.regex.directionsCount()
		} else {
			total += 1
		}
	}
	return total
}

func readRegex(r *bufio.Reader) *Regex{
	var regex Regex
	for {
		if c, _, err := r.ReadRune(); err != nil { //end of file
			return &regex
		} else {
			if c == ')' || c == '\n' {
				return &regex
			}
			if c == '|' {
				regex = append(regex, readBranch(r))
			} else {
				r.UnreadRune()
				regex = append(regex, readBranch(r))
			}
		}
	}
}

func getRoomId(p Point, rs *Rooms) int{
	id,ok := rs.rooms[p]
	if ok {
		return id
	} else {
		rs.rooms[p] = rs.total
		rs.total++
		return rs.total - 1
	}
}

func loadBranchToGraph(
	graph *graph.Mutable,
	branch *Branch,
	origin Point,
	rooms *Rooms) {
	for _,s := range *branch {
		if s.regex != nil {
			loadRegexToGraph(graph, s.regex,origin, rooms)
		} else {
			originId := getRoomId(origin, rooms)
			switch s.direction {
			case 'N':
				origin = Point{origin.x,origin.y+1}
			case 'S':
				origin = Point{origin.x,origin.y-1}
			case 'E':
				origin = Point{origin.x+1,origin.y}
			case 'W':
				origin = Point{origin.x-1,origin.y}
			}
			graph.AddBothCost(originId, getRoomId(origin, rooms), 1)
		}
	}
}

func loadRegexToGraph(
	graph *graph.Mutable,
	regex *Regex,
	origin Point,
	rooms *Rooms) {
	for _,b := range *regex {
		loadBranchToGraph(graph, b, origin, rooms)
	}
}

func main(){
	r := bufio.NewReader(os.Stdin)
	regex := readRegex(r)
	var rooms Rooms
	rooms.rooms = make(map[Point]int)
	paths := graph.New(regex.directionsCount() + 1)
	loadRegexToGraph(paths, regex, Point{0,0}, &rooms)
	_, distances := graph.ShortestPaths(paths, 0)
	var max int64
	max = -1
	for _, distance := range distances {
		if distance > max {
			max = distance
		}
	}
	fmt.Println(max)
}
