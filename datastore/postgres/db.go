package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// db details

const (
	postgres_host     = "dpg-d0lu1dogjchc739bfbbg-a.oregon-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "postgres_admin"
	postgres_password = "a5aVK165BRjdRWDqt2ohxW1EvMtzTx5B"
	postgres_dbname   = "my_db_le7m"
)

// create pointer variable Db which points to sql driver
var Db *sql.DB

func init() {
	// creating a database connection string
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)
	fmt.Println(db_info)

	// var err error
	// establish a connection to PostgreSQL server using the driver
	var err error
	Db, err = sql.Open("postgres", db_info)
	// handle error
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database successfully connected")
	}
}
