package main
// Autor: A. Cyriacus, M. Seiss, P. Liehm, B. Schneider
// Datum: 15.06.2023
// Zweck: DBP - LWB - Datenbank
//--------------------------------------------------------------------

import ( 	."gfx"
			"fmt"
			"sync"
			//"time"
			"felder"
			"../Klassen/buttons"
			"SQL"
			"../Klassen/textboxTabelle"
			"../Klassen/sqlTabelle"
		)

var Mutex sync.Mutex					// erstellt Mutex
	
var Knoepfe, Suchknoepfe []buttons.Button		// Slices für alle erstellten Knöpfe / die Suchfelder
var Suchfelder []felder.Feld
var BuZurueck,BuEintrag,BuEnde buttons.Button
var Akt bool = true							// True gdw. Raum gewechselt wurde _----- NICHT BENÖTIGT?!
var Raumnummer uint8						// Raumnummer des momentanen Raumes
var Knopftexte, Suchknopftexte []string
var conn SQL.Verbindung

var font string = "../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf"


func main () {
	Fenster (1200, 700)
	Fenstertitel(" ###  LWB - Datenbank  ###")
	SetzeFont ("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",20)
	
	Knopftexte = append(Knopftexte, "", "Dozenten", "Minispiele", "Fachgebiete", "Semester")
	Suchknopftexte = append(Suchknopftexte, "Vorlesung", "Thema", "SWS", "Raum","Dozent/in")
	
	BuZurueck 	= buttons.New(10,620,200,70, 230,50,100, true, " zurück")				// zurück
	BuEintrag 	= buttons.New(400,620,350,70, 230,50,100, true, " Erstelle Eintrag")	// Erstelle Eintrag
	BuEnde 	= buttons.New(1000,620,180,70, 230,50,100, true, " Beenden")				// Beenden
	
	ErstelleKnoepfe()
	ErstelleSuchknoepfe()
	ErstelleSuchFelder()
	
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
	for { }
	TastaturLesen1 ()
}

func ErstelleKnoepfe() {
	bu1			:= buttons.New(100,150,300,70, 230,50,100, true, Knopftexte[1])		// Dozenten
	bu2			:= buttons.New(100,250,300,70, 230,50,100, true, Knopftexte[2])		// Minispiele
	bu3			:= buttons.New(100,350,300,70, 230,50,100, true, Knopftexte[3])		// Fachgebiete
	bu4			:= buttons.New(100,450,300,70, 230,50,100, true, Knopftexte[4])		// Semester
	Knoepfe = append(Knoepfe, bu1, bu2, bu3, bu4)
}
func ErstelleSuchknoepfe() {
	su1		:= buttons.New(100,150,260,50, 230,50,100, true, Suchknopftexte[0])		// 
	su2		:= buttons.New(390,150,180,50, 230,50,100, true, Suchknopftexte[1])		// 
	su3		:= buttons.New(590,150,80,50, 230,50,100, true, Suchknopftexte[2])		// 
	su4		:= buttons.New(700,150,80,50, 230,50,100, true, Suchknopftexte[3])		// 
	su5		:= buttons.New(850,150,150,50, 230,50,100, true, Suchknopftexte[4])		// 
	Suchknoepfe = append(Suchknoepfe, su1,su2,su3,su4,su5)
}
func ErstelleSuchFelder() {
	fe1 := felder.New (110, 160, 30, 'l', Suchknopftexte[0])		
	fe2 := felder.New (400, 160, 20, 'l', Suchknopftexte[1])
	fe3 := felder.New (600, 160, 5, 'l', Suchknopftexte[2])
	fe4 := felder.New (710, 160, 5, 'l', Suchknopftexte[3])
	fe5 := felder.New (860, 160, 15, 'l', Suchknopftexte[4])
	Suchfelder = append(Suchfelder, fe1,fe2,fe3,fe4,fe5)
}
func ZeichneKnoepfe() {
	for _,bu := range Knoepfe {
		bu.ZeichneButton()
	}
}
func AktiviereKnoepfe() {
	for _,bu := range Knoepfe {
		bu.AktiviereButton()
	}
}
func DeaktiviereKnoepfe() {
	for _,bu := range Knoepfe {
		bu.DeaktiviereButton()
	}
}
func ZeichneSuchknoepfe() {
	for _,bu := range Suchknoepfe {
		bu.ZeichneButton()
	}
}
func AktiviereSuchknoepfe() {
	for _,bu := range Suchknoepfe {
		bu.AktiviereButton()
	}
}
func DeaktiviereSuchknoepfe() {
	for _,bu := range Suchknoepfe {
		bu.DeaktiviereButton()
	}
}

