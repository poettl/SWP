package gamestore

type Game struct {
	hasAvaiable   IState
	hasBought     IState
	hasInstall    IState
	isBeingPlayed IState

	currentState IState
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

	g.SetState(hasAvaiableState)
	g.hasAvaiable = hasAvaiableState
	g.hasBought = hasBoughtState
	g.hasInstall = hasInstalledState
	g.isBeingPlayed = isBeingPlayedState
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
	buyCommand := &installCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: buyCommand,
	}
	execute.exec()
}

func (v *Game) PlayGame() {
	buyCommand := &playCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: buyCommand,
	}
	execute.exec()
}

func (v *Game) StopGame() {
	buyCommand := &stopCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: buyCommand,
	}
	execute.exec()
}

func (v *Game) UninstallGame() {
	buyCommand := &uninstallCommand{
		game: v.currentState,
	}
	execute := &execute{
		command: buyCommand,
	}
	execute.exec()
}

func (v *Game) SetState(s IState) {
	v.currentState = s
}
