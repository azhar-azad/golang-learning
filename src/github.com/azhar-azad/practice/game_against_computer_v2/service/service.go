package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/azhar-azad/practice/game_against_computer_v2/util"
)

func ShowSinglePlayerGameRules() {
	fmt.Println("Welcome to the Game")
	fmt.Println("=======================")
	fmt.Println("GAME RULES: Single Player")
	fmt.Println("## You will play against the Computer.")
	fmt.Println("## Choice available for both you and the Computer is 1 or 2 and initially 'Game Status' will be 0.")
	fmt.Println("## First, any player (you or computer) will choose (1 or 2).")
	fmt.Println("## 'Game Status' will then be updated by adding the choice to the previous value of 'Game Status'.")
	fmt.Println("## Then, another player will choose (1 or 2).")
	fmt.Println("## 'Game Status' will be updated again.")
	fmt.Println("## 'Game Status' will be increasing and finally if it reaches 20, game will be over")
	fmt.Println("## If your choice make the 20 then you win, otherwise, you lose")
	fmt.Println("")
	fmt.Println("")
}

func GetDifficultyChoice() int {
	fmt.Println("Choose the difficulty level.")
	fmt.Println("## EASY (press 1) -- [ If you are a human, you will win !! ]")
	fmt.Println("## MEDIUM (press 2) -- [ You need a brain to win !! ]")
	fmt.Println("## HARD (press 3) -- [ I dare you to win !! ]")

	var difficultyChoice int
	for {
		fmt.Print("Enter game difficulty choice (1, 2 or 3): ")
		difficultyChoice = util.GetIntegerInput("Negative/Non-integer input is invalid.")
		if difficultyChoice >= 1 && difficultyChoice <= 3 {
			break
		}
		fmt.Println("Invalid difficulty choice. Please try again.")
	}
	return difficultyChoice
}

func GamePlay(difficultyLevel string) {
	println("\nGame Difficulty: " + difficultyLevel + "\n")

	tossWinner := runCoinToss()

	gameStatus := 0

	for {
		if tossWinner == 'c' {
			// comp wins the toss
			compChoice := getCompChoice(difficultyLevel, gameStatus)

			gameStatus = showGameStatus("Computer's", compChoice, gameStatus)
			if gameStatus >= 20 {
				println("You've lost the game !!")
				break
			}
			fmt.Println("")

			playerChoice := getPlayerChoice()

			gameStatus = showGameStatus("Your", playerChoice, gameStatus)
			if gameStatus >= 20 {
				println("You've WON !!")
				break
			}
			fmt.Println("")

		} else {
			// player wins the toss
			playerChoice := getPlayerChoice()

			gameStatus = showGameStatus("Your", playerChoice, gameStatus)
			if gameStatus >= 20 {
				println("You've WON !!")
				break
			}
			fmt.Println("")

			compChoice := getCompChoice(difficultyLevel, gameStatus)

			gameStatus = showGameStatus("Computer's", compChoice, gameStatus)
			if gameStatus >= 20 {
				println("You've lost the game !!")
				break
			}
			fmt.Println("")
		}
	}
}

func runCoinToss() rune {

	fmt.Println("Coin toss is running....")
	fmt.Println("Call HEAD(1) or TAIL(2)")

	var tossChoice int
	for {
		fmt.Print("Enter (1 or 2): ")
		tossChoice = util.GetIntegerInput("Negative/Non-integer input is invalid.")
		if tossChoice == 1 || tossChoice == 2 {
			break
		}
		fmt.Println("Invalid choice. Please try again.")
	}

	fmt.Println("")

	tossResult := util.GetRandomChoice()

	if tossChoice == tossResult {
		fmt.Println("You've won the toss. You will choose first.")
		return 'p'
	}
	fmt.Println("You've lost the toss. Computer will choose first.")
	return 'c'

}

func showGameStatus(name string, choice int, status int) int {
	fmt.Println(name + " choice: " + strconv.Itoa(choice))
	status += choice
	fmt.Println("Game Status: " + strconv.Itoa(status))
	return status
}

func getPlayerChoice() int {
	var playerChoice int
	for {
		fmt.Print("Enter your choice (1 or 2): ")
		playerChoice = util.GetIntegerInput("Negative/Non-integer input is invalid.")
		if playerChoice == 1 || playerChoice == 2 {
			break
		}
		fmt.Println("Invalid choice. Please try again.")
	}
	return playerChoice
}

// Based on the game's difficulty, computer will use this method to make it's choice
// If EASY: returns random
// If MEDIUM or HARD: returns based on winChoices list
// About winChoices, check getWinChoiceList(String) method
func getCompChoice(difficultyLevel string, gameStatus int) int {
	if strings.ToUpper(difficultyLevel) == "EASY" {
		return util.GetRandomChoice()
	}

	winChoices := getWinChoiceList(difficultyLevel)

	for _, value := range winChoices {
		if (value - 1) == gameStatus {
			return 1
		}
		if (value - 2) == gameStatus {
			return 2
		}
	}

	return util.GetRandomChoice()
}

// Returns the checkpoints list to win the game, based on the the difficultyLevel.
// If MEDIUM: this logic will be activated at the half of the game.
// If HARD: this logic will be activated at the start of the game.
func getWinChoiceList(difficultyLevel string) []int {
	var winChoiceList []int

	var firstItem int
	if strings.ToUpper(difficultyLevel) == "MEDIUM" {
		firstItem = 11
	} else {
		firstItem = 2
	}

	for i := firstItem; i <= 20; i += 3 {
		winChoiceList = append(winChoiceList, i)
	}

	return winChoiceList
}
