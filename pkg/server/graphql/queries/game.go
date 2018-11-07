package queries

import (
	"github.com/friendsofgo/go-api-structure-example/pkg"
	"github.com/friendsofgo/go-api-structure-example/pkg/server/graphql/types"
	"github.com/graphql-go/graphql"
)

// GamesQuery represent the query for listing all games
func GamesQuery(gR pkg.GameRepository) *graphql.Field {
	gamesQueryField := &graphql.Field{
		Type: graphql.NewList(types.Game()),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "The Id of the game to search",
			},
			"name": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "The name of the game to search",
			},
			"genre": &graphql.ArgumentConfig{
				Type:        graphql.String,
				Description: "The genre of the games to search",
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ID := p.Args["id"].(string)
			return gR.Find(ID)
		},
	}

	return gamesQueryField
}
