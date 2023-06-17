package main
// Autor: A. Cyriacus, M. Seiss, P. Liehm, B. Schneider
// Datum: 15.06.2023
// Zweck: DBP - LWB - Datenbank
//--------------------------------------------------------------------

import ( 	."gfx"
			"fmt"
			"sync"
			"time"
			"felder"
			"../Klassen/buttons"
			"SQL"
			"../Klassen/textboxTabelle"
			"../Klassen/sqlTabelle"
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
*/

var Mutex sync.Mutex					// erstellt Mutex
	
var Katknoepfe, Hinzuknoepfe []buttons.Button		// Slices für alle erstellten Knöpfe / die Suchfelder
var DurchsucheFeld felder.Feld
var VeranstaltungFelder, MinispielFelder []felder.Feld
var BuZurueck,BuEintrag,BuEnde, BuSuche buttons.Button
var Akt bool = true							// True gdw. Raum gewechselt wurde _----- NICHT BENÖTIGT?!
var Anfrage, Suchwort string							// Durchsuchen/Suchwort-String
var Raumnummer uint8						// Raumnummer des momentanen Raumes
var Katknopftexte, Hinzuknopftexte,VeranstaltungFeldtexte []string
var conn SQL.Verbindung

var font string = "../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf"


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

	/*
	vname := felder.New (100,  10, 30, 'l', "Vorname")	
	vn := vname.Edit ()
	fmt.Println(vn)
	
	var vname,nname,str,plz,ort,leer1,leer2 felder.Feld
	var s string


	vname = felder.New (10,  10, 30, 'l', "Vorname")
	nname = felder.New (10,  50, 30, 'r', "Nachname")
	str   = felder.New (10,  90, 30, 'z', "Straße/Hausnummer")
	plz   = felder.New (10, 130,  5, 'l', "PLZ")
	plz.SetzeErlaubteZeichen (felder.Digits)
	ort   = felder.New (10, 170, 30, 'l', "Ort")

	leer1 = felder.New (400, 10, 30, 'l', "")
	leer2 = felder.New (400, 50, 30, 'l', "")

	// Editieren der Eingabefelder
	// gelieferte Zeichenketten werden nicht entgegengenommen ...
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
	*/
	for { time.Sleep(1e9) }
	TastaturLesen1 ()
}

func ErstelleTexte() {
	Katknopftexte = append(Katknopftexte, "", "Veranstaltungen", "Highscores", "Dummy", "LWB-Übersicht")
	Hinzuknopftexte = append(Hinzuknopftexte, "Veranstaltung hinzufügen", "Minispiel hinzufügen", "SWS hinzufügen", "Raum hinzufügen","Dozent/in hinzufügen")
	VeranstaltungFeldtexte = append(VeranstaltungFeldtexte, "Vorlesung", "Thema", "SWS", "Raum","Dozent/in")
}

func ErstelleKnoepfe() {
	BuSuche		= buttons.New(100,110,500,50, 0,255,0, true, 	"   Durchsuche")			// durchsuche
	BuZurueck 	= buttons.New(10,620,200,70, 230,50,100, true, 	" zurück")				// zurück
	BuEintrag 	= buttons.New(400,620,350,70, 230,50,100, true, " Erstelle Eintrag")	// Erstelle Eintrag
	BuEnde 		= buttons.New(1000,620,180,70, 230,50,100, true," Beenden")				// Beenden
	
	Katknoepfe = append(Katknoepfe,
				buttons.New(100,150,300,70, 230,50,100, true, Katknopftexte[1]),
				buttons.New(100,250,300,70, 230,50,100, true, Katknopftexte[2]),		
				buttons.New(100,350,300,70, 230,50,100, true, Katknopftexte[3]),		
				buttons.New(100,450,300,70, 230,50,100, true, Katknopftexte[4])	)
	
	
	Hinzuknoepfe = append(Hinzuknoepfe,
				buttons.New(100,120,1000,70, 230,50,100, true, Hinzuknopftexte[0]),		// 
				buttons.New(100,220,1000,70, 230,50,100, true, Hinzuknopftexte[1]),		// 
				buttons.New(100,320,1000,70, 230,50,100, true, Hinzuknopftexte[2]),		// 
				buttons.New(100,420,1000,70, 230,50,100, true, Hinzuknopftexte[3]),		// 
				buttons.New(100,520,1000,70, 230,50,100, true, Hinzuknopftexte[4])	)	// 
}

