package main

import (
	"path/filepath"
	"runtime"
)

func main() {
	err := runPkg()
	if err != nil {
		panic(err)
	}
}

func getCurrentDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("could not determine current directory")
	}
	return filepath.Dir(filename)
}
