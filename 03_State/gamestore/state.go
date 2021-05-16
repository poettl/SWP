package gamestore

type IState interface {
	buyGame()
	installGame()
	playGame()
	stopGame()
	uninstallGame()
}
