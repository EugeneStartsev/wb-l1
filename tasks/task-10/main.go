package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	res := make(map[int][]float64)

	for _, val := range temperatures {
		keyRange := int(val/10) * 10

		res[keyRange] = append(res[keyRange], val)
	}

	fmt.Println(res)
}
