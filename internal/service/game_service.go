package service

import (
	"fmt"
	"guess-number-app/internal/repository"
	"math/rand"
	"time"
)

type GameService struct {
	repo repository.GameRepository
}

func NewGameService(repo repository.GameRepository) *GameService {
	return &GameService{repo: repo}
}

func (s *GameService) PlayGame(difficulty string, chances int) {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	attempts := 0
	startTime := time.Now()

	fmt.Printf("Guess the number between 1 and 100. You have %d chances!\n", chances)
	for attempts < chances {
		fmt.Print("Enter your guess: ")
		var guess int
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		attempts++
		if guess == target {
			duration := time.Since(startTime)
			fmt.Printf("🎉 Correct! You guessed the number in %d attempts and %.2f seconds.\n", attempts, duration.Seconds())
			s.repo.SaveHighScore(difficulty, attempts)
			return
		} else if guess < target {
			fmt.Println("Incorrect! The number is greater.")
		} else {
			fmt.Println("Incorrect! The number is smaller.")
		}
	}
	fmt.Printf("❌ Game Over! The correct number was %d.\n", target)
}
