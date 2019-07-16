package two_pointers

/*
Problem Statement:
Given an array ‘A’ of sorted integers and another non negative integer k,
find if there exists 2 indices i and j such that A[i] - A[j] = k, i != j.
Return 0 / 1 ( 0 for false, 1 for true ) for this problem

Sample Input:
[1 3 5] & 4

Sample Output:
1
*/

//Check if two numbers exist in the sorted array with difference K
func DiffKPossible(A []int , B int )  int {
	i, j := 0, 0

	for i < len(A) {
		//j continued from last value, but increment j if it is same as i
		if j == i {
			j = i + 1
		}
		//Check if solution possible from i
		for j < len(A) && A[j]-A[i] < B {
			j++
		}
		//If solution found return
		if j < len(A) && A[j] - A[i] == B {
			return 1
		}
		i++
	}
	return 0
}