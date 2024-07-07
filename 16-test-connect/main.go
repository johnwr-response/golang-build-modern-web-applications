package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// connect to a database
	conn, err := sql.Open("pgx", "host=localhost port=5433 dbname=test_connect user=postgres password=exampleDOT33")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Fatal(fmt.Sprintf("Unable to close connection: %v\n", err))
		}
	}(conn)
	log.Println("Connected to database")

	// test my connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database!")
	}
	log.Println("Pinged database")

	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// insert a row
	query := `INSERT INTO users(first_name,last_name) VALUES($1, $2)`
	_, err = conn.Exec(query, "Jack", "Brown")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created user")

	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// update a row
	stmt := `UPDATE users SET first_name = $1 WHERE first_name = $2`
	_, err = conn.Exec(stmt, "Jackie", "Jack")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Updated user(s)")

	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// get one row by id
	query = `SELECT id, first_name, last_name FROM users WHERE id = $1`
	var firstName, lastName string
	var id int
	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("User %v %v %v", id, firstName, lastName)

	// delete a row
	query = `DELETE FROM users WHERE id = $1`
	_, err = conn.Exec(query, 6)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted user 6")

	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("SELECT id, first_name, last_name FROM users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Unable to close rows: %v\n", err)
		}
	}(rows)

	var firstName, lastName string
	var id int
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Printf("Unable to scan row: %v\n\n", err)
			return err
		}
		fmt.Printf("Record of id %v is: \t First name: %s,\t Last name: %s\n", id, firstName, lastName)
	}
	if err = rows.Err(); err != nil {
		log.Fatalf("Error scanning rows: %v\n", err)
	}
	fmt.Println("______________________________________")

	return nil
}
