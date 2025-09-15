package main

import (
	"fmt"
)

func AppendRange(min, max int) []int {
	var app []int
	for i := min; i < max; i++ {
		app = append(app, i)
	}
	return app
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
