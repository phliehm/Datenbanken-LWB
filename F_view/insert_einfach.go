// Autor: Philipp Liehm (basierend auf O. Schäfers ProgrammCode)
// Datum: Juni 2023
// Zweck: Insert INTO PSQL Datenbank

package main

import (
  "gfx"
  "SQL"
  "fmt"
  "felder"
  "../Klassen/sqlTabelle"
  "../Klassen/textboxTabelle"
  //"time"
)

var font string = "../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf"

func main () {
	gfx.Fenster (1200, 700)
	var conn SQL.Verbindung
	//var erg SQL.Ergebnismenge
	
	conn = SQL.PgSQL ("user=lewein dbname=lewein")
	defer conn.Beenden ()
	
	if conn == nil {
		println ("Verbindungstest fehlgeschlagen")
	} else {
		println ("Verbindungstest erfolgreich")
	}
	
	fügeHinzuVeranst(conn)
	//updateLieblingsgetraenk(conn)
	//prüfeUNDfügeHinzuNPCs(conn)
	//löscheAusTabelle(conn,löscheNPC,"NPC")
	//löscheDoz(conn)
	//prüfeUNDfügeHinzuNPCs(conn)
	//fügeDozentHinzu(conn)
	gfx.TastaturLesen1()
	
}

//
func fügeHinzuVeranst(conn SQL.Verbindung) {
	//Name, Thema, Kürzel, Dozentin, SWS, Semester, Raum
	// 
	zeigeVeranst(conn)
	var eingabe string
	var vnrS,vnameS,gebietnameS,kuerzelS,npcnameS,swsS,semesterS,raumnrS string
	var gebietnrS,npcnrS string // um gebietnummer zu finden
	var eintragWarVorhanden bool
	
	vnameS = "Lolo"
	gebietnameS = "Lala"
	kuerzelS = "Kon"
	npcnameS = "Ben Schneider"
	swsS = "7"
	semesterS ="1"
	raumnrS = "4"
	
	fmt.Println(vnameS,gebietnameS,kuerzelS,npcnameS,swsS,semesterS,raumnrS)
	// mögliche Probleme
	// existiert Dozent? Existiert Thema?
	
	// 1. Gebiet prüfen ob vorhanden, sonst neu erstellen
	// --> in veranstaltunen eintragen
	// 2. neuer Eintrag in unterricht --> vnr, npcnr,raumnr
	// Da sollte noch kein Eintrag existieren, weil die Veranstaltung ja noch nicht existiert.
	
	// GEBIET
	eintragWarVorhanden,gebietnrS = prüfeObVorhandenFindeNr(conn,"themengebiete", "gebietname", gebietnameS, "gebietnr" )
	if eintragWarVorhanden == false {
		// Einfügen in themengebiete
		eingabe = fmt.Sprintf(`
		INSERT INTO themengebiete
		VALUES (%s,'%s');`, gebietnrS,gebietnameS)
	    conn.Ausfuehren(eingabe)
		fmt.Println("Neue Werte wurden eingefügt!")
	}
	
	// VERANSTALTUNGEN
	eintragWarVorhanden, vnrS = prüfeObVorhandenFindeNr(conn,"veranstaltungen", "vname", vnameS, "vnr")
	if eintragWarVorhanden {return}		// Wenn es den Eintrag schon gab, mache nichts
	// Trage neue Veranstaltung ein
	eingabe = fmt.Sprintf(`
		INSERT INTO veranstaltungen 
		VALUES ('%s','%s','%s','%s','%s','%s');`,vnrS,vnameS, kuerzelS,swsS,semesterS,gebietnrS)
	conn.Ausfuehren(eingabe)
	
	// NPC
	eintragWarVorhanden, npcnrS = prüfeObVorhandenFindeNr(conn,"npcs", "npcname", npcnameS, "npcnr")
	// Wenn es den NPC noch nicht gab, füge ihn hinzu
	if eintragWarVorhanden == false {
		eingabe = fmt.Sprintf(`
		INSERT INTO npcs
		VALUES ('%s','%s');`,npcnrS,npcnameS)
	conn.Ausfuehren(eingabe)
	}
	
	// DOZENTIN
	//_ Da man ja schon eine npcnr hat, die aber nicht überschrieben werde soll
	eintragWarVorhanden, _ = prüfeObVorhandenFindeNr(conn,"dozent_innen", "npcnr", npcnrS, "npcnr")
	if eintragWarVorhanden == false {
		eingabe = fmt.Sprintf(`
		INSERT INTO dozent_innen
		VALUES ('%s','?');`,npcnrS)
	conn.Ausfuehren(eingabe)
	}
	
	// UNTERRICHTEN
	eingabe = fmt.Sprintf(`
		INSERT INTO unterricht
		VALUES ('%s','%s','%s');`,vnrS,npcnrS,raumnrS)
	conn.Ausfuehren(eingabe)
	// Nochmal Veranstaltungen zeichnen
	zeigeVeranst(conn)
}


