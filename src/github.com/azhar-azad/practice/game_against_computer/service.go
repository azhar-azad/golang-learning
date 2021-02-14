package main

import (
	"strconv"
	"strings"
)

type Service struct {
	util Util
}

func (service *Service) showSinglePlayerGameRules() {
	println("Welcome to the Game")
	println("=======================")
	println("GAME RULES: Single Player")
	println("## You will play against the Computer.")
	println("## Choice available for both you and the Computer is 1 or 2 and initially 'Game Status' will be 0.")
	println("## First, any player (you or computer) will choose (1 or 2).")
	println("## 'Game Status' will then be updated by adding the choice to the previous value of 'Game Status'.")
	println("## Then, another player will choose (1 or 2).")
	println("## 'Game Status' will be updated again.")
	println("## 'Game Status' will be increasing and finally if it reaches 20, game will be over")
	println("## If your choice make the 20 then you win, otherwise, you lose")
	println("")
	println("")
}

func (service *Service) getDifficultyChoice() int8 {
	println("Choose the difficulty level.")
	println("## EASY (press 1) -- [ If you are a human, you will win !! ]")
	println("## MEDIUM (press 2) -- [ You need a brain to win !! ]")
	println("## HARD (press 3) -- [ I dare you to win !! ]")

	var difficultyChoice int8
	for {
		print("Enter game difficulty choice (1, 2 or 3): ")
		difficultyChoice = int8(service.util.getIntegerInput("Negative/Non-integer input is invalid."))
		if difficultyChoice >= 1 && difficultyChoice <= 3 {
			break
		}
		println("Invalid difficulty choice. Please try again.")
	}
	return difficultyChoice
}

func (service *Service) gamePlay(difficultyLevel string) {
	println("\nGame Difficulty: " + difficultyLevel + "\n")

	var tossWinner rune
	if service.runCoinToss() {
		tossWinner = 'p'
	} else {
		tossWinner = 'c'
	}

	gameStatus := 0

	for {
		if tossWinner == 'c' {
			// comp wins the toss
			compChoice := service.getCompChoice(difficultyLevel, gameStatus)

			gameStatus = service.showGameStatus("Computer's", compChoice, gameStatus)
			if gameStatus >= 20 {
				println("You've lost the game !!")
				break
			}
			println("")

			playerChoice := service.getPlayerChoice()

			gameStatus = service.showGameStatus("Your", playerChoice, gameStatus)
			if gameStatus >= 20 {
				println("You've WON !!")
				break
			}
			println("")

		} else {
			// player wins the toss
			playerChoice := service.getPlayerChoice()

			gameStatus = service.showGameStatus("Your", playerChoice, gameStatus)
			if gameStatus >= 20 {
				println("You've WON !!")
				break
			}
			println("")

			compChoice := service.getCompChoice(difficultyLevel, gameStatus)

			gameStatus = service.showGameStatus("Computer's", compChoice, gameStatus)
			if gameStatus >= 20 {
				println("You've lost the game !!")
				break
			}
			println("")
		}
	}
}

func (service *Service) runCoinToss() bool {

	println("Coin toss is running....")
	println("Call HEAD(1) or TAIL(2)")

	var tossChoice int
	for {
		print("Enter (1 or 2): ")
		tossChoice = service.util.getIntegerInput("Negative/Non-integer input is invalid.")
		if tossChoice == 1 || tossChoice == 2 {
			break
		}
		println("Invalid choice. Please try again.")
	}

	println("")

	tossResult := service.getCoinTossResult()

	if tossChoice == tossResult {
		println("You've won the toss. You will choose first.")
		return true
	} else {
		println("You've lost the toss. Computer will choose first.")
		return false
	}
}

func (service *Service) getCoinTossResult() int {
	return service.util.getRandomChoice()
}

func (service *Service) showGameStatus(name string, choice int, status int) int {
	println(name + " choice: " + strconv.Itoa(choice))
	status += choice
	println("Game Status: " + strconv.Itoa(status))
	return status
}

func (service *Service) getPlayerChoice() int {
	var playerChoice int
	for {
		print("Enter your choice (1 or 2): ")
		playerChoice = service.util.getIntegerInput("Negative/Non-integer input is invalid.")
		if playerChoice == 1 || playerChoice == 2 {
			break
		}
		println("Invalid choice. Please try again.")
	}
	return playerChoice
}

// Based on the game's difficulty, computer will use this method to make it's choice
// If EASY: returns random
// If MEDIUM or HARD: returns based on winChoices list
// About winChoices, check getWinChoiceList(String) method
func (service *Service) getCompChoice(difficultyLevel string, gameStatus int) int {
	if strings.ToUpper(difficultyLevel) == "EASY" {
		return service.util.getRandomChoice()
	}

	winChoices := service.getWinChoiceList(difficultyLevel)

	for _, value := range winChoices {
		if (value - 1) == gameStatus {
			return 1
		}
		if (value - 2) == gameStatus {
			return 2
		}
	}

	return service.util.getRandomChoice()
}

// Returns the checkpoints list to win the game, based on the the difficultyLevel.
// If MEDIUM: this logic will be activated at the half of the game.
// If HARD: this logic will be activated at the start of the game.
func (service *Service) getWinChoiceList(difficultyLevel string) []int {
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
