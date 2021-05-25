package gamestore

type Game struct {
	bought        bool
	hasAvaiable   IState
	hasBought     IState
	hasInstall    IState
	isBeingPlayed IState
	hasBorrowed   IState
	hasLend       IState
	currentState  IState
}

func NewGame() *Game {
	g := &Game{}
	hasAvaiableState := &HasAvaiableState{
		game: g,
	}
	hasBoughtState := &HasBoughtState{
		game: g,
	}
	hasInstalledState := &HasInstallState{
		game: g,
	}
	isBeingPlayedState := &IsBeingPlayedState{
		game: g,
	}

	hasBorrowedState := &HasBorrowedState{
		game: g,
	}

	hasLendState := &HasLendState{
		game: g,
	}

	g.SetState(hasAvaiableState)
	g.hasAvaiable = hasAvaiableState
	g.hasBought = hasBoughtState
	g.hasInstall = hasInstalledState
	g.isBeingPlayed = isBeingPlayedState
	g.hasBorrowed = hasBorrowedState
	g.hasLend = hasLendState
	return g
}

func (v *Game) BuyGame() {
	executable := getExecutable(v.currentState, "buy")
	executable.exec()
}

func (v *Game) InstallGame() {
	executable := getExecutable(v.currentState, "install")
	executable.exec()
}

func (v *Game) PlayGame() {
	executable := getExecutable(v.currentState, "play")
	executable.exec()
}

func (v *Game) StopGame() {
	executable := getExecutable(v.currentState, "stop")
	executable.exec()
}

func (v *Game) UninstallGame() {
	executable := getExecutable(v.currentState, "uninstall")
	executable.exec()
}

func (v *Game) BorrowGame() {
	executable := getExecutable(v.currentState, "borrow")
	executable.exec()
}
func (v *Game) LendGame() {
	executable := getExecutable(v.currentState, "lend")
	executable.exec()
}

func (v *Game) ReclaimGame() {
	executable := getExecutable(v.currentState, "reclaim")
	executable.exec()
}

func (v *Game) ReturnGame() {
	executable := getExecutable(v.currentState, "return")
	executable.exec()
}

func (v *Game) OneClickPlay() {
	commands := []string{"buy", "install", "play"}
	execList(commands, v)
}

func (v *Game) SetState(s IState) {
	v.currentState = s
}

func getExecutable(game IState, mode string) execute {
	var genExecute execute
	switch mode {
	case "buy":
		buyCommand := &buyCommand{
			game: game,
		}
		genExecute = execute{
			command: buyCommand,
		}
	case "install":
		installCommand := &installCommand{
			game: game,
		}
		genExecute = execute{
			command: installCommand,
		}
	case "play":
		playCommand := &playCommand{
			game: game,
		}
		genExecute = execute{
			command: playCommand,
		}
	case "stop":
		stopCommand := &stopCommand{
			game: game,
		}
		genExecute = execute{
			command: stopCommand,
		}
	case "uninstall":
		uninstallCommand := &uninstallCommand{
			game: game,
		}
		genExecute = execute{
			command: uninstallCommand,
		}
	case "borrow":
		borrowCommand := &borrowCommand{
			game: game,
		}
		genExecute = execute{
			command: borrowCommand,
		}
	case "lend":
		lendCommand := &lendCommand{
			game: game,
		}
		genExecute = execute{
			command: lendCommand,
		}
	case "reclaim":
		reclaimCommand := &reclaimCommand{
			game: game,
		}
		genExecute = execute{
			command: reclaimCommand,
		}
	case "return":
		returnCommand := &returnCommand{
			game: game,
		}
		genExecute = execute{
			command: returnCommand,
		}
	}
	return genExecute
}
