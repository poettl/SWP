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

func (i *IsBeingPlayedState) borrowGame() {
	fmt.Println("You already own this game!")
}

func (i *IsBeingPlayedState) lendGame() {
	fmt.Println("It's not possible to lend games while playing")
}

func (i *IsBeingPlayedState) reclaimGame() {
	fmt.Println("You don't lend this game!")
}

func (i *IsBeingPlayedState) returnGame() {
	if i.game.bought {
		fmt.Println("You own this game")
	} else {
		fmt.Println("Please stop and uninstall it to return the game")
	}
}
