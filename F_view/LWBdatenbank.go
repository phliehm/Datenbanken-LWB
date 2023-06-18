package main
// Autor: A. Cyriacus, M. Seiss, P. Liehm, B. Schneider
// Datum: 15.06.2023
// Zweck: DBP - LWB - Datenbank
//--------------------------------------------------------------------

import ( 	."gfx"
			"fmt"
			"sync"
			"time"
			"strings"
			"felder"
			"../Klassen/buttons"
			"SQL"
			"../Klassen/textboxTabelle"
			"../Klassen/sqlTabelle"
			"../Klassen/textboxen"
			"os"
			"path/filepath"
			"os/exec"
		)
/*
EINTRÄGE HINZUFÜGEN:
-Veranstaltung hinzufügen:
Name, Thema, Kürzel, Dozentin, SWS, Semester, Raum
-Dozentin anlegen:
Name, Lieblingsgetränk

EDIT
-Veranstaltungen: Raumnummer, Dozentin

ORDER:
- Highscores: Sortieren/Eingrenzen:
"Notenbereich", "Punktebereich"
minNote, maxNote, minPunkte,maxPunkte
- um Note/Punkte Button für Sortierung

---
-- SQL-ANFRAGEN
-- ------------

-- 1a. Welche Räume gibt es in der LWB-Adventure-World?
\C '1a. Welche Räume gibt es in der LWB-Adventure-World?'
SELECT * FROM raeume;

-- 1b. Welche Aufgaben haben die sonstigen NPCs im LWB-Adventure?
\C '1b. Welche Aufgaben haben die sonstigen NPCs im LWB-Adventure?'
SELECT aufgabe FROM sonstigenpcs;


-- 2a. Welche Lehrveranstaltungen haben 6 SWS?
\C '2a. Welche Lehrveranstaltungen haben 6 SWS?'
SELECT * FROM veranstaltungen WHERE sws = 6;

-- 2b. Welche Lehrveranstaltungen gibt es im 4. Semester?
\C '2b. Welche Lehrveranstaltungen gibt es im 4. Semester?'
SELECT * FROM veranstaltungen WHERE semester = 4;

-- 2c. Welche Minigames gibt es im 4. Semester?
\C '2c. Welche Minigames gibt es im 4. Semester?'
SELECT * FROM minigames NATURAL JOIN veranstaltungen WHERE semester = 4;


-- 3a. Wie heißen die Spieler_innen, die bisher das LWB-Adventure gespielt haben?
\C '3a. Wie heißen die Spieler_innen, die bisher das LWB-Adventure gespielt haben?'
SELECT spname FROM spieler_innen;

-- 3b. Wie heißen die Dozenten im LWB-Adventure?
\C '3b. Wie heißen die Dozenten im LWB-Adventure?'
SELECT npcname FROM dozent_innen NATURAL JOIN npcs;

-- 3c. Welche Aufgabe hat NPC 'Heidi'?
\C '3c. Welche Aufgabe hat NPC "Heidi"?'
SELECT aufgabe FROM sonstigenpcs NATURAL JOIN npcs WHERE npcname = 'Heidi';


-- 4a. Welche Lehrveranstaltungen gehören zum Themengebiet 'Programmierung'?
\C '4a. Welche Lehrveranstaltungen gehören zum Themengebiet "Programmierung"?'
SELECT * FROM veranstaltungen NATURAL JOIN themengebiete WHERE gebietname = 'Programmierung';

-- 4b. Welche Lehrveranstaltungen haben etwas mit 'Daten' oder 'Programmierung' zu tun?
\C '4b. Welche Lehrveranstaltungen haben etwas mit "Daten" oder "Programmierung" zu tun?'
SELECT * FROM veranstaltungen WHERE vname LIKE '%Daten%' OR vname LIKE '%Programmierung%';

-- oder
\C 'oder mit Suche in ThemengebietName:'
SELECT * FROM veranstaltungen NATURAL JOIN themengebiete WHERE gebietname LIKE '%Daten%' OR gebietname LIKE '%Programmierung%';


-- 5. Was ist das Lieblingsgetränk von Darth Schmidter?
\C '5. Was ist das Lieblingsgetränk von Darth Schmidter?'
SELECT lieblingsgetraenk FROM dozent_innen NATURAL JOIN npcs WHERE npcname = 'Darth Schmidter';


-- 6. Welche Lehrveranstaltungen finden nicht in der 'FU Berlin' statt?
\C '6. Welche Lehrveranstaltungen finden nicht in der "FU Berlin" statt?'
SELECT vname, semester, ort FROM raeume NATURAL JOIN unterricht NATURAL JOIN veranstaltungen WHERE ort != 'FU Berlin';

-- alternativ (falls man nicht sicher ist, ob 'FU Berlin' die vollständige Orts-Bezeichnung ist:
\C 'alternativ mit Wildcards (falls man nicht sicher ist, ob "FU Berlin" die vollständige Orts-Bezeichnung ist:'
SELECT vname, semester, ort FROM raeume NATURAL JOIN unterricht NATURAL JOIN veranstaltungen WHERE ort NOT LIKE '%FU%Berlin%';


-- 7. Welche Dozenten sind in der LWB nur leitend tätig und machen keine Assistenz?
\C '7. Welche Dozenten sind in der LWB nur leitend tätig und machen keine Assistenz?'
SELECT npcname FROM npcs NATURAL JOIN (SELECT npcnr FROM dozent_innen EXCEPT SELECT npcnr FROM assistenz) AS xyz;
-- Kommentar: 	Hier braucht es einen Alias, damit der NATURAL JOIN mit der Unterabfrage funktioniert.
--				Die Bezeichnung ist jedoch egal, da nur auf den NPCNamen projiziert wird.


-- Anfragen, die nur mit erweiterter relationaler Algebra beschrieben werden können:
-- ---------------------------------------------------------------------------------

-- 8. Wieviele Mini-Games gibt es in der LWB-Adventure-World? (Ausgaben-Titel: AnzahlMinigames)
\C '8. Wieviele Mini-Games gibt es in der LWB-Adventure-World? (Ausgaben-Titel: AnzahlMinigames)'
SELECT COUNT(*) AS AnzahlMinigames FROM minigames;


-- 9. Wieviele SWS müssen in der LWB ingesamt absolviert werden? (Ausgaben-Titel: GesamtanzahlSWS)
\C '9. Wieviele SWS müssen in der LWB ingesamt absolviert werden? (Ausgaben-Titel: GesamtanzahlSWS)'
SELECT SUM(sws) AS GesamtanzahlSWS FROM veranstaltungen;


-- 10. Wie heißt die Veranstaltung mit den meisten SWS?
\C '10. Wie heißt die Veranstaltung mit den meisten SWS?'
SELECT vname FROM veranstaltungen WHERE sws = (SELECT MAX(sws) FROM veranstaltungen);


-- 11. Gesucht sind die Namen, Semester und SWS aller Veranstaltungen von Winnie the K absteigend sortiert nach SWS-Anzahl!
\C '11. Gesucht sind die Namen, Semester und SWS aller Veranstaltungen von Winnie the K absteigend sortiert nach SWS-Anzahl!'
SELECT vname, sws, semester FROM veranstaltungen NATURAL JOIN unterricht NATURAL JOIN npcs WHERE npcname = 'Winnie the K' ORDER BY sws DESC;


-- 12. Wieviele Veranstaltungen gibt es pro Standort?
\C '12. Wieviele Veranstaltungen gibt es pro Standort?'
SELECT ort, COUNT(*) AS AnzahlVeranstaltungen FROM raeume NATURAL JOIN unterricht GROUP BY ort ORDER BY COUNT(*);


--13. Welche Spieler_innen haben einen Gesamt-Notendurchschnitt, der nicht zwischen 2.0 und 4.0 liegt? (Sortierung nach Gesamt-Notendurchschnitt aufsteigend, also bester Schnitt zuerst)
\C '13. Welche Spieler_innen haben einen Gesamt-Notendurchschnitt, der nicht zwischen 2.0 und 4.0 liegt? (Sortierung nach Gesamt-Notendurchschnitt aufsteigend, also bester Schnitt zuerst)'
SELECT SpName FROM spieler_innen NATURAL JOIN spielstaende GROUP BY spname HAVING AVG(note) NOT BETWEEN 2.0 AND 4.0 ORDER BY AVG(note),spname;

-- oder mit Ausgabe des jeweiligen Notendurchschnitts:
\C 'oder mit Ausgabe des jeweiligen Notendurchschnitts:'
SELECT SpName, ROUND(AVG(note),2) AS Notendurchschnitt FROM spieler_innen NATURAL JOIN spielstaende GROUP BY spname HAVING AVG(note) NOT BETWEEN 2.0 AND 4.0 ORDER BY AVG(note),spname;
*/

