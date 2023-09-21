package merge

import "msis/pkg/sort"

type client struct {
}

func New() sort.Service {
	return &client{}
}

func (c *client) Sort(n int, arr []*int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []*int, begin int, end int) {
	if begin >= end {
		return
	}

	mid := begin + (end-begin)/2
	mergeSort(arr, begin, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, begin, mid, end)
}

func merge(arr []*int, left int, mid int, right int) {
	subArrayOne := mid - left + 1
	subArrayTwo := right - mid

	leftArray := make([]*int, subArrayOne)
	rightArray := make([]*int, subArrayTwo)

	for i := 0; i < subArrayOne; i++ {
		leftArray[i] = arr[left+i]
	}
	for j := 0; j < subArrayTwo; j++ {
		rightArray[j] = arr[mid+1+j]
	}

	indexOfSubArrayOne := 0
	indexOfSubArrayTwo := 0
	indexOfMergedArray := left

	for indexOfSubArrayOne < subArrayOne && indexOfSubArrayTwo < subArrayTwo {
		if *leftArray[indexOfSubArrayOne] <= *rightArray[indexOfSubArrayTwo] {
			arr[indexOfMergedArray] = leftArray[indexOfSubArrayOne]
			indexOfSubArrayOne++
		} else {
			arr[indexOfMergedArray] = rightArray[indexOfSubArrayTwo]
			indexOfSubArrayTwo++
		}
		indexOfMergedArray++
	}

	for indexOfSubArrayOne < subArrayOne {
		arr[indexOfMergedArray] = leftArray[indexOfSubArrayOne]
		indexOfSubArrayOne++
		indexOfMergedArray++
	}

	for indexOfSubArrayTwo < subArrayTwo {
		arr[indexOfMergedArray] = rightArray[indexOfSubArrayTwo]
		indexOfSubArrayTwo++
		indexOfMergedArray++
	}
}
