package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	histogramtool "github.com/aybabtme/uniplot/histogram"
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
func frecuenciescalc(words []string) []int {
	// frecuencies := make(map[int]int)
	var fre []int
	var d int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			d = distance(words[i], words[j])
			// if val, ok := frecuencies[d]; ok {
			// 	frecuencies[d] = val + 1
			// } else {
			// 	frecuencies[d] = 1
			// }
			fre = append(fre, d)
		}
	}
	return fre
}
func createhist(f []int) {
	var new_bound []float64
	for i := 1; i < len(f); i++ {
		new_bound = append(new_bound, float64(f[i]))
	}
	hist := histogramtool.Hist(10, new_bound)
	if err := histogramtool.Fprint(os.Stdout, hist, histogramtool.Linear(5)); err != nil {
		panic(err)
	}
}

func main() {
	// fmt.Println(distance("luis", "luiso"))
	datacsharp, err := ioutil.ReadFile("./languages/csharp.txt")
	wordssharp := strings.Split(string(datacsharp), " ")
	fmt.Println("---------------------------C#---------------------------")
	f := frecuenciescalc(wordssharp)
	createhist(f)
	datacpp, err := ioutil.ReadFile("./languages/cpp.txt")
	wordscpp := strings.Split(string(datacpp), " ")
	fmt.Println("---------------------------C++---------------------------")
	f = frecuenciescalc(wordscpp)
	createhist(f)
	datajava, err := ioutil.ReadFile("./languages/java.txt")
	wordsjava := strings.Split(string(datajava), " ")
	fmt.Println("---------------------------Java---------------------------")
	f = frecuenciescalc(wordsjava)
	createhist(f)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

}
