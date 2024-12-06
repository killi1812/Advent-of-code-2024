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
	printArr(pmap)
	rez1 := count(pmap)
	rez1++
	fmt.Printf("rez1: \n%v\n", rez1)
}

func count(arr [][]rune) int {
	sum := 0
	for _, row := range arr {
		for _, char := range row {
			if char == passed {
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
	fmt.Printf("playerLocation: %v\n", playerLocation)
	pchar := arr[playerLocation.Y][playerLocation.X]
	csprite := slices.Index(playerSprite, rune(pchar))
	csprite++
	if csprite >= len(playerSprite) {
		csprite = 0
	}
	fmt.Printf("csprite: %v\n", csprite)
	fmt.Printf("pchar: %v\n", pchar)

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

const object = '#'
const passed = 'X'

var up = [2]int{0, 1}
var left = [2]int{-1, 0}
var down = [2]int{0, -1}
var right = [2]int{1, 0}

func goDirection(arr [][]rune, bounds location, plocation *location, direction [2]int) int {
	fmt.Printf("bounds: %v\n", bounds)
	if plocation.Y-direction[1] == -1 || plocation.Y-direction[1] == bounds.Y || plocation.X-direction[0] == -1 || plocation.X-direction[0] == bounds.X {
		return -1
	}

	if arr[plocation.Y-direction[1]][plocation.X-direction[0]] == object {
		return 0
	}

	tmp := arr[plocation.Y][plocation.X]
	arr[plocation.Y][plocation.X] = passed

	plocation.X -= direction[0]
	plocation.Y -= direction[1]
	arr[plocation.Y][plocation.X] = tmp

	fmt.Printf("plocation: %v\n", plocation)

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
