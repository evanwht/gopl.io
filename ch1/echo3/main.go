// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//!+
func main() {
	start := time.Now()
	indexed := make([]string, len(os.Args))
	for i, arg := range os.Args {
		indexed[i] = strconv.Itoa(i) + " " + arg
	}
	s := strings.Join(indexed, "\n")
	secs := time.Since(start).Seconds()
	fmt.Println(s)
	fmt.Printf("time: %f\n", secs)
}

//!-
