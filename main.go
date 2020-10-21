package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}
func getindexletter(l string, p int) int {
	return int(l[p] - 97)
}
func distance(a string, b string) int {
	// create the matrix
	d := make([][]int, len(a)+2)
	for i := range d {
		d[i] = make([]int, len(b)+2)
	}
	// fill the matrix
	maxdist := len(a) + len(b)
	d[0][0] = maxdist
	for i := 1; i < len(a)+2; i++ {
		d[i][0] = maxdist
		d[i][1] = i - 1
	}
	for j := 1; j < len(b)+2; j++ {
		d[0][j] = maxdist
		d[1][j] = j - 1
	}
	// lastrow
	lastrow := make(map[string]int)
	// fmt.Println(d)
	for i := 1; i < len(a)+1; i++ {
		cha := string([]rune(a)[i-1])
		lastmatchcol := 0
		for j := 1; j < len(b)+1; j++ {
			chb := string([]rune(b)[j-1])
			var lastmatchingrow int = 0
			if val, ok := lastrow[chb]; ok {
				lastmatchingrow = val
			}
			cost := 1
			if cha == chb {
				cost = 0
			}
			d[i+1][j+1] = MinOf(d[i][j]+cost, // Substitution
				d[i+1][j]+1, // Addition
				d[i][j+1]+1, // Deletion
				d[lastmatchingrow][lastmatchcol]+(i-lastmatchingrow-1)+1+(j-lastmatchcol-1))

			if cost == 0 {
				lastmatchcol = j
			}
		}
		lastrow[cha] = i
	}
	// fmt.Println(d)
	// fmt.Println(lastrow)
	return d[len(a)][len(b)]
}
func main() {
	// fmt.Println(distance("luis", "luiso"))
	data, err := ioutil.ReadFile("./languages/csharp.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	words := strings.Split(string(data), " ")
	fmt.Println(words)
}
