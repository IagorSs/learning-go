package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("Go OS:", runtime.GOOS)
	fmt.Println("Go Architecture:", runtime.GOARCH)
}
