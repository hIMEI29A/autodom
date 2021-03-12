package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	searchTextFlag = flag.String("t", "no sound", "Text to search")
	urlFlag        = flag.String("u", "http://localhost:8080", "Text to search")
	countFlag      = flag.Int("c", 1, "Number of results")
)

type Obj struct {
	SearchText  string `json:"searchText"`
	AnswerCount int    `json:"answerCount"`
}

func main() {
	flag.Parse()

	client := &http.Client{}

	obj := Obj{SearchText: *searchTextFlag, AnswerCount: *countFlag}
	jsonStr, err := json.Marshal(obj)

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", *urlFlag, bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "text/html")

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// DEBUG
	fmt.Printf("%s\n%s\n", "response Body:", string(body))
}
