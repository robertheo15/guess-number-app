package repository

import (
	"fmt"
	"os"
)

type GameRepository interface {
	SaveHighScore(difficulty string, attempts int)
	GetHighScore(difficulty string) int
}

type fileGameRepository struct {
	filePath string
}

func NewGameRepository(filePath string) GameRepository {
	return &fileGameRepository{filePath: filePath}
}

func (r *fileGameRepository) SaveHighScore(difficulty string, attempts int) {
	highScore := r.GetHighScore(difficulty)
	if highScore == 0 || attempts < highScore {
		file, err := os.Create(r.filePath)
		if err != nil {
			fmt.Println("Error saving high score:", err)
			return
		}
		defer file.Close()
		_, _ = file.WriteString(fmt.Sprintf("%s:%d\n", difficulty, attempts))
	}
}

func (r *fileGameRepository) GetHighScore(difficulty string) int {
	file, err := os.Open(r.filePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	var storedDifficulty string
	var score int
	for {
		_, err := fmt.Fscanf(file, "%s:%d\n", &storedDifficulty, &score)
		if err != nil {
			break
		}
		if storedDifficulty == difficulty {
			return score
		}
	}
	return 0
}
