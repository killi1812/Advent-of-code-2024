package main

import (
	"fmt"
	"os"
	"strings"
)

type mapa struct {
	m             [][]rune
	height, width int
}

type cords struct {
	X int
	Y int
}

type antennasPair struct {
	first  cords
	second cords
}

func newMapa(arr [][]rune) *mapa {
	return &mapa{
		m:      arr,
		height: len(arr),
		width:  len(arr[0]),
	}
}

func main() {
	input := read()
	m := newMapa(stringsToRunes(input))
	m.printMap()
	pairs := m.findAntenaPairs()
	for _, pair := range pairs {
		pair.print(*m)
	}
}

func (this *antennasPair) print(m mapa) {
	char, err := m.at(this.first)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Pair: %s : %v, %v\n", string(char), this.first, this.second)
}

func stringsToRunes(strArray []string) [][]rune {
	var runeArray [][]rune
	for _, str := range strArray {
		runeArray = append(runeArray, []rune(str))
	}

	return runeArray
}

func (this *mapa) printMap() {
	for _, row := range this.m {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func (this *mapa) at(loc cords) (rune, error) {
	if this.height <= loc.Y || this.width <= loc.X {
		return '.', fmt.Errorf("Out of bounds")
	}
	return this.m[loc.Y][loc.X], nil

}

const empty = '.'

func (this *mapa) findAntenaPairs() []antennasPair {
	ret := []antennasPair{}

	for i, row := range this.m {
		for j, char := range row {
			if char == empty {
				continue
			}
			pairs := this.findPairsForAntena(char, cords{X: j, Y: i})
			ret = append(ret, pairs...)
		}
	}

	return ret
}

func (this *mapa) findPairsForAntena(find rune, loc cords) []antennasPair {
	ret := []antennasPair{}
	offset := loc.Y + 1
	for i, row := range this.m[offset:] {
		for j, char := range row {
			if char == find {
				ret = append(ret, antennasPair{first: loc, second: cords{X: j, Y: i + offset}})
			}
		}
	}
	return ret
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
