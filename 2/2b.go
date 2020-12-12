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

		nums := strings.Split(split[0], "-")
		first, err := strconv.Atoi(nums[0])
		check(err)
		second, err := strconv.Atoi(nums[1])
		check(err)

		char := strings.ReplaceAll(split[1], ":", "")
		password := split[2]

		count := strings.Count(password, char)
		fmt.Println(split)
		fmt.Println(count)
		var firstcheck, secondcheck bool
		if char == string(password[first-1]) {
			firstcheck = true
		} else {
			firstcheck = false
		}

		if char == string(password[second-1]) {
			secondcheck = true
		} else {
			secondcheck = false
		}

		if (firstcheck || secondcheck) && !(firstcheck && secondcheck) {
			valid++
		}

	}
	fmt.Println(valid)
}
