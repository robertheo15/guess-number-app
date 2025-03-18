package main

import (
	"fmt"

	"number-guessing-game/repository"
	"number-guessing-game/service"
)

func main() {
	repo := repository.NewGameRepository("highscores.txt")
	gameService := service.NewGameService(repo)

	for {
		fmt.Println("\nWelcome to the Number Guessing Game!")
		fmt.Println("Select a difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")
		fmt.Print("Enter your choice: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Invalid choice. Please enter a number between 1 and 3.")
			continue
		}

		difficulties := map[int]string{1: "Easy", 2: "Medium", 3: "Hard"}
		chances := map[int]int{1: 10, 2: 5, 3: 3}

		fmt.Printf("Great! You have selected %s difficulty.\n", difficulties[choice])
		gameService.PlayGame(difficulties[choice], chances[choice])

		fmt.Print("Do you want to play again? (y/n): ")
		var playAgain string
		fmt.Scan(&playAgain)
		if playAgain != "y" {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}
