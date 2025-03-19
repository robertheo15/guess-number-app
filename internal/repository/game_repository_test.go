package repository_test

import (
	"guess-number-app/internal/repository"
	"os"
	"testing"
)

func TestGameRepository(t *testing.T) {
	// Create a temporary file for testing
	testFile := "test_highscores.txt"
	defer os.Remove(testFile) // Cleanup after the test

	repo := repository.NewGameRepository(testFile)

	// Test saving a high score
	repo.SaveHighScore("Easy", 5)
	savedScore := repo.GetHighScore("Easy")
	if savedScore != 5 {
		t.Errorf("Expected high score 5, got %d", savedScore)
	}

	// Test updating a lower high score
	repo.SaveHighScore("Easy", 3) // Lower attempts should replace the previous score
	savedScore = repo.GetHighScore("Easy")
	if savedScore != 3 {
		t.Errorf("Expected high score 3, got %d", savedScore)
	}

	// Test not updating a higher high score
	repo.SaveHighScore("Easy", 7) // Higher attempts should NOT replace the previous score
	savedScore = repo.GetHighScore("Easy")
	if savedScore != 3 {
		t.Errorf("Expected high score 3, got %d", savedScore)
	}

	// Test retrieving high score for a difficulty that has not been set
	savedScore = repo.GetHighScore("Medium")
	if savedScore != 0 {
		t.Errorf("Expected high score 0 for Medium, got %d", savedScore)
	}
}
