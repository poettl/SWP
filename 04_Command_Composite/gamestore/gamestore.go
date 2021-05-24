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
	v.currentState.buyGame()
}

func (v *Game) InstallGame() {
	v.currentState.installGame()
}

func (v *Game) PlayGame() {
	v.currentState.playGame()
}

func (v *Game) StopGame() {
	v.currentState.stopGame()
}

func (v *Game) UninstallGame() {
	v.currentState.uninstallGame()
}

func (v *Game) SetState(s IState) {
	v.currentState = s
}
