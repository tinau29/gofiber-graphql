package resolver

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{DB: db}
}

// HelloResolver resolves the hello query
func HelloResolver(p graphql.ResolveParams) (interface{}, error) {
	return "Hello, world!", nil
}
