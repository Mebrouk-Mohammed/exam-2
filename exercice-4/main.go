package main

import "fmt"

func CountIf(f func(string) bool, tab []string) int {
	countif := 0
	for _, a := range tab {
		if f(a) {
			countif++

		}
	}
	return countif
}

func main() {
	fmt.Println(CountIf("Mohammed "))
	fmt.Println(CountIf("Mebrouk "))
}
