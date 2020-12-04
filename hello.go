package main

import (
	"fmt"
	"os"
	"strings"
)

const name = "John Doe"

func main() {
	fmt.Println("Hello World 的", name)
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}

	var x = 0

	for {
		x++
		if x == 10 {
			break
		}
		fmt.Println(x)
	}

	fmt.Println(strings.Join(os.Args[1:], ":"))
}
