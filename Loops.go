package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	for k := 1; k <= 5; k++ {
		for l := 1; l <= k; l++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
