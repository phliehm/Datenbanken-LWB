// Autor: Oliver Schäfer
// Datum: Sa 26. Jan 16:03:20 CET 2019
// Zweck: Demo des felder-Paketes, das den ADT Feld exportiert.

package main

import (
  "gfx"
  "SQL"
  "fmt"
  "felder"
)

func main () {
	var conn SQL.Verbindung
	var erg SQL.Ergebnismenge
	
	conn = SQL.PgSQL ("user=lewein dbname=lewein")
	defer conn.Beenden ()
	
	if conn == nil {
		println ("Verbindungstest fehlgeschlagen")
	} else {
		println ("Verbindungstest erfolgreich")
	}
	
  var vname,nname,str,plz,ort,kund felder.Feld
  var kid,nn,vn,stra,post,or string

  gfx.Fenster (800, 210)
  vname = felder.New (10,  10, 30, 'l', "Vorname")		// Position 10/10; Länge von 30 Zeichen; linksbündig; Name des Feldes
  nname = felder.New (10,  50, 30, 'l', "Nachname")
  str   = felder.New (10,  90, 30, 'l', "Straße")
  plz   = felder.New (10, 130,  5, 'l', "PLZ")
  plz.SetzeErlaubteZeichen (felder.Digits)
  ort   = felder.New (10, 170, 30, 'l', "Ort")

  kund = felder.New (400, 10, 30, 'l', "Kunden-ID")
  // leer2 = felder.New (400, 50, 30, 'l', "")

  // Editieren der Eingabefelder
  // gelieferte Zeichenketten werden HIER nicht entgegengenommen ...
  
  vn = vname.Edit ()
  nn = nname.Edit ()
  stra = str.Edit ()
  post = plz.Edit ()
  or = ort.Edit ()
  kid = kund.Edit()
	
	anfrage := fmt.Sprintf(`
			SELECT * 
			FROM kunden 
			WHERE kid=%s ;`, kid)
	
	// fmt.Println(anfrage)
	// gfx.TastaturLesen1()
	
	erg = conn.Anfrage(anfrage)
	
	if erg.GibtTupel() {
		fmt. Println("Kunden-ID existiert bereits!")
	} else {
		eingabe := fmt.Sprintf(`
			INSERT INTO kunden 
			VALUES (%s,'%s','%s','%s','%s','%s');`, kid,nn,vn,stra,post,or)
		// fmt.Println(eingabe)
		
		conn.Ausfuehren (eingabe)
		println ("Neue Werte wurden eingefügt!")
	}
	
	gfx.TastaturLesen1 ()
}
