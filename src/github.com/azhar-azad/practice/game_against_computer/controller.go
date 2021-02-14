package main

import "fmt"

type Controller struct {
	service Service
}

func (controller *Controller) runGame() {
	fmt.Println("controller - rungame")

	controller.service.showSinglePlayerGameRules()

	difficultyChoice := controller.service.getDifficultyChoice()

	switch difficultyChoice {
	case 1:
		controller.service.gamePlay("EASY")
	case 2:
		controller.service.gamePlay("MEDIUM")
	case 3:
		controller.service.gamePlay("HARD")
	default:

	}
}
