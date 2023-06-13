// Autor: Oliver Schäfer
// Datum: Sa 26. Jan 15:48:24 CET 2019
// Zweck: Test der Datenbankverbindung zum lokalen PgSQL-Server

package main
import "SQL"

func main () {
  var conn SQL.Verbindung

  conn = SQL.PgSQL ("user=postgres dbname=postgres")
  defer conn.Beenden ()

  if conn == nil {
    println ("Verbindungstest fehlgeschlagen")
  } else {
    println ("Verbindungstest erfolgreich")
  }
}
