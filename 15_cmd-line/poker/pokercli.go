package poker

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// PokerCLI define PokerCli struct
type PokerCLI struct {
	playerStore PlayerStore
	in          *bufio.Reader
}

func NewPokerCLI(store PlayerStore, in io.Reader) *PokerCLI {
	return &PokerCLI{
		playerStore: store,
		in:          bufio.NewReader(in),
	}
}

// PlayPoker plays poker
func (cli *PokerCLI) PlayPoker() {

	userInput, _ := cli.in.ReadString('\n')
	cli.playerStore.RecordWin(extractWinner(userInput))
	fmt.Println(cli.playerStore)
}

func extractWinner(userInput string) string {
	winner := strings.Replace(string(userInput), " wins\n", "", -1)
	return winner
}
