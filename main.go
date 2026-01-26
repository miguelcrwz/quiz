package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Questions struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Name      string
	Questions []Questions
}

func (g *GameState) Init() {
	fmt.Println("\033[35;1mWELCOME TO F1 QUIZ! ⚒︎\033[0m")
	time.Sleep(time.Millisecond * 350)
	fmt.Printf("\033[37;1m Enter your NAME: \033[0m")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	name = strings.TrimSpace(name)
	g.Name = name
	time.Sleep(time.Millisecond * 350)
	fmt.Printf("\033[35;1mHello, %s! LET'S START!\033[0m\n", g.Name)
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
	game.Init()
}

func toInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	}

	return i, nil
}
