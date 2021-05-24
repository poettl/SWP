package gamestore

import (
	"fmt"
	"math/rand"
)

type HasAvaiableState struct {
	game *Game
}

func (i *HasAvaiableState) buyGame() {
	i.game.bought = true
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

func (i *HasAvaiableState) borrowGame() {
	random := rand.Intn(2 - 0)
	if random == 1 {
		i.game.bought = false
		hasBorrowedState := &HasBorrowedState{
			game: i.game,
		}
		i.game.SetState(hasBorrowedState)
		fmt.Println("You borrowed this game!")
	} else {
		fmt.Println("Request was rejected!")

	}
}

func (i *HasAvaiableState) lendGame() {
	fmt.Println("You don't own this game!")
}

func (i *HasAvaiableState) reclaimGame() {
	fmt.Println("You don't own this game!")
}

func (i *HasAvaiableState) returnGame() {
	fmt.Println("You don't borrowed this game!")
}
