package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"

	chart "github.com/wcharczuk/go-chart"
)

var expertTab = CalculateExpertTable(Experts)

func CalculateExpertTable(expertNo int) *ExpertTable {
	expertTab := ExpertTable{
		ExpertsNo: expertNo,
	}
	expertTab.FillResultTable()
	expertTab.CalculateMembershipFunc()
	expertTab.FuzzyCore()
	expertTab.FuzzySupport()
	expertTab.FuzzyLimit()

	expertTab.ExpertsNo = 3
	return &expertTab
}

func calculateProb(heights []int) []int {
	result := make([]int, len(heights))
	for i := range result {
		lim := 0.01
		if math.Abs(float64(heights[i])-175.0) <= 5.0 {
			lim = 0
		} else if math.Abs(float64(heights[i])-175.0) <= 10.0 {
			lim = 0.1
		} else if math.Abs(float64(heights[i])-175.0) <= 15.0 {
			lim = 0.5
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

func drawMembershipFunction(res http.ResponseWriter, req *http.Request) {
	seriesMembership := drawContinuousSeries(fmt.Sprintf("Membership function; Experts: %d", expertTab.ExpertsNo),
		MapKeys(expertTab.Membership),
		MapValues(expertTab.Membership))

	seriesCore := drawContinuousSeries("Core",
		MapKeys(expertTab.Core),
		MapValues(expertTab.Core))

	seriesSupport := drawContinuousSeries("Support",
		MapKeys(expertTab.Support),
		MapValues(expertTab.Support))

	seriesLimits := drawContinuousSeries("Limits",
		MapKeys(expertTab.Limit),
		MapValues(expertTab.Limit))
	series := make([]chart.Series, 0)
	series = append(series, seriesMembership, seriesCore, seriesLimits, seriesSupport)

	graph := drawChart(series)
	renderChart(graph, res)
}

func drawApproximationTrapeze(res http.ResponseWriter, req *http.Request) {
	series := drawContinuousSeries("Trapeze approximation function",
		expertTab.HeightsAsFloat(),
		trapec(expertTab.HeightsAsFloat(), 165, 170, 180, 185))

	graph := drawChart([]chart.Series{series})
	renderChart(graph, res)
}

func drawApproximationTriangle(res http.ResponseWriter, req *http.Request) {
	series := drawContinuousSeries("Triangle approximation function",
		expertTab.HeightsAsFloat(),
		trian(expertTab.HeightsAsFloat(), 165, 175, 185))

	graph := drawChart([]chart.Series{series})
	renderChart(graph, res)
}
