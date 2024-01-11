package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "postgres"
)

func createtable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE data(
			id SERIAL Primary key,
			Name varchar(200),
			Email varchar(200) );
		`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Created Sucessfully")
}
func connection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		defer db.Close()
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected")
	return db
}

// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
// if err != nil {
// 	log.Fatal(err)
// }
// defer db.Close()

// //create the table if it doesn't exist
// _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")

// if err != nil {
// 	log.Fatal(err)
// }
