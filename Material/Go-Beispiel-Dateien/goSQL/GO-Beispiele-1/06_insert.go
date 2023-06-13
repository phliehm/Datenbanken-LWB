// Autor: Oliver Schäfer, Sebastian Herker
// Datum: Sa 26. Jan 15:52:11 CET 2019, zuletzt geändert: 16.01.2022
// Zweck: Erstellen einer Tabelle und Einfügen von Daten mit GO

package main
import ("SQL")

func main () {
  var (
    conn SQL.Verbindung
    n    int64
  )

  conn = SQL.PgSQL ("user=lewein dbname=lewein")
  defer conn.Beenden ()

  println ("Verbindungstest erfolgreich")
	conn.Ausfuehren (`CREATE TABLE test (
                      id    INTEGER,
                      name  VARCHAR
                     );`)
  println ("Tabelle 'test' erstellt")
  n = conn.Ausfuehren ("INSERT INTO test VALUES (1,'Oliver');")
  println (n, "Zeilen eingefügt")
  n = conn.Ausfuehren ("INSERT INTO test VALUES (2,'Sebastian');")
  println (n, "Zeilen eingefügt")
  n = conn.Ausfuehren ("DELETE FROM test;")
  println (n, "Zeilen gelöscht")
  n = conn.Ausfuehren ("DROP TABLE test CASCADE;")
  println ("Tabelle 'test' gelöscht")
}
