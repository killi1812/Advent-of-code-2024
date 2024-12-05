package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := read()
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
