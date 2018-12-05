package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Coord Silence
type Coord struct {
	x int
	y int
}

func newCoord(coords string) Coord {

	var c Coord

	split := strings.Split(coords, ",")

	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	c.x = x
	c.y = y

	return c

}

// Area Silence
type Area struct {
	width  int
	height int
}

func newArea(area string) Area {

	var a Area

	split := strings.Split(area, "x")

	w, _ := strconv.Atoi(split[0])
	h, _ := strconv.Atoi(split[1])

	a.width = w
	a.height = h

	return a

}

func main() {

	b, err := ioutil.ReadFile("input_3.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)

	claims := strings.Split(str, "\n")

	var fabric [1000][1000]int

	for r := 0; r < 29; r++ {
		for c := 0; c < 29; c++ {
			fabric[r][c] = 0
		}
	}

	for _, claim := range claims {

		segs := strings.Split(claim, " ")

		origin := newCoord(strings.TrimRight(segs[2], ":"))
		area := newArea(segs[3])

		for y := origin.y; y < origin.y+area.height; y++ {
			for x := origin.x; x < origin.x+area.width; x++ {
				fabric[y][x]++
			}
		}
	}

	count := 0

	for r := 0; r < 1000; r++ {
		for c := 0; c < 1000; c++ {
			if fabric[r][c] >= 22 {
				count++
			}
		}
	}

	for _, claim := range claims {
		segs := strings.Split(claim, " ")

		origin := newCoord(strings.TrimRight(segs[2], ":"))
		area := newArea(segs[3])

		ok := true

		for y := origin.y; y < origin.y+area.height; y++ {
			for x := origin.x; x < origin.x+area.width; x++ {
				if fabric[y][x] > 1 {
					ok = false
				}
			}
		}

		if ok {
			fmt.Println(claim)
		}
	}

	fmt.Println(count)

}
