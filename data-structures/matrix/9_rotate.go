package matrix

/*
Problem Statement:
You are given an nxn 2D matrix representing an image.
You need to rotate the image by 90 degrees (clockwise), in place.

Sample Input:
[
    [1, 2],
    [3, 4]
]

Sample Output:
[
    [3, 1],
    [4, 2]
]
*/

//Method to rotate an NxN matrix by 90 degrees clockwise
func rotate(A [][]int) [][]int {

	N := len(A)
	X, Y := N/2, N/2
	if N%2 == 1 {
		X++
	}
	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			A[i][j], A[N-j-1][i], A[N-i-1][N-j-1], A[j][N-i-1] = A[N-j-1][i], A[N-i-1][N-j-1], A[j][N-i-1], A[i][j]
		}
	}
	return A
}
