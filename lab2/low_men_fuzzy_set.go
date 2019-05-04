package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"

	chart "github.com/wcharczuk/go-chart"
)

var lowMenTable = CalculatelowMenEstimatele(Experts)

// CalculatelowMenEstimatele tests 
func CalculatelowMenEstimatele(expertNo int) *ExpertTable {
	table := ExpertTable{
		ExpertsNo: expertNo,
	}
	table.FillHeights(150, 180, 5)
	table.FillResultTable(lowMenEstimate)
	table.CalculateMembershipFunc()
	table.FuzzyCore()
	table.FuzzySupport()
	table.FuzzyLimit()

	return &table
}

func lowMenEstimate(heights []int) []int {
	result := make([]int, len(heights))
	for i := range result {
		lim := 0.01
		if heights[i] <= 165 {
			lim = 0
		} else if math.Abs(float64(heights[i])-165.0) <= 5.0 {
			lim = 0.3
		} else if math.Abs(float64(heights[i])-165.0) <= 10.0 {
			lim = 0.5
		} else if math.Abs(float64(heights[i])-165.0) <= 15.0 {
			lim = 0.7
		} else {
			lim = 1.0
		}

		if rand.Float64() > lim {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}
	return result
}

func drawLowMenMembershipFunction(res http.ResponseWriter, req *http.Request) {
	seriesMembership := drawContinuousSeries(fmt.Sprintf("Membership function; Experts: %d", lowMenTable.ExpertsNo),
		MapKeys(lowMenTable.Membership),
		MapValues(lowMenTable.Membership))

	seriesCore := drawContinuousSeries("Core",
		MapKeys(lowMenTable.Core),
		MapValues(lowMenTable.Core))

	seriesSupport := drawContinuousSeries("Support",
		MapKeys(lowMenTable.Support),
		MapValues(lowMenTable.Support))

	seriesLimits := drawContinuousSeries("Limits",
		MapKeys(lowMenTable.Limit),
		MapValues(lowMenTable.Limit))
	series := make([]chart.Series, 0)
	series = append(series, seriesMembership, seriesCore, seriesLimits, seriesSupport)

	graph := drawChart(series)
	renderChart(graph, res)
}
