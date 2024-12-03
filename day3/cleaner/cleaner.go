package cleaner

import (
	"fmt"
	"strings"
)

func CleanInput(input string) []string {
	arr := strings.Split(input, "mul(")
	doCount := true
	var ret []string
	for i := 0; i < len(arr); i++ {
		bef, _, success := strings.Cut(arr[i], ")")
		if !success {
			//fmt.Println(arr[i])
		}
		if success && doCount {
			if isValidMul(bef) {
				ret = append(ret, bef)
			} else {
				println(arr[i])
			}
		}
		cm, err := doesCount(arr[i])
		if err == nil {
			doCount = cm
		}
		//fmt.Printf("%t\n", doCount)
	}
	return ret
}

func doesCount(input string) (bool, error) {
	if input == "" {
		return false, fmt.Errorf("Empty ")
	}
	doCount := strings.Contains(input, "do()")
	dontCount := strings.Contains(input, "don't()")
	if doCount && dontCount {
		panic(2)
	}
	if !doCount && !dontCount {
		//fmt.Printf("no command in (%s)\n", input)
		return false, fmt.Errorf("no command")
	}
	if doCount {
		//fmt.Printf("do command in (%s)\n", input)
		return true, nil
	}

	if dontCount {
		//fmt.Printf("don't command in (%s)\n", input)
		return false, nil
	}
	return true, nil
}

func isValidMul(input string) bool {
	counter, comma := 0, false
	for _, v := range input {
		if _, err := rtoi(v); err == nil {
			counter++
		} else if v == ',' {
			comma = counter > 0 && counter < 4
			counter = 0
		} else {
			return false
		}

	}

	return comma && counter > 0 && counter < 4
}

func rtoi(r rune) (int, error) {
	br := int(r)
	if br < int('0') || br > int('9') {
		return 0, fmt.Errorf("Error parsing int %s", r)
	}
	return br, nil
}
