/**
 * Problem Statement:
 *
 *  Given an array of distinct integers, find the pair of integers whose sum is equal to the target number, if any
 *
 * Arguments:
 *
 *  1. Array of distinct integers
 *  2. Target sum number
 *
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// O(n^2) time | O(1) space
// func twoNumberSum(numbersArray []int, targetSum int) []int {
// 	for i := 0; i < len(numbersArray); i++ {
// 		firstNumber := numbersArray[i]
// 		for j := i + 1; j < len(numbersArray); j++ {
// 			secondNumber := numbersArray[j]
// 			if firstNumber+secondNumber == targetSum {
// 				return []int{firstNumber, secondNumber}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// O(n) time | O(n) space
// func twoNumberSum(numbersArray []int, targetSum int) []int {
// 	numbersMap := make(map[int]bool)
// 	for i := 0; i < len(numbersArray); i++ {
// 		number := numbersArray[i]
// 		match := targetSum - number
// 		if _, ok := numbersMap[match]; ok {
// 			return []int{match, number}
// 		} else {
// 			numbersMap[number] = true
// 		}
// 	}
// 	return []int{}
// }

// O(nlogn) time | O(1) space
func twoNumberSum(numbersArray []int, targetSum int) []int {
	sort.Ints(numbersArray)
	left := 0
	right := len(numbersArray) - 1
	for left < right {
		currentSum := numbersArray[left] + numbersArray[right]
		if currentSum == targetSum {
			return []int{numbersArray[left], numbersArray[right]}
		} else if currentSum < targetSum {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

func castToIntSlice(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

/*
 input: go run twoNumberSum.go
 2 5 3 7
 10
 output: [3 7]
*/
func main() {
	var numbersArray []int
	var targetSum int
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2 && scanner.Scan(); i++ {
		switch i {
		case 0:
			numbersArray = castToIntSlice(scanner.Text())
		case 1:
			sum, err := strconv.Atoi(scanner.Text())
			if err == nil {
				targetSum = sum
			}
		}
	}
	result := twoNumberSum(numbersArray, targetSum)
	fmt.Println(result)
}
