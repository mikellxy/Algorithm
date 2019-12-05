package daily_exercise

func getOnesNum(val uint64) int {
	var num int
	for val != 0 {
		if val & 1 == 1 {
			num++
		}
		val >>= 1
	}
	return num
}
