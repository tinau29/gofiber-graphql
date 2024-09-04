package schema

import (
	"api/resolver"
	"fmt"

	"github.com/graphql-go/graphql"
)

var employeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"position": &graphql.Field{
			Type: graphql.String,
		},
		"address": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"updated_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"employee": &graphql.Field{
			Type: employeeType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			// Resolve: resolver.FindEmployeeByID(params),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				employee, err := resolver.FindEmployeeByID(params)
				fmt.Println("ok employee :", employee)
				return employee, err
				// 	id, ok := params.Args["id"].(string)
				// 	if !ok {
				// 		return nil, nil
				// 	}
				// 	fmt.Println("id :", id)
				// 	var employee model.Employee
				// 	if err := services.DB.Model(&employee).Where(`id = ?`, id).First(&employee).Error; nil != err {
				// 		return nil, fmt.Errorf(fmt.Sprintf(`%s`, err.Error()))
				// 	}

				// 	fmt.Println("employee query :", employee)
				// 	// err := services.DB.QueryRow("SELECT id, name, position FROM employees WHERE id = $1", id).Scan(&employee.ID, &employee.Name, &employee.Position)
				// 	// if err != nil {
				// 	// 	log.Fatal(err)
				// 	// }
				// 	return employee, nil
			},
		},
	},
})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"addEmployee": &graphql.Field{
			Type: employeeType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"position": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			// Resolve: resolver.AddEmployee(params),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				employee, err := resolver.AddEmployee(params)
				return employee, err
			},
		},
		// Define other mutations like updateEmployee, deleteEmployee here
	},
})

var EMPLOYEE, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: mutationType,
})