// Zeigt die wesentlichen Attribute von Veranstaltungen						
func zeigeVeranst(conn SQL.Verbindung) {
	anfrage := "SELECT vname,gebietname,kuerzel,npcname,sws,semester,raumnr FROM veranstaltungen"+
				" NATURAL JOIN unterricht NATURAL JOIN npcs NATURAL JOIN themengebiete;"
	textboxTabelle.ZeichneAnfrage(conn,anfrage,10,200,true,0,0,0,0,0,255,16,font) 
				
}

// Prüft ob ein Attributswert existiert, wenn nicht, ist eine freie Nummer zurückgeliefert, 
// Wenn ja, ist die zugehörige Nummer geliefert.
func prüfeObVorhandenFindeNr(conn SQL.Verbindung,tabelle, attributName, attributsWert, nrName string) (bool, string) {
	var nrWert string
	var eintragWarVorhanden bool
	// 1. Tabelle prüfen ob Attributswert vorhanden, sonst neu erstellen
	anfrage := "SELECT "+ nrName + " FROM " + tabelle + " WHERE "+attributName +"='"+ attributsWert +"';"
	sT:= sqlTabelle.New(conn,anfrage)
	//fmt.Println(sT.GibTabelle())
	// Lese den Wert für die Zahl aus
	nrWert = sT.GibTabelle()[0][0]
	// Prüfe ob es den Wert gab
	if nrWert == "Kein Ergebnis für diese Suchanfrage" {
		// Nr herausfinden die benutzt werden kann
		anfrage = "SELECT * FROM " + tabelle + ";"
		sT := sqlTabelle.New(conn, anfrage)
		// Transponieren um an Spalte zu kommen
		transponiert := transponiere(sT.GibTabelle())
		// Finde freie Nummer
		for i:=1;i<1000;i++ {
			if !enthalten(transponiert[0],fmt.Sprint(i)) {		// [0] weil da die nr steht
				nrWert = fmt.Sprint(i)
				//fmt.Println("Nummer gefunden:",fmt.Sprint(nrWert))
				break
			}
		}
		eintragWarVorhanden = false
		return eintragWarVorhanden,nrWert
	}
	eintragWarVorhanden = true
	return eintragWarVorhanden,nrWert
}

// Vor.: NPC ist vorhanden
func findeNpcNummer(conn SQL.Verbindung, npcname string)string {
	anfrage := "SELECT npcnr FROM npcs WHERE npcname='" + npcname + "';"
	return sqlTabelle.New(conn,anfrage).GibTabelle()[0][0]
}

// Funktion zum Updaten/Ändern des Lieblingsgetränks
func updateLieblingsgetraenk(conn SQL.Verbindung) {
	// Wie Löschen, nur, dass am Ende UPDATE ausgeführt wird anstatt DELETE
	
	var npcname felder.Feld
	var npcnameS string
	
	// Tabelle anzeigen
	anfrage := "SELECT npcname FROM npcs NATURAL JOIN dozent_innen;"
	textboxTabelle.ZeichneAnfrage(conn,anfrage,20,200,true,0,0,0,0,0,255,16,font)
	
	felder.Voreinstellungen(230,230,230,20)
	npcname = felder.New (10,  50, 20, 'l', "Name")	
	// Eingabe in das Feld
	npcnameS = npcname.Edit ()
	
	var npcnrS string //int64
	var lieblingsgetraenk  felder.Feld
	var lieblingsgetraenkS string
	
	lieblingsgetraenk   = felder.New (10,  90, 20, 'l', "Lieblingsgetränk")
	lieblingsgetraenkS = lieblingsgetraenk.Edit()
	// 2. Mit Anfrage in npcs die npcnr herausbekommt 
	anfrage = "SELECT npcnr FROM npcs WHERE npcname = '"
	anfrage += npcnameS
	anfrage += "';"
	sT := sqlTabelle.New(conn,anfrage)		// Sende Anfrage
	npcnrS = sT.GibTabelle()[0][0]		// Tabelle sollte nur ein Ergebnis haben
	
	// 3. 
	// Lieblingsgetränk ändern
	anfrage = "UPDATE dozent_innen SET lieblingsgetraenk ='"+ lieblingsgetraenkS + "' WHERE npcnr =" + npcnrS + ";"
	conn.Ausfuehren(anfrage)
	
	// Namen und Lieblingsgetränk anzeigen
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	anfrage = "SELECT npcname,lieblingsgetraenk FROM npcs NATURAL JOIN dozent_innen;"
	textboxTabelle.ZeichneAnfrage(conn,anfrage,20,200,true,0,0,0,0,0,255,16,font) 
	
	gfx.TastaturLesen1()
}

