package gamestore

import "fmt"

type IsBeingPlayedState struct {
	game *Game
}

func (i *IsBeingPlayedState) buyGame() {
	fmt.Println("You already own this game!")
}

func (i *IsBeingPlayedState) installGame() {
	fmt.Println("Already installed!")
}

func (i *IsBeingPlayedState) playGame() {
	fmt.Println("You are already playing!")
}

func (i *IsBeingPlayedState) stopGame() {
	hasInstallState := &HasInstallState{
		game: i.game,
	}
	i.game.SetState(hasInstallState)
	fmt.Println("You left the game!")
}

func (i *IsBeingPlayedState) uninstallGame() {
	fmt.Println("You cannot uninstall while playing")
}
