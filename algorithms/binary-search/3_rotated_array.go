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

//Find Minimum number in a rotated array with unique values
func MinRotatedArray(A []int) int {

	//If array is empty return
	if len(A) == 0 {
		return -1
	}

	//If array is not rotated return
	if A[0] < A[len(A)-1] {
		return A[0]
	}

	//binary search
	start, end, mid := 0, len(A) - 1, (len(A) - 1)/2
	for start < end {
		if A[mid]  > A[mid+1] {
			mid = mid+1
			break
		} else if A[start] < A[mid] {
			start = mid+1
		} else {
			end = mid
		}
		mid = (start+end)/2
	}

	return A[mid]
}

//Find Index of Minimum number in a rotated array with unique values
func MinRotatedArrayIndex(A []int) int {

	//If array is empty return
	if len(A) == 0 {
		return -1
	}

	//If array is not rotated return
	if A[0] < A[len(A)-1] {
		return 0
	}

	//binary search
	start, end, mid := 0, len(A) - 1, (len(A) - 1)/2
	for start < end {
		if A[mid]  > A[mid+1] {
			mid = mid+1
			break
		} else if A[start] < A[mid] {
			start = mid+1
		} else {
			end = mid
		}
		mid = (start+end)/2
	}

	return mid
}