package gamestore

import "fmt"

type HasInstallState struct {
	game *Game
}

func (i *HasInstallState) buyGame() {
	fmt.Println("You already own this game!")
}

func (i *HasInstallState) installGame() {
	fmt.Println("Already installed!")
}

func (i *HasInstallState) playGame() {
	isBeingPlayed := &IsBeingPlayedState{
		game: i.game,
	}
	i.game.SetState(isBeingPlayed)
	fmt.Println("Starting game!")
}

func (i *HasInstallState) stopGame() {
	fmt.Println("You are not playing!")
}

func (i *HasInstallState) uninstallGame() {
	hasBoughtState := &HasBoughtState{
		game: i.game,
	}
	i.game.SetState(hasBoughtState)
	fmt.Println("Uninstall succesful!")
}
