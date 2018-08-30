package main

import (
	"fmt"
	"parkinglot"
)

func main() {
	foo, _ := parkinglot.NewParkingLot(42, 32, 3)

	fmt.Println(foo)
	foo.PrintLevels()
	bar, _ := foo.BuyMonthlyPass(144, true)
	fmt.Println(bar)

	}