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
	buyCommand := &buyCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: buyCommand,
	}
	execute.exec()
}

func (v *Game) InstallGame() {
	installCommand := &installCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: installCommand,
	}
	execute.exec()
}

func (v *Game) PlayGame() {
	playCommand := &playCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: playCommand,
	}
	execute.exec()
}

func (v *Game) StopGame() {
	stopCommand := &stopCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: stopCommand,
	}
	execute.exec()
}

func (v *Game) UninstallGame() {
	uninstallCommand := &uninstallCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: uninstallCommand,
	}
	execute.exec()
}

func (v *Game) BorrowGame() {
	borrowCommand := &borrowCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: borrowCommand,
	}
	execute.exec()
}
func (v *Game) LendGame() {
	lendCommand := &lendCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: lendCommand,
	}
	execute.exec()
}

func (v *Game) ReclaimGame() {
	reclaimCommand := &reclaimCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: reclaimCommand,
	}
	execute.exec()
}

func (v *Game) ReturnGame() {
	returnCommand := &returnCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: returnCommand,
	}
	execute.exec()
}

func (v *Game) SetState(s IState) {
	v.currentState = s
}
