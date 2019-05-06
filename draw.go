package fuzzy

import (
	"fmt"
	"net/http"

	chart "github.com/wcharczuk/go-chart"
)

func drawChart(series []chart.Series) chart.Chart {
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 20,
			},
		},
		Series: series,
	}

	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}
	return graph
}

func drawContinuousSeries(name string, xValues, yValues []float64) chart.Series {
	series := chart.ContinuousSeries{
		Name:    name,
		XValues: xValues,
		YValues: yValues,
		Style: chart.Style{
			Show:        true,
			StrokeWidth: 3.0,
		},
	}
	return series
}

func renderChart(graph chart.Chart, res http.ResponseWriter) {
	res.Header().Set("Content-Type", chart.ContentTypeSVG)
	err := graph.Render(chart.SVG, res)
	if err != nil {
		fmt.Printf("Error rendering pie chart: %v\n", err)
	}
}
