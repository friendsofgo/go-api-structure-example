package pkg

// Game defines the properties of a game
type Game struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Genre       GameGenre `json:"genre"`
	Status      bool      `json:"status"`
}

// GameGenre defines the genres which can have a game
// Normally this its define with a entity but here use the type for simplicity
type GameGenre string

const (
	// AdventureGenre define the adventure genre
	AdventureGenre GameGenre = "adventure"

	// ActionGenre define the action genre
	ActionGenre = "action"

	// RolGenre define the rol genre
	RolGenre = "rol"

	// ShooterGenre define the shooter genre
	ShooterGenre = "shooter"
)

// GameRepository provides access to game repository
type GameRepository interface {
	Find(ID string) ([]*Game, error)
}
