package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo1()
	echo3()
}

func echo1() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}

func echo3() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}
