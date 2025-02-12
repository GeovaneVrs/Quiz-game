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

type Question struct {
	Text    string
	Options []string
	Answer  int
}

type GameState struct {
	Name     string
	Points   int
	Question []Question
}

func (g *GameState) Init() {
	fmt.Println("Seja Bem-vindo(a) ao quiz")
	fmt.Print("Escreva seu nome: ")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		panic("Erro ao ler o nome")
	}

	g.Name = strings.TrimSpace(name)

	fmt.Println("Escolha o tema do quiz:")
	fmt.Println("1 - Matemática")
	fmt.Println("2 - História")
	fmt.Println("3 - Perguntas Gerais")
	fmt.Print("Digite o número do tema escolhido: ")

	var choice string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice = strings.TrimSpace(scanner.Text())

	var fileName string
	switch choice {
	case "1":
		fileName = "./csv/matematica.csv"
	case "2":
		fileName = "./csv/historia.csv"
	case "3":
		fileName = "./csv/geral.csv"
	default:
		fmt.Println("Escolha inválida, saindo do jogo.")
		os.Exit(1)
	}

	fmt.Printf("Vamos ao jogo, %s!\n\n", g.Name)
	g.ProcessCSV(fileName)
}

func (g *GameState) ProcessCSV(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic("Erro ao ler arquivo CSV")
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler CSV")
	}

	for index, record := range records {
		if index > 0 {
			correctAnswer, _ := ToInt(record[5])
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}
			g.Question = append(g.Question, question)
		}
	}
}

func (g *GameState) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for index, question := range g.Question {
		fmt.Printf("\033[33m%d. %s\033[0m\n", index+1, question.Text)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Print("Digite uma alternativa (você tem 60 segundos): ")

		answerCh := make(chan int)
		timeout := time.After(60 * time.Second)

		go func() {
			for {
				scanner.Scan()
				input := strings.TrimSpace(scanner.Text())
				answer, err := ToInt(input)
				if err != nil {
					fmt.Println(err.Error())
					fmt.Print("Digite um número válido: ")
					continue
				}
				answerCh <- answer
				return
			}
		}()

		var answer int
		select {
		case answer = <-answerCh:
		case <-timeout:
			fmt.Println("\nTempo esgotado! Resposta errada.")
			answer = -1
		}

		if answer == question.Answer {
			fmt.Println("Parabéns, você acertou!!")
			g.Points += 10
		} else {
			fmt.Println("Ops! Errou!")
		}

		fmt.Println("------------------")
	}

	fmt.Printf("Pontuação final: %d\n", g.Points)
	if g.Points >= 50 {
		fmt.Println("Parabéns, você foi APROVADO!")
	} else {
		fmt.Println("Infelizmente, você foi REPROVADO.")
	}
}

func main() {
	game := &GameState{}
	game.Init()
	game.Run()
}

func ToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("entrada inválida, digite apenas números")
	}
	return i, nil
}
