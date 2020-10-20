package main

import (
	"fmt"
)

func getindexletter(l string, p int) int {
	return int(l[p] - 97)
}

func main() {
	fmt.Println(getindexletter("u", 0))
}
