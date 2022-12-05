package main

import (
	"flag"

	"github.com/masl/aoc22/utils"
)

func main() {
	// flag defining the day of the aoc
	dayFlag := flag.Uint("day", 1, "day of the advent of code puzzle")
	flag.Parse()

	utils.CreateStructure(*dayFlag)
}
