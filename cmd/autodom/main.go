package main

import (
	"autodom/env"
	"autodom/model"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//https://reqbin.com/

var (
	userFlag     = flag.String("u", "", "DB user")
	passFlag     = flag.String("p", "", "DB password")
	dbNameFlag   = flag.String("d", "", "DB name")
	httpAddrFlag = flag.String("http.addr", ":8080", "HTTP listen address")
)

type Request struct {
	SearchText  string `json:"searchText"`
	AnswerCount int    `json:"answerCount"`
}

type Env struct {
	DB *sql.DB
}

func NewEnv(db *sql.DB) *Env {
	return &Env{
		DB: db,
	}
}

///***
func (e *Env) ProblemsHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req Request

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &req)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(req.SearchText))
	text := string(req.SearchText)
	count := req.AnswerCount

	res := QueryDB(e.DB, text, count)

	if res == nil {
		res = append(res, model.Entity{Category: "Common", Title: "Unknown problem", Description: "solutions not found"})
	}

	//entity := Entity{Category: "hardware", Title: "some problem", Description: "some text"}

	js, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)

	if err != nil {
		panic(err)
	}
}

func QueryDB(db *sql.DB, text string, num int) []model.Entity {
	query := getSelectQuery(text)
	rows, err := db.Query("SELECT * FROM cases WHERE MATCH (title) AGAINST (? IN NATURAL LANGUAGE MODE) LIMIT ?", text, num)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	entities := []model.Entity{}

	for rows.Next() {
		e := model.Entity{}
		err := rows.Scan(&e.Category, &e.Title, &e.Description)

		if err != nil {
			fmt.Println(err)
			continue
		}

		entities = append(entities, e)
	}

	if len(entities) > num {
		return entities[0 : len(entities)-(len(entities)-num)]
	}

	return entities
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

	env := NewEnv(db)

	http.HandleFunc("/", env.ProblemsHandler)
	fmt.Println("Start...")
	log.Fatal(http.ListenAndServe(*httpAddrFlag, nil))

}
