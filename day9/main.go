package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	input := read()
	input = input[:len(input)-1]
	trs := transform(input)
	comp := move(trs)
	rez := checkSum(comp)
	fmt.Printf("rez: %v\n", rez)
	comp2 := moveWhole(trs)
	rez2 := checkSum(comp2)
	fmt.Printf("rez: %v\n", rez2)

}

func backFile(arr []string, size int) []string {
	i := len(arr) - 1
	last := arr[i]
	beginFile := 0
	for ; i >= 0; i-- {
		if arr[i] == empty {
			continue
		} else if beginFile == 0 {
			beginFile = i
			last = arr[i]
		}

		fmt.Printf("file size: %d, required: %d\n", beginFile-i, size)
		if arr[i] != last && beginFile-i <= size {
			fmt.Printf("arr: %v\n", arr[i+1:beginFile+1])
			return arr[i+1 : beginFile+1]
		}
		last = arr[i]
	}
	return []string{}
}

func findLastConsecutiveSame(arr []string, n int) (int, int) {
	if len(arr) == 0 || n <= 0 {
		return -1, -1
	}

	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] == arr[i-1] {
			startIndex := i - 1
			for startIndex > 0 && arr[startIndex] == arr[startIndex-1] {
				startIndex--
			}
			if i-startIndex+1 <= n {
				return startIndex, i
			}
		}
	}

	return -1, -1
}

func moveWhole(input []string) []string {
	fmt.Println("Zad02")
	arr := []string{}
	j := 0
	for i, char := range input {
		if char == empty {
			j = i
			for input[j] == empty {
				j++
			}

			fmt.Printf("input: %v\n", input)
			b, e := findLastConsecutiveSame(input, j-i)
			fmt.Printf("b: %v,e %d\n", b, e)
			if b == -1 || e == -1 {
				fmt.Printf("Continuing indexes bad\n")
				continue
			}
			back := input[b : e+1]
			fmt.Printf("back: %v\n", back)
			input = slices.Replace(input, i, i+(e-b)+1, back...)
			back = slices.Delete(back, 0, len(back))
			//arr = append(arr, input[j])
			j--
		} else {
			arr = append(arr, char)
		}
		//fmt.Printf("arr: %v\n", arr)
	}

	return arr
}

func checkSum(input []string) int64 {
	var sum int64 = 0
	for i, char := range input {
		br, err := strconv.Atoi(char)
		//fmt.Printf("char: %v\n", char)
		if err != nil {
			continue
		}
		//fmt.Printf("%d * %d\n", i, br)
		sum += int64(i) * int64(br)
	}
	return sum
}

const empty = "."

func move(input []string) []string {
	arr := []string{}
	j := len(input) - 1
	for i, char := range input {
		if char == empty {
			for input[j] == empty {
				j--

				if j <= i {
					return arr
				}
			}
			arr = append(arr, input[j])
			j--
		} else {
			arr = append(arr, char)
		}
		//fmt.Printf("arr: %v\n", arr)
		if j <= i {
			return arr
		}
	}

	return arr
}

func transform(input string) []string {
	var id int = 0
	inc := false
	ret := []string{}
	for i, char := range input {
		br, _ := strconv.Atoi(string(char))
		for j := 0; j < br; j++ {
			if i%2 == 1 {
				ret = append(ret, empty)
			} else {
				inc = true
				ret = append(ret, fmt.Sprint(id))
			}
		}
		if inc {
			id++
			inc = false
		}

	}
	fmt.Printf("id: %v\n", id)
	return ret
}

func read() string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(data)
}
