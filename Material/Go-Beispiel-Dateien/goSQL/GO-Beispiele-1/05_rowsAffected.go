//  Autor:  Oliver Schäfer
//  Datum:  Sa 26. Jan 15:52:28 CET 2019
//  Zweck:  Anzahl betroffener Zeilen ermitteln

package main
import "SQL"

func main() {
  var conn SQL.Verbindung
  var n    int64

  conn = SQL.PgSQL ("user=lewein dbname=lewein")
  defer conn.Beenden ()

  conn.Ausfuehren ("CREATE TABLE test (id INTEGER, name VARCHAR);")
  conn.Ausfuehren ("INSERT INTO test VALUES (1, 'Oliver');")
  n = conn.Ausfuehren ("DELETE FROM test WHERE id = 1;")
  println ("Anzahl entfernter Tupel:", n)
  conn.Ausfuehren ("DROP TABLE test CASCADE;")
}
