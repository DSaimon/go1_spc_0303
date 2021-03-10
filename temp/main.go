package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var (
		globalWidth  = 35
		globalLength = 45
	)
	var age int
	age = 33
	var uid int = 123123
	var infer = "Bob"
	var length, width int = 30, 40
	count := 0
	a, b := 122.5, 122.6

	fmt.Printf("My age is %d\n", age)
	fmt.Printf("And my UID is %d\n", uid)
	fmt.Printf("Infer is %s\n", infer)
	fmt.Printf("Perimeter is %d\n", (length+width)*2)
	fmt.Printf("Global Perimeter is %d\n", (globalLength+globalWidth)*2)
	fmt.Printf("Count now is %d\n", count)
	count++
	count = 200
	fmt.Printf("Count after increment %d\n", count)
	fmt.Printf("Minimum value is %.3f\n", math.Min(a, b))
	fmt.Println("Hello world!")

	fmt.Print("Hey, this is ust a Print()")
	fmt.Print("Hey, this is another Print()\n")

	var (
		figureName string
		widthScan  int
		lengthScan int
	)

	fmt.Scan(&widthScan, &lengthScan, &figureName)

	fmt.Printf("Scanning figure is %s\n", figureName)
	fmt.Printf("Now we have width:%d and length:%d\n", widthScan, lengthScan)

	var (
		circleRadius int
		circleColor  string
	)

	fmt.Fscan(os.Stdin, &circleRadius, &circleColor)
	fmt.Printf("Figure color: %s\nRadius : %d\nArea : %.2f\n", circleColor, circleRadius, math.Pi*float32(circleRadius)*float32(circleRadius))
}
