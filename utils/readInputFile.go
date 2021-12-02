package utils

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
