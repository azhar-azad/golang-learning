package controller

import "github.com/azhar-azad/practice/game_against_computer_v2/service"

func RunGame() {
	service.ShowSinglePlayerGameRules()

	difficultyChoice := service.GetDifficultyChoice()

	switch difficultyChoice {
	case 1:
		service.GamePlay("EASY")
	case 2:
		service.GamePlay("MEDIUM")
	case 3:
		service.GamePlay("HARD")
	default:

	}
}
