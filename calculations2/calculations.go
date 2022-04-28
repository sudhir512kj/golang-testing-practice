package calculations

import (
	"fmt"
	"math"
)

func calculateSI(p float64, r float64, t float64) string {
	return fmt.Sprintf("%.2f", (p*r*t)/100)
}

func calculateCI(p float64, r float64, t float64) string {
	return fmt.Sprintf("%.2f", (p*float64(math.Pow((1+r/100.00), t)))-p)
}
