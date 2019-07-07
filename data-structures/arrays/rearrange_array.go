package arrays

/*Problem Statement:
Rearrange a given array so that Arr[i] becomes Arr[Arr[i]] with O(1) extra space.

Sample Input:
[1,0]

Sample Output:
[0,1]
*/

func rearrange(A []int) []int {

	for i := 0; i < len(A); i++ {
		if A[i] >= len(A) {
			continue
		}

		j := i
		k := A[i]
		for k != i {
			A[j] = A[k] + len(A)
			k = j
		}
	}

	for i := 0; i < len(A); i++ {
		A[i] -= len(A)
	}
	return A
}