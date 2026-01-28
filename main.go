package main

import (
	"bufio"
	"encoding/csv"
	"errors"
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
	Points    int
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

func (g *GameState) Run() {
	for index, question := range g.Questions {
		time.Sleep(time.Millisecond * 350)
		fmt.Printf("\033[33;1m\n %d. %s\033[0m\n", index+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Printf("\033[33;1m Enter an ALTERNATIVE: \033[0m")

		var answer int
		var err error

		for {
			time.Sleep(time.Millisecond * 350)
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')

			answer, err = toInt(read[:len(read)-1])

			if err != nil {
				fmt.Println(err.Error())
				fmt.Printf("\n\033[33;1m Enter an ALTERNATIVE: \033[0m")
				continue
			}

			if answer < 1 || answer > len(question.Options) {
				fmt.Println("\n\033[31;1m✘ INVALID OPTION, PLEASE SELECT A VALID OPTION.\033[0m")
				fmt.Printf("\n\033[33;1m Enter an ALTERNATIVE: \033[0m")
				continue
			}

			break
		}

		if answer == question.Answer {
			time.Sleep(time.Millisecond * 350)
			fmt.Print("\n\033[35;1m✔ CORRECT ANSWER\033[0m\n")
			g.Points++
		} else {
			time.Sleep(time.Millisecond * 350)
			fmt.Printf("\n\033[31;1m✘ INCORRECT ANSWER\033[0m / \033[35;1mRIGHT ANSWER: [%d] %s\033[0m\n", question.Answer, question.Options[question.Answer-1])
		}
	}
}

func main() {
	game := &GameState{}
	game.ProcessCSV()
	game.Init()
	time.Sleep(time.Millisecond * 350)
	game.Run()
}

func toInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, errors.New("\n\033[31;1m✘ INVALID VALUE, PLEASE SELECT A VALID OPTION.\033[0m")
	}

	return i, nil
}
