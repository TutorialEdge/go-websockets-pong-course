package game

// Position
type Position struct {
	X int
	Y int
}

// Game - represents a game
type Game struct {
	Player1 Player
	Player2 Player
	Ball    Ball
}
