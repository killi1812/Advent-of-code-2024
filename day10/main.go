package main

import (
	"fmt"
	"os"
	"strings"
)

type cords struct {
	X int
	Y int
}

func (this cords) add(loc cords) cords {
	return cords{
		X: this.X + loc.X,
		Y: this.Y + loc.Y,
	}
}

func (this cords) subtract(loc cords) cords {
	return cords{
		X: this.X - loc.X,
		Y: this.Y - loc.Y,
	}
}

type HikeMap struct {
	m             [][]rune
	length, width int
	starts, ends  []cords
}

func (this *HikeMap) printMap() {
	for _, row := range this.m {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func (this *HikeMap) at(loc cords) (rune, error) {
	if this.length <= loc.Y || this.width <= loc.X || loc.X < 0 || loc.Y < 0 {
		return '.', fmt.Errorf("Out of bounds")
	}
	return this.m[loc.Y][loc.X], nil

}

func newHikeMap(arr [][]rune) *HikeMap {
	return &HikeMap{
		m:      arr,
		length: len(arr),
		width:  len(arr[0]),
	}
}

func (this HikeMap) findStarts() []cords {
	starts := []cords{}
	for x := 0; x < this.width; x++ {
		if this.m[0][x] == '0' {
			starts = append(starts, cords{X: x, Y: 0})
		}
		if this.m[this.length-1][x] == '0' {
			starts = append(starts, cords{X: x, Y: this.length - 1})

		}
	}

	for y := 1; y < this.length-1; y++ {
		if this.m[y][0] == '0' {
			starts = append(starts, cords{X: 0, Y: y})
		}
		if this.m[y][this.width-1] == '0' {
			starts = append(starts, cords{Y: y, X: this.width - 1})
		}
	}

	return starts
}

func (this HikeMap) findEnds() []cords {
	ends := []cords{}
	for x, row := range this.m {
		for y, char := range row {
			if char == '9' {
				ends = append(ends, cords{X: x, Y: y})
			}
		}
	}
	return ends
}

var directions = [4]cords{{X: -1, Y: 0}, {X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}}

func (this HikeMap) nextSteap(poz, dir cords) *cords {
	a, _ := this.at(poz)
	height := rtoi(a)
	tmp := poz.add(dir)
	rn, _ := this.at(tmp)
	br := rtoi(rn)
	//fmt.Printf("height %v, tmp %v, rn %v\n", height, tmp, br)
	if br != height+1 {
		return nil
	}
	fmt.Printf("next step to %v: %v\n", tmp, br)
	return &tmp
}

func rtoi(char rune) int {
	return int(char - '0')
}

type path struct {
	start  cords
	end    cords
	steaps []cords
	length int
}

// returns -1 if smaller, 0 if same or different, 1 if bigger
func (this path) comp(other path) int {
	if this.start != other.start || this.end != other.end {
		return 0
	}

	diff := this.length - other.length
	if diff < 0 {
		return -1
	} else if diff > 0 {
		return 1
	}
	return 0
}

func findPaths(m HikeMap) []path {
	paths := []path{}
	for _, path := range m.starts {
		fmt.Printf("path: %v\n", path)
		moveNext(m, path)
	}
	return paths
}

func moveNext(m HikeMap, loc cords) {
	for _, dir := range directions {
		newLoc := m.nextSteap(loc, dir)
		if newLoc == nil {
			continue
		}
		m.nextSteap(*newLoc, dir)
	}
	return
}

func main() {
	input := read()
	mp := newHikeMap(stringsToRunes(input))
	mp.printMap()
	mp.starts = mp.findStarts()
	mp.ends = mp.findEnds()
	findPaths(*mp)
	fmt.Printf("mp.starts: %v\n", mp.starts)
	fmt.Printf("mp.ends: %v\n", mp.ends)
}

func stringsToRunes(strArray []string) [][]rune {
	var runeArray [][]rune
	for _, str := range strArray {
		if str == "" {
			continue
		}
		runeArray = append(runeArray, []rune(str))
	}

	return runeArray
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
