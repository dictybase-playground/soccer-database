package main

import(
  "fmt"
  "net/http"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(rw, "Home")
}
func CountryHandler(rw http.ResponseWriter, r *http.Request){
  var db = OpenDB()
  rows, err := db.Query("SELECT countryid, country FROM countries LIMIT $1", 3)
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
func CountryInfoHandler(rw http.ResponseWriter, r *http.Request){
  var db = OpenDB()
  params := Params(r)
  cid := params.ByName("name")
  rows, err := db.Query("SELECT countryid, country FROM countries WHERE country=$1", cid)
  if err != nil{
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
func CountryInsertHandler(rw http.ResponseWriter, r *http.Request){
//  var db = OpenDB()
  //params := Params(r)

}

func PlayersHandler(rw http.ResponseWriter, r *http.Request){
  var db = OpenDB()
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
