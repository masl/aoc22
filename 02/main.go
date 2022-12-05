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
		calculateOutcome(strategy[0], strategy[1], &score)
	}

	fmt.Println("Part 1:", score)

	score = 0
	for _, line := range input {
		strategy := strings.Split(line, " ")
		calculateOutcomeTwo(strategy[0], strategy[1], &score)
	}

	fmt.Println("Part 2:", score)
}

func calculateOutcomeTwo(opponent string, outcome string, score *int) {
	scoreTable := map[string]int{
		"A": 1, // Rock
		"B": 2, // Paper
		"C": 3, // Scissors
		"X": 1, // Lose
		"Y": 2, // Draw
		"Z": 3, // Win
	}

	myTotalScore := 0
	opponentValue := scoreTable[opponent]
	outcomeValue := scoreTable[outcome]

	// Rock
	if opponentValue == 1 {
		switch outcomeValue {

		// Lose - Scissors
		case 1:
			myTotalScore += 3
			break

		// Draw - Rock
		case 2:
			myTotalScore += 3
			myTotalScore += 1
			break

		// Win - Paper
		case 3:
			myTotalScore += 6
			myTotalScore += 2
			break
		}
		fmt.Println(outcomeValue, myTotalScore)
	}

	// Paper
	if opponentValue == 2 {
		switch outcomeValue {

		// Lose - Rock
		case 1:
			myTotalScore += 1
			break

		// Draw - Paper
		case 2:
			myTotalScore += 3
			myTotalScore += 2
			break

		// Win - Scissors
		case 3:
			myTotalScore += 6
			myTotalScore += 3
			break

		}
		fmt.Println(outcomeValue, myTotalScore)
	}

	// Scissors
	if opponentValue == 3 {
		switch outcomeValue {

		// Lose - Paper
		case 1:
			myTotalScore += 2
			break

		// Draw - Scissors
		case 2:
			myTotalScore += 3
			myTotalScore += 3
			break

		// Win - Rock
		case 3:
			myTotalScore += 6
			myTotalScore += 1
			break
		}
		fmt.Println(outcomeValue, myTotalScore)
	}

	*score += myTotalScore
}

func calculateOutcome(opponent string, me string, score *int) {
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
