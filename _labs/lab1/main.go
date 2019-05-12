package main

import (
	"fmt"
	"github.com/svehera/go-fuzzy-logic/appr"
	"log"
	"net/http"
	//	util "github.com/svehera/go-fuzzy-logic/util"
)

const Experts = 100

func main() {
	//	fmt.Printf("%v\n", expertTab.Affilation)
	//	fmt.Printf("%v\n", trap(expertTab.HeightsAsFloat(), 170, 175, 180))
	//	fmt.Printf("%v\n", expertTab.ResultTable)
	//	fmt.Printf("%v\n", expertTab.Core)
	var trian appr.Triangular
	trian = "Triangular"
	//util.SliceAsTable(expertTab.ResultTable)
	//	fmt.Printf("%v\n", expertTab.Membership)
	fmt.Printf("%v\n", expertTab.Membership)
	fmt.Printf("%v\n", appr.SortMapByValue(expertTab.Membership))
	//fmt.Printf("%v\n", appr.Triangular(expertTab.Membership))
	http.HandleFunc("/", drawMembershipFunction)
	//	http.HandleFunc("/trian", drawApproximationTriangle)
	http.HandleFunc("/trian", drawApproximation(expertTab.Membership, trian, "Test"))
	http.HandleFunc("/trap", drawApproximationTrapeze)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
