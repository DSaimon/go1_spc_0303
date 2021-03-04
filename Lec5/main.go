package main

import (
	"fmt"
	"strings"
)

func main() {
	//Классический условный оператор
	// 	if condition {
	// 		//body
	// 	}
	// }
	//	Условный оператор с блоком else
	// if condition {

	// } else {

	// }
	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("The number", value, "is even")
	} else {
		fmt.Println("The number", value, "is odd")
	}
	// if condition1 {

	// } else if condition2 {

	// } else if ... {

	// } else {

	// }

	var color string
	fmt.Scan(&color)
	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}
}
