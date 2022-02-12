package main

import "fmt"

func qualy(x float64) string {
	if x < 6 {
		return "Don't pass"
	}
	return "Pass"
}

func main() {
	var number float64 = 4.5
	fmt.Println(qualy(number))
}
