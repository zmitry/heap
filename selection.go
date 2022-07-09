package heap

import (
	"golang.org/x/exp/constraints"
)

// split the array into two parts: left and right, right is bigger than pivotValue and left is smaller than pivotValue
func partition[N constraints.Ordered](list []N, left, right, pivotIndex int) int {
	pivotValue := list[pivotIndex]
	// swap pivot index so pivot values goes to the end, for instance, [3,4,6,3,7,4] and pivot index 2, value 6 goes to end and makes [3,4,4,3,7,6]
	// we need this since we don't know at which index our pivot value might end up.
	list[right], list[pivotIndex] = list[pivotIndex], list[right]
	storeIndex := left
	for i := left; i < right; i++ {
		if list[i] < pivotValue {
			// [3,4,6,3,7,4]
			//    ^ pivot index indicated which value is already smaller that pivot value
			list[storeIndex], list[i] = list[i], list[storeIndex]
			storeIndex++
		}
	}
	// once we reached the end of the loop, we know where our pivot value ends up
	list[storeIndex], list[right] = list[right], list[storeIndex]
	return storeIndex
}

func QuickSelect[N constraints.Ordered](list []N, left, right, k int) N {
	for left <= right {
		if left == right {
			return list[left]
		}

		pivotIndex := (left + right) / 2
		pivotIndex = partition(list, left, right, pivotIndex)
		if pivotIndex == k {
			return list[pivotIndex]
		}
		// example [3,4,6*,3,7,4] and k=2, and pivot index is 4, result will be [3,4,4,3] [6,7], * is selected index, and we select first half
		if pivotIndex > k {
			right = pivotIndex - 1
		} else {
			// example [3,4,6,3*,7,4] and k=2, and pivot index is 1, result will be [3,3], [4,6,7,4], and here we select second half
			left = pivotIndex + 1
		}

	}
	var t N
	return t
}
