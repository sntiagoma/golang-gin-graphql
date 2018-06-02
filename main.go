package main

func main() {
	r := setupRouter()
	r.POST("/graphql", GetGraphQLGinHandler())
	r.GET("/graphql", GetGraphQLGinHandler())
	r.Run(":3001")
}
