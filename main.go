package main

import (
	//p "sample-app/pass_fail"
	//g "sample-app/guess_go"
	"fmt"
	pc "sample-app/paint_calculator"
)

func main() {
	//p.Grade()
	//g.GuessGo()
	var amount, total float64
	amount, err := pc.PaintCalc(4.2, -3.0)
	fmt.Println(err)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	amount, err = pc.PaintCalc(5.2, 3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)
}
