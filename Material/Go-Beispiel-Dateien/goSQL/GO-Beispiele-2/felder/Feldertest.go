// Autor: Oliver Schäfer
// Datum: Sa 26. Jan 16:03:20 CET 2019
// Zweck: Demo des felder-Paketes, das den ADT Feld exportiert.

package main

import (
  "gfx"
  "felder"
)

func main () {
  var vname,nname,str,plz,ort,leer1,leer2 felder.Feld
  var s string

  gfx.Fenster (800, 210)
  vname = felder.New (10,  10, 30, 'l', "Vorname")		// Position 10/10; Länge von 30 Zeichen; linksbündig; Name des Feldes
  nname = felder.New (10,  50, 30, 'r', "Nachname")
  str   = felder.New (10,  90, 30, 'z', "Straße/Hausnummer")
  plz   = felder.New (10, 130,  5, 'l', "PLZ")
  plz.SetzeErlaubteZeichen (felder.Digits)
  ort   = felder.New (10, 170, 30, 'l', "Ort")

  leer1 = felder.New (400, 10, 30, 'l', "")
  leer2 = felder.New (400, 50, 30, 'l', "")

  // Editieren der Eingabefelder
  // gelieferte Zeichenketten werden HIER nicht entgegengenommen ...
  vname.Edit ()
  nname.Edit ()
  str.Edit ()
  plz.Edit ()
  ort.Edit ()

  // ... dieser schon
  s = leer1.Edit ()
  // ... in das zweite leere Feld geschrieben
  leer2.Schreibe (s)

  // Bereits verwendete Felder lassen sich editieren
  vname.Edit ()
  nname.Edit ()
  // oder als Ausgabefelder verwenden
  ort.Schreibe ("Zehlendorf")
  str.Schreibe ("Straßenname ist zu lang und wird gekürzt")
  ort.Edit ()

  gfx.TastaturLesen1 ()
}
