package main

import(
	"os"
	"bufio"
	"regexp"
	"fmt"
	"strconv"
	"image"
	"image/color"
	"image/gif"
)

var validEntry = regexp.MustCompile(`(x|y)=(-?\d+)?, (x|y)=(-?\d+)\.\.(-?\d+)`)

type point struct {
	x int
	y int
}

var minx = 1000000
var maxx = -1000000
var miny = 1000000
var maxy = -100000

// optimization:
var clayAtLevel = make(map[int][]point)

var clay = make(map[point]bool)
var flowing = make(map[point]bool)
var stalled = make(map[point]bool)

func printMap(filename string){
	fo, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	fw := bufio.NewWriter(fo)

	width := maxx - minx + 3
	height := maxy + 1

	palette := []color.Color{color.White, color.Black, color.CMYK{255,0,0,0}, color.CMYK{0,255,0,0}}
	rect := image.Rect(0, 0,
		width,
		height)
	img := image.NewPaletted(rect, palette)
	for k, _ := range clay {
		img.SetColorIndex((k.x - minx) + 1, k.y, 1)
	}

	for k, _ := range flowing {
		img.SetColorIndex((k.x - minx) + 1, k.y, 2)
	}

	for k, _ := range stalled {
		img.SetColorIndex((k.x - minx) + 1, k.y, 3)
	}


	anim := gif.GIF{Delay: []int{0}, Image: []*image.Paletted{img}}
	gif.EncodeAll(fw, &anim)
}

func readMap(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			match := validEntry.FindStringSubmatch(line)
			vertical := match[1] == "x"
			root, _ := strconv.Atoi(match[2])
			from, _ := strconv.Atoi(match[4])
			to, _ := strconv.Atoi(match[5])
			for i := from; i <= to; i++ {
				p := point{i,root}
				if vertical {
					p = point{root,i}
				}
				//
				clay[p] = true
				clayAtLevel[p.y] = append(clayAtLevel[p.y],p)
				if p.x < minx {
					minx = p.x
				}
				if p.x > maxx {
					maxx = p.x
				}
				if p.y > maxy {
					maxy = p.y
				}
				if p.y < miny {
					miny = p.y
				}
			}
		}
	}
	flowing[point{500,0}] = true
}

func empty(p point) bool {
	return !clay[p] && !stalled[p] && !flowing[p]
}

func hard(p point) bool {
	return clay[p] || stalled[p]
}

func water(p point) bool {
	return stalled[p] || flowing[p]
}

// could use some optimization
func flow() bool {
	flowen := false
	for p,_ := range flowing {
		if flowing[p] {
			left := point{p.x - 1, p.y}
			right := point{p.x + 1, p.y}
			down := point{p.x, p.y + 1}
			if (empty(down) && down.y <= maxy) {
				flowing[down] = true
				flowen = true
			}
			if hard(down) {
				if empty(left) {
					flowing[left] = true
					flowen = true
				}
				if empty(right) {
					flowing[right] = true
					flowen = true
				}
			}
		}
	}
	return flowen
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func waterWithinClay(p point) (bool,int,int) {
	closestLeftClayX := -1
	closestRightClayX := -1
	rightClayMinDistance := 100000000
	leftClayMinDistance := 100000000
	for _,explored := range clayAtLevel[p.y]{
		if clay[explored] {
			distance := Abs(p.x - explored.x)
			if explored.x < p.x {
				if distance < leftClayMinDistance {
					closestLeftClayX = explored.x
					leftClayMinDistance = distance
				}
			} else {
				if distance < rightClayMinDistance {
					closestRightClayX = explored.x
					rightClayMinDistance = distance
				}

			}
		}
	}
	if closestRightClayX == -1 || closestLeftClayX == -1  {
		return false,0,0
	}
	for i := closestLeftClayX + 1; i < closestRightClayX; i++ {
		if !water(point{i,p.y}) {
			return false,0,0
		}
	}
	return true, closestLeftClayX, closestRightClayX
}

func stall() bool {
	gotStalled := false
	for p,_ := range flowing {
		if flowing[p] { // could have changed!
			down := point{p.x, p.y + 1}
			if hard(down) {
				withinClay,leftClay,rightClay := waterWithinClay(p)
				if withinClay {
					for j := leftClay + 1; j < rightClay; j++ {
						p2 := point{j,p.y}
						delete(flowing, p2)
						stalled[p2] = true
						gotStalled = true
					}
				}
			}
		}
	}
	return gotStalled
}

func waterCount() int {
	return len(stalled)
}

func main(){
	readMap()
	changed := true
	for changed {
		changed = flow() || stall()
	}
	printMap("map.gif")
	fmt.Println(waterCount())
}
