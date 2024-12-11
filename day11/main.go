package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type key struct {
	stone string
	left  int
}

func main() {
	input := read()
	sum := 0
	fmt.Printf("sum: %v\n", sum)
}

func t2(arr []string) {
	cache := make(map[key]int, 0)
	for _, v := range arr {
		transform([]string{v})
	}
}

func transform(arr []string) []string {
	ret := make([]string, len(arr))
	extra := 0
	for i, stone := range arr {
		if stone == "0" {
			ret[i+extra] = "1"
		} else if size := len(stone); size%2 == 0 {
			half := size / 2
			br, _ := strconv.Atoi(stone[:half])
			ret[i+extra] = fmt.Sprint(br)
			ret = append(ret, ".")
			extra++
			br, _ = strconv.Atoi(stone[half:])
			ret[i+extra] = fmt.Sprint(br)
		} else {
			br, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			ret[i+extra] = fmt.Sprint(br * 2024)
		}

	}
	transform(ret)
	return ret
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data[:len(data)-1]), " ")
}
