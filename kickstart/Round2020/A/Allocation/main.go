package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	T := 0
	fmt.Fscan(reader, &T)
	for i := 0; i < T; i++ {
		N, B := 0, 0
		fmt.Fscan(reader, &N, &B)

		A := make([]int, N)
		for j := 0; j < N; j++ {
			var a int
			fmt.Fscan(reader, &a)
			A[j] = a
		}
		c := 0
		sort.Ints(A)
		for k := 0; k < N; k++ {
			if A[k] <= B {
				B -= A[k]
				c++
			}
		}
		fmt.Printf("Case #%v: %v\n", i+1, c)
	}
}
