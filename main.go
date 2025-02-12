package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Name     string
	Points   string
	Question []Question
}

func (g *GameState) Init() {
	fmt.Println("Seja Bem vindo(a) ao quiz")
	fmt.Println(("Escreva seu nome:"))
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string")
	}

	g.Name = name

	fmt.Printf("Vamos ao jogo %s", g.Name)
}

func (g *GameState) ProcessCSV() {
	f, err := os.Open("quizgo.csv")
	if err != nil {
		panic("erro ao ler arquivo")
	}

	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Error ao ler csv")
	}

	for index, record := range records {
		fmt.Println(index, record)

		if index > 0 {
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  ToInt(record[5]),
			}

			g.Question = append(g.Question, question)
		}
	}
}

func main() {
	game := &GameState{}
	go game.ProcessCSV()
	game.Init()
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}
