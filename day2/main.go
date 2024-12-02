package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := read()
	zad1(input)
}

func zad1(arr []string) {
	safeConut := 0
	for _, report := range arr {
		lvls := strings.Split(report, " ")
		if checkLvls(lvls) {
			safeConut++
		}
	}
	println("Safe reports: ")
	println(safeConut)
}

func checkLvls(arr []string) bool {
	fmt.Print("Report: ")
	for _, v := range arr {
		fmt.Printf("%s ", v)
	}
	fmt.Println()

	last, _ := strconv.Atoi(arr[0])
	second, _ := strconv.Atoi(arr[1])
	asc := last > second
	for i, val := range arr {
		if i == 0 {
			continue
		}
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		if num > last == asc {
			fmt.Printf("Report unsafe becouse of uneaveness \n")
			return false
		}

		diff := abs(num - last)
		if diff > 3 || diff < 1 {
			fmt.Printf("Report unsafe becouse of step %d \n", diff)
			return false
		}
		last = num
	}
	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
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
