package main

import "fmt"

func main() {
	fmt.Println(reverseByDefer("Hello, 気 world! 気"))
	fmt.Println(reverse("Hello, 気 world! 気"))
}

func reverseByDefer(s string) (ret string) {
	for _, v := range s {
		defer func(r rune) { ret += string(r) }(v)
	}
	return
}

func reverse(s string) string {
	res := []rune(s)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
