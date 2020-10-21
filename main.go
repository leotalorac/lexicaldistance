package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"

	histogramtool "github.com/aybabtme/uniplot/histogram"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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
func frecuenciescalc(words []string) ([]int, float64) {
	// frecuencies := make(map[int]int)
	var fre []int
	var d int
	sum := 0
	c := 0
	for i := 0; i < len(words); i++ {
		for j := i; j < len(words); j++ {
			d = distance(words[i], words[j])
			sum = sum + d
			c++
			fre = append(fre, d)
		}
	}
	var avg = (float64(sum)) / (float64(c))
	return fre, avg
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
func nehistogram(f []int, filename string) {
	rand.Seed(int64(0))
	v := make(plotter.Values, len(f))
	for i := range f {
		v[i] = float64(f[i])
	}
	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Histogram"
	h, err := plotter.NewHist(v, 10)
	if err != nil {
		panic(err)
	}
	p.Add(h)
	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, filename); err != nil {
		panic(err)
	}
}
func getinfolanguage(name string, file string) {
	datacsharp, err := ioutil.ReadFile("./languages/" + file + ".txt")
	wordssharp := strings.Split(string(datacsharp), " ")
	fmt.Println("---------------------------" + name + "---------------------------")
	f, avg := frecuenciescalc(wordssharp)
	fmt.Println(avg)
	createhist(f)
	nehistogram(f, "./hist/"+file+".png")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
}

func main() {
	getinfolanguage("C#", "csharp")
	getinfolanguage("C++", "cpp")
	getinfolanguage("Java", "java")
}
