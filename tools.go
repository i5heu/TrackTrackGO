package main

import (
	"regexp"
	"strings"
)

func ReplaceSpecialChars(s string) (sc string) {
	chars := []string{"]", "^", "\\\\", "[", ".", "(", ")", "<", ">", "+", "/", "#", "?", "=", "ß", "*", "'", "´", "\"", "%", ";", ":", "&", " "}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	sc = re.ReplaceAllString(s, "-")
	return
}

func ReplaceSpecialCharsWith_(s string) (sc string) {
	chars := []string{"]", "^", "\\\\", "[", ".", "(", ")", "<", ">", "+", "-", "/", "#", "?", "=", "ß", "*", "'", "´", "\"", "%", ";", ":", "&", " "}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	sc = re.ReplaceAllString(s, "_")
	return
}

func ReplaceSpecialCharsWithSpaceSpaceALLOWED(s string) (sc string) {
	chars := []string{"]", "^", "\\\\", "[", ".", "(", ")", "<", ">", "+", "-", "/", "#", "?", "=", "ß", "*", "'", "´", "\"", "%", ";", ":", "&", "\n"}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	sc = re.ReplaceAllString(s, " ")
	return
}

func mysql1rowInt1(comand string) (back int) {
	row, err := db.Query(comand)
	checkErr(err)
	defer row.Close()

	row.Next()
	_ = row.Scan(&back)

	return
}
