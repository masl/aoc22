package main

import (
	"fmt"
	"strings"

	"github.com/masl/aoc22/utils"
)

func main() {
	input := utils.ReadInput(2)
	score := 0

	for _, line := range input {
		strategy := strings.Split(line, " ")
		calulcateOutcome(strategy[0], strategy[1], &score)
	}

	fmt.Println("Part 1:", score)
}

func calulcateOutcome(opponent string, me string, score *int) {
	scoreTable := map[string]int{
		"A": 1, // Rock
		"B": 2, // Paper
		"C": 3, // Scissors
		"X": 1, // Rock
		"Y": 2, // Paper
		"Z": 3, // Scissors
	}
	myTotalScore := 0
	opponentValue := scoreTable[opponent]
	myValue := scoreTable[me]

	// Rock
	if myValue == 1 {
		switch opponentValue {
		// Rock - Tie
		case 1:
			fmt.Println("Rock - Rock")
			myTotalScore += myValue
			myTotalScore += 3
			break

		// Paper - Lose
		case 2:
			fmt.Println("Rock <- Paper")
			myTotalScore += myValue
			break

		// Scissors - Win
		case 3:
			fmt.Println("Rock -> Scissors")
			myTotalScore += myValue
			myTotalScore += 6
			break

		}
	}

	// Paper
	if myValue == 2 {
		switch opponentValue {
		// Rock - Win
		case 1:
			fmt.Println("Paper -> Rock")
			myTotalScore += myValue
			myTotalScore += 6
			break

		// Paper - Tie
		case 2:
			fmt.Println("Paper - Paper")
			myTotalScore += myValue
			myTotalScore += 3
			break

		// Scissors - Lose
		case 3:
			fmt.Println("Paper <- Scissors")
			myTotalScore += myValue
			break

		}
	}

	// Scissors
	if myValue == 3 {
		switch opponentValue {
		// Rock - Lose
		case 1:
			fmt.Println("Scissors <- Rock")
			myTotalScore += myValue
			break

		// Paper - Win
		case 2:
			fmt.Println("Scissors -> Paper")
			myTotalScore += myValue
			myTotalScore += 6
			break

		// Scissors - Tie
		case 3:
			fmt.Println("Scissors - Scissors")
			myTotalScore += myValue
			myTotalScore += 3
			break
		}
	}

	*score += myTotalScore
}
