package dynamic_programming

/*Problem Statement:
Find the longest increasing subsequence of a given sequence / array.
In other words, find a subsequence of array in which the subsequence’s elements are in strictly increasing order, and in which the subsequence is as long as possible.
This subsequence is not necessarily contiguous, or unique.
In this case, we only care about the length of the longest increasing subsequence.

Sample Input:
[0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15]

Sample Output:
6 ([0, 2, 6, 9, 13, 15])
*/

func LongestIncreasingSubsequence(A []int)  int {
	if len(A) == 0 {
		return 0
	}
	dp := make([]int, len(A))
	dp[0] = 1

	max := 1
	for i := 1; i < len(A); i++ {
		for j := i-1; j >= 0; j-- {
			if A[j] < A[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j]+1
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
