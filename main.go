package main

func longestLine(M [][]int) int {
	if len(M) == 0 {
		return 0
	}

	ones := 0

	// dp represents an array for each element of M[0]
	// where the 4 keys represent counts of V,H,Diag up, Diag down
	dp := make([][]int, len(M[0]))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 4)
	}

	for i := 0; i < len(M); i++ {
		old := 0
		for j := 0; j < len(M[0]); j++ {
			if M[i][j] == 1 {
				dp[j][0] = horizontal(j, dp)
				dp[j][1] = vertical(i, j, dp)

				prev := dp[j][2]
				dp[j][2] = diag(i, j, old, dp)
				old = prev
				l := len(M[0]) - 2
				dp[j][3] = diag2(i, j, l, dp)

				ones = max(ones, max(max(dp[j][0], dp[j][1]), max(dp[j][2], dp[j][3])))
			} else {
				old = dp[j][2]
				dp[j][0] = 0
				dp[j][1] = 0
				dp[j][2] = 0
				dp[j][3] = 0
			}
		}
	}

	return ones
}

// horizontal counts are stored at dp[0]
func horizontal(j int, dp [][]int) int {
	// nil guard
	if j > 0 {
		return dp[j-1][0] + 1
	} else {
		return 1
	}
}

// vertical counts are stored at dp[1]
func vertical(i, j int, dp [][]int) int {
	// nil guard
	if i > 0 {
		return dp[j][1] + 1
	} else {
		return 1
	}
}

// diag counts are stored at dp[2]
func diag(i, j, old int, dp [][]int) int {
	// nil guard
	if i > 0 && j > 0 {
		return old + 1
	} else {
		return 1
	}

}

// diag2 counts are stored at dp[3]
func diag2(i, j, l int, dp [][]int) int {
	// nil guard
	if i > 0 && j < l {
		return dp[j+1][3] + 1
	} else {
		return 1
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
