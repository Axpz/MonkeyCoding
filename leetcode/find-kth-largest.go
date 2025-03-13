package leetcode

func findKthLargest(nums []int, k int) int {
	target := len(nums) - k
	l, r := 0, len(nums)-1

	for l <= r {
		pivotIndex := partition(nums, l, r)
		if pivotIndex == target {
			return nums[pivotIndex]
		} else if pivotIndex > target {
			r = pivotIndex - 1
		} else {
			l = pivotIndex + 1
		}
	}
	return -1
}

func getPivotIndex(nums []int, l, r int) int {
	mid := l + (r-l)/2
	if nums[l] > nums[mid] {
		nums[l], nums[mid] = nums[mid], nums[l]
	}
	if nums[l] > nums[r] {
		nums[l], nums[r] = nums[r], nums[l]
	}
	if nums[mid] > nums[r] {
		nums[mid], nums[r] = nums[r], nums[mid]
	}
	return mid
}

func partition(nums []int, l, r int) int {
	pivotIndex := getPivotIndex(nums, l, r)
	pivot := nums[pivotIndex]

	// swap to the right end
	nums[pivotIndex], nums[r] = nums[r], nums[pivotIndex]

	storeIndex := l
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
			storeIndex++
		}
	}

	// swap back
	nums[storeIndex], nums[r] = nums[r], nums[storeIndex]
	return storeIndex
}
