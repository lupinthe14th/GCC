package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := NewReader()
	T := readInt(reader)
	for i := 0; i < T; i++ {
		NB := readInts(reader)
		N, B := NB[0], NB[1]
		A := readInts(reader)
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

func NewReader() *bufio.Reader {
	var fp *os.File = os.Stdin
	return bufio.NewReaderSize(fp, 1024)
}

func readInt(reader *bufio.Reader) int {
	buf := make([]byte, 0, 1024)

	for {
		l, p, _ := reader.ReadLine()

		buf = append(buf, l...)
		if !p {
			break
		}
	}

	num, err := strconv.Atoi(string(buf))
	if err != nil {
		panic(err)
	}
	return num
}

func readInts(reader *bufio.Reader) []int {
	buf := make([]byte, 0, 1024)

	for {
		l, p, _ := reader.ReadLine()

		buf = append(buf, l...)
		if !p {
			break
		}
	}

	texts := strings.Split(string(buf), " ")
	out := make([]int, len(texts))
	for i := range texts {
		num, err := strconv.Atoi(texts[i])
		if err != nil {
			panic(err)
		}
		out[i] = num
	}
	return out
}
