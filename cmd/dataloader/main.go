package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	createCases = `
	CREATE TABLE cases (category VARCHAR(255) NOT NULL,
	title VARCHAR(255) NOT NULL,                    
	description VARCHAR(1024) NOT NULL,
	FULLTEXT (title)                    
);
`
)

var (
	userFlag   = flag.String("u", "", "DB user")
	passFlag   = flag.String("p", "", "DB password")
	dbNameFlag = flag.String("d", "", "DB name")
)

type dbEntity struct {
	category string
	symptom  string
	solution string
}

func newDbEntity(cat, smp, sol string) dbEntity {
	return dbEntity{
		category: cat,
		symptom:  smp,
		solution: sol,
	}
}

func prepareDB(db *sql.DB, statement string) error {
	stmt, err := db.Prepare(statement)

	if err != nil {
		return err
	}
	_, err = stmt.Exec()

	if err != nil {
		return err
	} else {
		fmt.Println("Successfully..")
	}

	return err
}

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *userFlag == "" || *passFlag == "" || *dbNameFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	credits := fmt.Sprintf("%s:%s@/%s", *userFlag, *passFlag, *dbNameFlag)

	db, err := sql.Open("mysql", credits)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	err = prepareDB(db, createCases)

	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Open("../../resources/sample.json")

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Opened json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	rawResult := make(map[string]interface{})
	result := []dbEntity{}

	err = json.Unmarshal([]byte(byteValue), &rawResult)

	if err != nil {
		panic(err)
	}

	for k, v := range rawResult {
		switch vv := v.(type) {
		case []interface{}:
			for _, l := range vv {
				switch ll := l.(type) {
				case map[string]interface{}:
					var symp, so string

					for _, p := range ll {
						switch lll := p.(type) {
						case string:
							symp = lll
						case []interface{}:
							for _, f := range lll {
								switch ff := f.(type) {
								case string:
									so = ff
									entity := newDbEntity(k, symp, so)
									fmt.Println("symptom", entity.symptom)
									result = append(result, entity)
								}
							}

						}
					}
				}
			}
		}
	}

	for _, q := range result {
		_, err := db.Exec("INSERT INTO autodom.cases (category, title, description) VALUES (?, ?, ?)",
			q.category, q.symptom, q.solution)

		if err != nil {
			panic(err)
		}

	}
}