func ErstelleFelder() {
	felder.Voreinstellungen(0,255,0,32)
	
	DurchsucheFeld = felder.New (110,  110, 30, 'l', "Durchsuche")		// durchsuchen
	
	VeranstaltungFelder = append( VeranstaltungFelder,
		felder.New (20, 160, 25, 'l', VeranstaltungFeldtexte[0]),	
		felder.New (450, 160, 20, 'l', VeranstaltungFeldtexte[1]),
		felder.New (770, 160, 3, 'l', VeranstaltungFeldtexte[2]),
		felder.New (820, 160, 4, 'l', VeranstaltungFeldtexte[3]),
		felder.New (880, 160, 15, 'l', VeranstaltungFeldtexte[4])	)
}
func ZeichneKatknoepfe() {
	for _,bu := range Katknoepfe {
		bu.ZeichneButton()
	}
}
func AktiviereKatknoepfe() {
	for _,bu := range Katknoepfe {
		bu.AktiviereButton()
	}
}
func DeaktiviereKatknoepfe() {
	for _,bu := range Katknoepfe {
		bu.DeaktiviereButton()
	}
}
func ZeichneHinzuknoepfe() {
	for _,bu := range Hinzuknoepfe {
		bu.ZeichneButton()
	}
}
func AktiviereHinzuknoepfe() {
	for _,bu := range Hinzuknoepfe {
		bu.AktiviereButton()
	}
}
func DeaktiviereHinzuknoepfe() {
	for _,bu := range Hinzuknoepfe {
		bu.DeaktiviereButton()
	}
}

func ZeichneRaum() {
	DeaktiviereKatknoepfe()
	DeaktiviereHinzuknoepfe()
	BuSuche.DeaktiviereButton()

	UpdateAus () 										// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
	Stiftfarbe(255,255,255)
	Cls()												// Cleart vollständigen Screen
	SetzeFont ("../Schriftarten/Ubuntu-B.ttf", 80 )
	Stiftfarbe(100,100,100)
	felder.Voreinstellungen(0,255,0,32)
	switch Raumnummer {
		case 0:	
		SchreibeFont(300,10,"LWB - Datenbank")
		AktiviereKatknoepfe()
		ZeichneKatknoepfe()
		BuEintrag.ZeichneButton()
		BuEnde.ZeichneButton()
		case 1:
		SchreibeFont(300,10,Katknopftexte[1])
		BuSuche.AktiviereButton() 
		BuSuche.ZeichneButton()
		
		textboxTabelle.ZeichneAnfrage(conn,sucheDozVer(""),20,170,true,0,0,0,0,0,255,16,font)
		BuZurueck.ZeichneButton()
		case 2:
		SchreibeFont(300,10,Katknopftexte[2])
		BuSuche.AktiviereButton() 
		BuSuche.ZeichneButton()
		
		textboxTabelle.ZeichneAnfrage(conn,sucheSpielerGamesScores(""),20,170,true,0,0,0,0,0,255,16,font)
		BuZurueck.ZeichneButton()
		case 3:
		SchreibeFont(300,10,Katknopftexte[3])
		BuZurueck.ZeichneButton()
		case 4:
		SchreibeFont(300,10,Katknopftexte[4])
		BuZurueck.ZeichneButton()
		case 10:										// Eintrag hinzufügen
		SchreibeFont(250,10,"Eintrag hinzufügen")
		AktiviereHinzuknoepfe()
		ZeichneHinzuknoepfe()
		BuZurueck.ZeichneButton()
	}
	UpdateAn () 										// Nun wird der gezeichnete Frame sichtbar gemacht!
}

