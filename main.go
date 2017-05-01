package main

import (
  "database/sql"
  "fmt"
  "net/http"
  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "rashadlaher"
  dbname   = "mydb"
)

func main() {
  //createDB()
  r := NewRouter()
  r.Get("/", HomeHandler)
  r.Get("/countries", CountryHandler)
  r.Get("/countries/:name", CountryInfoHandler)
  r.Get("/players", PlayersHandler)
  r.Post("/countries/insert/:name", CountryInsertHandler)
  fmt.Println("Starting server on :8080")
  http.ListenAndServe(":8080", r.Router)


}

func OpenDB() *sql.DB{
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

func LookupCountryId(cname string) int{
  var db = OpenDB()
  countrySQL := `
  SELECT countryid
  FROM countries
  WHERE country=$1`

  var cid int
  row := db.QueryRow(countrySQL, cname)
  switch err := row.Scan(&cid); err{
  case sql.ErrNoRows:
    fmt.Println("no rows were returned")
    panic(err)
  case nil:
    return cid
  default:
    panic(err)
  }

}

func LookupClubId(cname string) int{
  var db = OpenDB()
  clubSQL := `
  SELECT clubid
  FROM clubs
  WHERE clubname=$1`

  var clubid int
  row := db.QueryRow(clubSQL, cname)
  switch err := row.Scan(&clubid); err{
  case sql.ErrNoRows:
    fmt.Println("no rows were returned")
    panic(err)
  case nil:
    return clubid
  default:
    panic(err)
  }

}
func createCountries(){
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

func createClubs(){
  var db=OpenDB()
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
func createPlayers(){
  var db=OpenDB()
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
func createDB(){
  createCountries()
  createClubs()
  createPlayers()
}
