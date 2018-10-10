package pkg

// Game defines the properties of a game
type Game struct {
	ID string
	Name string
	ShortDscr string
	Genre GameGenre
	Status bool
}

// GameGenre defines the genres which can have a game
// Normally this its define with a entity but here use the type for simplicity
type GameGenre string

const (
	// AdventureGenre define the adventure genre
	AdventureGenre GameGenre = "adventure"

	// ActionGenre define the action genre
	ActionGenre ="action"

	// RolGenre define the rol genre
	RolGenre = "rol"

	// ShooterGenre define the shooter genre
	ShooterGenre = "shooter"
)