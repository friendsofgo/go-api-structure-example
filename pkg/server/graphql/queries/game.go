package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/ubeep/go-api-structure-example/pkg"
	"github.com/ubeep/go-api-structure-example/pkg/server/graphql/types"
)

// GamesQuery represent the query for listing all games
func GamesQuery(gR pkg.GameRepository) *graphql.Field {
	gamesQueryField := &graphql.Field{
		Type: graphql.NewList(types.Game()),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			game, ok := p.Source.(pkg.Game)
			if !ok {
				return nil, nil
			}
			return gR.Find(game)
		},
	}

	return gamesQueryField
}
