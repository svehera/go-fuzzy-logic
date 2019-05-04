package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	chart "github.com/wcharczuk/go-chart"
)

const Experts = 100

var lowMenComparisonTable = CalculatelowMenEstimate()

func CalculatelowMenEstimate() *ComparisonTable {
	table := ComparisonTable{}
	table.FillHeights(150, 180, 5)
	table.FillResultTable(lowMenComparison)
	table.CalculateMembershipFunc()
	table.FuzzyCore()
	table.FuzzySupport()
	table.FuzzyLimit()

	return &table
}

func absSubstract(a, b float64) float64 {
	return math.Abs(a - b)
}

func lowMenComparison(heights []int) [][]float64 {
	result := make([][]float64, len(heights))
	for i := range result {
		result[i] = make([]float64, len(heights))
	}
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights); j++ {
			result[i][j] = -1.0
		}
	}
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights); j++ {
			if i == j {
				result[i][j] = 1.0
				continue
			}
			if result[i][j] != -1.0 {
				continue
			}
			comp := absSubstract(float64(heights[i]), float64(heights[j]))
			if comp >= 30.0 {
				result[i][j] = 9
				result[j][i] = 1.0 / 9.0
			} else if comp >= 25.0 {
				result[i][j] = 8
				result[j][i] = 1.0 / 8.0
			} else if comp >= 20.0 {
				result[i][j] = 7
				result[j][i] = 1.0 / 7.0
			} else if comp >= 15.0 {
				result[i][j] = 6
				result[j][i] = 1.0 / 6.0
			} else if comp >= 10.0 {
				result[i][j] = 5
				result[j][i] = 1.0 / 5.0
			} else if comp >= 5.0 {
				result[i][j] = 3
				result[j][i] = 1.0 / 3.0
			}
		}
	}
	return result
}

func drawCompMembershipFunction(res http.ResponseWriter, req *http.Request) {
	seriesMembership := drawContinuousSeries("Test",
		MapKeys(lowMenComparisonTable.Membership),
		MapValues(lowMenComparisonTable.Membership))

	seriesCore := drawContinuousSeries("Core",
		MapKeys(lowMenComparisonTable.Core),
		MapValues(lowMenComparisonTable.Core))

	seriesSupport := drawContinuousSeries("Support",
		MapKeys(lowMenComparisonTable.Support),
		MapValues(lowMenComparisonTable.Support))

	seriesLimits := drawContinuousSeries("Limits",
		MapKeys(lowMenComparisonTable.Limit),
		MapValues(lowMenComparisonTable.Limit))
	series := make([]chart.Series, 0)
	series = append(series, seriesMembership, seriesCore, seriesLimits, seriesSupport)

	graph := drawChart(series)
	renderChart(graph, res)
}

func main() {
	//	fmt.Printf("%v\n", tallMenTable.Affilation)
	//	fmt.Printf("%v\n", trap(tallMenTable.HeightsAsFloat(), 170, 175, 180))
	//	fmt.Printf("%v\n", tallMenTable.ResultTable)
	//	fmt.Printf("%v\n", tallMenTable.Core)

	//	fmt.Printf("%v\n", tallMenTable.Membership)
	for i := 0; i < len(lowMenComparisonTable.ResultTable); i++ {
		for j := 0; j < len(lowMenComparisonTable.ResultTable); j++ {
			fmt.Printf("%f  ", lowMenComparisonTable.ResultTable[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Printf("%v\n", lowMenComparisonTable.Membership)
	//fmt.Printf("%v\n", lowMenComparisonTable.ResultTable)
	//fmt.Printf("%f\n", sqrn(8, 3))
	//	http.HandleFunc("/tall", drawTallMenMembershipFunction)
	//http.HandleFunc("/low", drawLowMenMembershipFunction)
	//	http.HandleFunc("/trian", drawApproximationTriangle)
	//	http.HandleFunc("/trap", drawApproximationTrapeze)
	http.HandleFunc("/", drawCompMembershipFunction)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
