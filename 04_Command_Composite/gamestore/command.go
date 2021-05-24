package gamestore

type execute struct {
	command command
}

func (b *execute) exec() {
	b.command.execute()
}

type command interface {
	execute()
}

type buyCommand struct {
	game IState
}

type playCommand struct {
	game IState
}

type installCommand struct {
	game IState
}

type stopCommand struct {
	game IState
}

type uninstallCommand struct {
	game IState
}

func (c *buyCommand) execute() {
	c.game.buyGame()
}

func (c *playCommand) execute() {
	c.game.playGame()
}

func (c *installCommand) execute() {
	c.game.installGame()
}

func (c *stopCommand) execute() {
	c.game.stopGame()
}

func (c *uninstallCommand) execute() {
	c.game.uninstallGame()
}
