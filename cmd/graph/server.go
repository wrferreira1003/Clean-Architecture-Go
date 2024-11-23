package graph

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/infra/graph"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/usecase"
)

type Resolver struct {
	CreateOrderUseCase usecase.OrderUseCaseInterface
	ListOrderUseCase   usecase.ListOrderUseCaseInterface
}

func NewServer(createOrderUseCase usecase.OrderUseCaseInterface, listOrderUseCase usecase.ListOrderUseCaseInterface) {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		OrderUseCase:     createOrderUseCase,
		ListOrderUseCase: listOrderUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:8082/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8082", nil)) // Certifique-se de que está assim

}
