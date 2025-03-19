package service_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

// Mock Repository
type MockGameRepository struct {
	highScores map[string]int
}

func (m *MockGameRepository) SaveHighScore(difficulty string, attempts int) {
	if m.highScores == nil {
		m.highScores = make(map[string]int)
	}
	if _, exists := m.highScores[difficulty]; !exists || attempts < m.highScores[difficulty] {
		m.highScores[difficulty] = attempts
	}
}

func (m *MockGameRepository) GetHighScore(difficulty string) int {
	if val, exists := m.highScores[difficulty]; exists {
		return val
	}
	return 0
}

func TestPlayGame(t *testing.T) {
	mockRepo := &MockGameRepository{}
	gameService := NewGam(mockRepo)

	// Mock random number generator for predictable testing
	rand.Seed(42)
	targetNumber := rand.Intn(100) + 1 // Simulate known target

	// Mock user input
	input := fmt.Sprintf("%d\n", targetNumber) // User guesses correctly on first attempt
	output := new(bytes.Buffer)

	// Redirect standard input/output
	oldStdout := *outputclear

	defer func() { *output = oldStdout }()

	// Run the game
	go func() {
		gameService.PlayGame("Easy", 10)
	}()

	// Validate output
	if !strings.Contains(output.String(), "🎉 Correct!") {
		t.Errorf("Expected success message, but got: %s", output.String())
	}

	// Check if high score was saved correctly
	highScore := mockRepo.GetHighScore("Easy")
	if highScore != 1 {
		t.Errorf("Expected high score 1, but got %d", highScore)
	}
}
