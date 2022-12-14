func binary_search(nums []int, target int) int { 
	low := 0
	high := len(nums)-1

	for low <= high {
		mid := (low + high) / 2

    if target == nums[mid] {
        return mid
    } else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
    return -1
}
