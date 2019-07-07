package binary_search

/*Problem Statement:
Compute and return the square root of x.
If x is not a perfect square, return floor(sqrt(x))

Sample Input:
11

Sample Output:
3
*/

//Find Square Root of an Integer
func SquareRoot(n int) int {
	if n <=0 {
		return 0
	}
	if n == 1 {
		return n
	}
	root := n

	start, end, mid := 0, n, n/2
	for start < end {
		if mid * mid == n || (mid * mid < n && (mid+1) * (mid+1) > n) {
			root = mid
			break
		} else if mid * mid > n {
			end = mid+1
		} else {
			start = mid
		}
		mid = (start + end)/2
	}
	return root
}