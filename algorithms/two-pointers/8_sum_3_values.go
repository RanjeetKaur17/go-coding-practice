package two_pointers

import "sort"

/*
Problem Statement:
Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target.
Return the sum of the three integers.

Sample Input:
[-1,2,1,-4] & 1

Sample Output:
2
*/

//Get closest sum to given target, using 3 values from the array
func ThreeSumClosest(A []int , B int )  int {

	//Return if len of array is less than or equal to 3
	if len(A) <= 3 {
		sum := 0
		for _, v := range A {
			sum += v
		}
		return sum
	}

	sort.Ints(A)

	sum := A[0] + A[1] + A[2]
	minDiff := abs(sum - B)

	for i := 0; i < len(A)-2; i++ {
		//use two pointer, one from start and one from end, on the remaining array, to calculate closest sum
		j := i+1
		k := len(A)-1

		for j < k {
			tSum := A[i]+A[j]+A[k]
			//if exact match found, return
			if tSum == B {
				return tSum
			} else if tSum < B {
				j++
			} else {
				k--
			}

			if abs(tSum - B) < minDiff {
				minDiff = abs(tSum - B)
				sum = tSum
			}
		}
	}

	return sum
}



//Get Absolute value
func abs(i int) int {
	if i <0 {
		return -i
	} else {
		return i
	}
}
