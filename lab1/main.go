package main

import (
	"fmt"
	"log"
	"net/http"
)

const Experts = 100

func main() {
	//	fmt.Printf("%v\n", expertTab.Affilation)
	//	fmt.Printf("%v\n", trap(expertTab.HeightsAsFloat(), 170, 175, 180))
	//	fmt.Printf("%v\n", expertTab.ResultTable)
	//	fmt.Printf("%v\n", expertTab.Core)

	fmt.Printf("%v\n", expertTab.Membership)

	http.HandleFunc("/", drawMembershipFunction)
	http.HandleFunc("/trian", drawApproximationTriangle)
	http.HandleFunc("/trap", drawApproximationTrapeze)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
