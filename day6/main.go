package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type location struct {
	X int
	Y int
}

var playerSprite = []rune{'^', '>', 'v', '<'}

func main() {
	input := read()
	pmap := stringsToRunes(input[:len(input)-1])
	for _, v := range input {
		fmt.Printf("%s\n", v)
	}

	walk(pmap)
	//printArr(pmap)
	rez1 := count(pmap, passed)
	rez1++
	fmt.Printf("rez1: \n%v\n", rez1)

	count(pmap, newObsticle)
	fmt.Printf("rez1: \n%v\n", countRout)
}

func count(arr [][]rune, obj rune) int {
	sum := 0
	for _, row := range arr {
		for _, char := range row {
			if char == obj {
				sum++
			}
		}
	}
	return sum
}

func findStart(arr [][]rune) location {
	for i, row := range arr {
		for j, char := range row {
			if slices.Contains(playerSprite, char) {
				return location{X: j, Y: i}
			}
		}
	}
	return location{}
}

func stringsToRunes(strArray []string) [][]rune {
	var runeArray [][]rune
	fmt.Println(len(strArray))
	fmt.Println(len(strArray[0]))
	for _, str := range strArray {
		runeArray = append(runeArray, []rune(str))
	}

	return runeArray
}

func walk(arr [][]rune) {
	bounds := location{X: len(arr[0]), Y: len(arr)}
	playerLocation := findStart(arr)
	for {
		rez := forward(arr, bounds, &playerLocation)
		if rez == -1 {
			break
		}
		rotate(arr, playerLocation)
	}
}

func rotate(arr [][]rune, playerLocation location) {
	pchar := arr[playerLocation.Y][playerLocation.X]
	csprite := slices.Index(playerSprite, rune(pchar))
	csprite++
	if csprite >= len(playerSprite) {
		csprite = 0
	}

	arr[playerLocation.Y][playerLocation.X] = rune(playerSprite[csprite])
}

func forward(arr [][]rune, bounds location, plocation *location) int {
	//TODO loop
	cont := 1
	for cont == 1 {
		switch arr[plocation.Y][plocation.X] {
		case playerSprite[0]:
			cont = goDirection(arr, bounds, plocation, up)
			break
		case playerSprite[1]:
			cont = goDirection(arr, bounds, plocation, left)
			break
		case playerSprite[2]:
			cont = goDirection(arr, bounds, plocation, down)
			break
		case playerSprite[3]:
			cont = goDirection(arr, bounds, plocation, right)
			break
		}
	}
	return cont
}

const obsticle = '#'
const passed = 'X'
const newObsticle = 'O'

var up = [2]int{0, 1}
var left = [2]int{-1, 0}
var down = [2]int{0, -1}
var right = [2]int{1, 0}
var countRout = 0

func nextDirection(direction [2]int) [2]int {
	switch direction {
	case up:
		return left

	case left:
		return down
	case down:
		return right
	case right:
		return up
	}
	return direction
}

func checkNext(arr [][]rune, bounds, plocation location, direction [2]int) int {
	if plocation.Y-direction[1] == -1 || plocation.Y-direction[1] == bounds.Y || plocation.X-direction[0] == -1 || plocation.X-direction[0] == bounds.X {
		return -1
	}

	if arr[plocation.Y-direction[1]][plocation.X-direction[0]] == obsticle {
		return 2
	}

	if arr[plocation.Y-direction[1]][plocation.X-direction[0]] == passed {
		return 1
	}

	return 0
}

// -1 out
// 0 nothing
// 1 passed
// 2 obsticle

func checkDirection(arr [][]rune, bounds, plocation location, direction [2]int) bool {
	val := -1
	loc := location{
		X: plocation.X,
		Y: plocation.Y,
	}
	for {
		loc.X = loc.X - direction[0]
		loc.Y = loc.Y - direction[1]
		old := val
		val = checkNext(arr, bounds, loc, direction)
		if val == -1 {
			return false
		}
		//TODO mora zavrsavati s XXX#
		if val == 2 && (old == 1 || old == -1) {
			break
		}
	}
	fmt.Printf("Found direction %v, \n", loc)
	printArr(arr)
	return true
}

func goDirection(arr [][]rune, bounds location, plocation *location, direction [2]int) int {
	if plocation.Y-direction[1] == -1 || plocation.Y-direction[1] == bounds.Y || plocation.X-direction[0] == -1 || plocation.X-direction[0] == bounds.X {
		return -1
	}
	if arr[plocation.Y-direction[1]][plocation.X-direction[0]] == obsticle {
		return 0
	}

	//TODO check for circular path must be x and end with #
	// and place a new object if true
	if checkDirection(arr, bounds, *plocation, nextDirection(direction)) {
		countRout++
		//arr[plocation.Y-direction[1]][plocation.X-direction[0]] = newObsticle
	}

	tmp := arr[plocation.Y][plocation.X]
	arr[plocation.Y][plocation.X] = passed

	plocation.X -= direction[0]
	plocation.Y -= direction[1]
	arr[plocation.Y][plocation.X] = tmp

	return 1
}

func printArr(arr [][]rune) {

	for _, row := range arr {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
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
