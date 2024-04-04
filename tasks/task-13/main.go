package main

import "fmt"

func changeArithmeticSum(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b

	return a, b
}

func changeBoolean(a, b int) (int, int) {
	a = a ^ b
	b = b ^ a
	a = a ^ b

	return a, b
}

func changeArithmeticMultiplication(a, b int) (int, int) {
	a = a * b
	b = a / b
	a = a / b

	return a, b
}

func changeByGo(a, b int) (int, int) {
	a, b = b, a

	return a, b
}

func main() {
	a := 64
	b := 128

	fmt.Println(a, b)

	fmt.Println(changeArithmeticSum(a, b))
	fmt.Println(changeBoolean(a, b))
	fmt.Println(changeArithmeticMultiplication(a, b))
	fmt.Println(changeByGo(a, b))
}
