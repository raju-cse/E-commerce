package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string{
	// user - > postgres
	// password -> password
	// host -> localhost
	// port - > 5432
	// db name -> ecommerce

	return "user=postgres password=password host=localhost port=5432 dbname=ecommerce"
}

func NewConnection() (*sqlx.DB, error){
	dbSource := GetConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbSource)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}