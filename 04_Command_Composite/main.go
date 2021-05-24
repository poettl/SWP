package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"swp/command/gamestore"
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
	fmt.Println("borrow")
	fmt.Println("lend")
	fmt.Println("reclaim")
	fmt.Println("return")
	fmt.Println("exit")
	fmt.Println()

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
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
		if strings.Compare("borrow", text) == 0 {
			gamestore.BorrowGame()
		}
		if strings.Compare("lend", text) == 0 {
			gamestore.LendGame()
		}
		if strings.Compare("reclaim", text) == 0 {
			gamestore.ReclaimGame()
		}
		if strings.Compare("return", text) == 0 {
			gamestore.ReturnGame()
		}
		if strings.Compare("exit", text) == 0 {
			break
		}

	}

}
