package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	T := 0
	fmt.Fscan(reader, &T)
	for i := 0; i < T; i++ {
		N, K, P := 0, 0, 0
		fmt.Fscan(reader, &N, &K, &P)

		M := make([][]int, N)
		for i := range M {
			M[i] = make([]int, K)
		}
		for i := 0; i < N; i++ {
			for j := 0; j < K; j++ {
				fmt.Fscan(reader, &M[i][j])
			}
		}
		dp := make([][]int, N+1)
		for i := range dp {
			dp[i] = make([]int, P+1)
		}
		dp[0][0] = 0
		for i := 0; i < N; i++ {
			// memcpy
			for j := 0; j < len(dp[0]); j++ {
				dp[i+1][j] = dp[i][j]
			}
			s := 0
			for j := 0; j < K; j++ {
				s += M[i][j]
				for l := 0; l+j+1 <= P; l++ {
					dp[i+1][j+l+1] = max(dp[i][l]+s, dp[i+1][j+l+1])
				}
			}
		}
		fmt.Printf("Case #%v: %v\n", i+1, dp[N][P])
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
