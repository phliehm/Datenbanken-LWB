package main

//  Autor:  Oliver Schäfer, Sebastian Herker
//  Datum:  Sa 26. Jan 15:53:12 CET 2019, zuletzt geändert 16.01.2022
//  Zweck:  Suchen in einer Kundendatenbank

import (
  "SQL"
  "fmt"
)

type kunde struct {
  kid     int
  nname   string
  vname   string
  strasse string
  plz     string
  ort     string
}

func main() {
  var (
    conn    SQL.Verbindung
    rs      SQL.Ergebnismenge
    tupel   kunde
    query   string
    eingabe string
  )

  conn = SQL.PgSQL ("user=lewein dbname=lewein")
  defer conn.Beenden ()

  fmt.Print ("Bitte den gesuchten Namen eingeben: ")
  fmt.Scanf ("%s\n", &eingabe)
  query = fmt.Sprintf (`SELECT  *
                        FROM    kunden
                        WHERE   nname='%s'`, eingabe)
  rs = conn.Anfrage (query)

  for rs.GibtTupel () {
    rs.LeseTupel (&tupel.kid, &tupel.nname, &tupel.vname, &tupel.strasse, &tupel.plz, &tupel.ort)
    fmt.Printf ("%2d %-10s %-10s %-10s %-5s %-10s\n",
                tupel.kid, tupel.nname, tupel.vname, tupel.strasse, tupel.plz, tupel.ort)
  }
  rs.Schliessen ()
}
