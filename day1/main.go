package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	l1, l2 := read()
	zad1(l1, l2)
	zad2(l1, l2)
}

func zad2(l1 []int, l2 []int) {
	rez := 0
	for _, num := range l1 {
		rez += count(l2, num)
	}
	println("Result 2: ")
	println(rez)
}

func count(list []int, n int) int {
	sum := 0
	for _, num := range list {
		if num == n {
			sum++
		} else if sum != 0 && num != n {
			return sum * n
		}
	}
	return sum * n
}

func zad1(l1 []int, l2 []int) {
	sum := 0
	for i := 0; i < len(l1); i++ {
		sum += abs(l1[i] - l2[i])
	}
	println("Result 1: ")
	println(sum)
}

func abs(this int) int {
	if this < 0 {
		return -this
	}
	return this
}

func read() ([]int, []int) {
	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)

	var l1, l2 []int
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		nums := strings.Split(line, "   ")
		num, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}
		l1 = append(l1, num)
		num, err = strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}
		l2 = append(l2, num)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done reading")

	fmt.Println("Sorting")
	slices.Sort(l1)
	slices.Sort(l2)
	fmt.Println("Done sorting")

	return l1, l2
}
