package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func GetGraphQLSchema() (graphql.Schema, error) {
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
	return graphql.NewSchema(schemaConfig)
}

func ExecuteGraphQLQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	return result
}

func GetGraphQLConfig() handler.Config {
	Schema, _ := GetGraphQLSchema()
	return handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	}
}

func GetGraphQLHTTPHandler() *handler.Handler {
	config := GetGraphQLConfig()
	return handler.New(&config)
}

func GetGraphQLGinHandler() gin.HandlerFunc {
	return gin.WrapH(GetGraphQLHTTPHandler())
}
