package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func treeTest(xRate, yRate int, field []string) {
	x, y, trees := 0, 0, 0
	for true {
		if y >= len(field) {
			break
		}

		if string(field[y][x]) == "#" {
			trees++
		}

		y += yRate
		x += xRate
		if x > 30 {
			x = x - 31
		}
		// fmt.Println(y, x)
	}
	fmt.Println(trees)
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	check(err)

	field := strings.Split(string(data), "\n")

	treeTest(1, 1, field)
	treeTest(3, 1, field)
	treeTest(5, 1, field)
	treeTest(7, 1, field)
	treeTest(1, 2, field)

}