// Fügt NPC hinzu, prüft selbstständig welche Nummer noch frei ist
func prüfeUNDfügeHinzuNPCs(conn SQL.Verbindung) {
	/*	Ziel: etwas in eine Tabelle eintragen. Schlüssel wird automatisch bestimmt
	 * 1. Lies ganze Tabelle und zeige diese
	 * 2. Prüfe welche Zahl als nächstes als npcnr verfügbar ist, beginnend bei 1
	 * 3. Schreibe neuen Datensatz in die Tabelle
	 * 4. Zeige neue Tabelle
	 */
	 
	 // 1. 
	 anfrage := "SELECT * FROM npcs;"
	 sT := sqlTabelle.New(conn,anfrage)
	 textboxTabelle.ZeichneAnfrage(conn,anfrage,20,200,true,0,0,0,0,0,255,16,font)
	 
	 // 2.
	 var transponiert [][]string
	 transponiert = transponiere(sT.GibTabelle())
	 fmt.Println(transponiert[0])
	 
	var npcname,lieblingsgetraenk felder.Feld
	var npcnameS,lieblingsgetraenkS string
	var npcnrS string //int64
	
	felder.Voreinstellungen(230,230,230,20)

	npcname = felder.New (10,  50, 20, 'l', "NPC Name")
	lieblingsgetraenk   = felder.New (10,  90, 20, 'l', "Lieblingsgetränk")
	
	
	// Eingabe in die Felder
	
	npcnameS = npcname.Edit ()
	lieblingsgetraenkS = lieblingsgetraenk.Edit ()
	
	// Bestimme npcnr
	for i:=1;i<1000;i++ {
		if !enthalten(transponiert[0],fmt.Sprint(i)) {
			npcnrS = fmt.Sprint(i)
			fmt.Println("Nummer gefunden:",fmt.Sprint(npcnrS))
			break
		}
	}
	
	// 3. 	
	// Senden der Anfragen 
	eingabe := fmt.Sprintf(`
		INSERT INTO npcs
		VALUES (%s,'%s');`, npcnrS,npcnameS)
	conn.Ausfuehren(eingabe)
	println ("Neue Werte wurden eingefügt!")
	
	eingabe = fmt.Sprintf(`
			INSERT INTO dozent_innen 
			VALUES (%s,'%s');`, npcnrS,lieblingsgetraenkS)
		// fmt.Println(eingabe)
		
	conn.Ausfuehren(eingabe)
	println ("Neue Werte wurden eingefügt!")
	
	// 4. Tabelle nochmal neu zeichnen
	
	textboxTabelle.ZeichneAnfrage(conn,anfrage,20,200,true,0,0,0,0,0,255,16,font)
	
	gfx.TastaturLesen1 ()
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	anfrage = "SELECT * FROM dozent_innen NATURAL JOIN npcs;"
	textboxTabelle.ZeichneAnfrage(conn,anfrage,20,200,true,0,0,0,0,0,255,16,font)
	gfx.TastaturLesen1 ()
}

