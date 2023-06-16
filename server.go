package main

import (
	"log"
	"net/http"
	"os"
	"samroehrich/training-freaks/db"
	"samroehrich/training-freaks/graph"
	"samroehrich/training-freaks/rest"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := db.CreateConnection()
	
	if err != nil {
		log.Fatal("Unable to establish database connection...")
	}
	
	var mb int64 = 1 << 20

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}}))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{
		MaxMemory: 32 * mb,
		MaxUploadSize: 50 * mb,
	})

	srv.Use(extension.Introspection{})
	
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.HandleFunc("/upload", rest.UploadFile)
	
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	defer db.Close()
}
