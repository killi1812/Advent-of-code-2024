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
	zad2(input)
}
func zad2(arr []string) {
	safeConut := 0
	for _, report := range arr {
		lvls := strings.Split(report, " ")
		success, index := checkLvls(lvls)
		if success {
			safeConut++
			continue
		}
		if tryCleen(lvls, index) {
			safeConut++
		}
	}
	println("Safe reports 2: ")
	println(safeConut)
}

func remove(slice []string, s int) []string {
	sliceCopy := make([]string, len(slice))
	copy(sliceCopy, slice)
	if s == 0 {
		return append(sliceCopy[s+1:])
	}
	return append(sliceCopy[:s], sliceCopy[s+1:]...)
}

func tryCleen(arr []string, index int) bool {
	success := false
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Removing element %d\n", i)
		redact := remove(arr, i)
		printReport(redact)
		s, _ := checkLvls(redact)
		success = success || s
	}

	if !success {
		printReport(arr)
		fmt.Println("Failed cleaning")
	}

	return success
}

func zad1(arr []string) {
	safeConut := 0
	for _, report := range arr {
		lvls := strings.Split(report, " ")
		success, _ := checkLvls(lvls)
		if success {
			safeConut++
		}
	}
	println("Safe reports: ")
	println(safeConut)
}

func printReport(arr []string) {
	fmt.Print("Report: ")
	for _, v := range arr {
		fmt.Printf("%s ", v)
	}
	fmt.Println()

}

func checkLvls(arr []string) (bool, int) {

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
			//printReport(arr)
			//fmt.Printf("Report unsafe becouse of uneaveness \n")
			return false, i
		}

		diff := abs(num - last)
		if diff > 3 || diff < 1 {
			//printReport(arr)
			//fmt.Printf("Report unsafe becouse of step %d \n", diff)
			return false, i
		}
		last = num
	}
	return true, -1
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
