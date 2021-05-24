package gamestore

import (
	"fmt"
)

type HasBoughtState struct {
	game *Game
}

func (i *HasBoughtState) buyGame() {
	fmt.Println("You already own this game!")
}

func (i *HasBoughtState) installGame() {
	hasInstallState := &HasInstallState{
		game: i.game,
	}
	i.game.SetState(hasInstallState)
	fmt.Println("Game is being installed.")
}

func (i *HasBoughtState) playGame() {
	fmt.Println("You need to install first!")
}

func (i *HasBoughtState) stopGame() {
	fmt.Println("You are not playing!")
}

func (i *HasBoughtState) uninstallGame() {
	fmt.Println("Game is not installed!")
}

func (i *HasBoughtState) borrowGame() {
	fmt.Println("You already own this game!")
}

func (i *HasBoughtState) lendGame() {
	hasLendState := &HasLendState{
		game: i.game,
	}
	i.game.SetState(hasLendState)
	fmt.Println("You lend this game!")
}

func (i *HasBoughtState) reclaimGame() {
	fmt.Println("You don't lend this game!")
}

func (i *HasBoughtState) returnGame() {
	fmt.Println("You  own this game!")
}
