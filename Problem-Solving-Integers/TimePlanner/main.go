package main

import "fmt"

func MeetingPlanner(slotsA [][]int, slotsB [][]int, dur int) []int {
	aCounter := 0
	bCounter := 0
	for aCounter < len(slotsA) && bCounter < len(slotsB) {
		if HasCommonDuration(slotsA[aCounter], slotsB[bCounter], dur) {
			start := Max(slotsA[aCounter][0], slotsB[bCounter][0])
			return []int{start, start + dur}
		}
		if slotsA[aCounter][1] < slotsB[bCounter][1] {
			aCounter++
		} else {
			bCounter++
		}
	}
	return []int{}
}

func HasCommonDuration(slotA, slotB []int, dur int) bool {
	start := Max(slotA[0], slotB[0])
	end := Min(slotA[1], slotB[1])
	if end-start >= dur {
		return true
	}
	return false
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	slotsA := [][]int{{10, 50}, {60, 120}, {140, 210}}
	slotsB := [][]int{{0, 15}, {20, 40}, {170, 250}}
	fmt.Println(MeetingPlanner(slotsA, slotsB, 35))
}
