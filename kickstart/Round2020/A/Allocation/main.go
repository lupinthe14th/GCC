package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := NewScanner()
	T := sc.ReadInt()
	for i := 0; i < T; i++ {
		N := sc.ReadInt()
		B := sc.ReadInt()
		A := make([]int, N)
		for j := 0; j < N; j++ {
			A[j] = sc.ReadInt()
		}
		c := 0
		sort.Ints(A)
		for j := 0; j < N; j++ {
			if A[j] <= B {
				B -= A[j]
				c++
			}
		}
		fmt.Printf("Case #%v: %v\n", i+1, c)
	}
}

type Scanner struct {
	bufScanner *bufio.Scanner
}

func NewScanner() *Scanner {
	bufSc := bufio.NewScanner(os.Stdin)
	bufSc.Split(bufio.ScanWords)
	bufSc.Buffer(nil, 100000000)
	return &Scanner{bufSc}
}

func (sc *Scanner) ReadInt() int {
	bufSc := sc.bufScanner
	bufSc.Scan()
	text := bufSc.Text()

	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num
}
