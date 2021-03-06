package main
import (
  "database/sql"
  "fmt"
  "log"
  _ "github.com/denisenkom/go-mssqldb"
)
func main() {
  condb, errdb := sql.Open("mssql", "server=;user id=;password=;database=")
  if errdb != nil {
    fmt.Println(" Error open db:", errdb.Error())
  }
  var (
    sqlversion string
  )
  rows, err := condb.Query("select status from users")

  if err != nil {
    log.Fatal(err)
  }
  for rows.Next() {
    err := rows.Scan(&sqlversion)
    if err != nil {
      log.Fatal(err)
    }
    log.Println(sqlversion)
  }
  defer condb.Close()
}