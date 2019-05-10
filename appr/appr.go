package appr

import (
	"sort"
)

type Interface interface {
	Approximate(map[float64]float64) []float64
}

func getParamsTrinagular(input map[float64]float64) (a, b, c float64) {
	sorted := SortMapByValue(input)
	a = sorted[2].Key
	b = sorted[len(sorted)/2].Key
	c = sorted[len(sorted)-3].Key
	return
}
func getParamsTrap(input map[float64]float64) (a, b, c, d float64) {
	sorted := SortMapByValue(input)
	a = sorted[2].Key
	b = sorted[len(sorted)/2-1].Key
	c = sorted[len(sorted)/2+1].Key
	d = sorted[len(sorted)-3].Key
	return
}

func SortMapByValue(wordFrequencies map[float64]float64) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	return pl
}

type Pair struct {
	Key   float64
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Key < p[j].Key }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
