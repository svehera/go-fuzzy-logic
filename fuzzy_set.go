package fuzzy

import (
	"sort"

	util "github.com/svehera/go-fuzzy-logic/util"
)

// NormilizeSetMembership check whether membership function height is equal to 1.0
// and normalizes it if not
func NormilizeSetMembership(membership map[float64]float64) (normilized map[float64]float64) {
	if max, isNorm := isNormilized(util.MapValues(membership)); isNorm == false {
		for k, v := range membership {
			normilized[k] = v / max
		}
		return normilized
	}
	return membership
}

// DetermineSetCore retuns the core of the membership function (mu == 1)
func DetermineSetCore(membership map[float64]float64) (core map[float64]float64) {
	core = make(map[float64]float64)
	for k, mu := range membership {
		if mu == 1.0 {
			core[k] = mu - 0.03 // minus " 0.03" used to draw series a bit lower that membership series
		}
	}
	return
}

// DetermineSetSupport retuns the support of the membership function (mu > 0)
func DetermineSetSupport(membership map[float64]float64) (support map[float64]float64) {
	support = make(map[float64]float64)
	for k, mu := range membership {
		if mu > 0 {
			support[k] = mu // minus " 0.03" used to draw series a bit above 0
		}
	}
	return
}

// DetermineSetLimits retuns the limits of the membership function (0 < mu < 1)
func DetermineSetLimits(membership map[float64]float64) (limits map[float64]float64) {
	limits = make(map[float64]float64)
	for k, mu := range membership {
		if mu > 0 && mu < 1 {
			limits[k] = mu - 0.03 // minus " 0.03" used to draw series a bit lower that membership series
		}
	}
	return
}

func isNormilized(slice []float64) (float64, bool) {
	sort.Float64s(slice)
	max := slice[len(slice)-1]
	if max == 1.0 {
		return max, true
	}
	return max, false
}
