// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + strconv.Itoa(i) + " " + os.Args[i]
		sep = "\n"
	}
	secs := time.Since(start).Seconds()
	fmt.Println(s)
	fmt.Printf("time: %f\n", secs)
}

//!-
// 1.1 add os.Args[0] to output
// 1.2 print index and value on separate lines
// 1.3 experiment with strings.Join for performance
