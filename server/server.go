package main

import (
	keep_server "keep-server"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
	"github.com/mitchellh/go-homedir"
	"github.com/rs/cors"
)

const defaultPort = "1333"

func main() {
	port := defaultPort

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	router := mux.NewRouter()
	router.HandleFunc("/", handler.Playground("GraphQL playground", "/query"))
	router.HandleFunc("/graphql", handler.GraphQL(keep_server.NewExecutableSchema(keep_server.Config{Resolvers: &keep_server.Resolver{}})))
	router.HandleFunc("/query", handler.GraphQL(keep_server.NewExecutableSchema(keep_server.Config{Resolvers: &keep_server.Resolver{}})))

	home, _ := homedir.Dir()

	if err := http.ListenAndServeTLS(":"+port, home+"/ssl/myself.crt", home+"/ssl/myself.key", c.Handler(router)); err != nil {
		log.Fatal("err", err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
