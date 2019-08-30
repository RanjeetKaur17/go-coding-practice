package matrix

/*
Problem Statement:
Given a 2D board containing 'X' and 'O', capture all regions surrounded by 'X'.
A region is captured by flipping all 'O's into 'X's in that surrounded region.

Sample Input:
X X X X
X O O X
X X O X
X O X X

Sample Output
X X X X
X X X X
X X X X
X O X X
*/

func CaptureRegions(A [][]rune) [][]rune {

	m := len(A)
	n := len(A[0])
	for i := 0; i < m; i++ {
		if A[i][0] == 'O' {
			markSafe(A, i, 0, m, n)
		}
		if A[i][n-1] == 'O' {
			markSafe(A, i, n-1, m, n)
		}
	}

	for j := 0; j < n; j++ {
		if A[0][j] == 'O' {
			markSafe(A, 0, j, m, n)
		}
		if A[m-1][j] == 'O' {
			markSafe(A, m-1, j, m, n)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 'O' {
				A[i][j] = 'X'
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 'S' {
				A[i][j] = 'O'
			}
		}
	}
	return A
}

func markSafe(A [][]rune, i, j, m, n int) {
	if i >= 0 && i < m && j >= 0 && j < n && A[i][j] == 'O' {
		A[i][j] = 'S'
		markSafe(A, i-1, j, m, n)
		markSafe(A, i+1, j, m, n)
		markSafe(A, i, j-1, m, n)
		markSafe(A, i, j+1, m, n)
	}
}