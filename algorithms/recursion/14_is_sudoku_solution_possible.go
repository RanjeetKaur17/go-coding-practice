package recursion

/*Problem Statement:
Given a partially completed Sudoku and validate if solution is possible.

Sample Input:


Sample Output:
[
 [0,6,0,1,0,4,0,5,0],
 [0,0,8,3,0,5,6,0,0],
 [2,0,0,0,0,0,0,0,1],
 [8,0,0,4,0,7,0,0,6],
 [0,0,6,0,0,0,3,0,0],
 [7,0,0,9,0,1,0,0,4],
 [5,0,0,0,0,0,0,0,2],
 [0,0,7,2,0,6,9,0,0],
 [0,4,0,5,0,8,0,7,0]
]

Sample Output:
[
 [9,6,3,1,7,4,2,5,8],
 [1,7,8,3,2,5,6,4,9],
 [2,5,4,6,8,9,7,3,1],
 [8,2,1,4,3,7,5,9,6],
 [4,9,6,8,5,2,3,1,7],
 [7,3,5,9,6,1,8,2,4],
 [5,8,9,7,1,3,4,6,2],
 [3,1,7,2,4,6,9,8,5],
 [6,4,2,5,9,8,1,7,3]
]
*/

//Check if solution is possible or not for a partial sudoku
func IsSolutionPossible(A [][]int) bool {
	solveSudoku(A,0,0)

	//Check if all values are set or not
	for ii := 0; ii < 9; ii++ {
		for jj := 0; jj < 9; jj++ {
			//If some position is empty, return false
			if A[ii][jj] == 0 {
				return false
			}
		}
	}

	//If solution valid, return true
	return true
}


//Return solved sudoku for a partial sudoku
func SudokuSolution(A [][]int) [][]int {
	solveSudoku(A,0,0)
	return A
}

//Solve a partial Sudoku
func solveSudoku(A [][]int, i, j int) bool {
	//if crossed last element, return
	if i > 8 || j > 8 {
		return true
	}

	//set next indexes for passing to recursive call
	ni, nj := i+1, 0
	if j < 8 {
		ni, nj = i, j+1
	}

	//if current value is not set
	if A[i][j] == 0 {
		//try all values from 1 to 9
		for k := 1; k <= 9; k++ {
			//if a value k can be used
			if isPossible(A, i, j, k) {
				//set k to current position
				A[i][j] = k
				//check if solution is possible with k, return true
				if solveSudoku(A, ni, nj) {
					return true
				}
				//if solution not found with k, reset current position
				A[i][j] = 0
			}
		}
	//if current value is set, continue to next position
	} else if solveSudoku(A, ni, nj) {
		//if solution found, return true
		return true
	}
	//if no solution found, return false
	return false
}

//Check if value k can be set on position i,j in Sudoku A
func isPossible(A [][]int, I, J, k int) bool {
	//Check if same value exists in row
	for i := 0; i < 9; i++ {
		if A[I][i] == k {
			return false
		}
	}
	//Check if same value exists in column
	for i := 0; i < 9; i++ {
		if A[i][J] == k {
			return false
		}
	}

	//Check if same value exists in 3 x 3 block
	si := (I/3) * 3
	sj := (J/3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j <3; j++ {
			if A[si + i][sj + j] == k {
				return false
			}
		}
	}
	return true
}