package main

import "fmt"

func simplifiedCountingSort(array []int) []int {
	maxValue := max(array)
	minValue := min(array)
	// calculate the offset
	offset := 0 - minValue
	counts := make([]int, 1+maxValue+offset)

	for _, num := range array {
		// add offset to to counts index
		counts[num+offset] += 1
	}

	sortedIdx := 0
	for element, frequency := range counts {
		for frequency > 0 {
			// subtract offset to to counts index to get current value
			array[sortedIdx] = element - offset
			sortedIdx += 1
			frequency -= 1
		}
	}
	return array
}

func countingSort(array []int) []int {
	maxValue := max(array)
	minValue := min(array)
	// calculate the offset
	offset := 0 - minValue
	counts := make([]int, 1+maxValue+offset)
	sortedArray := make([]int, len(array))

	for _, num := range array {
		// add offset to to counts index
		counts[num+offset] += 1
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	for i := len(array) - 1; i >= 0; i-- {
		currentElement := array[i]
		// add offset to to counts indexs
		position := counts[currentElement+offset]
		sortedIdx := position - 1
		sortedArray[sortedIdx] = currentElement
		// add offset to to counts index
		counts[currentElement+offset] -= 1
	}

	return sortedArray
}

func max(nums []int) int {
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	return maxNum
}

func min(nums []int) int {
	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}
	return minNum
}

func main() {
	array := []int{9, 5, -2, 4, 2, 8, 5}
	fmt.Println("simplified counting sort", simplifiedCountingSort(array))
	fmt.Println("counting sort", countingSort(array))
}

/* output
simplified counting sort [-2 2 4 5 5 8 9]
counting sort [-2 2 4 5 5 8 9]
*/
