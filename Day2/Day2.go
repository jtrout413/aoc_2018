package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Running Day2...")
	//part1()
	//part2()
}

func part1() {

	b, err := ioutil.ReadFile("input_2_1.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)

	ids := strings.Split(str, "\n")

	doubleCount := 0
	tripleCount := 0

	for _, id := range ids {

		letterCounts := make(map[string]int)

		for _, letter := range id {

			if c, ok := letterCounts[string(letter)]; ok {
				letterCounts[string(letter)] = c + 1
			} else {
				letterCounts[string(letter)] = 1
			}

		}

		hasDouble := false
		hasTriple := false

		for _, v := range letterCounts {
			if v == 2 {
				hasDouble = true
			} else if v == 3 {
				hasTriple = true
			}

		}

		if hasDouble {
			doubleCount++
		}

		if hasTriple {
			tripleCount++
		}

	}

	checkSum := doubleCount * tripleCount
	fmt.Println(checkSum)

}

func part2() {

	b, err := ioutil.ReadFile("input_2_1.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)

	ids := strings.Split(str, "\n")

	for i, id1 := range ids {

		for j := 0; j < len(ids); j++ {

			if j != i {

				id2 := ids[j]

				var letters1 []string
				for _, l := range id1 {
					letters1 = append(letters1, string(l))
				}

				var letters2 []string
				for _, l := range ids[j] {
					letters2 = append(letters2, string(l))
				}

				count := 0

				var similarLetters []string

				for k, l := range id2 {
					if letters1[k] == string(l) {
						similarLetters = append(similarLetters, string(l))
						count++
					}
				}

				if count == len(letters1)-1 {
					fmt.Println(letters1, letters2)
					fmt.Println(strings.Join(similarLetters, ""))
				}

			}

		}

	}

}
