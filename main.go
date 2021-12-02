package main

import (
	"fmt"

	"github.com/garcialuis/2021-AdventOfCode/day1"
)

func main() {

	// DAY 1:
	textLines := day1.ReadInputAsText("./day1/sonarSweepInput.txt")
	nums := day1.StringToIntSlice(textLines)
	// Test Data: // nums := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	increasingDepths := day1.NumOfIncreasingDepths(nums)
	increasingDepthsByWindow := day1.NumOfIncreasingWindows(nums)
	fmt.Printf("DAY 1 RESULTS: \n\t(part1) -> %d increasing depths \n\t(part2) -> %d increasing depths in windows of 3.\n",
		increasingDepths, increasingDepthsByWindow)

}