// Es folgt die Maus-Komponente 1 --- Kein Bestandteil der Welt, also unabhängig -----
func maussteuerung () {

	for {
		_, status, mausX, mausY := MausLesen1()
		// maus.SetzeKoordinaten(mausX,mausY)								// Aktualisiert Maus-Koordinaten
		
		if status==1 { 													// Maustaste gedrückt
			
			// ------------------------------- RAUM-WECHSEL ---------------------------------------------------- AB HIER
			for _,knopf := range Katknoepfe { 									// überprüft Knöpfe im Array
				if knopf.TesteXYPosInButton(mausX,mausY) {
					fmt.Println("Knopf gedrückt: ", knopf.GibBeschriftung() )
					switch knopf.GibBeschriftung() {
						case Katknopftexte[1]: Raumnummer = 1
						case Katknopftexte[2]: Raumnummer = 2						
						case Katknopftexte[3]: Raumnummer = 3
						case Katknopftexte[4]: Raumnummer = 4						
					}
					ZeichneRaum()					// Raum wurde gewechselt und muss neu gezeichnet werden
				}
			}
			if BuZurueck.TesteXYPosInButton(mausX,mausY) {
				Raumnummer = 0
				ZeichneRaum()					// Raum wurde gewechselt und muss neu gezeichnet werden
			} else if BuEintrag.TesteXYPosInButton(mausX,mausY) {
				Raumnummer = 10
				ZeichneRaum()					// Raum wurde gewechselt und muss neu gezeichnet werden
			}
			// ------------------------------- RAUM-WECHSEL ---------------------------------------------------- BIS HIER
			
			// ------------------------------- KNOPF-FUNKTIONEN ------------------------------------------------ AB HIER
			Stiftfarbe(255,255,255)
			if BuSuche.TesteXYPosInButton(mausX,mausY) {
				Vollrechteck(100,110,500,60)
				Suchwort = DurchsucheFeld.Edit()
				
				Stiftfarbe(255,255,255)
				Vollrechteck(20,170,1180,530)
				BuZurueck.ZeichneButton()
				
				switch Raumnummer {
					case 1: Anfrage = sucheDozVer(Suchwort)
					case 2: Anfrage = sucheSpielerGamesScores(Suchwort)
				}
				textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
			}
			for _,suchknopf := range Hinzuknoepfe { 									// überprüft SUCHKNÖPFE im Array
				if suchknopf.TesteXYPosInButton(mausX,mausY) {
					fmt.Println("Einfüge-Knopf gedrückt: ", suchknopf.GibBeschriftung() )
					switch suchknopf.GibBeschriftung() {
						case Hinzuknopftexte[0]: 
						Vollrechteck(100,120,1000,70)
						
						for _,feldwert := range VeranstaltungFelder {
							feldwert.Edit()
						}
						case Hinzuknopftexte[1]:	
						
						case Hinzuknopftexte[2]: 
						Vollrechteck(590,150,80,50)
						
						case Hinzuknopftexte[3]:	
						Vollrechteck(700,150,80,50)
						
						case Hinzuknopftexte[4]:	
						Vollrechteck(850,150,150,50)
							
					}
				}
			}
			
		}
	}
}

func zeichneAnfrage(conn SQL.Verbindung) {
	Stiftfarbe(255,255,255)
	//Cls()
	sT := sqlTabelle.New(conn,Anfrage)
	//fmt.Println(sT.GibTabelle())
	
	// Nur zum Testen auch SQL Anfrage anzeigen
	Stiftfarbe(0,0,0)
	Schreibe(400,200,Anfrage)
	
	// Textbox Tabelle
	tbT := textboxTabelle.New(sT.GibTabelle(),sT.GibKopf(),400,250)
	tbT.SetzeFarbeTabelle(0,0,0)
	tbT.SetzeZeilenAbstand(1)
	tbT.SetzeSchriftgrößeTabelle(20)
	tbT.SetzeSpaltenAbstand(20)
	tbT.SetzeFarbeKopf(0,0,255)
	tbT.SetzeFontKopf("../Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	tbT.SetzeFontTabelle("../Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf")
	tbT.Zeichne()
	//TastaturLesen1()
	//TastaturLesen1()
	//tbT.VariableBreite()
	/*gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	tbT.Zeichne()
	gfx.TastaturLesen1()
	* */
}

// Funktion für Suchfeld für Dozentinnen und Veranstaltungen
func sucheDozVer(suchwort string) string {
	Anfrage = "SELECT vname AS Vorlesung,gebietname AS Thema,sws,raumnr AS Raumnummer,npcname AS DozentIn FROM veranstaltungen NATURAL JOIN dozent_innen NATURAL JOIN npcs NATURAL JOIN unterricht NATURAL JOIN themengebiete  WHERE CONCAT(npcname,gebietname,vname) LIKE '%"
	Anfrage += suchwort
	Anfrage += "%';"
	return Anfrage
}

// Funktion für SpielerInnen und Games, Scores
func sucheSpielerGamesScores(suchwort string) string {
	Anfrage := "SELECT spname AS SpielerIn,gamename AS MiniGame,vname AS Vorlesung,note AS Note,punkte AS Punkte"+
	 " FROM spielstaende NATURAL JOIN minigames NATURAL JOIN veranstaltungen NATURAL JOIN spieler_innen WHERE CONCAT(spname,gamename,vname,note,punkte) LIKE '%"
	Anfrage += suchwort
	Anfrage += "%';"
	return Anfrage
}





