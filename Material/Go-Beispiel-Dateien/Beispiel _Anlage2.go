// Autor: Benjamin Schneider
// Datum: Mo 27. Feb 2023
// Zweck: Tabellen mit GO

package main
import "SQL";"fmt"

func main () {
  var conn SQL.Verbindung

  conn = SQL.PgSQL ("user=lewein dbname=lewein")
  defer conn.Beenden ()

  if conn == nil {
    println ("Verbindungstest fehlgeschlagen")
  } else {
    println ("Verbindungstest erfolgreich")
    conn.Ausfuehren ("DROP TABLE test CASCADE;")
    println ("Tabelle 'test' gelöscht")
  }
}
