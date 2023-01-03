package main

import (
	"bf_compiler/compiler"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	compiler.New().Compile(b)
}
