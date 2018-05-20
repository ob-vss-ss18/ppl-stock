package pplStock

import "github.com/graphql-go/graphql"


var	PPLStockSchema graphql.Schema


/*
* Automatically called by go to initialize variables defined here.
*/
func init() {

	/*
	* Create a query object type with field message by using GraphQLObjectTypeConfig:
	* 	- Name: name of the type to be created
	* 	- Fields: a map of created by using GraphQLFields
	* 	(these are the fields that can be requested in a query)
	* We use GraphQLFieldConfig to define/configure the message field:
	*	- Type: type of Field (graphgl.String)
	*	- Resolve: function which will return the 'Hello World!' message
 	*/
	query := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"message": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return "Hello World!", nil
					},
				},
				"greet": &graphql.Field{
					Type: graphql.String,
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(input graphql.ResolveParams) (interface{}, error) {
						return "Hello " + input.Args["name"].(string), nil
					},
				},
			},
		})

	/*
	* Now create a new schema which has the hello world query created above as
	* top level query.
	*/
	PPLStockSchema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: query,
		},
	)

}