// Schön wenn diese Funktin mit verschiedenen Fkt, funktionieren würde, derzeit nur mit Löschen
// Man könnte jetzt aber verschiedene Löschszenarien (mit 1 Feld) definieren
func löscheAusTabelle(conn SQL.Verbindung,Änderung func(SQL.Verbindung,string),feldname string) {
	var feld felder.Feld
	var feldS string
	
	// Tabelle anzeigen
	anfrage := "SELECT npcname FROM npcs NATURAL JOIN dozent_innen;"
	textboxTabelle.ZeichneAnfrage(conn,anfrage,20,200,true,0,0,0,0,0,255,16,font)
	
	felder.Voreinstellungen(230,230,230,20)
	feld = felder.New (10,  50, 20, 'l', feldname)	
	// Eingabe in das Feld
	feldS = feld.Edit ()
	
	Änderung(conn,feldS)
	// Namen der Npcs anzeigen
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	anfrage = "SELECT npcname FROM npcs NATURAL JOIN dozent_innen;"
	textboxTabelle.ZeichneAnfrage(conn,anfrage,20,200,true,0,0,0,0,0,255,16,font) 
	
	gfx.TastaturLesen1()
	
}


// Löscht dozent_innen aus allen verlinkten relationen
func löscheNPC(conn SQL.Verbindung,npcnameS string) {
	// 1. Dozentennamen entgegennehmen
	// 2. Mit Anfrage in npcs die npcnr herausbekommt 
	// 3. Jetzt in dozent_innen, aufenthaltsorte,unterricht, assistenz löschen
	// 4. Namen der Npcs anzeigen
	 
	 // 1. Dozentennamen entgegennehmen
	// Variable für Nummer
	var npcnrS string //int64
	
	// 2. Mit Anfrage in npcs die npcnr herausbekommt 
	anfrage := "SELECT npcnr FROM npcs WHERE npcname = '"
	anfrage += npcnameS
	anfrage += "';"
	sT := sqlTabelle.New(conn,anfrage)		// Sende Anfrage
	npcnrS = sT.GibTabelle()[0][0]		// Tabelle sollte nur ein Ergebnis haben
	
	// 3. 
	// Die Reihenfolge ist wichtig, man kann nicht zuerst aus npcs löschen
	var löschenListe []string = []string{"unterricht","assistenz","dozent_innen","aufenthaltsorte","npcs"}
	for _,t := range löschenListe {
		anfrage = "DELETE FROM " + t +" WHERE npcnr =" + npcnrS + ";"
		conn.Ausfuehren(anfrage)
	}	 
}



/*
func fügeDozentHinzu(conn SQL.Verbindung) {
	var npcnr,npcname,lieblingsgetraenk felder.Feld
	var npcnameS,lieblingsgetraenkS string
	var npcnrS string //int64
	felder.Voreinstellungen(230,230,230,20)
	
	npcnr = felder.New (10,  10, 20, 'l', "NPC Nummer")		// Position 10/10; Länge von 30 Zeichen; linksbündig; Name des Feldes
	//vname.SetzeHintergrundfarbe(255,0,0)
	npcname = felder.New (10,  50, 20, 'l', "NPC Name")
	lieblingsgetraenk   = felder.New (10,  90, 20, 'l', "Lieblingsgetränk")
	//plz.SetzeErlaubteZeichen (felder.Digits)
	
	
	// Eingabe in die Felder
	
	npcnrS = npcnr.Edit ()
	npcnameS = npcname.Edit ()
	lieblingsgetraenkS = lieblingsgetraenk.Edit ()
	
	// Senden der Anfragen 
	eingabe := fmt.Sprintf(`
		INSERT INTO npcs
		VALUES (%s,'%s');`, npcnrS,npcnameS)
	conn.Ausfuehren(eingabe)
	println ("Neue Werte wurden eingefügt!")
	
	eingabe = fmt.Sprintf(`
			INSERT INTO dozent_innen 
			VALUES (%s,'%s');`, npcnrS,lieblingsgetraenkS)
		// fmt.Println(eingabe)
		
	conn.Ausfuehren(eingabe)
	println ("Neue Werte wurden eingefügt!")
	
	gfx.TastaturLesen1 ()
}
*/


// transponiert einen 2d-Slice aus Strings
func transponiere(tabelle [][]string) [][]string{
	var spalten int
	spalten = len(tabelle[0])
	
	transponiert := make([][]string,spalten) 
	
	for i:=0;i<len(tabelle);i++ {
		for j:=0;j<len(tabelle[0]);j++ {
			transponiert[j] = append(transponiert[j], tabelle[i][j]) 
		}
	}
	return transponiert
}

// Prüft ob ein String in einer Liste von Strings enthalten ist, wenn ja, true
func enthalten(liste []string,e string) bool{
	//fmt.Println("Element: ",e)
	for _,s := range liste {
		//fmt.Println(s)
		if s == e {return true}
	}
	return false
}
