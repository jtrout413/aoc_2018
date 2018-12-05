package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {

	b, err := ioutil.ReadFile("input_4.txt")

	if err != nil {
		fmt.Println(err)
	}
	str := string(b)

	records := strings.Split(str, "\n")
	sort.Strings(records)

	guardSleep := make(map[string][]int)

	var cg string
	var sleepStart int

	for _, r := range records {

		segs := strings.Split(r, " ")

		if len(segs) == 6 {
			// New Guard
			cg = segs[3][1:]
		} else {
			// Falls asleep or wakes up

			switch segs[2] {
			case "falls":
				time, _ := strconv.Atoi(strings.TrimRight(segs[1][3:], "]"))
				sleepStart = time
			case "wakes":
				time, _ := strconv.Atoi(strings.TrimRight(segs[1][3:], "]"))

				for i := sleepStart; i < time; i++ {
					guardSleep[cg] = append(guardSleep[cg], i)
				}
			}
		}
	}

	// Find guard with most sleep
	maxGuard := ""
	max := 0
	for guard := range guardSleep {
		if len(guardSleep[guard]) > max {
			maxGuard = guard
			max = len(guardSleep[guard])
		}
	}

	sort.Ints(guardSleep[maxGuard])

	prevMin := -1
	maxMin := -1
	maxMinCount := -1

	indCount := -1

	for _, min := range guardSleep[maxGuard] {

		if min == prevMin {
			indCount++
			if indCount > maxMinCount {
				maxMinCount = indCount
				maxMin = min
			}
		} else {
			indCount = 0
			prevMin = min
		}
	}

	fmt.Println(maxGuard, maxMin)

	fmt.Println("--")

	prevMin = -1
	maxMin = -1
	maxGuard = ""
	maxMinCount = -1
	indCount = 0

	for guard := range guardSleep {

		asleepTimes := guardSleep[guard]
		sort.Ints(asleepTimes)

		for _, t := range asleepTimes {
			if t == prevMin {
				indCount++
				if indCount > maxMinCount {
					maxMin = t
					maxMinCount = indCount
					maxGuard = guard
				}
			} else {
				indCount = 0
				prevMin = t
			}
		}

	}

	fmt.Println(maxGuard, maxMin)

}
