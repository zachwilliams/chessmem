package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// ColorGuess command prompt game.
// requres a valid chess board [8][8]square and prompts user to guess colors
// valiudates and keeps track of score
func ColorGuess(board [8][8]square) {
	instructions := "guess the correct color d=dark, l=light, s=score"
	fmt.Println(instructions)
	score, count := 0, 0
	var text string
	var s square
	redo := false
loop:
	for {
		if !redo {
			s = getRandomSquare(board)
		}
		redo = false
		fmt.Print(s.loc + " -> ")
		fmt.Scanf("%s", &text)

		switch {
		case strings.Compare(s.color, text) == 0:
			score++
			count++
			break
		case strings.Compare("m", text) == 0:
			fmt.Printf("score: %v/%v\n", score, count)
			break loop
		case strings.Compare("s", text) == 0:
			fmt.Printf("score: %v/%v\n", score, count)
			redo = true
			break
		case strings.Compare("q", text) == 0:
			os.Exit(0)
		default:
			fmt.Printf(":(\n")
			count++
		}
	}
}

func getRandomSquare(board [8][8]square) square {
	return board[rand.Intn(8)][rand.Intn(8)]
}
