package gamestore

import (
	"fmt"
)

type HasBorrowedState struct {
	game *Game
}

func (i *HasBorrowedState) buyGame() {
	i.game.bought = true
	hasBoughtState := &HasBoughtState{
		game: i.game,
	}
	i.game.SetState(hasBoughtState)
	fmt.Println("You bought this game!")
}

func (i *HasBorrowedState) installGame() {
	hasInstallState := &HasInstallState{
		game: i.game,
	}
	i.game.SetState(hasInstallState)
	fmt.Println("Game is being installed.")
}

func (i *HasBorrowedState) playGame() {
	fmt.Println("You need to buy and install first!")
}

func (i *HasBorrowedState) stopGame() {
	fmt.Println("You are not playing!")
}

func (i *HasBorrowedState) uninstallGame() {
	fmt.Println("Game is not installed!")
}

func (i *HasBorrowedState) borrowGame() {
	fmt.Println("You already borrowed this game!")
}

func (i *HasBorrowedState) lendGame() {
	fmt.Println("You don't own this game!")
}

func (i *HasBorrowedState) reclaimGame() {
	fmt.Println("You don't own this game!")
}

func (i *HasBorrowedState) returnGame() {
	hasAvaiableState := &HasAvaiableState{
		game: i.game,
	}
	i.game.SetState(hasAvaiableState)
	fmt.Println("Game is being returned.")
}
