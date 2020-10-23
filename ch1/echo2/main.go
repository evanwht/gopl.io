// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for i, arg := range os.Args[0:] {
		s += sep + strconv.Itoa(i) + " " + arg
		sep = "\n"
	}
	secs := time.Since(start).Seconds()
	fmt.Println(s)
	fmt.Printf("time: %f\n", secs)
}

//!-
