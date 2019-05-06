package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"

	chart "github.com/wcharczuk/go-chart"
)

var tallMenTable = CalculatetallMenEstimatele(Experts)

func CalculatetallMenEstimatele(expertNo int) *ExpertTable {
	table := ExpertTable{
		ExpertsNo: expertNo,
	}
	table.FillHeights(170, 200, 5)
	table.FillResultTable(tallMenEstimate)
	table.CalculateMembershipFunc()
	table.FuzzyCore()
	table.FuzzySupport()
	table.FuzzyLimit()

	return &table
}

func tallMenEstimate(heights []int) []int {
	result := make([]int, len(heights))
	for i := range result {
		lim := 0.01
		if heights[i] >= 190 {
			lim = 0
		} else if math.Abs(float64(heights[i])-190.0) <= 5.0 {
			lim = 0.3
		} else if math.Abs(float64(heights[i])-190.0) <= 10.0 {
			lim = 0.5
		} else if math.Abs(float64(heights[i])-190.0) <= 15.0 {
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

func drawTallMenMembershipFunction(res http.ResponseWriter, req *http.Request) {
	seriesMembership := drawContinuousSeries(fmt.Sprintf("Membership function; Experts: %d", tallMenTable.ExpertsNo),
		MapKeys(tallMenTable.Membership),
		MapValues(tallMenTable.Membership))

	seriesCore := drawContinuousSeries("Core",
		MapKeys(tallMenTable.Core),
		MapValues(tallMenTable.Core))

	seriesSupport := drawContinuousSeries("Support",
		MapKeys(tallMenTable.Support),
		MapValues(tallMenTable.Support))

	seriesLimits := drawContinuousSeries("Limits",
		MapKeys(tallMenTable.Limit),
		MapValues(tallMenTable.Limit))
	series := make([]chart.Series, 0)
	series = append(series, seriesMembership, seriesCore, seriesLimits, seriesSupport)

	graph := drawChart(series)
	renderChart(graph, res)
}
