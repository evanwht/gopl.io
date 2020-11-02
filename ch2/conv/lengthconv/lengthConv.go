package lengthconv

import (
	"fmt"
)

type Feet float64
type Inches float64

func (f Feet) String() string {
	return fmt.Sprintf("%f'", f)
}

func (i Inches) String() string {
	return fmt.Sprintf("%f\"", i)
}

func FToI(f Feet) Inches {
	return Inches(f * 12)
}

func IToF(i Inches) Feet {
	return Feet(i / 12)
}
