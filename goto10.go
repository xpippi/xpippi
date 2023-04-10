package main

import "fmt"

func main() {

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println("goto10")

	for n := 1; n <= 5; n += 2 {
		fmt.Println(n)
	}
}
