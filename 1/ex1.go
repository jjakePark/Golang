package main

import "fmt"

func main() {
	fmt.Println("Hello, world")

	for i := 0; i < 100; i++ {
		if i %2 == 0 {
			fmt.Println(i)
		}
	}
}
