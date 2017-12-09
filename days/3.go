package days

import (
	"fmt"
)

type gridItem struct {
	x     int
	y     int
	value int
}

var thirdFirstInput = 325489
var thirdSecondInput = 325489

func ThirdFirst() {

	coordx := 0
	coordy := 0

	currentMove := "right"
	circle := 1
	i := 1
	for {
		switch currentMove {
		case "right":
			coordx++
			if coordx == circle {
				currentMove = "up"
			}
		case "up":
			coordy++
			if coordy == circle {
				currentMove = "left"
			}
		case "left":
			coordx--
			if coordx == circle*-1 {
				currentMove = "down"
			}
		case "down":
			coordy--
			if coordy == circle*-1 {
				currentMove = "right"
				circle++
			}
		}
		i++
		if i == thirdFirstInput {
			break
		}
	}

	if coordx < 0 {
		coordx = coordx * -1
	}
	if coordy < 0 {
		coordy = coordy * -1
	}
	fmt.Println("Steps: ", coordx+coordy)
}

func sumNeighbours(values map[string]int, x int, y int) int {
	sum := 0
	for i := (x - 1); i <= (x + 1); i++ {
		for j := (y + 1); j >= (y - 1); j-- {
			v, ok := values[fmt.Sprintf("%dx%d", i, j)]
			if ok {
				sum += v
			}
		}
	}
	return sum
}
func ThirdSecond() {

	coordx := 0
	coordy := 0
	sum := 0

	values := make(map[string]int)
	currentMove := "right"
	circle := 1
	i := 1
	values["0x0"] = i

	for {
		switch currentMove {
		case "right":
			coordx++

			if coordx == circle {
				currentMove = "up"
			}
		case "up":
			coordy++
			if coordy == circle {
				currentMove = "left"
			}
		case "left":
			coordx--
			if coordx == circle*-1 {
				currentMove = "down"
			}
		case "down":
			coordy--
			if coordy == circle*-1 {
				currentMove = "right"
				circle++
			}
		}
		sum = sumNeighbours(values, coordx, coordy)
		values[fmt.Sprintf("%dx%d", coordx, coordy)] = sum
		if sum > thirdSecondInput {
			break
		}
	}

	fmt.Println("Steps: ", sum)
}
