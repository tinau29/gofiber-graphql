// package main

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/graphql-go/graphql"
// )

// // Define your GraphQL schema
// var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
// 	Query: graphql.NewObject(graphql.ObjectConfig{
// 		Name: "Query",
// 		Fields: graphql.Fields{
// 			"hello": &graphql.Field{
// 				Type: graphql.String,
// 				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 					return "Hello, world!", nil
// 				},
// 			},
// 		},
// 	}),
// })

// func main() {
// 	app := fiber.New()

// 	app.Post("/graphql", func(c *fiber.Ctx) error {
// 		query := c.FormValue("query")
// 		result := executeQuery(query, schema)
// 		return c.JSON(result)
// 	})

// 	app.Listen(":8000")
// }

// func executeQuery(query string, schema graphql.Schema) graphql.Result {
// 	params := graphql.Params{Schema: schema, RequestString: query}
// 	result := graphql.Do(params)
// 	return *result
// }

package main

import (
	"fmt"
	"log"

	"api/schema"
	"api/services"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

func main() {
	services.InitDB()
	app := fiber.New()

	app.Post("/graphql", func(c *fiber.Ctx) error {
		var params struct {
			Query string `json:"query"`
		}

		if err := c.BodyParser(&params); err != nil {
			return err
		}

		fmt.Println("query :", params)
		result := executeQuery(params.Query, schema.EMPLOYEE)
		return c.JSON(result)
	})

	log.Fatal(app.Listen(":8000"))
}

func executeQuery(query string, schema graphql.Schema) graphql.Result {
	params := graphql.Params{Schema: schema, RequestString: query}
	result := graphql.Do(params)
	return *result
}

// {
//     "query": "{\"mutation\": \"addEmployee\": {\"name\": \"John Doe\", \"position\": \"Software Engineer\", \"address\": \"Jakarta\"}}"
// }
