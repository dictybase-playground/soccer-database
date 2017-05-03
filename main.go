package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "rashadlaher"
	dbname = "mydb"
)

type Env struct {
	db *sql.DB
}

func main() {
	//Set environment variables
	os.Setenv("DBNAME", "mydb")
	os.Setenv("DBUSER", "rashadlaher")
	env := &Env{db: OpenDB()}
	//createDB()
	/*
		clubs := NewClubsStore(env.db)
		Barcelona, err := NewClubs("Barcelona")
		err = clubs.Insert(Barcelona)
		if err != nil {
			panic(err)
		}

		countries := NewCountriesStore(env.db)
		Spain, err := NewCountries("Argentina")
		err = countries.Insert(Spain)
		if err != nil {
			panic(err)
		}

		players := NewPlayersStore(env.db)
		Messi, err := NewPlayers("Messi", "Barcelona", "Argentina")
		err = players.Insert(Messi)
		if err != nil {
			panic(err)
		}
	*/
	r := NewRouter()
	r.Get("/", HomeHandler)
	r.Get("/countries", env.CountryHandler)
	r.Get("/countries/:name", env.CountryInfoHandler)
	r.Get("/players", env.PlayersHandler)
	r.Get("/clubs", env.ClubHandler)
	r.Post("/players", env.PlayerInsertHandler)
	r.Post("/countries", env.CountryInsertHandler)
	r.Post("/clubs", env.ClubInsertHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r.Router)

}

func OpenDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func InsertCountry(cname string, database *sql.DB) {
	countries := NewCountriesStore(database)
	country, err := NewCountries(cname)
	err = countries.Insert(country)
	if err != nil {
		panic(err)
	}
}

func InsertClub(cname string, database *sql.DB) {
	clubs := NewClubsStore(database)
	club, err := NewClubs(cname)
	err = clubs.Insert(club)
	if err != nil {
		panic(err)
	}
}

func InsertPlayer(pname string, clubname string, countryname string, database *sql.DB) {
	players := NewPlayersStore(database)
	Messi, err := NewPlayers(pname, clubname, countryname)
	err = players.Insert(Messi)
	if err != nil {
		panic(err)
	}
}

func LookupCountryId(cname string) int64 {
	var db = OpenDB()
	countrySQL := `
  SELECT countryid
  FROM countries
  WHERE countryname=$1`

	var cid int64
	row := db.QueryRow(countrySQL, cname)
	switch err := row.Scan(&cid); err {
	case sql.ErrNoRows:
		fmt.Println("no rows were returned")
		panic(err)
	case nil:
		return cid
	default:
		panic(err)
	}

}

func LookupClubId(cname string) int64 {
	var db = OpenDB()
	clubSQL := `
  SELECT clubid
  FROM clubs
  WHERE clubname=$1`

	var clubid int64
	row := db.QueryRow(clubSQL, cname)
	switch err := row.Scan(&clubid); err {
	case sql.ErrNoRows:
		fmt.Println("no rows were returned")
		panic(err)
	case nil:
		return clubid
	default:
		panic(err)
	}

}
func createCountries() {
	//connect to database
	var db = OpenDB()
	//insert 2 countries
	insertCountry := `
INSERT INTO countries (country)
VALUES ($1)
RETURNING countryid`

	id := 0
	var err = db.QueryRow(insertCountry, "Argentina").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New country ID is:", id)
	err = db.QueryRow(insertCountry, "Brazil").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New country ID is:", id)

}

func createClubs() {
	var db = OpenDB()
	//insert 2 clubs to start
	insertClub := `
  INSERT INTO clubs (clubname)
  VALUES ($1)
  RETURNING clubid`

	clubid := 0
	var err = db.QueryRow(insertClub, "FC Barcelona").Scan(&clubid)
	if err != nil {
		panic(err)
	}
	fmt.Println("New club ID is:", clubid)
}
func createPlayers() {
	var db = OpenDB()
	//insert players

	insertPlayer := `
  INSERT INTO players (pname, clubid, countryid)
  VALUES ($1, $2, $3)
  RETURNING pid`

	clubid := LookupClubId("FC Barcelona")
	countryid := LookupCountryId("Argentina")

	pid := 0
	var err = db.QueryRow(insertPlayer, "Messi", clubid, countryid).Scan(&pid)
	if err != nil {
		panic(err)
	}
	fmt.Println("New player ID is:", pid)
}
func createDB() {
	createCountries()
	createClubs()
	createPlayers()
}
