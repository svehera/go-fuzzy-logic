package main

import (
	"fmt"
	"net/http"

	util "github.com/svehera/go-fuzzy-logic/util"
)

const Experts = 5

func main() {
	//	fmt.Printf("%v\n", expertTab.Affilation)
	//	fmt.Printf("%v\n", trap(expertTab.HeightsAsFloat(), 170, 175, 180))
	//	fmt.Printf("%v\n", expertTab.ResultTable)
	//	fmt.Printf("%v\n", expertTab.Core)

	util.SliceAsTable(expertTab.ResultTable)
	fmt.Printf("%v\n", expertTab.Membership)

	http.HandleFunc("/", drawMembershipFunction)
	http.HandleFunc("/trian", drawApproximationTriangle)
	http.HandleFunc("/trap", drawApproximationTrapeze)
	//log.Fatal(http.ListenAndServe(":8000", nil))
}
