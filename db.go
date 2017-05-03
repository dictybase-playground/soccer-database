package main

import kallax "gopkg.in/src-d/go-kallax.v1"

//go:generate kallax gen

type Players struct {
	kallax.Model `table:"players"`
	PID          int64 `pk:"autoincr"`
	Pname        string
	Clubid       int64 `fk:"clubid"`
	Countryid    int64 `fk:"countryid"`
}

type Clubs struct {
	kallax.Model `table:"clubs"`
	CLUBID       int64 `pk:"autoincr"`
	Clubname     string
}

type Countries struct {
	kallax.Model `table:"countries"`
	COUNTRYID    int64 `pk:"autoincr"`
	Countryname  string
}

func newClubs(clubname string) (*Clubs, error) {

	return &Clubs{Clubname: clubname}, nil
}

func newCountries(countryname string) (*Countries, error) {
	return &Countries{Countryname: countryname}, nil
}

func newPlayers(playername string, clubname string, countryname string) (*Players, error) {

	return &Players{Pname: playername, Clubid: LookupClubId(clubname), Countryid: LookupCountryId(countryname)}, nil
}
