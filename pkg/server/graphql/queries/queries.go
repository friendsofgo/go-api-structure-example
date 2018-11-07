package queries

import (
	"github.com/friendsofgo/go-api-structure-example/pkg"
	"github.com/graphql-go/graphql"
)

// GetRootFields returns all the available queries.
func GetRootFields(gR pkg.GameRepository) graphql.Fields {
	return graphql.Fields{
		"games": GamesQuery(gR),
	}
}
