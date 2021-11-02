package main

import "fmt"

func main() {
	var name string = "Pradeep Kumar"
	const pi float64 = 3.1415139
	win := true
	x := 5

	fmt.Println(len(name))
	fmt.Println(name + " Dasari")
	fmt.Println(pi)
	fmt.Printf("%f \n", pi)
	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", name) //Returns Data type
	fmt.Printf("%t \n", win)  //Return Boolean
	fmt.Printf("%d \n", x)
	fmt.Printf("%b \n", 3)
	fmt.Printf("%c \n", 3)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%e \n", pi)
}
