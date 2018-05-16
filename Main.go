package main

import (
	"fmt"
	"net/http"
	"github.com/graphql-go/graphql"
	"encoding/json"
)

func main() {

	http.HandleFunc("/graphql", func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		// httpRequest.URL.Query().Get("query"):
		// returns "message" (without the qoutes) for the example below
		result := executeQuery(httpRequest.URL.Query().Get("query"))
		json.NewEncoder(httpResponseWriter).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={message}'")
	http.ListenAndServe(":8080", nil)

}

/*
	Create a query object type with field message by using GraphQLObjectTypeConfig:
		- Name: name of the type to be created
		- Fields: a map of created by using GraphQLFields
		(these are the fields that can be requested in a query)
	We use GraphQLFieldConfig to define/configure the message field:
		- Type: type of Field (graphgl.String)
		- Resolve: function which will return the 'Hello World!' message
 */
var query = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"message": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello World!", nil
				},
			},
		},
	})

/*
	Now create a new schema which has the hello world query created above as
	top level query.
 */
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: query,
	},
)


/*
	A small function to which will call graphql.Do(), which will handle resolving the
	request and returns the results.
 */
func executeQuery(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		// this variable is set to "message" for the example below
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
