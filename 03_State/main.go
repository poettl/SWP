package main

import (
	"swp/state/gamestore"
)

func main() {
	gamestore := gamestore.NewGame()
	gamestore.BuyGame()
	gamestore.BuyGame()
	gamestore.PlayGame()
	gamestore.InstallGame()
	gamestore.PlayGame()
	gamestore.UninstallGame()

	// err = vendingMachine.insertMoney(10)

}
