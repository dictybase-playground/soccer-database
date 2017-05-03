package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Country struct {
	Country string `json:"country"`
}
type Club struct {
	Club string `json:"club"`
}
type Player struct {
	Player  string `json:"player"`
	Club    string `json:"club"`
	Country string `json:"country"`
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Home")
}
func (env *Env) CountryHandler(rw http.ResponseWriter, r *http.Request) {
	var db = env.db
	rows, err := db.Query("SELECT countryid, countryname FROM countries LIMIT $1", 3)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var countryid int
		var country string
		err = rows.Scan(&countryid, &country)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(countryid, country)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func (env *Env) CountryInfoHandler(rw http.ResponseWriter, r *http.Request) {
	var db = env.db
	params := Params(r)
	cid := params.ByName("name")
	rows, err := db.Query("SELECT countryid, countryname FROM countries WHERE countryname=$1", cid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var countryid int
		var country string
		err = rows.Scan(&countryid, &country)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(countryid, country)
	}
}
func (env *Env) CountryInsertHandler(rw http.ResponseWriter, r *http.Request) {
	var c Country
	if r.Body == nil {
		http.Error(rw, "please send a request body", http.StatusInternalServerError)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	fmt.Println(c.Country)
	InsertCountry(c.Country, env.db)
}

func (env *Env) PlayersHandler(rw http.ResponseWriter, r *http.Request) {
	var db = env.db
	rows, err := db.Query("SELECT pid, pname, clubid, countryid FROM players")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var pid int
		var pname string
		var clubid int
		var countryid int
		err = rows.Scan(&pid, &pname, &clubid, &countryid)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(pid, pname, clubid, countryid)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func (env *Env) PlayerInsertHandler(rw http.ResponseWriter, r *http.Request) {
	var p Player
	if r.Body == nil {
		http.Error(rw, "please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	fmt.Println(p)
	InsertPlayer(p.Player, p.Club, p.Country, env.db)
}

func (env *Env) ClubHandler(rw http.ResponseWriter, r *http.Request) {
	var db = env.db
	rows, err := db.Query("SELECT clubid, clubname FROM clubs LIMIT $1", 3)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var clubid int
		var club string
		err = rows.Scan(&clubid, &club)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(clubid, club)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func (env *Env) ClubInsertHandler(rw http.ResponseWriter, r *http.Request) {

	var c Club
	if r.Body == nil {
		http.Error(rw, "please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	InsertClub(c.Club, env.db)
	fmt.Println(c.Club)
}
