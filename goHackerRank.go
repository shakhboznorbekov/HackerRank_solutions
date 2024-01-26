package main

import (
	"fmt"
	"sort"
	"time"
)

//1
func plusMinus(arr []int32) {
    total := float64(len(arr))
    pos, neg, zero := 0, 0, 0

    for _, num := range arr {
        switch {
        case num > 0:
            pos++
        case num < 0:
            neg++
        default:
            zero++
        }
    }

    ratioPos := float64(pos) / total
    ratioNeg := float64(neg) / total
    ratioZero := float64(zero) / total

    fmt.Printf("%.6f\n%.6f\n%.6f\n", ratioPos, ratioNeg, ratioZero)
}
//2
func staircase(n int32) {
    for i := int32(1); i <= n; i++ {
        spaces := n - i
        stairs := i

        for j := int32(0); j < spaces; j++ {
            fmt.Print(" ")
        }

        for k := int32(0); k < stairs; k++ {
            fmt.Print("#")
        }

        fmt.Println()
    }
}
//3
func miniMaxSum(arr []int32) {
    // Convert the array to int for sorting
    intArr := make([]int, len(arr))
    for i, v := range arr {
        intArr[i] = int(v)
    }

    // Sort the array
    sort.Ints(intArr)

    // Calculate the minimum sum (excluding the largest element)
    minSum := int64(intArr[0] + intArr[1] + intArr[2] + intArr[3])

    // Calculate the maximum sum (excluding the smallest element)
    maxSum := int64(intArr[1] + intArr[2] + intArr[3] + intArr[4])

    fmt.Printf("%d %d\n", minSum, maxSum)
}
//4
func birthdayCakeCandles(candles []int32) int32 {
    var maxHeight int32
    var count int32

    // Find the maximum height of candles
    for _, height := range candles {
        if height > maxHeight {
            maxHeight = height
            count = 1
        } else if height == maxHeight {
            count++
        }
    }

    return count
}

//5
func timeConversion(s string) string {
	// Parse the input time string in 03:04:05PM format
	t, err := time.Parse("03:04:05PM", s)
	if err != nil {
		fmt.Println("Error:", err)
	return ""
}

// Format the time in 24-hour format (15:04:05)
return t.Format("15:04:05")
}

//6
func dayOfProgrammer(year int32) string {
	var day, month int32

	isLeapYear := func(y int32) bool {
		if y < 1918 {
			return y%4 == 0
		} else {
			return (y%400 == 0) || (y%4 == 0 && y%100 != 0)
		}
	}

	if year == 1918 {
		// Transition year (1918)
		day += 13 // 13 days were skipped during the transition
	} else {
		// Non-leap year, so February has 28 days
		if !isLeapYear(year) {
			if day > 31+28 {
				day -= 31 + 28
				month = 9
			}
		} else {
			// Leap year, so February has 29 days
			if day > 31+29 {
				day -= 31 + 29
				month = 9
			}
		}
	}

	return fmt.Sprintf("%02d.%02d.%04d", day, month, year)
}
func main(){
// Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with  places after the decimal.

// Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with absolute error of up to  are acceptable.

// Example

// There are  elements, two positive, two negative and one zero. Their ratios are ,  and . Results are printed as:

// 0.400000
// 0.400000
// 0.200000
arr :=[]int32{2,3,0,-1,-2,-6,1}

plusMinus(arr)

//second 
// Staircase detail

// This is a staircase of size :

//    #
//   ##
//  ###
// ####
// Its base and height are both equal to . It is drawn using # symbols and spaces. The last line is not preceded by any spaces.

// Write a program that prints a staircase of size .
var n int32 = 4

staircase(n)

//three
// Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

// Example

// The minimum sum is  and the maximum sum is . The function prints

// 16 24
arr2 := []int32{1,2,3,4,9,8}
miniMaxSum(arr2)

//four
// You are in charge of the cake for a child's birthday. You have decided the cake will have one candle for each year of their total age. They will only be able to blow out the tallest of the candles. Count how many candles are tallest.

// Example

// candles =[4,4,1,3]
// The maximum height candles are  units high. There are  of them, so return .

candles := []int32{4,4,1,3}
fmt.Println(birthdayCakeCandles(candles))

//five
// Given a time in -hour AM/PM format, convert it to military (24-hour) time.

// Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
// - 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

// Example
// s='12:01:00PM'
// Return '12:01:00'.
// s='12:01:00AM'
// Return '00:01:00'.

// Function Description
// Complete the timeConversion function in the editor below. It should return a new string representing the input time in 24 hour format.
s:="03:04:05PM"
k:="12:04:05AM"
fmt.Println(timeConversion(s))
fmt.Println(timeConversion(k))

//six
// Example usage with year 1984
result := dayOfProgrammer(1984)
fmt.Println(result)

}