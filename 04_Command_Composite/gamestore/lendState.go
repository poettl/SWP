package gamestore

import (
	"fmt"
)

type HasLendState struct {
	game *Game
}

func (i *HasLendState) buyGame() {
	fmt.Println("You already own this game!")
}

func (i *HasLendState) installGame() {
	fmt.Println("You already lend this game!")

}

func (i *HasLendState) playGame() {
	fmt.Println("You already lend this game!")
}

func (i *HasLendState) stopGame() {
	fmt.Println("You are not playing!")
}

func (i *HasLendState) uninstallGame() {
	fmt.Println("Game is not installed!")
}

func (i *HasLendState) borrowGame() {
	fmt.Println("You already own this game!")
}

func (i *HasLendState) lendGame() {
	fmt.Println("You already lend this game!")
}

func (i *HasLendState) reclaimGame() {
	i.game.bought = true
	hasBoughtState := &HasBoughtState{
		game: i.game,
	}
	i.game.SetState(hasBoughtState)
	fmt.Println("You got your game back!")
}

func (i *HasLendState) returnGame() {
	fmt.Println("You own this game")
}
