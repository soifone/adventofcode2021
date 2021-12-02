package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type direction string

const (
	forward = direction("forward")
	up      = direction("up")
	down    = direction("down")
)

type Command struct {
	D     direction
	Total int
}

func main() {
	b, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	l := buildFromInput(b)
	depth, horizontal := 0, 0
	for _, com := range l {
		switch com.D {
		case forward:
			horizontal += com.Total
		case up:
			depth -= com.Total
		case down:
			depth += com.Total
		}
	}

	fmt.Println("Part 1", depth, horizontal, depth*horizontal)

	depth, horizontal, aim := 0, 0, 0
	for _, com := range l {
		switch com.D {
		case forward:
			horizontal += com.Total
			depth += aim * com.Total
		case up:
			aim -= com.Total
		case down:
			aim += com.Total
		}
	}
	fmt.Println("Part 2", depth, horizontal, depth*horizontal)
}

func buildFromInput(b []byte) []Command {
	ls := strings.Split(string(b), "\n")
	cc := make([]Command, 0)
	for _, l := range ls {
		p := strings.Split(l, " ")
		if len(p) < 2 {
			continue
		}

		i, _ := strconv.Atoi(p[1])
		cc = append(cc, Command{
			D:     direction(p[0]),
			Total: i,
		})
	}

	return cc
}
