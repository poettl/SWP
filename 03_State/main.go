package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"swp/state/gamestore"
)

func main() {
	gamestore := gamestore.NewGame()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Gamestore")
	fmt.Println("---------------------")
	fmt.Println("Available Commands:")
	fmt.Println()
	fmt.Println("buy")
	fmt.Println("install")
	fmt.Println("play")
	fmt.Println("stop")
	fmt.Println("uninstall")
	fmt.Println("exit")
	fmt.Println()

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("buy", text) == 0 {
			gamestore.BuyGame()
		}
		if strings.Compare("install", text) == 0 {
			gamestore.InstallGame()
		}
		if strings.Compare("play", text) == 0 {
			gamestore.PlayGame()
		}
		if strings.Compare("stop", text) == 0 {
			gamestore.StopGame()
		}
		if strings.Compare("uninstall", text) == 0 {
			gamestore.UninstallGame()
		}
		if strings.Compare("exit", text) == 0 {
			break
		}

	}

}
