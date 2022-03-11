package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolv(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in, want string
	}{
		{in: "Mollaristan", want: "Bob"},
		{in: "Auritania", want: "Alice"},
		{in: "Zizily", want: "nobody"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solv(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	fd, _ := os.Open(filepath.Join("testdata", "sample"))
	orgStdin := os.Stdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output:
	// Case #1: Mollaristan is ruled by Bob.
	// Case #2: Auritania is ruled by Alice.
	// Case #3: Zizily is ruled by nobody.
}
