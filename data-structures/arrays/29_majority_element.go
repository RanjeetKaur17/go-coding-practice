package arrays

/*
Problem Statement:
Given an array of size n, find the majority element. The majority element is the element that appears more than floor(n/2) times.
You may assume that the array is non-empty and the majority element always exist in the array.

Sample Input:
[2, 1, 2]

Sample Output:
2
*/

func MajorityElement(A []int )  (int) {
	if len(A) == 0 {
		return -1
	} else if len(A) == 1 {
		return A[0]
	}

	p := -1
	C := 0

	for I := 0; I < len(A); I++ {
		if C == 0 {
			p = A[I]
			C = 1
		} else if A[I] == p {
			C++
		} else {
			C--
		}
	}

	return p
}