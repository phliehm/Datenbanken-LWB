//  Autor:  Oliver Schäfer, Sebastian Herker
//  Datum:  Sa 26. Jan 15:51:43 CET 2019, zuletzt geändert: 16.01.2022
//  Zweck:  SQL-Statements

package main
import "SQL"

func main() {
  var conn SQL.Verbindung

  conn = SQL.PgSQL ("user=lewein dbname=lewein")
  defer conn.Beenden ()

  conn.Ausfuehren (`CREATE TABLE test (
                      id    INTEGER,
                      name  VARCHAR
                    );`)
}
