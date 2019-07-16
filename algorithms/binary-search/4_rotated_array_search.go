package binary_search

/*
Problem Statement:
Suppose a sorted array is rotated at some pivot unknown to you beforehand.
You are given a target value to search. If found in the array, return its index, otherwise return -1.
You may assume no duplicate exists in the array.

Sample Input:
[4,5,6,7,0,1,2] & 4

Sample Output:
0
*/

//Search number in a rotated array
func SearchRotatedArray(A []int, N int) int {
	if len(A) == 0 {
		return -1
	}

	//find Point of rotation
	i := MinRotatedArrayIndex(A)

	//Search in part where value exists
	if A[len(A)-1] >= N {
		return Search(A[i:], N)
	} else {
		return Search(A[:i+1], N)
	}
}