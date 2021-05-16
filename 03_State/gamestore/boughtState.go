package gamestore

import "fmt"

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
