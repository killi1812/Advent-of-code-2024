package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := read()
	input = input[:len(input)-1]
	input = transform(input)
	fmt.Printf("input: %v\n", input)

	comp := move(input)
	fmt.Printf("input: %v\n", comp)
	rez := checkSum(comp)
	fmt.Printf("rez: %v\n", rez)
}

func checkSum(input string) int {
	sum := 0
	for i, char := range input {
		br, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		fmt.Printf("%d * %d\n", i, br)
		sum += i * br
	}
	return sum
}

const empty = '.'

func move(input string) string {
	str := strings.Builder{}
	j := len(input) - 1
	for i, char := range input {
		if char == empty {
			for input[j] == empty {
				j--
			}
			str.WriteByte(input[j])
			j--
		} else {
			str.WriteRune(char)
		}
		if j <= i {
			return str.String()
		}
	}

	return ""
}

func transform(input string) string {
	id := 0
	inc := false
	str := strings.Builder{}
	for i, char := range input {
		br, _ := strconv.Atoi(string(char))
		for j := 0; j < br; j++ {
			if i%2 == 1 {
				str.WriteRune(empty)
			} else {
				inc = true
				str.WriteString(fmt.Sprint(id))
			}
		}
		if inc {
			id++
			inc = false
		}

	}
	return str.String()
}

func read() string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(data)
}
