package main

import (
	"fmt"
	"strconv"
	"strings"
)

type square struct {
	loc   string
	color string
}

func main() {
	board := generateBoard()
	mainMenu(board)
}

func generateBoard() [8][8]square {
	var color, loc string
	var board [8][8]square
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			color = "l"
			if i%2 == j%2 {
				color = "d"
			}

			loc = letters[i] + strconv.Itoa(j+1)
			board[i][j] = square{loc, color}
		}
	}
	return board
}

func mainMenu(board [8][8]square) {

	intro := `Welcome to chessmem
main menu:
	1. color guess
	2. move to 
	m. this menu
	q. quit
`
loop:
	for {
		fmt.Println(intro)
		fmt.Print("-> ")
		var text string
		fmt.Scanf("%s", &text)
		switch {
		case strings.Compare("q", text) == 0:
			break loop
		case strings.Compare("1", text) == 0:
			ColorGuess(board)
			break
		case strings.Compare("2", text) == 0:
			MoveTo(board)
			break
		}

	}
}
