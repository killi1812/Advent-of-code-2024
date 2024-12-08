package main

import (
	"fmt"
	"os"
	"strings"
)

const empty = '.'
const antinode = '#'

type mapa struct {
	m             [][]rune
	height, width int
}

func (this *mapa) printMap() {
	for _, row := range this.m {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func (this *mapa) count(a rune) int {
	count := 0
	for _, row := range this.m {
		for _, char := range row {
			if char == a {
				count++
			}
		}
	}
	return count
}

func (this *mapa) at(loc cords) (rune, error) {
	if this.height <= loc.Y || this.width <= loc.X || loc.X < 0 || loc.Y < 0 {
		return '.', fmt.Errorf("Out of bounds")
	}
	return this.m[loc.Y][loc.X], nil

}

func (this *mapa) change(char rune, loc cords) error {
	if this.height <= loc.Y || this.width <= loc.X || loc.X < 0 || loc.Y < 0 {
		return fmt.Errorf("Out of bounds")
	}
	this.m[loc.Y][loc.X] = char
	return nil
}

func newMapa(arr [][]rune) *mapa {
	return &mapa{
		m:      arr,
		height: len(arr),
		width:  len(arr[0]),
	}
}

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

type antennasPair struct {
	first  cords
	second cords
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
		if str == "" {
			continue
		}
		runeArray = append(runeArray, []rune(str))
	}

	return runeArray
}

func diffParis(pair antennasPair) cords {
	return cords{
		X: pair.second.X - pair.first.X,
		Y: pair.second.Y - pair.first.Y,
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
	zad01(pairs, m)
	m.printMap()
	fmt.Printf("overlapCount: %v\n", overlapCount)
	fmt.Printf("Count of antinodes is: %d\n", m.count(antinode)+overlapCount)
}

var overlapCount = 0

func zad01(arr []antennasPair, m *mapa) {
	for _, pair := range arr {
		diff := diffParis(pair)
		a, b := pair.first.subtract(diff), pair.second.add(diff)
		if tmp, _ := m.at(a); tmp == empty || tmp == antinode {
			if tmp == antinode {
				overlapCount++
			}
			m.change(antinode, a)
		}
		if tmp, _ := m.at(b); tmp == empty || tmp == antinode {
			if tmp == antinode {
				overlapCount++
			}
			m.change(antinode, b)
		}
	}
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
