// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/evanwht/gopl.io/ch2/conv/lengthconv"
	"gopl.io/ch2/tempconv"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Fprintf(os.Stderr, "Too few args. Must declare type of conversion and at least one number to convert\n")
		os.Exit(1)
	}
	conversion := os.Args[1]
	switch conversion {
	case "cf":
		temp(os.Args[2:])
	case "fi":
		length(os.Args[2:])
	case "pk":
		weight(os.Args[2:])
	}
}

func length(args []string) {
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)
			os.Exit(1)
		}
		f := lengthconv.Feet(t)
		i := lengthconv.Inches(t)
		fmt.Printf("%s = %s, %s = %s\n", f, lengthconv.FToI(f), i, lengthconv.IToF(i))
	}
}

func feetToInches(feet float64) float64 {
	return feet * 12
}

func weight(args []string) {
}

func temp(args []string) {
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

//!-
