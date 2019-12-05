package daily_exercise

import "fmt"

func reversedPrint(n *node) {
	if n == nil {
		return
	}
	reversedPrint(n.next)
	fmt.Println(n.val)
}