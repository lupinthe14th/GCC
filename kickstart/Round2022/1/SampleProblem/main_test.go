package main

import (
	"os"
	"path/filepath"
)

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
	// Case #1: 1
	// Case #2: 5
}
