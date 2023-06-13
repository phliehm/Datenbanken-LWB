package main

//  Autor:  Oliver Schäfer
//  Datum:  Sa 26. Jan 15:53:32 CET 2019
//  Zweck:  Suchen in einer Kundendatenbank
//          Verbesserung: Eingelesene Zeichenkette darf auch Leerzeichen enthalten.
//          von: http://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line

import (
  "SQL"
  "fmt"
  "bufio"
  "strings"
  "os"
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

  fmt.Print ("Name eingeben: ")
  // neu!!!!
  scanner := bufio.NewReader (os.Stdin)
  eingabe,_ = scanner.ReadString ('\n')
  eingabe = strings.TrimSuffix (eingabe, "\n")
  //
  query = fmt.Sprintf (`
    SELECT  *
    FROM    kunden
    WHERE   nname='%s';`, eingabe)
  fmt.Printf ("%s\n\n", query)
  rs = conn.Anfrage (query)

  for rs.GibtTupel () {
    rs.LeseTupel (&tupel.kid, &tupel.nname, &tupel.vname, &tupel.strasse, &tupel.plz, &tupel.ort)
    fmt.Printf ("%2d %-10s %-10s %-10s %-5s %-10s\n",
                tupel.kid, tupel.nname, tupel.vname, tupel.strasse, tupel.plz, tupel.ort)
  }
  rs.Schliessen ()
}
