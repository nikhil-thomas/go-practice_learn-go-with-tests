package poker_test

import (
	"io"
	"strings"
	"testing"

	"github.com/nikhil-thomas/go-practice_learn-go-with-tests/15_cmd-line/poker"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewPokerCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewPokerCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("do not read beyond tje first newline", func(t *testing.T) {
		in := failOnEndReader{t, strings.NewReader("Chris wins\n hello there")}

		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewPokerCLI(playerStore, in)
		cli.PlayPoker()
	})
}

type failOnEndReader struct {
	t   *testing.T
	rdr io.Reader
}

func (m failOnEndReader) Read(p []byte) (n int, err error) {
	n, err = m.rdr.Read(p)

	if n == 0 || err == io.EOF {
		m.t.Fatal("read to the end when you shouldn't have")
	}
	return n, err
}