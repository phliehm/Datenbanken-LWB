package main

//  Autor:  Oliver Schäfer
//  Datum:  Sa 26. Jan 16:13:28 CET 2019
//  Zweck:  Suchen in einer Kundendatenbank
//          Ausgabe über das gfx-Paket von Stefan Schmidt,
//          Einlesen der Benutzereingaben und Ausgabe über das felder-Paket

import (
  "SQL"
  "gfx"
  "felder"
  "strconv"
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
    suche     felder.Feld
    kid,nname,vname,strasse,ort felder.Feld
  )

  conn = SQL.PgSQL ("user=lewein dbname=lewein")
  defer conn.Beenden ()

  gfx.Fenster (600, 480)

  // Eingabefeld
  suche =   felder.New (10,  10, 30, 'l', "Nachname eingeben:")
  // Ausgabefelder
  kid =     felder.New (10,  50, 30, 'l', "KID:")
  nname =   felder.New (10,  90, 30, 'l', "Nachname:")
  vname =   felder.New (10, 130, 30, 'l', "Vorname:")
  strasse = felder.New (10, 170, 30, 'l', "Straße:")
  ort =     felder.New (10, 210, 30, 'l', "PLZ/Ort:")

  eingabe = suche.Edit ()
  query = fmt.Sprintf (`
    SELECT  *
    FROM    kunden
    WHERE   nname LIKE '%s';`, eingabe)
  fmt.Printf ("%s\n\n", query)
  rs = conn.Anfrage (query)

  for rs.GibtTupel () {
    rs.LeseTupel (&tupel.kid, &tupel.nname, &tupel.vname, &tupel.strasse, &tupel.plz, &tupel.ort)
    kid.Schreibe (strconv.Itoa (tupel.kid))
    nname.Schreibe (tupel.nname)
    vname.Schreibe (tupel.vname)
    strasse.Schreibe (tupel.strasse)
    ort.Schreibe (tupel.plz + " " + tupel.ort)
    gfx.TastaturLesen1 ()
    gfx.TastaturLesen1 ()
  }
  rs.Schliessen ()
}
