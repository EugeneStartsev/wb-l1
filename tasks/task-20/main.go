package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(flipString("snow dog cat"))
}

func flipString(s string) string {
	res := strings.Split(s, " ")

	var b strings.Builder

	b.WriteString(res[len(res)-1])

	for i := len(res) - 2; i >= 0; i-- {
		b.WriteString(" ")
		b.WriteString(res[i])
	}

	return b.String()
}
