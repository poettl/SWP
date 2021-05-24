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
	if i.game.bought {
		hasBoughtState := &HasBoughtState{
			game: i.game,
		}
		i.game.SetState(hasBoughtState)
	} else {
		hasBorrowedState := &HasBorrowedState{
			game: i.game,
		}
		i.game.SetState(hasBorrowedState)
	}

	fmt.Println("Uninstall succesful!")

}

func (i *HasInstallState) borrowGame() {
	fmt.Println("You already own this game!")
}

func (i *HasInstallState) lendGame() {
	hasLendState := &HasLendState{
		game: i.game,
	}
	i.game.SetState(hasLendState)
	fmt.Println("You lend this game!")
}

func (i *HasInstallState) reclaimGame() {
	fmt.Println("You don't lend this game!")
}

func (i *HasInstallState) returnGame() {
	if i.game.bought {
		fmt.Println("You own this game")
	} else {
		fmt.Println("Please uninstall to return the game")
	}
}
