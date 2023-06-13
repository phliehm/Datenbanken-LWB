// Autor: Oliver Schäfer, Sebastian Herker
// Datum: Sa 26. Jan 15:49:04 CET 2019, zuletzt geändert: 16.01.2022
// Zweck: Löschen einer Tabelle mit GO

package main
import "SQL"

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
