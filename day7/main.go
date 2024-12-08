package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equ struct {
	rez  int
	nums []int
}

func main() {
	input := read()
	equs := parse(input[:len(input)-1])
	zad1(equs)
}

func zad1(equs []equ) {
	sum := 0
	for _, eq := range equs {
		if eq.IsTrue() {
			sum += eq.rez
		}
		go func() {
			fmt.Printf("Func solved %v\n", eq)
		}()
	}
	fmt.Printf("sum: %v\n", sum)
}

func nextPermutation(arr []int) bool {
	// Start from the right side of the array
	for i := len(arr) - 1; i >= 0; i-- {
		// Try to increment the value
		if arr[i] < 2 { // if it's less than 2, we can increment it
			arr[i]++
			// After incrementing, set all the elements to the right of it to the smallest value (0)
			for j := i + 1; j < len(arr); j++ {
				arr[j] = 0
			}
			return true // valid next permutation found
		}
	}
	// If no more valid permutation is found (all elements are 2), return false
	return false
}

func (this *equ) IsTrue() bool {
	mask := make([]int, len(this.nums)-1)
	for {
		//fmt.Printf("this.nums: %v\n", this.nums)
		//fmt.Printf("mask: %v\n", mask)
		if this.sum(mask) {
			return true
		}
		if !nextPermutation(mask) {
			break
		}
	}
	return false
}

func (this *equ) sum(mask []int) bool {
	ctn := 0
	sum := this.nums[0]
	for _, v := range this.nums[1:] {
		//fmt.Printf("mask[ctn]: %v\n", mask[ctn])
		if mask[ctn] == 0 {
			sum *= v
		} else if mask[ctn] == 1 {
			sum += v
		} else {
			sum, _ = strconv.Atoi(fmt.Sprintf("%d%d", sum, v))
		}
		ctn++
	}
	//fmt.Printf("sum: %v\n", sum)
	//fmt.Printf("this.rez: %v\n", this.rez)
	return this.rez == sum
}

func parse(arr []string) []equ {
	ret := []equ{}
	for _, v := range arr {
		vals := strings.Split(v, ": ")
		rez, _ := strconv.Atoi(vals[0])
		tmp := equ{
			rez:  rez,
			nums: conv(strings.Split(vals[1], " ")),
		}

		ret = append(ret, tmp)
	}
	return ret
}

func conv(arr []string) []int {
	ret := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		ret[i], _ = strconv.Atoi(arr[i])
	}
	return ret
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
