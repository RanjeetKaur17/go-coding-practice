package arrays

import "math"

/*Problem Statement:
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

Sample Input:
[-2,1,-3,4,-1,2,1,-5,4]

Sample Output:
6
*/

//Find Max Contigous Subarray Sum
func MaxSubArraySum(A []int) int {
	max := math.MinInt32

	tSum := 0
	for _, v := range A {
		tSum += v
		if tSum > max {
			max = tSum
		}

		if tSum < 0 {
			tSum = 0
		}
	}

	return max
}

//Method to return Max Sum Contigous Subarray
func MaxSumSubArray(A []int) []int {
	start, end, max := 0,0,math.MinInt32
	//-2,1,-3,4,-1,2,1,-5,4}
	tStart, tEnd, tSum := 0, 0, 0
	for i, v := range A {
		tSum += v
		if tSum > max {
			max = tSum
			start, end = tStart, tEnd
		}

		if tSum < 0 {
			tSum = 0
			tStart = i+1
			tEnd = i+1
		} else {
			tEnd++
		}
	}

	return A[start:end+1]
}