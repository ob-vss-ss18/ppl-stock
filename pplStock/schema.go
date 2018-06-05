package pplStock

import (
	"github.com/graphql-go/graphql"
)

/*
 * This is the only exported type from this file. It defines the
 * schema of our api.
 */
var PPLStockSchema graphql.Schema

type ski struct {
	id      int
	useCase string // langlauf, piste, park
	brand   string // Fischer, Völkl
}

var exampleSki ski

/*
 * Automatically called by go to initialize variables defined here.
 */
func init() {

	// example data
	exampleSki = ski{
		id:      10,
		useCase: "Langlauf",
		brand:   "Fischer",
	}

	/*
	 * Defines how to resolve the type "skiType". Each of the
	 * resolve functions expect as input a ski struct.
	 */
	skiType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Ski",
		Description: "A ski.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the ski",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(ski); ok {
						return ski.id, nil
					}
					return nil, nil
				},
			},
			"useCase": &graphql.Field{
				Type:        graphql.String,
				Description: "The useCase (Langlauf, Piste, Park) of the ski.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(ski); ok {
						return ski.useCase, nil
					}
					return nil, nil
				},
			},
			"brand": &graphql.Field{
				Type:        graphql.String,
				Description: "The brand of the ski.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(ski); ok {
						return ski.brand, nil
					}
					return nil, nil
				},
			},
		},
	})

	/*
	 * The query type which will define the fields which can be requested
	 * from this api. This is basically the definition of the api.
	 */
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				/*
				 * If someone requests field "ski" an id must be provided with the
				 * request. The requested subselections will be resolved using the
				 * definitions is "skiType" and the input param which is given
				 * by the resolve function.
				 *
				 * This could be a good spot to load the data from the database into
				 * a local struct. That is why there is a method called "loadSkiFrom
				 * Database".
				 */
				"Ski": &graphql.Field{
					Type: skiType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Description: "Id of the requested ski.",
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
						return loadSkiFromDatabase(parameter.Args["id"].(int)), nil
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
			Query: queryType,
		},
	)

}
/*
 * Dummy method this could load a ski given by id from a database.
 */
func loadSkiFromDatabase(id int) ski {
	result := ski{}
	if (exampleSki.id == id) {
		result = exampleSki
	}
	return result
}
