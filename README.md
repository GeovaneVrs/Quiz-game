# Quiz Game em Go

Este projeto é um jogo de perguntas e respostas desenvolvido em **Go**. O jogador pode escolher entre três temas: **Matemática, História ou Perguntas Gerais**. Para ser aprovado, é necessário atingir **50 pontos**. Cada pergunta tem um limite de tempo de **60 segundos** para ser respondida.

## 🚀 Como Executar o Jogo

1. **Clone o repositório** ou copie os arquivos necessários.
2. **Crie os arquivos CSV** para armazenar as perguntas:
   - `matematica.csv`
   - `historia.csv`
   - `geral.csv`
3. **Compile e execute o jogo** com o comando:
   ```sh
   go run main.go
   ```
4. **Digite seu nome** quando solicitado.
5. **Escolha o tema do quiz** digitando o número correspondente.
6. **Responda às perguntas** dentro do tempo limite.

## 📜 Regras do Jogo

- Cada tema contém **10 perguntas**.
- O jogador tem **60 segundos** para responder cada pergunta.
- Cada resposta correta vale **10 pontos**.
- Para ser aprovado, o jogador precisa atingir **50 pontos** ou mais.
- Caso contrário, será reprovado.

## 📂 Estrutura dos Arquivos CSV
Cada arquivo CSV deve conter perguntas no seguinte formato:
```csv
Pergunta,Opção1,Opção2,Opção3,Opção4,Resposta
Qual é a capital da França?,Londres,Berlim,Paris,Madrid,3
```

## 🛠 Tecnologias Utilizadas
- **Go** para o desenvolvimento do jogo
- **CSV** para armazenar as perguntas
- **Bufio** e **time** para entrada do usuário e temporizador

## 📌 Funcionalidades Futuras
- Adicionar mais categorias de perguntas
- Implementar um sistema de pontuação avançado
- Criar um ranking de jogadores

---
Desenvolvido por Geovane 🎮