func ZeichneRaum() {
	UpdateAus () 										// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
	Stiftfarbe(255,255,255)
	Cls()												// Cleart vollständigen Screen
	SetzeFont ("../Schriftarten/Ubuntu-B.ttf", 80 )
	Stiftfarbe(100,100,100)
	switch Raumnummer {
		case 0:	
		SchreibeFont(300,50,"LWB - Datenbank")
		ZeichneKnoepfe()
		BuEintrag.ZeichneButton()
		BuEnde.ZeichneButton()
		case 1:
		SchreibeFont(300,50,Knopftexte[1])
		
		var suche felder.Feld
		var suchwort string
		felder.Voreinstellungen(0,255,0,32)
	
		suche = felder.New (10,  10, 30, 'l', "Suche")
		
		suchwort = suche.Edit()
	
		anfrage := sucheDozVer(suchwort)
		textboxTabelle.ZeichneAnfrage(conn,anfrage,100,200,true,0,0,0,0,0,255,16,font)
	
		BuZurueck.ZeichneButton()
		case 2:
		SchreibeFont(300,50,Knopftexte[2])
		BuZurueck.ZeichneButton()
		case 3:
		SchreibeFont(300,50,Knopftexte[3])
		BuZurueck.ZeichneButton()
		case 4:
		SchreibeFont(300,50,Knopftexte[4])
		BuZurueck.ZeichneButton()
		case 10:										// Eintrag hinzufügen
		SchreibeFont(300,50,"Eintrag hinzufügen")
		ZeichneSuchknoepfe()
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
			for _,knopf := range Knoepfe { 									// überprüft Knöpfe im Array
				if knopf.TesteXYPosInButton(mausX,mausY) {
					fmt.Println("Knopf gedrückt: ", knopf.GibBeschriftung() )
					switch knopf.GibBeschriftung() {
						case Knopftexte[1]: 
						DeaktiviereKnoepfe()
						Raumnummer = 1
						case Knopftexte[2]:	
						DeaktiviereKnoepfe()
						Raumnummer = 2
						case Knopftexte[3]: 
						DeaktiviereKnoepfe()
						Raumnummer = 3
						case Knopftexte[4]:	
						DeaktiviereKnoepfe()
						Raumnummer = 4						
					}
					ZeichneRaum()					// Raum wurde gewechselt und muss neu gezeichnet werden
				}
			}
			for _,suchknopf := range Suchknoepfe { 									// überprüft SUCHKNÖPFE im Array
				if suchknopf.TesteXYPosInButton(mausX,mausY) {
					fmt.Println("Suchknopf gedrückt: ", suchknopf.GibBeschriftung() )
					Stiftfarbe(255,255,255)
					switch suchknopf.GibBeschriftung() {
						case Suchknopftexte[0]: 
						Vollrechteck(100,150,260,50)
						Suchfelder[0].Edit()
						case Suchknopftexte[1]:	
						Vollrechteck(390,150,180,50)
						Suchfelder[1].Edit()
						case Suchknopftexte[2]: 
						Vollrechteck(590,150,80,50)
						Suchfelder[2].Edit()
						case Suchknopftexte[3]:	
						Vollrechteck(700,150,80,50)
						Suchfelder[3].Edit()
						case Suchknopftexte[4]:	
						Vollrechteck(850,150,150,50)
						Suchfelder[4].Edit()						
					}
				}
			}
			if BuZurueck.TesteXYPosInButton(mausX,mausY) {
				DeaktiviereSuchknoepfe()
				AktiviereKnoepfe()
				Raumnummer = 0
				ZeichneRaum()					// Raum wurde gewechselt und muss neu gezeichnet werden
			} else if BuEintrag.TesteXYPosInButton(mausX,mausY) {
				DeaktiviereKnoepfe()
				AktiviereSuchknoepfe()
				Raumnummer = 10
				ZeichneRaum()					// Raum wurde gewechselt und muss neu gezeichnet werden
			}
		}
	}
}

func zeichneAnfrage(conn SQL.Verbindung,anfrage string) {
	Stiftfarbe(255,255,255)
	//Cls()
	sT := sqlTabelle.New(conn,anfrage)
	//fmt.Println(sT.GibTabelle())
	
	// Nur zum Testen auch SQL Anfrage anzeigen
	Stiftfarbe(0,0,0)
	Schreibe(400,200,anfrage)
	
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





