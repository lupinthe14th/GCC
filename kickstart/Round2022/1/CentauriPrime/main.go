package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solv(s string) string {
	out := ""
	t := strings.ToLower(s[len(s)-1:])
	switch t {
	case "y":
		out = "nobody"
	case "a", "e", "i", "o", "u":
		out = "Alice"
	default:
		out = "Bob"
	}
	return out
}

func main() {
	r := bufio.NewReader(os.Stdin)
	T := 0
	fmt.Fscan(r, &T)
	for i := 0; i < T; i++ {
		K := ""
		fmt.Fscan(r, &K)
		fmt.Printf("Case #%v: %v is ruled by %v.\n", i+1, K, solv(K))
	}
}
