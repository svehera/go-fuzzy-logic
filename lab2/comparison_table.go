package main

import "sort"

type ComparisonTable struct {
	Heights     []int
	ResultTable [][]float64
	Membership  map[float64]float64
	Core        map[float64]float64
	Support     map[float64]float64
	Limit       map[float64]float64
}

func (ct *ComparisonTable) FillHeights(minHeight, maxHeight, inc int) {
	for h := minHeight; h <= maxHeight; h += inc {
		ct.Heights = append(ct.Heights, h)
	}
}

func (ct *ComparisonTable) FillResultTable(estimate func([]int) [][]float64) {
	ct.ResultTable = estimate(ct.Heights)
}

func (ct *ComparisonTable) CalculateMembershipFunc() {
	membershipFuncValues := ComparisonTableMembersip(ct.ResultTable)
	//membershipFuncValues = normalize(membershipFuncValues)
	ct.Membership = make(map[float64]float64)
	temp := ct.HeightsAsFloat()
	sort.Float64s(temp)
	for i := 0; i < len(ct.Heights); i++ {
		ct.Membership[temp[i]] = membershipFuncValues[i]
	}
	ct.Membership = normalize2(ct.Membership)
}

func (ct *ComparisonTable) FuzzyCore() {
	ct.Core = make(map[float64]float64)

	for k, v := range ct.Membership {
		if v == 1 {
			ct.Core[k] = v - 0.03
		}
	}
}

func (ct *ComparisonTable) FuzzySupport() {
	ct.Support = make(map[float64]float64)

	for k, v := range ct.Membership {
		if v > 0 {
			ct.Support[k] = 0.03
		}
	}
}

func (ct *ComparisonTable) FuzzyLimit() {
	ct.Limit = make(map[float64]float64)

	for k, v := range ct.Membership {
		if v > 0 && v < 1 {
			ct.Limit[k] = v - 0.03
		}
	}
}

func (ct *ComparisonTable) HeightsAsFloat() []float64 {
	result := make([]float64, len(ct.Heights))
	for i := range result {
		result[i] = float64(ct.Heights[i])
	}
	return result
}
