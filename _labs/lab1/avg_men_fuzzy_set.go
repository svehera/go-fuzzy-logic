package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"

	fuzzy "github.com/svehera/go-fuzzy-logic"
	appr "github.com/svehera/go-fuzzy-logic/appr"
	util "github.com/svehera/go-fuzzy-logic/util"
	chart "github.com/wcharczuk/go-chart"
)

var expertTab = CalculateExpertTable(Experts)

func CalculateExpertTable(experts int) *fuzzy.ExpertTable {
	heights, _ := util.SetHeights(150, 200, 5)
	expertTable := fuzzy.NewExpertTable(experts, heights)
	expertTable.FillResultTable(calculateProb)
	expertTable.CalculateMembership()

	expertTable.Core = fuzzy.DetermineSetCore(expertTable.Membership)
	expertTable.Support = fuzzy.DetermineSetSupport(expertTable.Membership)
	expertTable.Limit = fuzzy.DetermineSetLimits(expertTable.Membership)

	return expertTable
}

func calculateProb(heights []float64) []float64 {
	result := make([]float64, len(heights))
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
	seriesMembership := fuzzy.DrawContinuousSeries(fmt.Sprintf("Membership function; Experts: %d", expertTab.Experts),
		util.MapKeys(expertTab.Membership),
		util.MapValues(expertTab.Membership))

	seriesCore := fuzzy.DrawContinuousSeries("Core",

		util.MapKeys(expertTab.Core),
		util.MapValues(expertTab.Core))

	seriesSupport := fuzzy.DrawContinuousSeries("Support",
		util.MapKeys(expertTab.Support),
		util.MapValues(expertTab.Support))

	seriesLimits := fuzzy.DrawContinuousSeries("Limits",
		util.MapKeys(expertTab.Limit),
		util.MapValues(expertTab.Limit))
	series := make([]chart.Series, 0)
	series = append(series, seriesMembership, seriesCore, seriesLimits, seriesSupport)

	graph := fuzzy.DrawChart(series)
	fuzzy.RenderChart(graph, res)
}

func drawApproximationTrapeze(res http.ResponseWriter, req *http.Request) {
	series := fuzzy.DrawContinuousSeries("Trapeze approximation function",
		expertTab.FuzzySet,
		appr.Trapezoidal(expertTab.FuzzySet, 165, 170, 180, 185))

	graph := fuzzy.DrawChart([]chart.Series{series})
	fuzzy.RenderChart(graph, res)
}

func drawApproximationTriangle(res http.ResponseWriter, req *http.Request) {
	series := fuzzy.DrawContinuousSeries("Triangle approximation function",
		util.MapKeys(expertTab.Membership),
		//	appr.Triangular(expertTab.FuzzySet, 160, 175, 190))
		appr.Triangular(expertTab.Membership))
	graph := fuzzy.DrawChart([]chart.Series{series})
	fuzzy.RenderChart(graph, res)
}
