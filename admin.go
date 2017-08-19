package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AdminSTRUCT struct {
	PWD string
}

type AdminData struct {
	Today int
	Month int
	Total int
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	decoder := json.NewDecoder(r.Body)
	var Sjson AdminSTRUCT
	errSearch := decoder.Decode(&Sjson)
	if errSearch != nil {
		fmt.Println(errSearch)
		fmt.Fprintf(w, "ERROR")
		checkErr(err)
		return
	}

	if personalpwd != Sjson.PWD {
		fmt.Fprintf(w, `{"Status":"Not Loged in"}`)
		fmt.Println("ApiHandler-Not Loged in:", time.Since(start))
		return
	}

	//today := mysql1rowInt1("SELECT Token,time,Domain,Path,Refferer FROM visits WHERE DATE(time) = CURDATE()")
	today := mysql1rowInt1("SELECT COUNT(*) AS a FROM visits WHERE DATE(time) = CURDATE()")

	montly := mysql1rowInt1("SELECT COUNT(*) AS a FROM visits WHERE MONTH(time) = MONTH(CURRENT_DATE()) AND YEAR(time) = YEAR(CURRENT_DATE())")

	total := mysql1rowInt1("SELECT COUNT(*) AS a FROM visits")

	data := AdminData{
		Today: today,
		Month: montly,
		Total: total,
	}

	foo, _ := json.Marshal(data)

	fmt.Fprintf(w, string(foo))
	fmt.Println("AdminHandler:", time.Since(start))
}