var Mutex sync.Mutex					// erstellt Mutex
	
var BuZurueck buttons.Button				// Spezille Knoepfe
var KatKnoepfe, HinzuKnoepfe, VeranstKnoepfe, SpielstKnoepfe, SQLAnfrKnoepfe []buttons.Button		// Slices für alle erstellten Knöpfe / die Suchfelder
var SQLAnfrFeld felder.Feld
var VeranstFelder, SpielstFelder, VeranstHinzuFelder, DozentHinzuFelder, MinispielHinzuFelder, HighscoreFelder []felder.Feld

var Ende bool = false							// True gdw. Programm beenden
var Anfrage, Suchwort string							// Durchsuchen/Suchwort-String
var Raumnummer uint8						// Raumnummer des momentanen Raumes
var MinNote,MaxNote,MinPunkte,MaxPunkte, Raumnr, Doz string
var Katknopftexte, Hinzuknopftexte,VeranstaltungFeldtexte []string
var conn SQL.Verbindung

var font string = "../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf"
var Aufgaben = []string{
		"1a. Welche Räume gibt es in der LWB-Adventure-World?",
		"1b. Welche Aufgaben haben die sonstigen NPCs im LWB-Adventure?",
		"2a. Welche Lehrveranstaltungen haben 6 SWS?",
		"2b. Welche Lehrveranstaltungen gibt es im 4. Semester",
		"2c. Welche Minigames gibt es im 4. Semester?",
		"3a. Wie heißen die Spieler_innen, die bisher das LWB-Adventure gespielt haben?",
		"3b. Wie heißen die Dozenten im LWB-Adventure?",
		"3c. Welche Aufgabe hat NPC 'Heidi'?",
		"4a. Welche Lehrveranstaltungen gehören zum Themengebiet 'Programmierung'?",
		"4b. Welche Lehrveranstaltungen haben etwas mit 'Daten' oder 'Programmierung' zu tun?",
		"5. Was ist das Lieblingsgetränk von Darth Schmidter?",
		"6. Welche Lehrveranstaltungen finden nicht in der 'FU Berlin' statt?",
		"7. Welche Dozenten sind in der LWB nur leitend tätig und machen keine Assistenz?",
		"8. Wieviele Mini-Games gibt es in der LWB-Adventure-World? (Ausgaben-Titel: AnzahlMinigames)",
		"9. Wieviele SWS müssen in der LWB ingesamt absolviert werden? (Ausgaben-Titel: GesamtanzahlSWS)",
		"10. Wie heißt die Veranstaltung mit den meisten SWS?",
		"11. Gesucht sind Namen, Semester und SWS aller Veranstaltungen von Winnie the K absteigend sortiert nach SWS-Anzahl!",
		"12. Wieviele Veranstaltungen gibt es pro Standort?",
		"13. Welche Spieler_innen haben einen Gesamt-Notendurchschnitt, der nicht zwischen 2.0 und 4.0 liegt?",
		"       (Sortierung nach Gesamt-Notendurchschnitt aufsteigend, also bester Schnitt zuerst)" }

