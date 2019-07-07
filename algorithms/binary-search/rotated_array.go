package binary_search

/*Problem Statement:
Suppose a sorted array A is rotated at some pivot unknown to you beforehand.
(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
Find the minimum element.

Sample Input:
[4,5,6,7,0,1,2]

Sample Output:
0
*/

//Find Minimum number in a rotated matrix with unique values
func MinRotatedMatrix(A []int) int {
	if len(A) == 0 {
		return -1
	}
	if A[0] < A[len(A)-1] {
		return A[0]
	}
	start, end, mid := 0, len(A) - 1, (len(A) - 1)/2
	for start < end {
		if A[mid]  > A[mid+1] {
			mid = mid+1
			break
		} else if A[start] < A[mid] {
			start = mid
		} else {
			end = mid
		}
		mid = (start+end)/2
	}

	return A[mid]
}