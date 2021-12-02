package main

import (
	"github.com/garcialuis/2021-AdventOfCode/day1"
)

func main() {

	// DAY 1:
	inputNums := day1.PrepareInput("./day1/sonarSweepInput.txt")
	increasingDepths := day1.NumOfIncreasingDepths(inputNums)
	increasingDepthsByWindow := day1.NumOfIncreasingWindows(inputNums)
	day1.DisplayResults(increasingDepths, increasingDepthsByWindow)
	// Test Data: // nums := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

}