func main () {
	Fenster (1200, 700)
	Fenstertitel(" ###  LWB - Datenbank  ###")
	SetzeFont ("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",20)
	
	ErstelleTexte()
	ErstelleKnoepfe()
	ErstelleFelder()
	
	
	ZeichneRaum()

	// --------------------- Verbindung zur Datenbank -----------------
	conn = SQL.PgSQL ("user=lewein dbname=lewein")
	defer conn.Beenden ()
	fmt.Println("Verbindung hergestellt.\n")
	
	// Das Hauptprogramm startet die View-Komponente als nebenläufigen Prozess!
	// go view_komponente()
	
	// Nebenläufig wird die Kontroll-Komponente für die Maus gestartet.
	go maussteuerung()
	
	for !Ende { time.Sleep(1e9) }
	//TastaturLesen1 ()
}

func ErstelleTexte() {
	Katknopftexte = append(Katknopftexte, 
				" Beenden", 
				"Veranstaltungen", 
				"  Spielstände", 
				" !!! RESET der Datenbank !!!", 
				" LWB-Übersicht", 
				" Aufgaben", 
				" freie SQL-Anfrage", 
				"  Neuer Listen-Eintrag" )
	
	Hinzuknopftexte = append(Hinzuknopftexte, 
			"        -> Veranstaltung <-     NEU   hinzufügen", 
			"          -> Dozent/in <-       NEU   hinzufügen",
			"          -> Minispiel <-       NEU   hinzufügen", 
			"          -> Spieler/in <-      NEU   hinzufügen", 
			"  Eine ->  kurze  <- Pause einlegen und prokrastinieren!" )
			
	// VeranstaltungFeldtexte = append(VeranstaltungFeldtexte, "NEUE Veranstaltung", "Thema", "SWS", "Raum","Dozent/in")
}

func ErstelleKnoepfe() {
	BuZurueck 	= buttons.New(20,20,200,70, 255,151,196, true, 	"  zurück")					// zurück
	

	KatKnoepfe = append(KatKnoepfe,
				buttons.New(1000,620,180,70, 255,151,196, true, Katknopftexte[0]),
				buttons.New(130,330,300,70, 246,109,237, true, Katknopftexte[1]),
				buttons.New(130,430,300,70, 246,109,237, true, Katknopftexte[2]),		
				buttons.New(100,570,600,80, 230,50,100, true, Katknopftexte[3]),		
				buttons.New(100,150,500,130, 255,193,46, true, Katknopftexte[4]),
				buttons.New(650,150,340,130, 100,230,50, true, Katknopftexte[5]),
				buttons.New(550,310,500,100, 50,100,230, true, Katknopftexte[6]),
				buttons.New(500,440,610,100, 255,248,23, true, Katknopftexte[7]) )
	
	SQLAnfrKnoepfe = append(SQLAnfrKnoepfe,
				buttons.New(20,110,550,50, 0,255,0, true, "     Neue SQL-Anfrage eingeben"),
				buttons.New(600,110,550,50, 0,255,0, true, "     Bestehende Listen anzeigen"),
				BuZurueck )
	
	HinzuKnoepfe = append(HinzuKnoepfe,
				buttons.New(20,140,1160,80, 255,132,198, true, Hinzuknopftexte[0]),		// 
				buttons.New(20,250,1160,80, 255,234,122, true, Hinzuknopftexte[1]),		// 
				buttons.New(20,360,1160,80, 182,249,148, true, Hinzuknopftexte[2]),		// 
				buttons.New(20,470,1160,80, 210,128,240, true, Hinzuknopftexte[3]),		// 
				buttons.New(20,580,1160,80, 135,250,223, true, Hinzuknopftexte[4])	)	// 
	
	SpielstKnoepfe = append(SpielstKnoepfe,
				buttons.New(20,105,350,55, 0,255,0, true, 	"Spielstände durchsuchen"),
				buttons.New(390,105,290,55, 0,255,0, true, 	"Highscores anzeigen"),
				buttons.New(690,105,210,55, 0,255,0, true, 	" Notenbereich"),
				buttons.New(910,105,220,55, 0,255,0, true, 	" Punktebereich"),
				BuZurueck )
	
	VeranstKnoepfe = append(VeranstKnoepfe,
				buttons.New(20,105,450,55, 0,255,0, true, 	"  Veranstaltungen durchsuchen"),
				buttons.New(520,105,290,55, 0,255,0, true, 	"   Eintrag ändern"),
				buttons.New(860,105,290,55, 0,255,0, true, 	"  Eintrag löschen"),
				BuZurueck )
}

