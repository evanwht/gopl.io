package weightconv

import (
	"fmt"
)

// Pounds pounds
type Pounds float64

// Kilos kilograms
type Kilos float64

func (p Pounds) String() string {
	return fmt.Sprintf("%f lbs", p)
}

func (k Kilos) String() string {
	return fmt.Sprintf("%f kg", k)
}

// PToK converts pounds to kilograms
func PToK(p Pounds) Kilos {
	return Kilos(p / 2.20462)
}

// KToP converts kilograms to pounds
func KToP(k Kilos) Pounds {
	return Pounds(k * 2.20462)
}
