package main

import "fmt"

func main() {
	roomLength := 6
	roomScope := 4
	roomMax := 2
	roomProfit := []int{1, 2, 3, 4, 5, 6}

	expectProfit, expectSelectedArr := someFn(roomLength, roomScope, roomMax, roomProfit)

	fmt.Printf("expectProfit: %v\nexpectProfit: %v\n", expectProfit, expectSelectedArr)
}

func someFn(roomLength, roomScope, roomMax int, roomProfit []int) (int, []int) {

	// slice every possile roomlength(4) -> 1,2,3,4-5,6 | 1,-2,3,4,5,-6 | 1,2-3,4,5,6
	slicedRoom := sliceEveryPossible(roomScope, roomProfit)

	// get max (limit = 2 room scope)
	maxRoom := getMaxByLimit(slicedRoom)

	// sum profit
	profit := sumProfit(maxRoom)

	return profit, nil
}

func sliceEveryPossible(roomScope int, roomProfit []int) [][][]int {
	return nil
}

func getMaxByLimit(slicedRoom [][][]int) [][]int {
	return nil
}

func sumProfit(maxRoom [][]int) int {
	return 0
}