func ErstelleFelder() {
	felder.Voreinstellungen(0,255,0,20)
	SQLAnfrFeld = felder.New (25,  120, 115, 'l', " Stelle neue SQL-Anfrage")
	
	VeranstHinzuFelder = append( VeranstHinzuFelder,
		felder.New (40, 160, 40, 'l', "NEUE Veranstaltung"),	
		felder.New (470, 160, 30, 'l', "Thema"),
		felder.New (790, 160, 2, 'l', "SWS"),
		felder.New (820, 160, 3, 'l', "Raum"),
		felder.New (900, 160, 25, 'l', "Dozent/in")	)
	
	DozentHinzuFelder = append( DozentHinzuFelder,
		felder.New (40, 160, 40, 'l', "NEUER Name Dozent/in"),	
		felder.New (470, 160, 30, 'l', "Lieblingsgetränk"),	)
	
	MinispielHinzuFelder = append( MinispielHinzuFelder,
		felder.New (40, 160, 40, 'l', "NEUER Minispiel-Name"),	
		felder.New (470, 160, 30, 'l', "zugeordnete Veranstaltung")	)
	
	felder.Voreinstellungen(0,255,0,32)
	VeranstFelder = append( VeranstFelder,
		felder.New (110,  110, 30, 'l', "Durchsuche Veranstaltungen"),	
		felder.New (80, 110, 30, 'l', "Bestehende Veranstaltung"),
		felder.New (590, 110, 6, 'l', "Neue Raumnummer"),
		felder.New (720, 110, 25, 'l', "Neue/r Dozent/in")	)
		
	SpielstFelder = append( SpielstFelder,
		felder.New (80,  110, 30, 'l', "Durchsuche Spielstände"),
		felder.New (700, 110, 3, 'l', "min. Note"),
		felder.New (790, 110, 3, 'l', "max. Note"),	
		felder.New (920, 110, 4, 'l', "min. Punkte"),
		felder.New (1000, 110, 4, 'l', "max. Punkte")	)

}
func ZeichneKatKnoepfe() {
	for _,bu := range KatKnoepfe {
		bu.ZeichneButton()
	}
}
func AktiviereKatKnoepfe() {
	for _,bu := range KatKnoepfe {
		bu.AktiviereButton()
	}
}
func DeaktiviereKatKnoepfe() {
	for _,bu := range KatKnoepfe {
		bu.DeaktiviereButton()
	}
}
func ZeichneHinzuKnoepfe() {
	for _,bu := range HinzuKnoepfe {
		bu.ZeichneButton()
	}
}
func AktiviereHinzuKnoepfe() {
	for _,bu := range HinzuKnoepfe {
		bu.AktiviereButton()
	}
}
func DeaktiviereHinzuKnoepfe() {
	for _,bu := range HinzuKnoepfe {
		bu.DeaktiviereButton()
	}
}

func ZeichneRaum() {
	DeaktiviereKatKnoepfe()
	DeaktiviereHinzuKnoepfe()

	UpdateAus () 										// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
	Stiftfarbe(255,255,255)
	Cls()												// Cleart vollständigen Screen
	SetzeFont ("../Schriftarten/Ubuntu-B.ttf", 80 )
	Stiftfarbe(100,100,100)
	felder.Voreinstellungen(0,255,0,32)
	switch Raumnummer {
		case 0:	
		SchreibeFont(300,10,"LWB - Datenbank")
		AktUndZeichne(KatKnoepfe)

		case 1:
		SchreibeFont(300,10,Katknopftexte[1])
		AktUndZeichne(VeranstKnoepfe)
		
		textboxTabelle.ZeichneAnfrage(conn,sucheDozVer(""),20,170,true,0,0,0,0,0,255,16,font)
		BuZurueck.ZeichneButton()
		case 2:
		SchreibeFont(300,10,Katknopftexte[2])
		AktUndZeichne(SpielstKnoepfe)
		
		textboxTabelle.ZeichneAnfrage(conn,sucheSpielerGamesScores(""),20,170,true,0,0,0,0,0,255,16,font)
		BuZurueck.ZeichneButton()
		case 3:
		SchreibeFont(300,10,Katknopftexte[3])
		resetDatenbank()
		SetzeFont(font,50)
		Stiftfarbe(255,0,0)
		SchreibeFont(100,300,"DIE DATENBANK IST JETZT WIEDER WIE NEU!")
		BuZurueck.ZeichneButton()
		case 4:
		SchreibeFont(300,10,Katknopftexte[4])
		// Räume	
		textboxTabelle.ZeichneAnfrage(conn,gibAnfrageRäume(),20,170,false,0,0,0,0,0,255,16,font)
		// DozentInnen
		textboxTabelle.ZeichneAnfrage(conn,gibAnfrageDozentInnen(),800,170,false,0,0,0,0,0,255,16,font)
		// sonstige NPCs
		textboxTabelle.ZeichneAnfrage(conn,gibAnfrageSonstigeNPCs(),800,400,false,0,0,0,0,0,255,16,font)
		// Minigames
		textboxTabelle.ZeichneAnfrage(conn,gibAnfrageMinigames(),20,400,false,0,0,0,0,0,255,16,font)
		BuZurueck.ZeichneButton()
		case 8:											// --> 	Aufgaben
		SchreibeFont(300,10,Katknopftexte[5])
		
		//SetzeFont ("../Schriftarten/Ubuntu-B.ttf", 18 )
		SetzeFont ("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",20)
		for i,aufgabe := range Aufgaben {
			SchreibeFont(20,120+27*uint16(i),aufgabe)
		}
		BuZurueck.ZeichneButton()
		case 9:											// --> freie SQL-Anfrage
		SchreibeFont(300,10,Katknopftexte[6])
		AktUndZeichne(SQLAnfrKnoepfe)
		case 10:										// --> Eintrag hinzufügen
		SchreibeFont(250,10,"Eintrag hinzufügen")
		AktiviereHinzuKnoepfe()
		ZeichneHinzuKnoepfe()
		BuZurueck.ZeichneButton()
	}
	UpdateAn () 										// Nun wird der gezeichnete Frame sichtbar gemacht!
}

