package schema

import (
"github.com/graphql-go/graphql"
"github.com/ob-vss-ss18/ppl-stock/models"
)

var stickType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Stick",
	Description: "A stick.",

	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The id of the stick",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.Id, nil
				}
				return nil, nil
			},
		},
		"usage": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The use case of the stick",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.Usage, nil
				}
				return nil, nil
			},
		},
		"usertype": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The usertype of the ski",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.Usertype, nil
				}
				return nil, nil
			},
		},
		"gender": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The gender by which the ski is intended to be used.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.Gender, nil
				}
				return nil, nil
			},
		},
		"manufactorer": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The manufactorer of the ski.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.Manufactorer, nil
				}
				return nil, nil
			},
		},
		"modell": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The model of the ski.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.Model, nil
				}
				return nil, nil
			},
		},
		"length": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The length of the ski.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.Length, nil
				}
				return nil, nil
			},
		},
		"bodyheight": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "The best bodyheight for using this ski.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.Bodyheight, nil
				}
				return nil, nil
			},
		},
		"color": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The color of the ski.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.Color, nil
				}
				return nil, nil
			},
		},
		"price_new": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "The new price of the ski.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.PriceNew, nil
				}
				return nil, nil
			},
		},
		"condition": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The condition of the ski.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.Condition, nil
				}
				return nil, nil
			},
		},
		"availability": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The status/availability of the ski.",
			Resolve: func(parameter graphql.ResolveParams) (interface{}, error) {
				if ski, ok := parameter.Source.(models.Stick); ok {
					return ski.Status, nil
				}
				return nil, nil
			},
		},
	},
})

