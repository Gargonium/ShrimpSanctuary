package screens

type Screen interface {
	HandleInput()
	Draw()
	drawButtons()
}
