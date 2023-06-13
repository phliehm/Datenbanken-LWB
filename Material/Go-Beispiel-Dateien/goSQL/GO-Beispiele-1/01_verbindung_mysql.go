//  Autor:  Oliver Schäfer
//  Datum:  26.01.2019
//  Zweck:  Verbindungstest mySQL
//          > erfordert die Installation eines MySQL-Servers
//          > in den Paketquellen von SL ist nur der kompatible Fork MariaDB enthalten

package main
import "SQL"

func main() {
  var conn SQL.Verbindung

  conn = SQL.MySQL ("lewein@/lewein")
  defer conn.Beenden ()

  if conn == nil {
    println ("Verbindung fehlgeschlagen")
  } else {
    println ("Verbindung erfolgreich")
  }
}
