package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiSTRUCT struct {
	Token    string
	Domain   string
	Path     string
	Refferer string
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	decoder := json.NewDecoder(r.Body)
	var Sjson ApiSTRUCT
	errSearch := decoder.Decode(&Sjson)
	if errSearch != nil {
		fmt.Println(errSearch)
		fmt.Fprintf(w, "ERROR - INVALID JSON")
		checkErr(err)
		return
	}

	fmt.Println(Sjson.Domain)

	switch Sjson.Domain {
	case "selbstbildung.org":
	case "localhost":
	default:
		fmt.Fprintf(w, `{"Status":"ERROR-456Asdi9"}`)
		return
	}

	db.Exec("INSERT INTO visits(Token,Domain,Path,Refferer) VALUES(?,?,?,?)", Sjson.Token, Sjson.Domain, Sjson.Path, Sjson.Refferer)

	fmt.Fprintf(w, `{"Status":"OK"}`)
	fmt.Println("ApiHandler:", time.Since(start))
}
