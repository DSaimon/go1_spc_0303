package main

import "fmt"

func main() {
	//Простейший вывод на косноль. println - это вывод аргумента + '\n'
	fmt.Println("Hello world", "Hello another")
	fmt.Println("Second line")
	//Функция print - простой вывод аргумента
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")
	//Форматированный вывод: Printf - стандартный вывод os.Stdout с флагами форматирования
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 42)
}
