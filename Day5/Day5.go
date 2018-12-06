package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// A - Z : 65-90
// a - z : 97-122

func main() {

	b, err := ioutil.ReadFile("input_5.txt")

	if err != nil {
		fmt.Println(err)
	}
	polymer := string(b)

	fmt.Println("Collapsed Length:", collapsedPolymerLength(polymer))

	removedLengths := make(map[string]int)

	for l := 97; l <= 122; l++ {

		mod := strings.Replace(polymer, string(l), "", -1)
		mod = strings.Replace(mod, strings.ToUpper(string(l)), "", -1)

		length := collapsedPolymerLength(mod)

		removedLengths[string(l)] = length

	}

	shortest := "a"

	for k := range removedLengths {
		if removedLengths[k] < removedLengths[shortest] {
			shortest = k
		}
	}

	fmt.Println("Max collapsed by removing", shortest, ":", removedLengths[shortest])
}

func collapsedPolymerLength(polymer string) int {

	madeCollapse := false

	for {

		for i := 0; i < len(polymer)-1; i++ {

			a := polymer[i]
			b := polymer[i+1]

			// Test if letters are opposites
			if (a >= 97 && b <= 90 && strings.ToUpper(string(a)) == string(b)) || (a <= 90 && b >= 97 && strings.ToLower(string(a)) == string(b)) {
				//fmt.Println("Valid", string(a), string(b))
				polymer = polymer[:i] + polymer[i+2:]
				madeCollapse = true
				//fmt.Println(polymer)
			}
		}

		if !madeCollapse {
			break
		} else {
			madeCollapse = false
		}

	}

	return len(polymer)
}
