package main

import (
		"gfx"
		"fmt"
		"../Klassen/textboxTabelle"
		"../Klassen/sqlTabelle"
		"SQL"
		)
		
var anfragenSlice []string = []string{
"SELECT note FROM spielstaende;",
"SELECT * FROM raeume;", 
"SELECT * FROM spielstaende;",
"SELECT * FROM veranstaltungen WHERE sws = 4;",
"SELECT npcname FROM dozent_innen NATURAL JOIN npcs;",
"SELECT * FROM veranstaltungen WHERE vname LIKE '%Programmierung';",
"SELECT lieblingsgetraenk FROM dozent_innen,npcs WHERE npcname = 'Herk';",
"SELECT COUNT(*) AS AnzahlMiniGames FROM minigames;",
"SELECT SUM(sws) AS GesamtanzahlSWS FROM veranstaltungen;",
"SELECT vname FROM veranstaltungen WHERE sws = (SELECT MAX(sws) FROM veranstaltungen);",
"SELECT ort, COUNT(*) AS AnzahlVeranstaltungen FROM raeume, veranstaltungen GROUP BY ort ORDER BY COUNT(*);",
"SELECT * FROM raeume, veranstaltungen WHERE ort NOT LIKE '%FU%';",
"SELECT * FROM spielstaende;",
}		

// "SELECT vname, raumnr, sws FROM (veranstaltungen,unterricht),(dozentinnen,npcs) WHERE npcname = 'Winnie the K' ORDER BY sws DESC;",
func main() {
	gfx.Fenster(1200,700)
	gfx.SetzeFont("../Schriftarten/terminus-font/Terminus-Bold.ttf",20)
	// Verbindungsaufbau
	var conn SQL.Verbindung
	conn = SQL.PgSQL ("user=lewein dbname=lewein")
	defer conn.Beenden ()
	fmt.Println("Verbindung hergestellt.\n")
	/*
	var anfrage string
	anfrage = "SELECT * FROM spielstaende ;"
	zeichneAnfrage(conn,anfrage)
	anfrage = "SELECT * FROM veranstaltungen;"
	zeichneAnfrage(conn,anfrage)

	anfrage = "SELECT * FROM aufenthalt;"
	zeichneAnfrage(conn,anfrage)
	anfrage = "SELECT * FROM npcs;"
	zeichneAnfrage(conn,anfrage)
	anfrage = "SELECT * FROM veranstaltungen WHERE vname LIKE '%Programmierung';"
	zeichneAnfrage(conn,anfrage)
	*/
	zeichneAlleAnfragen(conn)
}

func zeichneAnfrage(conn SQL.Verbindung,anfrage string) {
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	sT := sqlTabelle.New(conn,anfrage)
	//fmt.Println(sT.GibTabelle())
	
	// Nur zum Testen auch SQL Anfrage anzeigen
	gfx.Stiftfarbe(0,0,0)
	gfx.Schreibe(10,10,anfrage)
	
	// Textbox Tabelle
	tbT := textboxTabelle.New(sT.GibTabelle(),sT.GibKopf(),50,50)
	tbT.SetzeFarbeTabelle(0,0,0)
	tbT.SetzeZeilenAbstand(1)
	tbT.SetzeSchriftgrößeTabelle(16)
	tbT.SetzeSpaltenAbstand(20)
	tbT.SetzeFarbeKopf(0,0,255)
	tbT.SetzeFontKopf("../Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	tbT.SetzeFontTabelle("../Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	tbT.Zeichne()
	gfx.TastaturLesen1()
	gfx.TastaturLesen1()
	//tbT.VariableBreite()
	/*gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	tbT.Zeichne()
	gfx.TastaturLesen1()
	* */
}

func zeichneAlleAnfragen(conn SQL.Verbindung) {
	for _,anfrage:=range anfragenSlice {
		zeichneAnfrage(conn,anfrage)
	}
}
