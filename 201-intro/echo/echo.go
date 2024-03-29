package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	// ECHO 1
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s64 := strconv.FormatFloat(time.Since(start).Seconds(), 'f', -1, 64)
		s += sep + s64 + " " + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)

	//ECHO 2
	// var s, sep string
	// for _, arg := range os.Args[1:] {
	//   s64 := strconv.FormatFloat(time.Since(start).Seconds(), 'f', -1, 64)
	//   s += sep + s64 + " " + arg
	//   sep = "\n"
	// }
	// fmt.println(s)

	// ECHO 3
	// fmt.println(strings.Join(os.Args[1:], " "))
	// fmt.println(os.Args[0:])
}
