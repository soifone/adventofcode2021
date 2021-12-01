package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	ls := strings.Split(string(b), "\n")
	li := convertStrListToInt(ls)
	fmt.Println("Input 1 increase", countOnIncrease(li))
	fmt.Println("Input 2 increase", countOnIncrease(slidingWindow(li)))
}

func convertStrListToInt(ls []string) []int {
	il := make([]int, 0)

	for _, l := range ls {
		i, _ := strconv.Atoi(l)
		il = append(il, i)
	}

	return il
}

func slidingWindow(ls []int) []int {
	s := 0
	nls := make([]int, 0)
	for i := 0; i < len(ls); i++ {
		if s+2 >= len(ls) {
			break
		}

		nls = append(nls, ls[s]+ls[s+1]+ls[s+2])
		s++
	}

	return nls
}

func countOnIncrease(ls []int) int {
	last := 0
	i := 0
	for _, l := range ls {
		if l > last && last != 0 {
			i++
		}
		last = l
	}

	return i
}
