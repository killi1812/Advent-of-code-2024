package main

import (
	"bufio"
	"day3/cleaner"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFromFile()
	vals := cleaner.CleanInput(input)
	fmt.Printf("\nFound %d multiplications\n\n", len(vals))
	sum := 0
	for _, v := range vals {
		//fmt.Println(v)
		//fmt.Printf("comand %s, ", v)
		sum += procesMul(v)
	}
	fmt.Printf("Result: \n%d\n", sum)
}

func procesMul(tmp string) int {
	nums := strings.Split(tmp, ",")
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])

	//fmt.Printf("Clean %s \n", tmp)
	return x * y
}

func rtoi(r rune) (int, error) {
	br := int(r)
	if br < int('0') || br > int('9') {
		return 0, fmt.Errorf("Error parsing int %s", r)
	}
	return br, nil
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
