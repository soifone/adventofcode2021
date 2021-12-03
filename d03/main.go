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

	ls := strings.Split(strings.TrimSuffix(string(b), "\n"), "\n")
	m, l := commonsDec(ls)
	fmt.Println("Part 1", m, l, m*l)

	o, s := oxygenAndScrubber(ls)
	fmt.Println("Part 2", o, s, o*s)
}

func binarySums(ls []string) []int {
	ts := make([]int, len(ls[0]))
	for i := 0; i < len(ls); i++ {
		cs := strings.Split(ls[i], "")
		for p, c := range cs {
			in, _ := strconv.Atoi(c)
			ts[p] += in
		}
	}

	return ts
}

func oxygenAndScrubber(ls []string) (int64, int64) {
	ts := binarySums(ls)
	oxygenMask, scrubberMask := generateBinaryStr(ts, len(ls))
	oxygen := reduceOxygen(ls, oxygenMask, 0)
	coscrubber := reduceScrubber(ls, scrubberMask, 0)

	o, _ := strconv.ParseInt(oxygen, 2, 64)
	c, _ := strconv.ParseInt(coscrubber, 2, 64)

	return o, c
}

func reduce(ls []string, mask string, lookup int) []string {
	find := mask[lookup : lookup+1]
	ll := make([]string, 0)

	for i := 0; i < len(ls); i++ {
		if find != ls[i][lookup:lookup+1] {
			continue
		}

		ll = append(ll, ls[i])
	}

	return ll
}

func reduceScrubber(ls []string, mask string, lookup int) string {
	ll := reduce(ls, mask, lookup)
	if len(ll) == 1 {
		return ll[0]
	}

	ts := binarySums(ls)
	_, mask = generateBinaryStr(ts, len(ls))
	return reduceScrubber(ll, mask, lookup+1)
}

func reduceOxygen(ls []string, mask string, lookup int) string {
	ll := reduce(ls, mask, lookup)
	if len(ll) == 1 {
		return ll[0]
	}

	ts := binarySums(ls)
	mask, _ = generateBinaryStr(ts, len(ls))
	return reduceOxygen(ll, mask, lookup+1)
}

func generateBinaryStr(ts []int, ll int) (string, string) {
	l := ""
	m := ""
	for _, t := range ts {
		if ll-t == ll/2 {
			m += "1"
			l += "0"
			continue
		}

		if t < ll-t {
			m += "0"
			l += "1"
		} else {
			m += "1"
			l += "0"
		}
	}

	return m, l
}

func commonsDec(ls []string) (int64, int64) {
	ts := binarySums(ls)
	m, l := generateBinaryStr(ts, len(ls))
	md, _ := strconv.ParseInt(m, 2, 64)
	ld, _ := strconv.ParseInt(l, 2, 64)
	return md, ld
}
