package main

import (
	"fmt"
	"math/big"
)

func main() {
	first, _ := big.NewInt(0).SetString("1234675347853489057893478", 10)
	second, _ := big.NewInt(0).SetString("3674589345789345789345376", 10)

	res := big.NewInt(0)

	fmt.Println(res.Add(first, second))
	fmt.Println(res.Div(second, first))
	fmt.Println(res.Sub(second, first))
	fmt.Println(res.Mul(first, second))
}
