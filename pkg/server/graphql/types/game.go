package types

import (
	"github.com/graphql-go/graphql"
)

const GameTypeName = "Game"

// GameType define the Graphql structure for Game entity
func Game() *graphql.Object {

	gameType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: GameTypeName,
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "The id of the game",
				},
				"name": &graphql.Field{
					Type:        graphql.String,
					Description: "The name of the game",
				},
				"description": &graphql.Field{
					Type:        graphql.String,
					Description: "The description of the game",
				},
				"genre": &graphql.Field{
					Type:        graphql.String,
					Description: "The genre of the game",
				},
				"status": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "The status of the game if is active or inactive",
				},
			},
		},
	)

	return gameType
}
