package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// db details
const (
	postgres_host     = "dpg-d88qcaugvqtc73bci18g-a.singapore-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "postgres_admin"
	postgres_password = "bWjvBsTMtraa96Zd1bk5aktOKjtGLTpx"
	postgres_dbname   = "my_db_3k99"
)

// create pointer variable Db which points to sql driver
var Db *sql.DB

// init() is always called before main() by the Go compiler
func init() {
	//creating a database connection string
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)

	var err error

	Db, err = sql.Open("postgres", db_info)

	//handle error
	if err != nil {
		panic(err)
	} else {
		log.Println("Database successfully connected")
	}
}
