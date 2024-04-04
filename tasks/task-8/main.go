package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strconv"
)

func setBitToZeroOrOne(bits string, i int, changer int) {
	res := make([]int, 0, len(bits))

	for _, v := range bits {
		num, _ := strconv.Atoi(string(v))
		res = append(res, num)
	}

	res[i] = changer

	fmt.Println(res)
}

func setBitByBooleanOperation(number int64, i int, changer int) {
	fmt.Println(number, i)

	if changer == 1 {
		number |= 1 << i
	} else {
		number &^= 1 << i
	}

	fmt.Println(number)
}

func main() {
	randNumber := int64(rand.IntN(math.MaxInt64))
	bits := strconv.FormatInt(randNumber, 2)

	i := rand.IntN(len(bits))
	changer := rand.IntN(2)

	setBitToZeroOrOne(bits, i, changer)
	setBitByBooleanOperation(randNumber, i, changer)
}
