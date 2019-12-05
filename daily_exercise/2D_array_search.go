package daily_exercise

func twoDArraySeach(arr [][]int, target int) bool {
	rowNum := len(arr)
	if rowNum == 0 {
		return false
	}
	colNum := len(arr[0])
	if colNum == 0 {
		return false
	}

	var firstCol []int
	for i := 0; i < rowNum; i++ {
		firstCol = append(firstCol, arr[i][0])
	}

	idx := seachSmallOrEqual(firstCol, target)
	if idx == -1 {
		return false
	} else if arr[idx][0] == target {
		return true
	}

	return searchEqual(arr[idx], target)
}

// 查找等于或小于target中最大的那个数
func seachSmallOrEqual(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			end = mid - 1
			if end >= 0 && arr[end] <= target {
				return end
			}
		} else {
			start = mid + 1
		}
	}

	if end == len(arr)-1 {
		return end
	}
	return -1
}

func searchEqual(arr []int, target int) bool {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}
