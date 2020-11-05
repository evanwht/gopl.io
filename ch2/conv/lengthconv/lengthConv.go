package lengthconv

import (
	"fmt"
)

// Feet feet
type Feet float64

// Inches inches
type Inches float64

func (f Feet) String() string {
	return fmt.Sprintf("%f'", f)
}

func (i Inches) String() string {
	return fmt.Sprintf("%f\"", i)
}

// FToI feet to inches
func FToI(f Feet) Inches {
	return Inches(f * 12)
}

// IToF inches to feet
func IToF(i Inches) Feet {
	return Feet(i / 12)
}
