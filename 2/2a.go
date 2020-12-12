package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	check(err)

	temp := strings.Split(string(data), "\n")

	valid := 0
	for i := 0; i < len(temp); i++ {
		split := strings.Split(temp[i], " ")

		minmax := strings.Split(split[0], "-")
		min, err := strconv.Atoi(minmax[0])
		check(err)
		max, err := strconv.Atoi(minmax[1])
		check(err)

		char := strings.ReplaceAll(split[1], ":", "")
		password := split[2]

		count := strings.Count(password, char)
		fmt.Println(split)
		fmt.Println(count)
		if count >= min && count <= max {
			valid++
		}
	}
	fmt.Println(valid)
}
