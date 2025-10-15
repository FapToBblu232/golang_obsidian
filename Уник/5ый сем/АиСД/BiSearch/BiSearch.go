package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func biSearch(nums []int, target int) int {
// 	low, high := 0, len(nums)-1
// 	answ := -1
// 	for low <= high {
// 		mid := low + int((high-low)/2)
// 		if nums[mid] < target {
// 			low = mid + 1
// 			continue
// 		} else if nums[mid] > target {
// 			high = mid - 1
// 		} else {
// 			answ = mid
// 			high = mid - 1
// 		}
// 	}
// 	return answ
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func biSearchRec(nums []int, target, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + int((high-low)/2)
	if nums[mid] == target {
		left := biSearchRec(nums, target, low, mid-1)
		if left == -1 {
			return mid
		}
	} else if nums[mid] < target {
		return biSearchRec(nums, target, mid+1, high)
	}
	return biSearchRec(nums, target, low, high-1)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	str := reader.Text()
	numbers_in_fields := strings.Fields(str)
	nums := make([]int, 0, len(numbers_in_fields))
	for i := range numbers_in_fields {
		val, _ := strconv.Atoi(numbers_in_fields[i])
		nums = append(nums, val)
	}

	for reader.Scan() {
		temp := reader.Text()
		fields := strings.Fields(temp)
		target, _ := strconv.Atoi(fields[1])
		fmt.Println(biSearchRec(nums, target, 0, len(nums)-1))
	}
}
