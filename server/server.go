package main

import (
	keep_server "keep-server"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/mitchellh/go-homedir"
)

const defaultPort = "1333"

func main() {
	port := defaultPort

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(keep_server.NewExecutableSchema(keep_server.Config{Resolvers: &keep_server.Resolver{}})))

	home, _ := homedir.Dir()

	err := http.ListenAndServeTLS(":"+port, home+"/ssl/myself.crt", home+"/ssl/myself.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
