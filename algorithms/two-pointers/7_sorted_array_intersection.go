package two_pointers

/*
Problem Statement:
Given 2 sorted arrays, find all the elements which occur in both the arrays.

Sample Input:
A : [1 2 3 3 4 5 6]
B : [3 3 5]

Sample Output:
[3 3 5]
 */

//Find the intersection of two sorted arrays.
func Intersect(A []int , B []int )  ([]int) {
	i,j := 0, 0
	var C []int
	for i < len(A) && j < len(B) {
		if A[i] == B[j] {
			//Intersecting element
			C = append(C, A[i])
			i++
			j++
		} else if A[i] < B[j] {
			i++
		} else {
			j++
		}
	}
	return C
}