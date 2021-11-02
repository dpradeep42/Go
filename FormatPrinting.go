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
	fmt.Printf("%.3f \n", pi) //Returns three decimal or floating values
	fmt.Printf("%T \n", name) //Returns Data type
	fmt.Printf("%t \n", win)  //Return Boolean
	fmt.Printf("%d \n", x)    //Returns Integer
	fmt.Printf("%b \n", 3)    //Returns Binaries
	fmt.Printf("%c \n", 3)    //Returns ASCII
	fmt.Printf("%x \n", 15)   //Returns Hex Code
	fmt.Printf("%e \n", pi)   //Retuens Scientific notations
}
