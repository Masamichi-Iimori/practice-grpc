package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Masamichi-Iimori/practice-grpc/article/client"
	"github.com/Masamichi-Iimori/practice-grpc/graph"
	"github.com/Masamichi-Iimori/practice-grpc/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	articleClient, err := client.NewClient("localhost:50051")
	log.Println(articleClient)
	if err != nil {
		log.Fatalf("Failed to create article client: %v\n", err)
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					ArticleClient: articleClient,
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
