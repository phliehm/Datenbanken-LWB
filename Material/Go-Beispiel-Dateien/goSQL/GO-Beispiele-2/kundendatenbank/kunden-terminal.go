package main

//  Autor:  Oliver Schäfer, Sebastian Herker
//  Datum:  Sa 26. Jan 16:13:28 CET 2019, zuletzt aktualisiert: 08.02.2021
//  Zweck:  Suchen in einer Kundendatenbank über das Terminalfenster

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
    conn      SQL.Verbindung
    rs        SQL.Ergebnismenge
    tupel     kunde
    query     string
    eingabe   string
  )

// Verbindung herstellen
  conn = SQL.PgSQL ("user=lewein dbname=lewein")
  defer conn.Beenden ()

// Kundenname abfragen
  fmt.Print ("Bitte geben Sie den Kundennamen ein: ")
  fmt.Scanln (&eingabe)

// SQL-Anfrage mittels Sprintf erstellen
  query = fmt.Sprintf (`
    SELECT  *
    FROM    kunden
    WHERE   nname LIKE '%s';`, eingabe)
  fmt.Printf ("%s\n\n", query)

// SQL-Anfrage stellen und Ergebnismenge ermitteln
  rs = conn.Anfrage (query)

// Ergebnistupel durchgehen und ausgeben
  for rs.GibtTupel () {
    rs.LeseTupel (&tupel.kid, &tupel.nname, &tupel.vname, &tupel.strasse, &tupel.plz, &tupel.ort)
    fmt.Println ("KID:",tupel.kid)
    fmt.Println ("Nachname:",tupel.nname)
    fmt.Println ("Vorname:",tupel.vname)
    fmt.Println ("Straße:",tupel.strasse)
    fmt.Println ("Ort:",tupel.plz + " " + tupel.ort)
    fmt.Println ()
  }
  rs.Schliessen () // Anfrage schließen
}
