package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/ubeep/go-api-structure-example/pkg"
)

// GetRootFields returns all the available queries.
func GetRootFields(gR pkg.GameRepository) graphql.Fields {
	return graphql.Fields{
		"games": GamesQuery(gR),
	}
}
