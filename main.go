package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/parakeety/typing-game/typing"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := input(os.Stdin)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	score := quiz(ctx, ch)
	fmt.Printf("Time up! Your score is: %d", score)
}

func quiz(ctx context.Context, ch <-chan string) (score int) {
	for {
		i := rand.Intn(len(typing.WordList) - 1)
		word := typing.WordList[i]
		fmt.Println(word)

		select {
		case typed := <-ch:
			if typed == word {
				score++
				fmt.Printf("correct! Your current score is %d \n", score)
				fmt.Println(">>>")
			}

		case <-ctx.Done():
			return
		}
	}
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
