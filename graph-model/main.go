package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "graph"
)

type Person struct {
	ID    string
	Name  string
	OrgID string
}

type Organization struct {
	ID   string
	Name string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	csvFile, err := os.Open("./data/person.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(csvFile)
	r.Read()
	personRecords, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range personRecords {
		person := Person{
			ID:    record[0],
			Name:  record[1],
			OrgID: record[2],
		}

		_, err := db.Exec(`INSERT INTO entity (entity_type, entity_id) VALUES ('person', $1)`, person.ID)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(`INSERT INTO relation (parent, child) VALUES ($1, $2)`, person.OrgID, person.ID)
		if err != nil {
			log.Fatal(err)
		}
	}

	csvFile, err = os.Open("./data/organization.csv")
	if err != nil {
		log.Fatal(err)
	}

	r = csv.NewReader(csvFile)
	r.Read()
	orgRecords, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range orgRecords {
		org := Organization{
			ID:   record[0],
			Name: record[1],
		}

		_, err := db.Exec(`INSERT INTO entity (entity_type, entity_id) VALUES ('organization', $1)`, org.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
}
