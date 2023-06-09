package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "admin"
	dbPassword = "admin"
	dbName     = "graph"
	port     = "5511"
)

type DBHandler struct {
	*sql.DB
}

func main() {
	db, err := connectDB(dbUser, dbPassword, dbName, port)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = processCSVFile(db, "./data/entity.csv", "INSERT INTO entity(node_id, entity_type, entity_id) VALUES ($1, $2, $3)")
	if err != nil {
		panic(err)
	}

	err = processCSVFile(db, "./data/relation.csv", "INSERT INTO relation(parent, child) VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}
}

func connectDB(user, password, dbname string, port string) (*DBHandler, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, dbname, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DBHandler{db}, nil
}

func processCSVFile(dbHandler *DBHandler, filename string, query string) error {
	records, err := parseCSV(filename)
	if err != nil {
		return err
	}
	return dbHandler.insertRecords(query, records)
}

func parseCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}
	return reader.ReadAll()
}

func (db *DBHandler) insertRecords(query string, records [][]string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, record := range records {
		values := make([]interface{}, len(record))
		for i, v := range record {
			values[i] = v
		}

		_, err := stmt.Exec(values...)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

