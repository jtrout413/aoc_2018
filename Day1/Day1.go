package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input_1_1.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)

	shifts := strings.Split(str, "\n")

	var results []int

	s := 0

	duplicateFound := false

	for {

		if duplicateFound == true {
			break
		}

		for _, e := range shifts {
			n, err := strconv.Atoi(e)
			if err != nil {
				fmt.Println("String convert error, trying to convert:", e)
			}
			s += n

			if contains(results, s) {
				fmt.Println("First duplicate:", s)
				duplicateFound = true
				break
			}
			results = append(results, s)

		}
	}

}

func contains(s []int, n int) bool {

	for _, x := range s {
		if x == n {
			return true
		}
	}
	return false

}
