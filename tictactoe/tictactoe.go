package main

import "tictactoe/board"

const welcomeStr = "Welcome! Here is your board:\n\n"
const p1 = "Player1"
const p2 = "Player2"

func main() {
	print(welcomeStr)
	testBoard := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
	print(board.ToString(testBoard))
}
