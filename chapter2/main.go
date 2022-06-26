package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mattn/go-tty"
)

const (
	FIELD_WIDTH  int = 160
	FIELD_HEIGHT int = 160
)

var field [FIELD_HEIGHT][FIELD_WIDTH]int = [FIELD_HEIGHT][FIELD_WIDTH]int{}

func main() {
	const patternWidth int = 10
	const patternHeight int = 8
	// var pattern [patternHeight][patternWidth]int = [patternHeight][patternWidth]int{
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	// 	{0, 0, 0, 0, 0, 1, 0, 1, 1, 0},
	// 	{0, 0, 0, 0, 0, 1, 0, 1, 0, 0},
	// 	{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
	// 	{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	// 	{0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// }
	var pattern [patternHeight][patternWidth]int = [patternHeight][patternWidth]int{
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 1, 0, 0, 0, 1},
		{1, 0, 1, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 1, 0, 1, 1, 1},
		{0, 0, 1, 0, 0, 1, 0, 0, 0, 1},
	}

	PatternTransfer(
		(FIELD_WIDTH/2)-(patternWidth/2),
		(FIELD_HEIGHT/2)-(patternHeight/2),
		patternWidth,
		patternHeight,
		pattern,
	)
	for {
		DrawField()
		StepSimulation()
	}
}

func DrawField() {
	ClearScreen()
	for y := 0; y < FIELD_HEIGHT; y++ {
		for x := 0; x < FIELD_WIDTH; x++ {
			fmt.Printf("%s", IsAlive(field[y][x]))
		}
		fmt.Printf("\n")
	}
	time.Sleep(time.Millisecond * 100)
}

func IsAlive(v int) string {
	if v == 1 {
		return "Ｏ"
	}
	return "　"
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func WaitInput() string {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	r, err := tty.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	return string(r)
}

func GetLivingCellsCount(_x int, _y int) int {
	count := 0
	for y := _y - 1; y <= _y+1; y++ {
		if (y < 0) || (y >= FIELD_HEIGHT) {
			continue
		}
		for x := _x - 1; x <= _x+1; x++ {
			if (x < 0) || (x >= FIELD_WIDTH) {
				continue
			}
			if (_y == y) && (_x == x) {
				continue
			}
			count += field[y][x]
		}
	}
	return count
}

func StepSimulation() {
	var nextField [FIELD_HEIGHT][FIELD_WIDTH]int = [FIELD_HEIGHT][FIELD_WIDTH]int{}
	for y := 0; y < FIELD_HEIGHT; y++ {
		for x := 0; x < FIELD_WIDTH; x++ {
			livingCellCount := GetLivingCellsCount(x, y)
			if livingCellCount <= 1 {
				nextField[y][x] = 0
			} else if livingCellCount == 2 {
				nextField[y][x] = field[y][x]
			} else if livingCellCount == 3 {
				nextField[y][x] = 1
			} else {
				nextField[y][x] = 0
			}
		}
	}
	field = nextField
}

func PatternTransfer(_destX int, _destY int, _srcWidth int, _srcHeight int, _pPattern [8][10]int) {
	for y := 0; y < _srcHeight; y++ {
		for x := 0; x < _srcWidth; x++ {
			field[_destY+y][_destX+x] = _pPattern[y][x]
		}
	}
}
