package game

// Player struct
type Player struct {
	Pos   Position
	Score int
}

// Position
type Position struct {
	X int
	Y int
}

// UpdatePosition - updates the players position on the map
func (p *Player) UpdatePosition(x, y int) {
	p.Pos.X += x
	p.Pos.Y += y
}

// GetScore - returns the players current score
func (p *Player) GetScore() int {
	return p.Score
}
