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
		for j := range M {
			M[j] = make([]int, K)
		}
		for j := 0; j < N; j++ {
			for k := 0; k < K; k++ {
				fmt.Fscan(reader, &M[j][k])
			}
		}
		dp := make([][]int, N+1)
		for i := range dp {
			dp[i] = make([]int, P+1)
		}
		dp[0][0] = 0
		for j := 0; j < N; j++ {
			// memcpy
			for k := 0; k < len(dp[0]); k++ {
				dp[j+1][k] = dp[j][k]
			}
			s := 0
			for k := 0; k < K; k++ {
				s += M[j][k]
				for l := 0; l+k+1 <= P; l++ {
					dp[j+1][k+l+1] = max(dp[j][l]+s, dp[j+1][k+l+1])
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
