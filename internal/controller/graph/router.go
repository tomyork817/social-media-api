package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"
	"social-media-api/internal/controller/graph/generated"
	"social-media-api/internal/usecase"
)

type Router struct {
	GraphQLServer *handler.Server
	Multiplexer   *http.ServeMux
}

func NewRouter(post usecase.Post, comment usecase.Comment, subscription usecase.Subscription) *Router {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: NewResolver(post, comment, subscription),
	}))
	srv.AddTransport(&transport.Websocket{})

	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	mux.Handle("/graphql", srv)

	return &Router{
		GraphQLServer: srv,
		Multiplexer:   mux,
	}
}
