// Author: Philipp Liehm
// Datum: Juni 2023
// Zweck: 	Eine SELECT Anfrage wird für eine bestehende Verbindung gestellt
//			und der Inhalt in eine Tabelle des Datentyps string geschrieben 

package sqlTabelle

import (
		"fmt"
		"SQL"
		)

type data struct {
	kopf  []string
	tabelle [][]string // Oder sogar [][][]string um noch die Länge der Zelle zu speichern? 
	zeilen	uint16
	spalten uint16
	conn 	SQL.Verbindung
	anfrage string		// hier steckt auch der Name der SQL-Tabelle drin
}

// Vor.: Die Verbindung zur SQL-Datenbank steht, die "anfrage" ist korrekt und liefert eine Tabelle
func New(conn SQL.Verbindung,anfrage string) *data{
	sTabelle := new(data)
	sTabelle.zeilen = uint16(anzahlZeilen(conn,anfrage))
	sTabelle.spalten = uint16(anzahlSpalten(conn,anfrage))
	sTabelle.conn = conn
	sTabelle.anfrage = anfrage
	sTabelle.kopf = leseHeader(conn,anfrage)
	sTabelle.tabelle = gibTabelle(conn,anfrage) 
	
	return sTabelle
}

// Gibt die Tabelle mit Strings zurück
func (sT *data) GibTabelle() [][]string {
	return sT.tabelle
}

func (sT *data) GibKopf() []string {
	return sT.kopf
}

//////////////////////
// HILFS-FUNKTIONEN //
//////////////////////


// Liefert den Tabellenkopf als []string
func leseHeader(conn SQL.Verbindung,anfrage string) []string {
	rs := conn.Anfrage(anfrage)		// Starte Anfrage
	return rs.Attribute()			// Tabellenkopf
}


// Liefert die Anzahl der Spalten einer SQL-Tabelle
func anzahlSpalten(conn SQL.Verbindung, anfrage string) int {
	var rs SQL.Ergebnismenge
	var n int
	rs = conn.Anfrage(anfrage)
	n = rs.AnzahlAttribute()
	rs.Schliessen()
	return n
}

// Liefert die Anzahl der Zeilen in einer Tabelle
func anzahlZeilen(conn SQL.Verbindung, anfrage string) int {
	var rs SQL.Ergebnismenge
	var n int
	rs = conn.Anfrage(anfrage)
	for rs.GibtTupel() {
		n+=1
	}
	return n
}

// Erstellt eine Tabelle aus strings mit einer gegebenen Anzahl an Spalten und Zeilen
func erstelleTabelle(zeilen,spalten int) [][]string {
	// erstelle für jede Zelle eine Textbox
	// Speichere den Zeiger auf die Textbox in ein feld
	var stringTabelle [][]string
	
	for i:=0;i<zeilen;i++ {
		zeile := make([]string,0) 	// Slice aus Textboxen
		for j:=0;j<spalten;j++ {
			var s string
			zeile = append(zeile,s)
		}
		stringTabelle = append(stringTabelle,zeile)
	}
	return stringTabelle
}

// Liest eine Tabelle einer SQL-Anfrage in einen [][]string ein
func gibTabelle(conn SQL.Verbindung, anfrage string) [][]string {
	var rs SQL.Ergebnismenge
	//var n int
	var ergebnisse []interface{}	// Slice, beliebig viele Ergebnisse von typen die ich nicht kenne
	var ergebnisAdressen []interface{}	// Slice mit den Adressen zu den Ergebnissen 

	
	// Stelle Anfrage
	rs = conn.Anfrage(anfrage)
	
		
	// generiere leere Tabelle
	tabelle := erstelleTabelle(anzahlZeilen(conn,anfrage),anzahlSpalten(conn,anfrage))
	// Tabelle ist leer
	if len(tabelle)<1 {return [][]string{{"Kein Ergebnis für diese Suchanfrage"}}}
	// generiert benötigte Anzahl an Variablen für eine Zeile (also Anzahl der Spalten)
	for i:=0;i<len(tabelle[0]);i++ {
		ergebnisse = append(ergebnisse,0)	// Dummy Wert für das leere Interface
	}
	// generiert einen Array für die Adressen zu den jeweiligen Werten, da GibTupel Adressen fordert
	for i:=0;i<len(tabelle[0]);i++ {
		ergebnisAdressen = append(ergebnisAdressen,&ergebnisse[i])	// "&" weil man bei rs.LeseTupel() Zeiger benötigt
	}
	
	i:=0		// Zeilen
	for rs.GibtTupel() {
		rs.LeseTupel(ergebnisAdressen...)	// Werte werden über die Adressen an die richtige Stelle gespeichert
		for j:=0;j<len(tabelle[0]);j++ {		// Für jeden Wert in einer Zeile, also für jede Spalte
			tabelle[i][j] = interfaceToString(ergebnisse[j])// Schreibe Inhalt in leere Tabelle
		}
		i+=1		// nächste Zeile
	}
	return tabelle	
}

// nötig um auch den Datentyp "note" bzw. []rune lesen zu können
func interfaceToString(i interface{}) string {
	switch v := i.(type) {
		case string:
			return v
		case []uint8:
			return string(v)
		case float64,int64:
			return fmt.Sprint(v)
		default:
			return ""
	}
}
