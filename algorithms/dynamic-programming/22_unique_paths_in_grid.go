package dynamic_programming

/*Problem Statement:
Given a grid of size m * n, lets assume you are starting at (1,1) and your goal is to reach (m,n).
At any instance, if you are on (x,y), you can either go to (x, y + 1) or (x + 1, y).

Now consider if some obstacles are added to the grids. How many unique paths would there be?
An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Sample Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]

Sample Output:
2
 */
func UniquePathsWithObstacles(A [][]int )  int {
	m,n := len(A),0

	//update obstacles to -1, for ease
	for i := range A {
		n = len(A[i])
		for j := range A[i] {
			if A[i][j] == 1 {
				A[i][j] = -1
			}
		}
	}

	//If source or destination is result, then no path possible
	if A[0][0] == -1 || A[m-1][n-1] == -1 {
		return 0
	}

	for i := range A {
		for j := range A[i] {
			//If current element is not an obstacle
			if A[i][j] == 0 {
				if i == 0 && j == 0 {
					//Set source value as 1
					A[i][j] = 1
				} else if j == 0 {
					//For all elements in 1st column, update as per above element
					if A[i-1][j] > 0 {
						A[i][j] = A[i-1][j]
					}
				} else if i == 0 {
					//For all elements in 1st row, update as per left element
					if A[i][j-1] > 0 {
						A[i][j] = A[i][j-1]
					}
				} else if A[i-1][j] > 0 && A[i][j-1] > 0 {
					//if current element can be reached from both above and left, sum up the possible paths
					A[i][j] = A[i-1][j] + A[i][j-1]
				} else if A[i-1][j] > 0 {
					//if current element can be reached from above update value
					A[i][j] = A[i-1][j]
				} else if A[i][j-1] > 0 {
					//if current element can be reached from left update value
					A[i][j] = A[i][j-1]
				}
			}
		}
	}

	//return possible paths
	return A[m-1][n-1]
}
