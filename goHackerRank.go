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

//seven
func sockMerchant(n int32, ar []int32) int32 {
    // Create a map to count the occurrences of each color
    colorCount := make(map[int32]int)

    // Count occurrences of each color
    for _, color := range ar {
        colorCount[color]++
    }

    // Count pairs based on occurrences
    pairCount := int32(0)
    for _, count := range colorCount {
        pairCount += int32(count / 2)
    }

    return pairCount
}

//eight
func bonAppetit(bill []int, k int, b int) {
	totalCost := 0

	// Calculate the total cost excluding the item Anna didn't eat
	for i, cost := range bill {
		if i != k {
			totalCost += cost
		}
	}

	// Calculate Anna's share
	annaShare := totalCost / 2

	// Check if Brian's calculation is correct
	if annaShare == b {
		fmt.Println("Bon Appetit")
	} else {
		// Calculate the difference and print the amount Brian owes Anna
		amountOwed := b - annaShare
		fmt.Println(amountOwed)
	}
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

//seven
// There is a large pile of socks that must be paired by color. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.

// Example
//n=7
//ar={1,2,1,3,1,2,1}

// There is one pair of color  and one of color . There are three odd socks left, one of each color. The number of pairs is .

// Function Description

// Complete the sockMerchant function in the editor below.

// sockMerchant has the following parameter(s):

// int n: the number of socks in the pile
// int ar[n]: the colors of each sock
// Returns

// int: the number of pairs
// Input Format

// The first line contains an integer , the number of socks represented in .
// The second line contains  space-separated integers, , the colors of the socks in the pile.
// Example usage:
	// You can replace the values with your input
	// n := int32(9)
	// ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	// result := sockMerchant(n, ar)
	// fmt.Println(result)

    //eight
//     Two friends Anna and Brian, are deciding how to split the bill at a dinner. Each will only pay for the items they consume. Brian gets the check and calculates Anna's portion. You must determine if his calculation is correct.

// For example, assume the bill has the following prices: . Anna declines to eat item  which costs . If Brian calculates the bill correctly, Anna will pay . If he includes the cost of , he will calculate . In the second case, he should refund  to Anna.

// Function Description

// Complete the bonAppetit function in the editor below. It should print Bon Appetit if the bill is fairly split. Otherwise, it should print the integer amount of money that Brian owes Anna.

// bonAppetit has the following parameter(s):

// bill: an array of integers representing the cost of each item ordered
// k: an integer representing the zero-based index of the item Anna doesn't eat
// b: the amount of money that Anna contributed to the bill
// Input Format

// The first line contains two space-separated integers  and , the number of items ordered and the -based index of the item that Anna did not eat.
// The second line contains  space-separated integers  where .
// The third line contains an integer, , the amount of money that Brian charged Anna for her share of the bill.

// Example usage:
	// You can replace the values with your input
	// bill := []int{3, 10, 2, 9}
	// k := 1
	// b := 12

	// bonAppetit(bill, k, b)
}