package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args
	fmt.Println(s)
	fmt.Println(s[1])
	fmt.Println(s[1][1])
}
