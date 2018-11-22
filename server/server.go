package main

import (
	log "log"
	http "net/http"
	os "os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	handler "github.com/99designs/gqlgen/handler"
	"github.com/lisiur/daydayup/graph/generated"
	"github.com/lisiur/daydayup/resolver"
)

const defaultPort = "8080"

var db *gorm.DB

func main() {
	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
