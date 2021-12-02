package day1

import (
	"fmt"

	"github.com/garcialuis/2021-AdventOfCode/utils"
)

func PrepareInput(filename string) (nums []int) {

	textLines := utils.ReadInputAsText(filename)
	nums = utils.StringToIntSlice(textLines)

	return nums
}

func NumOfIncreasingDepths(depths []int) int {

	count := 0

	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}

	return count
}

func NumOfIncreasingWindows(depths []int) int {

	var count, currentWindow int

	prevWindow := depths[0] + depths[1] + depths[2]
	start, end := 0, 2

	for end < len(depths)-1 {

		currentWindow = prevWindow - depths[start] + depths[end+1]

		if prevWindow < currentWindow {
			count++
		}

		prevWindow = currentWindow

		start++
		end++
	}

	return count
}

func DisplayResults(increasingDepths, increasingDepthsByWindow int) {
	fmt.Printf("DAY 1 RESULTS: \n\t(part1) -> %d increasing depths \n\t(part2) -> %d increasing depths in windows of 3.\n",
		increasingDepths, increasingDepthsByWindow)
}
