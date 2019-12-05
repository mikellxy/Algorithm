package daily_exercise

func evenOddPartition(arr []int) []int {
	n := len(arr)
	var i, j int
	for i < n {
		if arr[i]%2 == 0 {
			if j <= i {
				j = i + 1
			}
			for j < n {
				if arr[j]%2 == 1 {
					break
				}
				j++
			}

			if j >= n {
				return arr
			}

			tmp := arr[j]
			for x := j; x > i; x -- {
				arr[x] = arr[x-1]
			}
			arr[i] = tmp
		}
		i++

	}
	return arr
}
