package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := readFromFile()
	zad01(input)
}

var directions = [][2]int{
	{-1, 0},  // Up
	{1, 0},   // Down
	{0, -1},  // Left
	{0, 1},   // Right
	{-1, -1}, // Up-left
	{-1, 1},  // Up-right
	{1, -1},  // Down-left
	{1, 1},   // Down-right
}

const word = "XMAS"

func searchFrom(grid [][]rune, row, col, dRow, dCol int) bool {
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < len(word); i++ {
		nRow, nCol := row+i*dRow, col+i*dCol
		if nRow < 0 || nRow >= rows || nCol < 0 || nCol >= cols || grid[nRow][nCol] != rune(word[i]) {
			return false
		}
	}
	return true
}

// Find all occurrences of "XMAS" in the grid
func findAllOccurrences(grid [][]rune) [][2]int {
	var results [][2]int
	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 'X' { // Potential start of the word
				for _, dir := range directions {
					if searchFrom(grid, row, col, dir[0], dir[1]) {
						results = append(results, [2]int{row, col})
					}
				}
			}
		}
	}
	return results
}

var vertical = [3][2]int{{-1, 0}, {0, 0}, {1, 0}}   // Vertical "MAS"
var diagonal1 = [3][2]int{{-1, -1}, {0, 0}, {1, 1}} // Diagonal "\" "MAS"
var diagonal2 = [3][2]int{{-1, 1}, {0, 0}, {1, -1}} // Diagonal "/" "MAS"

func isMAS(grid [][]rune, startRow, startCol int, offsets [3][2]int) bool {
	rows := len(grid)
	cols := len(grid[0])

	for i, offset := range offsets {
		r, c := startRow+offset[0], startCol+offset[1]
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != rune("MAS"[i]) {
			return false
		}
	}
	return true
}

// Find all X-MAS patterns in the grid
func findXMAS(grid [][]rune) [][2]int {
	var results [][2]int
	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Check if the center is 'A'
			if grid[row][col] == 'A' {
				// Check vertical and both diagonal MAS sequences
				if isMAS(grid, row, col, vertical) && isMAS(grid, row, col, diagonal1) && isMAS(grid, row, col, diagonal2) {
					results = append(results, [2]int{row, col})
				}
			}
		}
	}
	return results
}

func zad01(input string) {
	input = input[:len(input)-1]
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	rez := findAllOccurrences(grid)
	fmt.Printf("Result 1:\n%d\n", len(rez))

	rez2 := findXMAS(grid)
	fmt.Printf("Result 2:\n%d\n", len(rez2))
}

func readFromFile() string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func read() []string {
	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)

	var input []string
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		input = append(input, line)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done reading")

	return input
}
