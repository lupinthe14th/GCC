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
		N, M := 0, 0
		fmt.Fscan(reader, &N, &M)
		sum := 0
		for j := 0; j < N; j++ {
			var c int
			fmt.Fscan(reader, &c)
			sum += c
		}
		fmt.Printf("Case #%v: %v\n", i+1, sum%M)
	}
}
