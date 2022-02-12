package main

import "fmt"

func main() {
	count := 0

	// Below init and post statements are optional

	for i := 0; i <= 100; i++ {
		count += i
	}
	fmt.Println(count)
}
