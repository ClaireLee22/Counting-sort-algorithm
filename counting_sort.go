package main

import "fmt"

func simplifiedCountingSort(array []int) []int {
	// step 1.a: get max value in the array
	maxValue := max(array)
	// step 1.b: initialize one auxiliary array
	counts := make([]int, 1+maxValue)

	// step2: count the frequency of every element in the array
	for _, num := range array {
		counts[num] += 1
	}

	// step3: rearrange the array to in a sorted manner
	sortedIdx := 0
	for element, frequency := range counts {
		for frequency > 0 {
			array[sortedIdx] = element
			sortedIdx += 1
			frequency -= 1
		}
	}
	return array
}

func countingSort(array []int) []int {
	// step 1.a: get max value in the array
	maxValue := max(array)
	// step 1.b: initialize two auxiliary arrays
	counts := make([]int, 1+maxValue)
	sortedArray := make([]int, len(array))

	// step2: count the frequency of every element in the array
	for _, num := range array {
		counts[num] += 1
	}

	// step3: compute the cumulative sum of counts
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	// step 4: loop backward to build up the sorted array
	for i := len(array) - 1; i >= 0; i-- {
		currentElement := array[i]
		// find the element's position
		position := counts[currentElement]
		sortedIdx := position - 1
		// place the element to the sorted array
		sortedArray[sortedIdx] = currentElement
		// decreament the corresponding value in counts by 1
		counts[currentElement] -= 1
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

func main() {
	array := []int{9, 5, 2, 4, 2, 8, 5}
	fmt.Println("simplified counting sort", simplifiedCountingSort(array))
	fmt.Println("counting sort", countingSort(array))
}

/* output
simplified counting sort [2 2 4 5 5 8 9]
counting sort [2 2 4 5 5 8 9]
*/
