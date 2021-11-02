package main

import "fmt"

func main() {
	var EvenNum [5]int

	EvenNum[0] = 0
	EvenNum[1] = 2
	EvenNum[2] = 4
	EvenNum[3] = 6
	EvenNum[4] = 8

	fmt.Println(EvenNum[2])

	OddNum := [5]int{1, 3, 5, 7, 9}
	fmt.Println(OddNum)

	for _, value := range OddNum {
		fmt.Println(value)
	}

	for i, value := range OddNum {
		fmt.Println(value, i)
	}

	//SLicing

	numSlice := []int{5, 4, 3, 2, 1}

	sliced := numSlice[3:5]
	fmt.Println(sliced)

	slice2 := make([]int, 5, 10)

	copy(slice2, numSlice)

	fmt.Println(slice2)
	fmt.Println(numSlice)

	slice3 := append(numSlice, 3, 0, -1)
	fmt.Println(slice3)
}
