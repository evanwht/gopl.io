// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// Celsius temparature type for the Celsius scale
type Celsius float64

// Fahrenheit temparature type for the Fahrenheit scale
type Fahrenheit float64

// Kelvin temparature type for the Kelvin scale
type Kelvin float64

const (
	// AbsoluteZeroC lowest tempurature on the Celsius scale
	AbsoluteZeroC Celsius = -273.15
	// FreezingC tempurature on the Celsius scale where water freezes
	FreezingC Celsius = 0
	// BoilingC tempurature on the Celsius scale where water boils
	BoilingC Celsius = 100

	// AbsoluteZeroK lowest tempurature on the Kelvin scale
	AbsoluteZeroK Kelvin = 0
	// FreezingK tempurature on the Kelvin scale where water freezes
	FreezingK Kelvin = 273.15
	// BoilingK tempurature on the kelvin scale where water boils
	BoilingK Kelvin = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Kelvin) String() string     { return fmt.Sprintf("%g°K", f) }

//!-
