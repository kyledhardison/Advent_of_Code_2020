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

	db := strings.Split(string(data), "\n\n")
	fmt.Println(string(db[0]))
	fmt.Println(strings.Contains(string(db[0]), "eyr:"))

}
