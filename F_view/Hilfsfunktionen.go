// Autor: Philipp Liehm
// Datum: Juni 2023
// Ziel: Paramter der Tabellen in der Datenbank abfragen (z.B. Spaltenanzahl, Zeilenanzahl)


package main

import (
		"fmt"
		"SQL"
		"./textboxen"
		"gfx"
		"time"
		)

func main() {
	gfx.Fenster(1200,700)
	gfx.SetzeFont("Terminus-Bold.ttf",20)
	var conn SQL.Verbindung
	conn = SQL.PgSQL ("user=lewein dbname=lewein")
	defer conn.Beenden ()
		fmt.Println("Verbindung hergestellt.\n")
	
	// Spalten
	fmt.Println(anzahlSpalten(conn,"raeume"))
	fmt.Println(anzahlSpalten(conn,"veranstaltungen"))
	
	// Zeilen
	fmt.Println(anzahlZeilen(conn,"raeume"))
	
	// erstelle Tabelle
	t := erstelleTabelle(anzahlSpalten(conn,"raeume"),anzahlZeilen(conn,"raeume"))
	fmt.Println(t)
	fülleTbTabelle(t,conn,"raeume")
	
	// Kurzer Text
	var tb textboxen.Textbox
	tb = textboxen.New(600,400,100,50)
	tb.SchreibeText("Hallo")
	tb.SetzeHintergrundFarbe(255,0,0)
	tb.Zeichne()
	
	
	gfx.TastaturLesen1()
}

// Liefert die Anzahl der Spalten einer SQL-Tabelle
func anzahlSpalten(conn SQL.Verbindung, tabelle string) int {
	var rs SQL.Ergebnismenge
	var n int
	rs = conn.Anfrage("SELECT * FROM " + tabelle + ";")
	n = rs.AnzahlAttribute()
	rs.Schliessen()
	return n
}

// Liefert die Anzahl der Zeilen in einer Tabelle
func anzahlZeilen(conn SQL.Verbindung, tabelle string) int {
	var rs SQL.Ergebnismenge
	var n int
	rs = conn.Anfrage("SELECT * FROM " + tabelle + ";")
	for rs.GibtTupel() {
		//rs.LeseTupel()		Zum Zählen muss nicht gelesen werden
		n+=1
	}
	return n
}

// Erstellt eine Tabelle aus Textboxen mit einer gegebenen Anzahl an Spalten und Zeilen
func erstelleTabelle(spalten, zeilen int) [][]textboxen.Textbox{
	// erstelle für jede Zelle eine Textbox
	// Speichere den Zeiger auf die Textbox in ein feld
	var textbTabelle [][]textboxen.Textbox
	for i:=0;i<zeilen;i++ {
		temp := make([]textboxen.Textbox,0) 	// Slice aus Textboxen
		y := 100+50*uint16(i)
		for j:=0;j<spalten;j++ {
			x:=50+200*uint16(j)	
			fmt.Println(x,y)
			t:=textboxen.New(x,y,200,40)
			temp = append(temp,t)
		}
		textbTabelle = append(textbTabelle,temp)
	}
	return textbTabelle
}


// Füllt eine Tabelle von Textboxen mit den Werten einer Datenbank-Tabelle
func fülleTbTabelle(tabelle [][]textboxen.Textbox,conn SQL.Verbindung,tabellenName string) [][]textboxen.Textbox {
	fmt.Println("Füllen startet")
	var rs SQL.Ergebnismenge
	var n int
	var ergebnisse []interface{}	// Slice, beliebig viele Ergebnisse von typen die ich nicht kenne
	var ergebnisAdressen []interface{}	// Slice mit den Adressen zu den Ergebnissen 
	//var empty interface{} // braucht man, um weitere Variablen von unbekannten Typen dem Slice hinzuzufügen
	
	rs = conn.Anfrage("SELECT * FROM " + tabellenName + ";")
	
	// generiert benötigte Anzahl an Variablen für eine Zeile
	for i:=0;i<len(tabelle[0]);i++ {
		ergebnisse = append(ergebnisse,0)	// Dummy Wert für das leere Interface
		//ergebnisAdressen = append(ergebnisAdressen,&ergebnisse[i]) // Schreibe Adressenarray
		fmt.Println("Spalte: ",i)
	}
	
	for i:=0;i<anzahlSpalten(conn,tabellenName);i++ {
		ergebnisAdressen = append(ergebnisAdressen,&ergebnisse[i])	// &empty weil man später Zeiger benötigt
	}
	fmt.Println(ergebnisAdressen)
	
	
	for i:=0;i<len(tabelle);i++ {			// Für alle Zeilen
		for rs.GibtTupel() {
			rs.LeseTupel(ergebnisAdressen...)	
			n+=1
			for j:=0;j<len(tabelle[0]);j++ {		// Für jeden Wert in einer Zeile, also für jede Spalte
				tabelle[i][j].SchreibeText(fmt.Sprint(ergebnisse[j]))
				tabelle[i][j].Zeichne()
				time.Sleep(8e8)
				gfx.Stiftfarbe(255,255,255)
				gfx.Cls()
				gfx.Stiftfarbe(0,0,0)
				//fmt.Println(tabelle[i][j].GibSchriftgröße())
				//fmt.Println(i)
			}
			fmt.Println(ergebnisse)
		}
		
		
	}
	return tabelle
}
