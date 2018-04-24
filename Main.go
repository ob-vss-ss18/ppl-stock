package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/graphql-go/graphql"
	"log"
	"encoding/json"
)

func doMagic(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello World!") // send data to client
	return
}

func main() {

	GraphQLTest();

	// get port from the environment variable (heroku sets this)
	var port string = os.Getenv("PORT");
	// or get an default value (e.g. for developing)
	if port == "" {
		port = "1337";
	}
	println("using port: " + port);

	http.HandleFunc("/", doMagic)  // map function to address
	err := http.ListenAndServe(":" + port, nil) // listen to http port not allowed
	//err := http.ListenAndServe(":1337", nil) // listen to http port not allowed
	if err != nil {
		fmt.Sprintf("An Error occured: %s\n", err) // Housten we have a problem
	}
}

func GraphQLTest() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			hello
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}