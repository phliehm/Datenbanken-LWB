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
	
var BuZurueck,BuEnde buttons.Button				// Spezille Knoepfe
var KatKnoepfe, HinzuKnoepfe, VeranstKnoepfe, SpielstKnoepfe []buttons.Button		// Slices für alle erstellten Knöpfe / die Suchfelder
var DurchsucheFeld felder.Feld
var VeranstFelder, SpielstFelder, VeranstHinzuFelder, HighscoreFelder []felder.Feld

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
	Katknopftexte = append(Katknopftexte, "", "Veranstaltungen", "  Spielstände", "   Dummy", " LWB-Übersicht", " Aufgaben", " freie SQL-Anfrage", " Neuer Listen-Eintrag")
	Hinzuknopftexte = append(Hinzuknopftexte, "       Neue Veranstaltung hinzufügen", "       Neues  -> Minispiel <-  hinzufügen", "  SWS hinzufügen", "  Raum hinzufügen","Dozent/in hinzufügen")
	VeranstaltungFeldtexte = append(VeranstaltungFeldtexte, "NEUE Veranstaltung", "Thema", "SWS", "Raum","Dozent/in")
}

func ErstelleKnoepfe() {
	BuZurueck 	= buttons.New(20,20,200,70, 230,50,100, true, 	" zurück")					// zurück
	BuEnde 		= buttons.New(1000,620,180,70, 230,50,100, true," Beenden")					// Beenden
	
	KatKnoepfe = append(KatKnoepfe,
				buttons.New(100,150,300,70, 230,50,100, true, Katknopftexte[1]),
				buttons.New(100,250,300,70, 230,50,100, true, Katknopftexte[2]),		
				buttons.New(100,350,300,70, 230,50,100, true, Katknopftexte[3]),		
				buttons.New(100,450,300,70, 230,50,100, true, Katknopftexte[4]),
				buttons.New(620,150,330,130, 100,230,50, true, Katknopftexte[5]),
				buttons.New(550,320,500,100, 50,100,230, true, Katknopftexte[6]),
				buttons.New(500,450,600,100, 130,150,100, true, Katknopftexte[7])   )
	
	
	HinzuKnoepfe = append(HinzuKnoepfe,
				buttons.New(20,120,1160,70, 230,50,100, true, Hinzuknopftexte[0]),		// 
				buttons.New(20,220,1160,70, 230,50,100, true, Hinzuknopftexte[1]),		// 
				buttons.New(20,320,1160,70, 230,50,100, true, Hinzuknopftexte[2]),		// 
				buttons.New(20,420,1160,70, 230,50,100, true, Hinzuknopftexte[3]),		// 
				buttons.New(20,520,1160,70, 230,50,100, true, Hinzuknopftexte[4])	)	// 
	
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
	felder.Voreinstellungen(0,255,0,32)
	
	DurchsucheFeld = felder.New (110,  110, 30, 'l', "Durchsuche")		// durchsuchen
	
	VeranstFelder = append( VeranstFelder,
		felder.New (110,  110, 30, 'l', "Durchsuche Veranstaltungen"),	
		felder.New (80, 110, 30, 'l', "Bestehende Veranstaltung"),
		felder.New (590, 110, 6, 'l', "Neue Raumnummer"),
		felder.New (720, 110, 25, 'l', "Neue/r Dozent/in")	)
		
	SpielstFelder = append( SpielstFelder,
		felder.New (110,  110, 30, 'l', "Durchsuche Spielstände"),	
		felder.New (80, 110, 30, 'l', "minPunkte"),
		felder.New (590, 110, 6, 'l', "maxPunkte"),
		felder.New (720, 110, 25, 'l', "minNote"),
		felder.New (720, 110, 25, 'l', "maxNote")	)
	
	VeranstHinzuFelder = append( VeranstHinzuFelder,
		felder.New (20, 160, 25, 'l', VeranstaltungFeldtexte[0]),	
		felder.New (450, 160, 20, 'l', VeranstaltungFeldtexte[1]),
		felder.New (770, 160, 3, 'l', VeranstaltungFeldtexte[2]),
		felder.New (820, 160, 4, 'l', VeranstaltungFeldtexte[3]),
		felder.New (880, 160, 15, 'l', VeranstaltungFeldtexte[4])	)
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
		
		BuEnde.ZeichneButton()
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
		BuZurueck.ZeichneButton()
		case 4:
		SchreibeFont(300,10,Katknopftexte[4])
		BuZurueck.ZeichneButton()
		case 8:											// --> 	Aufgaben
		SchreibeFont(300,10,Katknopftexte[5])
		BuZurueck.ZeichneButton()
		case 9:											// --> freie SQL-Anfrage
		SchreibeFont(300,10,Katknopftexte[6])
		BuZurueck.ZeichneButton()
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
		// maus.SetzeKoordinaten(mausX,mausY)								// Aktualisiert Maus-Koordinaten
		
		if status==1 { 													// Maustaste gedrückt
			
			// ------------------------------- RAUM-WECHSEL ---------------------------------------------------- AB HIER
			for _,knopf := range KatKnoepfe { 									// überprüft Knöpfe im Array
				if knopf.TesteXYPosInButton(mausX,mausY) {
					fmt.Println("Knopf gedrückt: ", knopf.GibBeschriftung() )
					switch knopf.GibBeschriftung() {
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
			
			// ------------------------------- KNOPF-FUNKTIONEN ------------------------------------------------ AB HIER
			Stiftfarbe(255,255,255)
			switch Raumnummer {
				case 1:																					// ------------------ Veranstaltungen-Raum
				if VeranstKnoepfe[0].TesteXYPosInButton(mausX,mausY) {				// -- Durchsuche die Liste
					Vollrechteck(20,105,1160,60)
					Suchwort = VeranstFelder[0].Edit()
					Anfrage = sucheDozVer(Suchwort)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(20,170,1180,530)
					
					Vollrechteck(20,105,1160,60)
					AktUndZeichne(VeranstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				} else if VeranstKnoepfe[1].TesteXYPosInButton(mausX,mausY) {		// -- Eintrag ändern
					Vollrechteck(20,105,1160,60)
					Suchwort = VeranstFelder[1].Edit()
					VeranstFelder[2].Edit()
					VeranstFelder[3].Edit()
					
					// ----------------------------------- HIER fehlt die SQL-Anfrage zum Ändern des Eintrags
					
					Anfrage = sucheDozVer(Suchwort)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(20,105,1160,60)
					AktUndZeichne(VeranstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				} else if VeranstKnoepfe[2].TesteXYPosInButton(mausX,mausY) {		// -- Eintrag löschen
					Vollrechteck(20,105,1160,60)
					Suchwort = VeranstFelder[1].Edit()
									
					// ----------------------------------- HIER fehlt die SQL-Anfrage zum Löschen des Eintrags
					
					Suchwort = ""
					Anfrage = sucheDozVer(Suchwort)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(20,105,1160,60)
					AktUndZeichne(VeranstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				}
				case 2: 																					// ------------------ Spielstände-Raum
				if VeranstKnoepfe[0].TesteXYPosInButton(mausX,mausY) {				// -- Durchsuche die Liste
					Vollrechteck(20,105,1160,60)
					Suchwort = VeranstFelder[0].Edit()
					Anfrage = sucheSpielerGamesScores(Suchwort)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(20,170,1180,530)
					
					Vollrechteck(20,105,1160,60)
					AktUndZeichne(SpielstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				} else if VeranstKnoepfe[1].TesteXYPosInButton(mausX,mausY) {		// -- Highscores
					Vollrechteck(20,105,1160,60)
					Suchwort = VeranstFelder[1].Edit()
					VeranstFelder[2].Edit()
					VeranstFelder[3].Edit()
					
					// ----------------------------------- HIER fehlt die SQL-Anfrage zum Ändern des Eintrags
					
					Anfrage = sucheDozVer(Suchwort)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(20,105,1160,60)
					AktUndZeichne(SpielstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				} else if VeranstKnoepfe[2].TesteXYPosInButton(mausX,mausY) {		// -- Notenbereich eingeben
					Vollrechteck(20,105,1160,60)
					Suchwort = VeranstFelder[1].Edit()
									
					// ----------------------------------- HIER fehlt die SQL-Anfrage zum Löschen des Eintrags
					
					Suchwort = ""
					Anfrage = sucheDozVer(Suchwort)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(20,105,1160,60)
					AktUndZeichne(SpielstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				} else if VeranstKnoepfe[3].TesteXYPosInButton(mausX,mausY) {		// -- Punktebereich eingeben
					Vollrechteck(20,105,1160,60)
					Suchwort = VeranstFelder[1].Edit()
									
					// ----------------------------------- HIER fehlt die SQL-Anfrage zum Löschen des Eintrags
					
					Suchwort = ""
					Anfrage = sucheDozVer(Suchwort)
					
					Stiftfarbe(255,255,255)
					Vollrechteck(20,105,1160,60)
					AktUndZeichne(SpielstKnoepfe)
					
					textboxTabelle.ZeichneAnfrage(conn,Anfrage,20,170,true,0,0,0,0,0,255,16,font)
				}
			
				case 10:
				for _,suchknopf := range HinzuKnoepfe { 									// überprüft SUCHKNÖPFE im Array
					if suchknopf.TesteXYPosInButton(mausX,mausY) {
						fmt.Println("Einfüge-Knopf gedrückt: ", suchknopf.GibBeschriftung() )
						switch suchknopf.GibBeschriftung() {
							case Hinzuknopftexte[0]: 
							Vollrechteck(100,120,1000,70)
							var veranstaltungsAttribute []string
							// Lies die einzelnen Eingabefelder aus und schreibe sie in einen Slice
							for _,feldwert := range VeranstHinzuFelder {
								veranstaltungsAttribute = append(veranstaltungsAttribute,feldwert.Edit())
							}
							// Füge Eintrag hinzu
							fügeHinzuVeranst(conn,veranstaltungsAttribute)
							
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


//////////////////////////////////
// EINFÜGEN VON VERANSTALTUNGEN //
//////////////////////////////////

func fügeHinzuVeranst(conn SQL.Verbindung,attribute []string) {
	//Name, Thema, Kürzel, Dozentin, SWS, Semester, Raum
	// 	VeranstaltungFeldtexte = append(VeranstaltungFeldtexte, "NEUE Veranstaltung", "Thema", "SWS", "Raum","Dozent/in")
	//zeigeVeranst(conn)
	var eingabe string
	var vnrS,vnameS,gebietnameS,kuerzelS,npcnameS,swsS,semesterS,raumnrS string
	var gebietnrS,npcnrS string // um gebietnummer zu finden
	var eintragWarVorhanden bool
	
	vnameS = attribute[0]
	gebietnameS = attribute[1]
	kuerzelS = "?"
	npcnameS = attribute[4]
	swsS = attribute[2]
	semesterS ="1"
	raumnrS = attribute[3]
	
	fmt.Println(vnameS,gebietnameS,kuerzelS,npcnameS,swsS,semesterS,raumnrS)
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
	//zeigeVeranst(conn)
}


// Zeigt die wesentlichen Attribute von Veranstaltungen						
func zeigeVeranst(conn SQL.Verbindung) {
	anfrage := "SELECT vname,gebietname,kuerzel,npcname,sws,semester,raumnr FROM veranstaltungen"+
				" NATURAL JOIN unterricht NATURAL JOIN npcs NATURAL JOIN themengebiete;"
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




