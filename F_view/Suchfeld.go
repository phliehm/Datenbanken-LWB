/* Autor: Philipp Liehm
 * Datum: Juni 2023
 * Zweck: Suchanfragen in einer Datenbank mit einem Suchfeld
 */

package main

import (
		"fmt"
		"SQL"
		"../Klassen/sqlTabelle"
		"../Klassen/textboxTabelle"
		"../Klassen/textboxen"
		"gfx"
		)
		
var relationenDozVorl []string = []string{"dozent_innen","npcs","raeume","themengebiete","unterricht","veranstaltungen"}
var listeAttributeDozVorl []string = []string{"vname","gebietname","sws","raumnr","npcname"}
/*
 * 1. Anfrage auf alles mit Platzhalter für Suchbegriff formulieren
 * 
 * 2. Feld einbinden, so dass auf Button-Klick das Feld gelesen wird
 * 
 * 3. Ausgabe (reicht auch erstmal die Konsole)
 * 
 * */ 

func main() {
	gfx.Fenster(1200,700)
	
	gfx.SetzeFont("../Schriftarten/terminus-font/Terminus-Bold.ttf",20)
	
	// Verbindungsaufbau
	var conn SQL.Verbindung
	conn = SQL.PgSQL ("user=lewein dbname=lewein")
	defer conn.Beenden ()
	fmt.Println("Verbindung hergestellt.\n")
	
	
	// Anfrage (diese geht)
	//anfrage := "SELECT * FROM dozent_innen NATURAL JOIN  npcs WHERE CONCAT(lieblingsgetraenk,npcname) LIKE '%Kaff%';"
	
	// Hinweis Natural Join: liefert Kartesisches Produkt wenn es keine übereinstimmenden Attribute gibt
	
	suchwort := "Kaff"
	//anfrage := erstelleLikeAnfrage(relationenDozVorl,listeAttributeDozVorl,suchwort)
	
	
	//anfrage := "SELECT * FROM assistenz,aufenthaltsorte ;"
	anfrage := sucheDozVer(suchwort)
	//sT := sqlTabelle.New(conn,anfrage)
	//fmt.Println(sT.GibTabelle())
	zeichneAnfrage(conn,anfrage)
	gfx.TastaturLesen1()
	
	suchwort = "Ba"
	anfrage = sucheSpielerGamesScores(suchwort)
	//sT = sqlTabelle.New(conn,anfrage)
	//fmt.Println(sT.GibTabelle())
	
	zeichneAnfrage(conn,anfrage)
	
	
	
}

// Funktion für Suchfeld für Dozentinnen und Veranstaltungen
func sucheDozVer(suchwort string ) string {
	anfrage := "SELECT vname AS Vorlesung,gebietname AS Thema,sws,raumnr AS Raumnummer,npcname AS DozentIn FROM veranstaltungen NATURAL JOIN dozent_innen NATURAL JOIN npcs NATURAL JOIN unterricht NATURAL JOIN themengebiete  WHERE CONCAT(npcname,gebietname,vname) LIKE '%"
	anfrage += suchwort
	anfrage += "%';"
	return anfrage
}

// Funktion für SpielerInnen und Games, Scores
func sucheSpielerGamesScores(suchwort string) string {
	anfrage := "SELECT spname AS SpielerIn,gamename AS MiniGame,vname AS Vorlesung,note AS Note,punkte AS Punkte"+
	 " FROM spielstaende NATURAL JOIN minigames NATURAL JOIN veranstaltungen NATURAL JOIN spieler_innen WHERE CONCAT(spname,gamename,vname,note,punkte) LIKE '%"
	anfrage += suchwort
	anfrage += "%';"
	return anfrage
}

func zeichneAnfrage(conn SQL.Verbindung,anfrage string) {
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	sT := sqlTabelle.New(conn,anfrage)
	//fmt.Println(sT.GibTabelle())
	
	// Nur zum Testen auch SQL Anfrage anzeigen
	gfx.Stiftfarbe(0,0,0)
	tbAnfrage := textboxen.New(10,10,1100,100)
	tbAnfrage.SetzeFont("../Schriftarten/terminus-font/Terminus-Bold.ttf")
	tbAnfrage.SetzeSchriftgröße(12)
	tbAnfrage.SchreibeText(anfrage)
	tbAnfrage.Zeichne()
	//gfx.Schreibe(10,10,anfrage)
	
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

//////////////////////////
// Doch nicht benötigt? //
//////////////////////////

func erstelleLikeAnfrage(tabellen []string,alleAttribute []string, suchwort string) string {
	anfrage := "SELECT * FROM "
	anfrage += tabellen[0]
	for i,t := range tabellen {
		if i>0 {		// Damit nur bei mehreren Attributen ein Natural Join gemacht wird
			anfrage += " , " + t 
		}
	}
	
	anfrage += " WHERE CONCAT ("
	anfrage += alleAttribute[0]
	for  i,at := range alleAttribute {
		if i>0 {
			anfrage += ","+ at
		}
	}
	anfrage += ") LIKE '%" + suchwort +"%';"
	
	fmt.Println(anfrage)
	return anfrage
}

// Gibt Attribute der ganzen Datenbank zurück, damit kann später eine Suchanfrage gestaltet werden
func liesAlleAttributeTabellen(conn SQL.Verbindung,relationen []string) []string{
	var alleAttribute []string
	for _,relation := range relationen {
		anfrage := "SELECT * FROM "+ relation + " ;"
		sT := sqlTabelle.New(conn,anfrage)
		kopf:= sT.GibKopf()
		// Prüft ob die gefundenen Attribute schon vorhanden sind, wenn nicht, werden sie hinzugefügt
		for _,e := range kopf {
			if beinhaltetWert(alleAttribute,e) == false {
				alleAttribute = append(alleAttribute,e)	// um Dopplungen in der späteren Tabelle zu vermeiden, das führt zu Fehlern
			}
		}
	}
	fmt.Println(alleAttribute)
	return alleAttribute
}

// Prüft ob ein string in dem string slice ist, gibt true oder false zurück
func beinhaltetWert(s []string, e string) bool {
    for _, a := range s {
        if a == e {return true}
    }
    return false
}
