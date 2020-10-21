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
)

//!+
func main() {
	indexed := make([]string, len(os.Args))
	for i, arg := range os.Args {
		indexed[i] = strconv.Itoa(i) + " " + arg
	}
	fmt.Println(strings.Join(indexed, "\n"))
}

//!-
