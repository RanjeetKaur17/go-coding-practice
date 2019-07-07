package binary_search

/*Problem Statement:
Search for a number in a sorted Array.

Sample Input:
[0,1,2,4,5,6,7] & 4

Sample Output:
3
*/

//Search number in a sorted array
func Search(A []int, N int) int {
	if len(A) == 0 {
		return -1
	}
	if len(A) == 1 {
		if A[0] == N {
			return 0
		}
		return -1
	}
	start, end, mid := 0, len(A) - 1, (len(A) - 1)/2
	for start < end {
		if A[mid] == N {
			break
		} else if N > A[mid] {
			start = mid+1
		} else {
			end = mid
		}
		mid = (start+end)/2
	}

	if A[mid] != N {
		return -1
	}
	return mid
}
