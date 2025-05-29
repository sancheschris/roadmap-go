package main

import (
	"fmt"
)

type Database interface {
	Connect() error
	Query(query string) string
}

type PostgresDB struct {
	ConnectionString string
	Connected bool
}

func (pg *PostgresDB) Connect() error {
	pg.Connected = true
	fmt.Println("Connected to PostgresDB at", pg.ConnectionString)
	return nil
}

func (pg *PostgresDB) Query(query string) string {
	if !pg.Connected {
		return "Error: Not connected to the database"
	}
	return fmt.Sprintf("Executing query on PostgresDB: %s", query)
}

type MongoDB struct {
	URI string
	Connected bool
}

func (m *MongoDB) Connect() error {
	m.Connected = true
	fmt.Println("Connected to MongoDB at", m.URI)
	return nil
}

func (m *MongoDB) Query(query string) string {
	if !m.Connected {
		return "Error: Not connected to the database"
	}
	return fmt.Sprintf("Executing query on MongoDB: %s", query)
}

func executeQuery(db Database, query string) {
	if err := db.Connect(); err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	result := db.Query(query)
	fmt.Println(result)
}


func main() {
	postgres := &PostgresDB{ConnectionString: "postgres://user:password@localhost:5432/mydb"}
	mongo := &MongoDB{URI: "mongodb://user:password@localhost:27017/mydb"}

	fmt.Println("Using PostgresDB:")
	executeQuery(postgres, "SELECT * FROM users")

	fmt.Println("\nUsing MongoDB:")
	executeQuery(mongo, "db.users.find()")
}