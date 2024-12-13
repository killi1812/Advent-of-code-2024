package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type btn struct {
	X float64
	Y float64
}

func (this btn) add(loc btn) btn {
	return btn{
		X: this.X + loc.X,
		Y: this.Y + loc.Y,
	}
}

type eq struct {
	A     btn
	B     btn
	Prize btn
}

func (this *eq) solve() int {

	this.Prize.X += 10000000000000
	this.Prize.Y += 10000000000000
	b := (this.Prize.Y - this.A.Y*(this.Prize.X/this.A.X)) / (-this.A.Y*(this.B.X/this.A.X) + this.B.Y)
	a := (this.Prize.X - this.B.X*b) / this.A.X

	a += 0.0000001
	b += 0.0000001
	a = math.Round(a*1000) / 1000
	b = math.Round(b*1000) / 1000

	if a < 0 || b < 0 || float64(int(a)) != math.Round(a*1000)/1000 || float64(int(b)) != math.Round(b*1000)/1000 {
		fmt.Printf("eq: %v\n", *this)
		fmt.Printf("Solution a: %f, b: %v\n", a, b)
		return 0
	}

	return int(3*a + b)
}

func main() {
	input := read()
	eqs := parse(input)
	sum := 0
	for _, v := range eqs {
		sum += v.solve()
	}
	fmt.Printf("sum: %v\n", sum)
	br := 15.000000000000004
	fmt.Println(float64(int(br)) == math.Round(br*1000)/1000)
}

func parse(arr []string) []eq {
	ret := []eq{}
	current := eq{}
	for _, v := range arr {
		if v == "" {
			ret = append(ret, current)
			current = eq{}
			continue
		}

		rest := strings.Split(v, ": ")
		if strings.Count(v, "A") == 1 {
			current.A = parseBtn(rest[1])
		} else if strings.Count(v, "B") == 2 {
			current.B = parseBtn(rest[1])
		} else {
			current.Prize = parseBtn(rest[1])
		}
	}

	return ret
}

func parseBtn(input string) btn {
	tmp := strings.Split(input, ", ")
	return btn{X: toInt(tmp[0][2:]), Y: toInt(tmp[1][2:])}
}

func toInt(str string) float64 {
	br, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return float64(br)
}

func read() []string {
	name := "input"
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
