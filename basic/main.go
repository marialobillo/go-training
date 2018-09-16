package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo")
	slice := []int{1, 2, 3, 4, 5}
	var resultado = sum(slice)
	fmt.Println(resultado)
}

func sum(sli []int) int {
	var sum int
	for _, item := range sli {
		sum += item
	}
	return sum
}

func isOdd(number int) bool {
	if number%2 == 0 {
		return false
	} else {
		return true
	}
}

func add(args ...int) int {

}
