package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	log "github.com/sirupsen/logrus"
	"github.com/ubeep/go-api-structure-example/pkg"
	"github.com/ubeep/go-api-structure-example/pkg/server/graphql/queries"
	"github.com/ubeep/graphiql"
)

// Server holds the dependencies for a HTTP server
type Server struct {
	gR     pkg.GameRepository
	logger *log.Logger
	router *mux.Router
}

// New returns a new HTTP server
func New(gR pkg.GameRepository, logger *log.Logger) *Server {
	s := &Server{
		gR:     gR,
		logger: logger,
	}

	graphiQL, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		s.logger.WithFields(
			log.Fields{
				"method": "server.New",
				"err":    err,
			},
		).Error("Error trying to launch graphiql client")
	}

	r := mux.NewRouter()
	r.Use(accessControl)
	r.HandleFunc("/graphiql", graphiQL.ServeHTTP)
	s.GraphQLServer(r)
	s.router = r

	return s
}

// ServeHTTP dispatches the handler registered in the matched route
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		h.ServeHTTP(w, r)
	})
}

func (s *Server) GraphQLServer(r *mux.Router) {
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Query",
			Fields: queries.GetRootFields(s.gR),
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	r.Handle("/graphql", h)
}
