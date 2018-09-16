package main

func main() {
	sum(1000)
}

func sum(number int) int {
	var result int
	for i := 1; i <= number; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	return result
}
