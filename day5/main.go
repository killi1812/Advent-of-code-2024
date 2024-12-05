package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	X int
	Y int
}

func main() {
	input := read()
	rules, list := parse(input)
	//printInput(&rules, list)
	validUpdates, invalidUpdates := filterLists(list, rules)
	printInput(nil, invalidUpdates)
	rez1 := sumMiddle(validUpdates)
	fmt.Printf("rez1: \n%v\n", rez1)
	rez2 := sumMiddle(order(invalidUpdates, rules))
	fmt.Printf("rez2: \n%v\n", rez2)
}

func order(updates [][]int, rules []Rule) [][]int {
	valid := [][]int{}
	for _, v := range updates {
		tmp := make([]int, len(v))
		copy(tmp, v)
		sort(v, rules)
		fmt.Printf("v: %v\n", tmp)
		if isValid(tmp, rules) {
			valid = append(valid, tmp)
		}
	}
	return valid
}

func sort(arr []int, rules []Rule) []int {

	return arr
}

func same(l1 []int, l2 []int) bool {
	//fmt.Printf("comparing l1: %v, l2: %v\n", l1, l2)
	if len(l1) != len(l2) {
		return false
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}

	return true
}

func sumMiddle(updates [][]int) int {
	sum := 0
	for _, v := range updates {
		sum += v[len(v)/2]
	}
	return sum
}

func filterLists(list [][]int, rules []Rule) ([][]int, [][]int) {
	valid := [][]int{}
	invalid := [][]int{}
	for _, v := range list {
		if isValid(v, rules) {
			valid = append(valid, v)
		} else {
			invalid = append(invalid, v)
		}
	}

	return valid, invalid
}

func isValid(list []int, rules []Rule) bool {
	for i, v := range list {
		relevantRules := filterRules(v, rules)
		for j := 0; j < i; j++ {
			if failsRule(list[j], relevantRules) {
				return false
			}
		}
	}
	return true
}

func failsRule(br int, rules []Rule) bool {
	for _, v := range rules {
		if br == v.Y {
			return true
		}
	}
	return false
}

func filterRules(br int, rules []Rule) []Rule {
	validRules := []Rule{}
	for _, v := range rules {
		if v.X == br {
			validRules = append(validRules, v)
		}
	}
	return validRules
}

func printInput(rules *[]Rule, list [][]int) {
	if rules != nil {
		for _, v := range *rules {
			fmt.Printf("X: %d, Y: %d\n", v.X, v.Y)
		}
	}

	for _, v := range list {
		for _, br := range v {
			fmt.Print(br, ",")
		}
		fmt.Println()

	}

}

func parse(input string) ([]Rule, [][]int) {
	lines := strings.Split(input, "\n")
	var rule []Rule
	var prt [][]int
	sencond := false

	for _, v := range lines {
		if v == "" {
			sencond = true
			continue
		}

		if !sencond {
			l := strings.Split(v, "|")
			rule = append(rule, Rule{X: atoi(l[0]), Y: atoi(l[1])})
		} else {
			row := strings.Split(v, ",")
			tmp := []int{}
			for _, br := range row {
				tmp = append(tmp, atoi(br))
			}
			prt = append(prt, tmp)
		}

	}
	return rule, prt
}

func atoi(a string) int {
	br, err := strconv.Atoi(a)
	if err != nil {
		fmt.Printf("Error parsing %s, %s", a, err)
	}
	return br
}

func read() string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(data)
}
