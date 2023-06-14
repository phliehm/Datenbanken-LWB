// Autor: Philipp Liehm
// Datum: Juni 2023
// Ziel: Einfache SELECT Anfragen an die LWB-Datenbank über die Go-Schnittstelle senden
// 		und die Antwort auf der Konsole ausgeben


package main

import (
		"fmt"
		"SQL"
		)

func main() {
	var (
	
	conn SQL.Verbindung		// Objekt der Klasse Verbindung mit SQL Datenbank
	rs SQL.Ergebnismenge	// Was ist das?  --> Objekt der Klasse Ergebnismenge
	n int64
	
	)
	
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
	
	var rname, ort string
	var rnr int					// Woher weiß ich welche Datentypen in der DB verwendet wurden?
	rs = conn.Anfrage("SELECT * FROM raeume;")
	fmt.Println("Es gibt ",rs.AnzahlAttribute()," Attribute.")
	fmt.Println("Die Attribute lauten:")
	fmt.Println(rs.Attribute())		// Liefert alle Attributnamen und muss vor dem Ende von GibtTupel() verwendet werden
	
	for rs.GibtTupel() {	// GibtTupel liefert True wenn noch eine Zeile nach der Cursor Position kommt
		rs.LeseTupel(&rnr,&rname,&ort)	// Liest von den Zeilen die einzelnen Attributswerte
		fmt.Println(rnr,rname,ort)	// Ein kompletter Datensatz	
	}
	rs.Schliessen() // Warum?

}
