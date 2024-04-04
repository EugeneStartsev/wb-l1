package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func calcDistance(p1 point, p2 point) float64 {
	xVariance := p1.x - p2.x
	yVariance := p1.y - p2.y

	return math.Sqrt(xVariance*xVariance + yVariance*yVariance)
}

func main() {
	p1 := point{
		15.0,
		12.4,
	}

	p2 := point{
		12.9,
		15.4,
	}

	fmt.Println(calcDistance(p1, p2))
}
