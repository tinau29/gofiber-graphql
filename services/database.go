package services

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/jackc/pgx/v5/stdlib"
// )

// var DB *sql.DB

// func InitDB() {
// 	var err error
// 	DB, err = sql.Open("pgx", "postgres://postgres:postgres@postgres:5432/postgres")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Automatically migrate the schema
// 	err = DB.AutoMigrate(&model.Employee{})
// 	if err != nil {
// 		log.Fatalf("failed to migrate database schema: %v", err)
// 	}

// }

import (
	"log"

	"api/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	if DB == nil {

		config := gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		}

		dsn := "host=postgres port=5432 user=postgres password=postgres dbname=postgres sslmode=disable TimeZone=Asia/Jakarta"
		db, err := gorm.Open(postgres.Open(dsn), &config)

		if db == nil {
			log.Println("database connection is nil")
		}

		// Automatically migrate the schema
		err = db.AutoMigrate(migrations.ModelMigrations...)
		if err != nil {
			log.Fatalf("failed to migrate database schema: %v", err)
		}

		DB = db

		return db
	}

	return DB
}
