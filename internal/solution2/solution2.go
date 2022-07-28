package solution2

import (
	"fmt"
	"strconv"
)

const SECOND = 1
const MINUTE = SECOND * 60
const HOUR = MINUTE * 60
const DAY = HOUR * 24
const WEEK = DAY * 7

// X - seconds
func Solution2(X int) string {
	// write your code in Go 1.4
	if X == 0 {
		return "0s"
	}

	remainSeconds := X % 60
	minutes := X / 60
	remainMinutes := minutes % 60
	hours := minutes / 60
	remainHours := hours % 24
	days := hours / 24
	remainDays := days % 7
	weeks := days / 7

	// fmt.Printf("remainSeconds: %v\n", remainSeconds)
	// fmt.Printf("minutes: %v\n", minutes)
	// fmt.Printf("remainMinutes: %v\n", remainMinutes)
	// fmt.Printf("hours: %v\n", hours)
	// fmt.Printf("remainHours: %v\n", remainHours)
	// fmt.Printf("days: %v\n", days)
	// fmt.Printf("remainDays: %v\n", remainDays)
	// fmt.Printf("weeks: %v\n", weeks)

	in := []int{weeks, remainDays, remainHours, remainMinutes, remainSeconds}

	if remainSeconds != 0 && (remainMinutes != 0 && remainHours != 0) {
		remainSeconds = 0
		remainMinutes += 1
	}
	if remainMinutes != 0 && (remainHours != 0 && remainDays != 0) {
		remainMinutes = 0
		remainHours += 1
	}
	if remainHours != 0 && (remainDays != 0 && weeks != 0) {
		remainHours = 0
		remainDays += 1
	}

	rounded := []int{weeks, remainDays, remainHours, remainMinutes, remainSeconds}
	result := convert(rounded)
	// result += strconv.Itoa(val) + getUnit(i)
	fmt.Printf("in: %v\n", in)
	fmt.Printf("convert(in): %v\n", convert(in))
	fmt.Printf("rounded: %v\n", rounded)
	fmt.Printf("result: %v\n", result)
	// return strconv.Itoa(seconds) + "s"

	return result
}

func getUnit(i int) string {
	if i == 0 {
		return "w"
	} else if i == 1 {
		return "d"
	} else if i == 2 {
		return "h"
	} else if i == 3 {
		return "m"
	} else if i == 4 {
		return "s"
	} else {
		panic("wrong input")
	}
}

func convert(input []int) string {
	result := ""

	for i, val := range input {
		if val == 0 {
			continue
		}
		result += strconv.Itoa(val) + getUnit(i)
	}

	return result
}
