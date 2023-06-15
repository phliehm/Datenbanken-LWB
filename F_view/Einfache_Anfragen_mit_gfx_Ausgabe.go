// Autor: Philipp Liehm
// Datum: Juni 2023
// Ziel: Einfache SELECT Anfragen an die LWB-Datenbank über die Go-Schnittstelle senden
// 		und die Antwort in einem gfx-Fenster ausgeben

package main

import (
		"fmt"
		"gfx"
		"SQL"
		"time"
		)	

const breite, hoehe uint16 = 1200,700
		
func main() {
	gfx.Fenster(breite,hoehe)		// Öffne gfx-Fenster
	
	// Variablen
	var (
	conn SQL.Verbindung		// Objekt der Klasse Verbindung mit SQL Datenbank
	rs SQL.Ergebnismenge	// Was ist das?  --> Objekt der Klasse Ergebnismenge
	n int64
	)
	
	// Verbindungsaufbau zur Datenbank
	conn = SQL.PgSQL("user=lewein dbname=lewein")
	defer conn.Beenden()	// Damit später beim Beenden des Programms die Verbindung geschlossen wird
	fmt.Println("Verbindungstest erfolgreich\n")
	
	// SELECT Anfrage
	zeigeRaeume(conn,rs)
	// INSERT
	n = conn.Ausfuehren("INSERT INTO raeume VALUES (10,'Spassraum','LalaLand');")
	fmt.Println(n," Zeile in raeume eingefügt\n")
	
	// PRÜFEN
	zeigeRaeume(conn,rs)
	
	// DELETE
	n = conn.Ausfuehren("DELETE FROM raeume WHERE raumnr=10;")
	fmt.Println(n," Zeilen gelöscht\n")
	
	// PRÜFEN
	zeigeRaeume(conn,rs)
	
}


// SELECT Anfrage für raeume
func zeigeRaeume(conn SQL.Verbindung,rs SQL.Ergebnismenge) {
	/*
	var rname, ort string
	var rnr int					// Woher weiß ich welche Datentypen in der DB verwendet wurden?
	*/
	//var rname,ort,rnr interface{}	// So muss man die Datentypen vorher nicht wissen
	// --> später kann man für die Ausgabe sowieso alles mit fmt.Sprint() in strings konvertieren
	
	var ergebnisse []interface{}
	var empty interface{}
	ergebnisse = append(ergebnisse,empty)
	ergebnisse = append(ergebnisse,empty)
	ergebnisse = append(ergebnisse,empty)
	
	var ergebnisAdressen []interface{}
	for i:=0;i<3;i++ {
		ergebnisAdressen = append(ergebnisAdressen,&ergebnisse[i])	// &empty weil man später Zeiger benötigt
	}
	
	rs = conn.Anfrage("SELECT * FROM raeume;")
	fmt.Println("Es gibt ",rs.AnzahlAttribute()," Attribute.")
	fmt.Println("Die Attribute lauten:")
	fmt.Println(rs.Attribute())		// Liefert alle Attributnamen und muss vor dem Ende von GibtTupel() verwendet werden
	
	gfx.Schreibe(100,100,"Raumnummer       Raumname         Ort")
	var abstand uint16 
	for rs.GibtTupel() {	// GibtTupel liefert True wenn noch eine Zeile nach der Cursor Position kommt
		//rs.LeseTupel(&rnr,&rname,&ort)	// Liest von den Zeilen die einzelnen Attributswerte
		rs.LeseTupel(ergebnisAdressen...)
		//gfx.Schreibe(100,140+abstand,fmt.Sprint(rnr)+"                "+fmt.Sprint(rname)+"       "+fmt.Sprint(ort))
		//fmt.Println(rnr,rname,ort)	// Ein kompletter Datensatz	
		fmt.Println(ergebnisse)
		abstand +=40 // Versetze in y-Richtung
		time.Sleep(1e9)
	}
	rs.Schliessen() // Warum?

}


