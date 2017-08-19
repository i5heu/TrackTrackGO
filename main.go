package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var fs = http.FileServer(http.Dir("static"))

var db *sql.DB
var err error
var TMPCACHE = make(map[string]template.HTML)
var TMPCACHECACHE = make(map[string]template.HTML)
var TMPCACHEWRITE bool = false
var TMPCACHECACHEWRITE bool = false

func main() {
	fmt.Println("HALLO")

	db, err = sql.Open("mysql", dblogin)
	db.SetConnMaxLifetime(time.Second * 2)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic
	}

	if err != nil {
		panic(err.Error())
	}
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	db.Exec("CREATE TABLE `visits` ( `id` int(10) unsigned NOT NULL AUTO_INCREMENT, `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `Token` varchar(30) NOT NULL DEFAULT 'No TOKEN', `Domain` varchar(100) NOT NULL DEFAULT 'NO DOMAIN', `Path` varchar(200) NOT NULL DEFAULT 'NO PATH', `Refferer` varchar(300) NOT NULL DEFAULT 'NO REFFERER', PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=latin1")

	http.HandleFunc("/admin", AdminHandler)
	http.HandleFunc("/api", ApiHandler)
	http.HandleFunc("/favicon.ico", FaviconHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8084", nil)
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/favicon/favicon.ico", 302)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("\033[0;31m", err, "\033[0m")
		err = nil
	}
}
