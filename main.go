package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type Questions struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Questions []Questions
}

func (g *GameState) ProcessCSV() {
	f, err := os.Open("quiz.csv")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	for index, record := range records {
		if index > 0 {
			correctAnswer, _ := toInt(record[5])
			question := Questions{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}
			g.Questions = append(g.Questions, question)
		}
	}
}

func main() {
	game := &GameState{}
	game.ProcessCSV()
}

func toInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	}

	return i, nil
}
