package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/ob-vss-ss18/ppl-stock/models"
)

/*
 * This is the only exported type from this file. It defines the
 * schema of our api.
 */
var PPLStockSchema graphql.Schema

var exampleSki models.Ski

/*
 * Automatically called by go to initialize variables defined here.
 */
func init() {

	// example data
	exampleSki = models.Ski{
		Id: 10,
		Usage: models.Langlauf,
		Category: models.Beginner,
		Usertype: models.Erwachsener,
		Gender: models.Male,
		Manufactorer: "Fischer",
		Model: "Super Ski 3000",
		Length: 2,
		Bodyheight: 3,
		Bodyweight: 4,
		Color: "Rot",
		PriceNew: 34.99,
		Condition: models.New,
		Status: models.Available,
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
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Id, nil
					}
					return nil, nil
				},
			},
			"usage": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The use case of the ski",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Usage, nil
					}
					return nil, nil
				},
			},
			"category": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The category of the ski",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Category, nil
					}
					return nil, nil
				},
			},
			"usertype": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The usertype of the ski",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Usertype, nil
					}
					return nil, nil
				},
			},
			"gender": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The gender by which the ski is intended to be used.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Gender, nil
					}
					return nil, nil
				},
			},
			"manufactorer": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The manufactorer of the ski.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Manufactorer, nil
					}
					return nil, nil
				},
			},
			"model": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The model of the ski.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Model, nil
					}
					return nil, nil
				},
			},
			"length": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The length of the ski.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Length, nil
					}
					return nil, nil
				},
			},
			"bodyheight": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The best bodyheight for using this ski.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Bodyheight, nil
					}
					return nil, nil
				},
			},
			"bodyweight": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The best bodyweight for using this ski.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Bodyweight, nil
					}
					return nil, nil
				},
			},
			"color": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The color of the ski.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Color, nil
					}
					return nil, nil
				},
			},
			"price_new": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Float),
				Description: "The new price of the ski.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.PriceNew, nil
					}
					return nil, nil
				},
			},
			"condition": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The condition of the ski.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Condition, nil
					}
					return nil, nil
				},
			},
			"availability": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The status/availability of the ski.",
				Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
					if ski, ok := parameter.Source.(models.Ski); ok {
						return ski.Status, nil
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
func loadSkiFromDatabase(id int) models.Ski {
	result := models.Ski{}
	if exampleSki.Id == uint32(id) {
		result = exampleSki
	}
	return result
}
