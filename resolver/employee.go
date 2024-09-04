package resolver

import (
	"api/lib"
	"api/model"
	"api/services"
	"fmt"

	"github.com/graphql-go/graphql"
)

func AddEmployee(params graphql.ResolveParams) (interface{}, error) {
	name := params.Args["name"].(string)
	fmt.Println("name :", name)
	position := params.Args["position"].(string)
	fmt.Println("position :", position)
	address := params.Args["address"].(string)

	fmt.Println("address : ", address)

	employee := model.Employee{
		Name:     &name,
		Position: &position,
		Address:  &address,
	}

	db := services.InitDB()

	if err := db.Create(&employee).Error; nil != err {
		fmt.Println("error disini :", err.Error())
		return nil, fmt.Errorf(fmt.Sprintf(`%s`, err.Error()))
	}

	var emp interface{}
	lib.Merge(&employee, &emp)
	return emp, nil
}

func FindEmployeeByID(params graphql.ResolveParams) (interface{}, error) {
	db := services.InitDB()

	id, ok := params.Args["id"].(string)
	if !ok {
		return nil, nil
	}
	fmt.Println("id :", id)
	var employee model.Employee
	if err := db.Model(&employee).Where(`id = ?`, id).First(&employee).Error; nil != err {
		return nil, fmt.Errorf(fmt.Sprintf(`%s`, err.Error()))
	}

	fmt.Println("employee query :", employee)
	var emp interface{}
	lib.Merge(&employee, &emp)

	return emp, nil
}
