package main

import "fmt"

type ChessBoard struct {
	x rune
	y int
}

func (b ChessBoard) get_golor() string {
	if b.y%2 == 1 && int(b.x)%2 == 1 {
		return "BLACK"
	}
	if b.y%2 == 0 && int(b.x)%2 == 0 {
		return "BLACK"
	}
	return "WHITE"
}

func main() {
	board := ChessBoard{}
	fmt.Scanf("%c%d", &board.x, &board.y)
	fmt.Print(board.get_golor())
}
