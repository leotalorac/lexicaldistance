package main

import "fmt"

func MinOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
func fillarray(arr []int, val int, n int) []int {
	for i := 0; i < n; i++ {
		arr = append(arr, val)
	}
	return arr
}
func getindexletter(l string, p int) int {
	return int(l[p] - 97)
}
func distance(a string, b string) int {
	var da [26]int
	for i := 0; i < len(da); i++ {
		da[i] = 0
	}
	d := make([][]int, len(a)+2)
	for i := range d {
		d[i] = make([]int, len(b)+2)
	}
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
	fmt.Println(d)
	for i := 1; i < len(a)+1; i++ {
		db := 0
		for j := 1; j < len(b)+1; j++ {
			k := da[getindexletter(b, j)]
			l := db
			if getindexletter(a, i) == getindexletter(b, j) {
				db = j
			} else {
				cost := 1
				d[i+1][j+1] = MinOf(d[i][j]+cost,
					d[i+1][j]+1,
					d[i][j+1]+1,
					d[k][l]+(i-k)+1+(j-l))
			}
		}
		da[getindexletter(a, i)] = i
	}
	fmt.Println(d)
	return d[len(a)][len(b)]
}
func main() {
	fmt.Println(distance("luis", "laiso"))
}
