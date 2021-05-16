package gamestore

import "fmt"

type HasAvaiableState struct {
	game *Game
}

func (i *HasAvaiableState) buyGame() {
	hasBoughtState := &HasBoughtState{
		game: i.game,
	}
	i.game.SetState(hasBoughtState)
	fmt.Println("You bought this game!")
}

func (i *HasAvaiableState) installGame() {
	fmt.Println("You need to buy first!")
}

func (i *HasAvaiableState) playGame() {
	fmt.Println("You need to buy and install first!")
}

func (i *HasAvaiableState) stopGame() {
	fmt.Println("You are not playing!")
}

func (i *HasAvaiableState) uninstallGame() {
	fmt.Println("Game is not installed!")
}