func AktUndZeichne(knoepfe []buttons.Button) {
	for _,knopf := range knoepfe {
		knopf.AktiviereButton()
		knopf.ZeichneButton()
	}
}

// Es folgt die Maus-Komponente 1 --- Kein Bestandteil der Welt, also unabhängig -----
func maussteuerung () {

	for {
		_, status, mausX, mausY := MausLesen1()
		
		if status==1 { 													// Maustaste gedrückt
			
			// ------------------------------------- KNOPF-FUNKTIONEN --------------------------------------------------------------------- AB HIER
			Stiftfarbe(255,255,255)
			switch Raumnummer {
				case 1:																			// ------------------ Veranstaltungen-Raum
				if VeranstKnoepfe[0].TesteXYPosInButton(mausX,mausY) {							// -- Durchsuche die Liste
					Vollrechteck(20,105,1160,60)
					Suchwort = VeranstFelder[0].Edit()
					Anfrage = sucheDozVer(Suchwort)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(20,170,1180,530)
					
					Vollrechteck(20,105,1160,60)
					AktUndZeichne(VeranstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				} else if VeranstKnoepfe[1].TesteXYPosInButton(mausX,mausY) {					// -- Eintrag ändern
					Vollrechteck(20,105,1160,60)
					Suchwort := VeranstFelder[1].Edit()			// veranstaltungsname
					Raumnr = VeranstFelder[2].Edit()			// raumnummer
					Doz = VeranstFelder[3].Edit()				// dozentIn
					
					ändereInVeranstaltungen(conn , Suchwort, Raumnr, Doz)
					
					Suchwort = ""
					Anfrage = sucheDozVer(Suchwort)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(0,105,1200,595)
					AktUndZeichne(VeranstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				} else if VeranstKnoepfe[2].TesteXYPosInButton(mausX,mausY) {		// -- Eintrag löschen
					Vollrechteck(20,105,1160,60)
					Suchwort = VeranstFelder[1].Edit()		// Veranstaltung
					
					löscheVeranstaltung(conn,Suchwort)
					
					Suchwort = ""
					Anfrage = sucheDozVer(Suchwort)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(0,105,1200,595)
					AktUndZeichne(VeranstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				}
				case 2: 																		// ------------------ Spielstände-Raum
				if SpielstKnoepfe[0].TesteXYPosInButton(mausX,mausY) {							// -- Durchsuche die Liste
					Vollrechteck(20,105,1160,60)
					Suchwort = SpielstFelder[0].Edit()
					Anfrage = sucheSpielerGamesScores(Suchwort)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(20,170,1180,530)
					Vollrechteck(20,105,1160,60)
					AktUndZeichne(SpielstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
					
				} else if SpielstKnoepfe[1].TesteXYPosInButton(mausX,mausY) {					// -- Highscores
					Vollrechteck(0,170,1200,530)				
					Anfrage = gibAnfrageHighscore()
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
					
				} else if SpielstKnoepfe[2].TesteXYPosInButton(mausX,mausY) {					// -- Notenbereich eingeben
					Vollrechteck(20,105,1160,60)
					
					MinNote = SpielstFelder[1].Edit()
					MaxNote = SpielstFelder[2].Edit()
					Anfrage = gibAnfrageScoresNotenbereich(MinNote,MaxNote)
										
					/*Suchwort = ""
					Anfrage = sucheSpielerGamesScores(Suchwort)
					*/
					Stiftfarbe(255,255,255)
					Vollrechteck(20,170,1180,530)
					AktUndZeichne(SpielstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				} else if SpielstKnoepfe[3].TesteXYPosInButton(mausX,mausY) {					// -- Punktebereich eingeben
					Vollrechteck(20,105,1160,60)
					
					MinPunkte = SpielstFelder[3].Edit()
					MaxPunkte = SpielstFelder[4].Edit()
					
					Anfrage = gibAnfrageScoresPunktebereich(MinPunkte,MaxPunkte)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
					/*
					Suchwort = ""
					Anfrage = sucheSpielerGamesScores(Suchwort)
					*/
					Stiftfarbe(255,255,255)
					Vollrechteck(20,170,1180,530)
					AktUndZeichne(SpielstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				}
				
				case 4:								// LWB-Übersicht
				/*
				// Räume	
				textboxTabelle.ZeichneAnfrage(conn,gibAnfrageRäume(),20,170,false,0,0,0,0,0,255,16,font)
				// DozentInnen
				textboxTabelle.ZeichneAnfrage(conn,gibAnfrageDozentInnen(),800,170,false,0,0,0,0,0,255,16,font)
				// sonstige NPCs
				textboxTabelle.ZeichneAnfrage(conn,gibAnfrageSonstigeNPCs(),800,400,false,0,0,0,0,0,255,16,font)
				// Minigames
				textboxTabelle.ZeichneAnfrage(conn,gibAnfrageMinigames(),20,400,false,0,0,0,0,0,255,16,font)
				*/
				case 9:
				if SQLAnfrKnoepfe[0].TesteXYPosInButton(mausX,mausY) {									// ------- freie SQL-Anfrage
					Stiftfarbe(255,255,255)
					Vollrechteck(20,110,1160,60)
							
					Anfrage = SQLAnfrFeld.Edit()
					
					if prüfeFreieSqlAnfrage(Anfrage) == false {
					// 	----------------- Fehler in der SQL Anfrage ----------
						SchreibeFont(100,200,"Diese Anfrage ist ungültig. Achten Sie auf korrekte Relationennamen und Schlüsselwörter.")
						break
						}
					
					Stiftfarbe(255,255,255)
					Vollrechteck(20,200,1160,500)
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(20,110,1160,60)
					SQLAnfrKnoepfe[0].ZeichneButton()
					SQLAnfrKnoepfe[1].ZeichneButton()
				} else if SQLAnfrKnoepfe[1].TesteXYPosInButton(mausX,mausY) {
					zeigeAlleRelationen()
					
				}
				case 10:
				for _,suchknopf := range HinzuKnoepfe { 									// überprüft EINFÜGEKNÖPFE im Array
					if suchknopf.TesteXYPosInButton(mausX,mausY) {

						switch suchknopf.GibBeschriftung() {
							
							case Hinzuknopftexte[0]: 
							Vollrechteck(20,140,1160,110)
							var veranstaltungsAttribute []string
							var veranstaltungsString string
							
							// Lies die einzelnen Eingabefelder aus und schreibe sie in einen Slice
							for i,feldwert := range VeranstHinzuFelder {
								if i == 2 || i == 3 {feldwert.SetzeErlaubteZeichen (felder.Digits)} // sws, raumnr
								wert := feldwert.Edit()
								veranstaltungsAttribute = append(veranstaltungsAttribute, wert)
								veranstaltungsString = veranstaltungsString + " , " + wert
							}
							// Füge Eintrag hinzu
							hinzugefügt := fügeHinzuVeranst(conn,veranstaltungsAttribute)
							
							Stiftfarbe(255,255,255)
							Vollrechteck(0,140,1200,560)
							
							Stiftfarbe(255,132,198)
							switch hinzugefügt{
								case true: SchreibeFont(40,220,"Eintrag hinzugefügt: " + veranstaltungsString[2:])
								case false: 
								Stiftfarbe(255,0,0)
								SchreibeFont(40,220,"Kein Eintrag hinzugefügt. Fehlerhafte Eingabe!")
							}
							
							AktUndZeichne(HinzuKnoepfe)
							
							case Hinzuknopftexte[1]:	
							Vollrechteck(20,250,1160,80)
							Stiftfarbe(255,234,122)
							
							case Hinzuknopftexte[2]: 
							Vollrechteck(20,360,1160,80)
							Stiftfarbe(182,249,148)
							
							case Hinzuknopftexte[3]:	
							Vollrechteck(20,470,1160,80)
							Stiftfarbe(210,128,240)
							
							case Hinzuknopftexte[4]:	
							Vollrechteck(20,580,1160,80)
							Stiftfarbe(135,250,223)
						}
					}
				}
			}
			// ------------------------------------- KNOPF-FUNKTIONEN --------------------------------------------------------------------- BIS HIER
			
			// ------------------------------- RAUM-WECHSEL ---------------------------------------------------- AB HIER
			for _,knopf := range KatKnoepfe { 									// überprüft Knöpfe im Array
				if knopf.TesteXYPosInButton(mausX,mausY) {
					switch knopf.GibBeschriftung() {
						case Katknopftexte[0]: Ende = true; return
						case Katknopftexte[1]: Raumnummer = 1
						case Katknopftexte[2]: Raumnummer = 2						
						case Katknopftexte[3]: Raumnummer = 3
						case Katknopftexte[4]: Raumnummer = 4
						case Katknopftexte[5]: Raumnummer = 8
						case Katknopftexte[6]: Raumnummer = 9
						case Katknopftexte[7]: Raumnummer = 10
					}
					ZeichneRaum()					// Raum wurde gewechselt und muss neu gezeichnet werden
				}
			}
			if BuZurueck.TesteXYPosInButton(mausX,mausY) {					// Zurück-Botton gedrückt
				Raumnummer = 0
				ZeichneRaum()					// Raum wurde gewechselt und muss neu gezeichnet werden
			}
			// ------------------------------- RAUM-WECHSEL ---------------------------------------------------- BIS HIER
		}
	}
}

// Funktion für Suchfeld für Dozentinnen und Veranstaltungen
func sucheDozVer(suchwort string) string {
	Anfrage = "SELECT vname AS Veranstaltung,gebietname AS Thema,sws,raumnr AS Raumnummer,npcname AS DozentIn FROM veranstaltungen NATURAL JOIN dozent_innen NATURAL JOIN npcs NATURAL JOIN unterricht NATURAL JOIN themengebiete  WHERE CONCAT(npcname,gebietname,vname,raumnr,sws) LIKE '%"
	Anfrage += suchwort
	Anfrage += "%' ORDER BY raumnr LIMIT 27;"
	return Anfrage
}

// Funktion für SpielerInnen und Games, Scores
func sucheSpielerGamesScores(suchwort string) string {
	Anfrage := "SELECT spname AS SpielerIn,gamename AS MiniGame,vname AS veranstaltung,note AS Note,punkte AS Punkte"+
	 " FROM spielstaende NATURAL JOIN minigames NATURAL JOIN veranstaltungen NATURAL JOIN spieler_innen WHERE CONCAT(spname,gamename,vname,note,punkte) LIKE '%"
	Anfrage += suchwort
	Anfrage += "%' LIMIT 27;"
	return Anfrage
}

// Scores mit Notenbereich
func gibAnfrageScoresNotenbereich(min,max string) string{
	Anfrage := "SELECT spname AS Spieler_in,gamename AS MiniGame,vname AS veranstaltung,note AS Note,punkte AS Punkte"+
	 " FROM spielstaende NATURAL JOIN minigames NATURAL JOIN veranstaltungen NATURAL JOIN spieler_innen WHERE CONCAT(spname,gamename,vname,note,punkte) LIKE '%"
	Anfrage += Suchwort
	Anfrage += "%' AND note>="+min+" AND note<="+max+" LIMIT 27;"
	return Anfrage
}

// Scores mit Notenbereich
func gibAnfrageScoresPunktebereich(min,max string) string{
	Anfrage := "SELECT spname AS Spieler_in,gamename AS MiniGame,vname AS veranstaltung,note AS Note,punkte AS Punkte"+
	 " FROM spielstaende NATURAL JOIN minigames NATURAL JOIN veranstaltungen NATURAL JOIN spieler_innen WHERE CONCAT(spname,gamename,vname,note,punkte) LIKE '%"
	Anfrage += Suchwort
	Anfrage += "%' AND punkte>="+min+" AND punkte<="+max+" LIMIT 27;"
	return Anfrage
}

// Gibt Highscore zurück, achtung, Dopplungen
func gibAnfrageHighscore() string {
		//anfrage := "SELECT gamename,note,punkte FROM minigames NATURAL JOIN spielstaende;"
		anfrage := `SELECT t.spname AS spieler_in, t.gamename AS minigame,t.vname AS veranstaltung, t.note,t.punkte
					FROM (minigames NATURAL JOIN spielstaende NATURAL JOIN spieler_innen NATURAL JOIN veranstaltungen) t
					INNER JOIN (
					  SELECT gamename, MIN(punkte) AS min_punkte
					  FROM minigames NaTURAL JOIN spielstaende NATURAL JOIN spieler_innen NATURAL JOIN veranstaltungen
					  GROUP BY gamename
					) AS subquery
					ON t.gamename = subquery.gamename AND t.punkte = subquery.min_punkte ORDER BY t.gamename;`
		return anfrage
}

// Gibt Anfrage für Räume
func gibAnfrageRäume() string {
	return "SELECT raumnr AS raum, raumname,ort,funktion FROM raeume;"
}
// Gibt Anfrage für npcs
func gibAnfrageDozentInnen() string {
	return "SELECT npcname AS name, lieblingsgetraenk FROM npcs NATURAL JOIN dozent_innen;"
}
// Gibt sonstige NPCs
func gibAnfrageSonstigeNPCs() string {
	return "SELECT npcname AS name, aufgabe FROM npcs NATURAL JOIN sonstigenpcs;"
}
// Gibt Minigames
func gibAnfrageMinigames() string {
	return "SELECT gamename AS minigame,vname,raumname FROM minigames NATURAL JOIN raeume NATURAL JOIN veranstaltungen NATURAL JOIN unterricht; "
}

//////////////////////////////////
// EINFÜGEN VON VERANSTALTUNGEN //
//////////////////////////////////

func fügeHinzuVeranst(conn SQL.Verbindung,attribute []string) bool {
	//Name, Thema, Kürzel, Dozentin, SWS, Semester, Raum
	// 	VeranstaltungFeldtexte = append(VeranstaltungFeldtexte, "NEUE Veranstaltung", "Thema", "SWS", "Raum","Dozent/in")
	//zeigeVeranst(conn)
	var eingabe string
	var vnrS,vnameS,gebietnameS,kuerzelS,npcnameS,swsS,semesterS,raumnrS string
	var gebietnrS,npcnrS string // um gebietnummer zu finden
	var eintragWarVorhanden bool
	
	// Prüfe ob in allen Attributen etwas steht
	if enthalten(attribute,"") {
		fmt.Println("Keine valider Eintrag")
		return false
	}
	
	vnameS = attribute[0]
	gebietnameS = attribute[1]
	kuerzelS = "?"
	npcnameS = attribute[4]
	swsS = attribute[2]
	semesterS ="1"
	raumnrS = attribute[3]
	
	//fmt.Println(vnameS,gebietnameS,kuerzelS,npcnameS,swsS,semesterS,raumnrS)
	// mögliche Probleme
	// existiert Dozent? Existiert Thema?
	
	
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
	if eintragWarVorhanden {return false}		// Wenn es den Eintrag schon gab, mache nichts
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
	//zeigeVeranst(conn)
	return true
}


// Zeigt die wesentlichen Attribute von Veranstaltungen						
func zeigeVeranst(conn SQL.Verbindung) {
	anfrage := "SELECT vname,gebietname,kuerzel,npcname,sws,semester,raumnr FROM veranstaltungen"+
				" NATURAL JOIN unterricht NATURAL JOIN npcs NATURAL JOIN themengebiete ORDER BY raumnr;"
	textboxTabelle.ZeichneAnfrage(conn,anfrage,10,200,true,0,0,0,0,0,255,16,font) 
				
}

// Prüft ob ein Attributswert existiert, wenn nicht, ist eine freie Nummer zurückgeliefert, 
// Wenn ja, ist die zugehörige Nummer geliefert.
// Vor.: Die benötigte Nummer ist in der ersten Spalte
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

func löscheVeranstaltung(conn SQL.Verbindung, veranstaltung string) {
	// 1. Aus unterrichten löschen, dann veranstaltungen
	
	vorhanden,vnr := prüfeObVorhandenFindeNr(conn ,"veranstaltungen", "vname", veranstaltung, "vnr")
	fmt.Println("vnr: ",vnr,vorhanden)
	// Wenn die Veranstaltung gar nicht existiert
	if vorhanden == false {return}
	
	// Es soll keine Vorlesung die mit einem Minigame verknüpft ist gelöscht werden
	vorhanden,_ = prüfeObVorhandenFindeNr(conn ,"minigames", "vnr", vnr, "vnr")
	fmt.Println("vnr: ",vnr,vorhanden)
	if vorhanden == true {return}
	
	// Es soll keine Vorlesung die mit einem Minigame verknüpft ist gelöscht werden
	vorhanden,_ = prüfeObVorhandenFindeNr(conn ,"assistenz", "vnr", vnr, "vnr")
	fmt.Println("vnr: ",vnr,vorhanden)
	if vorhanden == true {return}
	
	// Lösche die Veranstaltung aus unterricht
	eingabe := "DELETE FROM unterricht WHERE vnr=" + vnr + ";"
	conn.Ausfuehren(eingabe)
	
	// Lösche aus veranstaltungen
	eingabe = "DELETE FROM veranstaltungen WHERE vnr=" + vnr + ";"
	conn.Ausfuehren(eingabe)

}

func ändereInVeranstaltungen(conn SQL.Verbindung, veranstaltung, raumnr,doz string) {
	var npcnr, eingabe string
	vorhanden,vnr := prüfeObVorhandenFindeNr(conn ,"veranstaltungen", "vname", veranstaltung, "vnr")
	// Wenn die Veranstaltung gar nicht existiert
	if vorhanden == false {return}
	// Prüfe ob Dozent vorhanden ist, aber nur wenn ein Name gelierfert wurde
	if len(doz)>0 {
		vorhanden,npcnr = prüfeObVorhandenFindeNr(conn ,"npcs", "npcname", doz, "npcnr")
		// Wenn die Dozent gar nicht existiert
		if vorhanden == false {return}
		// Ändere dozentIn
		eingabe = "UPDATE unterricht SET npcnr=" + npcnr + " WHERE vnr= "+ vnr +";"
		conn.Ausfuehren(eingabe)
	}
	if len(raumnr)>0 {
		vorhanden,raumnr = prüfeObVorhandenFindeNr(conn ,"raeume", "raumnr", raumnr, "raumnr")
		// Wenn es den Raum gar nicht gibt, mache nichts
		if vorhanden == false {return}
		// Ändere Raumnr in unterricht
		fmt.Println(raumnr,vnr)
		eingabe = "UPDATE unterricht SET raumnr=" + raumnr + " WHERE vnr="+vnr +";"
		conn.Ausfuehren(eingabe)
		
	}
}

// Liefert false wenn die Anfrage offensichtlich falsch ist, sonst true
func prüfeFreieSqlAnfrage(anfrage string) bool {
	if len(anfrage)==0 {return false}
	tabellen := []string{"aufenthaltsorte","dozent_innen","minigames","npcs","raeume","sonstigenpcs",
						"spieler_innen","spielstaende","themengebiete","unterricht","veranstaltungen"}
	schlüsselWörter := []string{"SELECT","UPDATE","DELETE","INSERT"}				
	
	// Teste ob überhaupt eine Tabelle verwendet wird
	var enthalten bool
	for _,t := range tabellen {
		enthalten = strings.Contains(anfrage,t)		
		if enthalten == true {break}
	}
	if enthalten == false {return false}

	// Teste ob ein Schlüsselwort verwendet wird
	for _,sw := range schlüsselWörter {
		enthalten = strings.Contains(anfrage,sw)
		if enthalten == true {break}
	}
	if enthalten == false {return false}
	
	return true
}

// Zeichnet eine Textbox mit allen Relationen und Attributen
func zeigeAlleRelationen() {
	tB := textboxen.New(600,200,1100,500)
	tB.SchreibeText(
	"Relationen (Attribute)\n\n"+
	"assistenz (vnr,npcnr);\n"+
	"aufenthaltsorte (npcnr,raumnr);\n"+
	"dozent_innen (npcnr,lieblingsgetraenk);\n"+
	"minigames (gamenr,gamename,vnr);\n"+
	"npcs (npcnr,npcname);\n"+
	"raeume (raumnr,raumname,ort,funktion);\n"+
	"sonstigenpcs (npcnr,aufgabe);\n"+
	"themengebiete (gebietnr,gebietname);\n"+
	"unterricht (vnr,npcnr,raumnr);\n"+
	"veranstaltungen (vnr,vname,kuerzel,sws,semester,gebietnr);\n")
	tB.SetzeZeilenAbstand(10)
	tB.Zeichne()
}

// Stellt die Datenbank wieder auf den Ursprungszustand her
func resetDatenbank() {
	wd,_ := os.Getwd()					// Get Pfad
	relativePath := "../D_sql"			// relativer Pfad wo ich hin möchte
	combinedPath := filepath.Join(wd,relativePath)	// Kombiniere beide Pfade
	absolutePath,_:= filepath.Abs(combinedPath)			// nimmt .. weg
	os.Chdir(absolutePath)				// Gehe zum neuen Pfad
	cmd := exec.Command("bash", "-c", "psql -U lewein -d lewein -f Install-LWBadventure.sql")	// erstelle ausführbares Objekt
	cmd.Stdout = os.Stdout			// sende output an Konsole
	cmd.Stderr = os.Stderr
	cmd.Run()						// Führe aus
	os.Chdir(wd)					// Setze Pfad auf Ausgangspfad zurück

}

