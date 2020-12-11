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

	s := make([]int, len(temp))

	for i := 0; i < len(temp); i++ {
		s[i], err = strconv.Atoi(temp[i])
		check(err)
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			for k := 0; k < len(s); k++ {
				if s[i]+s[j]+s[k] == 2020 {
					fmt.Println(s[i])
					fmt.Println(s[j])
					fmt.Println(s[k])
					fmt.Println(s[i] * s[j] * s[k])
					return
				}
			}
		}
	}
}
