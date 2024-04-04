package main

import (
	"fmt"
	"strings"
)

var justString string

func main() {
	someFunc()
}

func createHugeString(size int) string {
	var b strings.Builder

	for i := 0; i < size; i++ {
		b.WriteString("気")
	}

	return b.String()
}

func someFunc() {
	v := createHugeString(1 << 10)

	// неправильно срежется слайс, так как руна может занимать от 1 до 4 байт
	justString = v[:100]

	fmt.Println(justString)

	res := []rune(v)

	justString = string(res[:100])

	fmt.Println(justString)
}
