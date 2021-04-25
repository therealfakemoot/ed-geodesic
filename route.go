package geodesic

import (
	"math"
)

func fuelUse(d, m, t float64) float64 {
	f := math.Round(10 + d/4*(1+(m+t)/25000))
	return f
}
