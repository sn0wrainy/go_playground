package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	var number int
	fmt.Print("How many numbers do you want?")
	fmt.Scanf("%v", &number)
	for i := 0; i < number; i++ {
		fmt.Println(f())
	}
}
