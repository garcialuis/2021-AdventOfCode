package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInputAsText(filepath string) []string {

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	return txtlines
}

func StringToIntSlice(txtlines []string) []int {

	nums := make([]int, len(txtlines))

	for i, eachline := range txtlines {
		//fmt.Println(eachline)
		num, err := strconv.Atoi(eachline)
		if err != nil {
			log.Fatalf("failed to convert values: %s", err)
		}
		nums[i] = num
	}

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

/*
//Alternate way to read file:
func ReadFile(filepath string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	fmt.Println("Contents of file: ", string(data))
}
*/
