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
	TodayHits     int
	MonthHits     int
	TotalHits     int
	TodayVisitors int
	MonthVisitors int
	TotalVisitors int
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

	todayVisitors := mysql1rowInt1("SELECT Count(Distinct Token) AS a FROM visits WHERE DATE(time) = CURDATE()")

	montlyVisitors := mysql1rowInt1("SELECT Count(Distinct Token) AS a FROM visits WHERE MONTH(time) = MONTH(CURRENT_DATE()) AND YEAR(time) = YEAR(CURRENT_DATE())")

	totalVisitors := mysql1rowInt1("SELECT Count(Distinct Token) AS a FROM visits")

	data := AdminData{
		TodayHits:     today,
		MonthHits:     montly,
		TotalHits:     total,
		TodayVisitors: todayVisitors,
		MonthVisitors: montlyVisitors,
		TotalVisitors: totalVisitors,
	}

	foo, _ := json.Marshal(data)

	fmt.Fprintf(w, string(foo))
	fmt.Println("AdminHandler:", time.Since(start))
}
