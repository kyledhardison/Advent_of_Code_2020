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

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	check(err)

	temp := strings.Split(string(data), "\n")

	// s := make([]int, len(temp))

	fmt.Println(temp[0])

	test := strings.Split(temp[0], " ")

	fmt.Printf("%q", test)
}
