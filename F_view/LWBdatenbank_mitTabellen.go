package main
// Autor: A. Cyriacus, M. Seiss, P. Liehm, B. Schneider
// Datum: 15.06.2023
// Zweck: DBP - LWB - Datenbank
//--------------------------------------------------------------------

import ( 	."gfx"
			"fmt"
			"sync"
			"time"
			//"felder"
			"../Klassen/buttons"
			"SQL"
			"../Klassen/textboxTabelle"
			"../Klassen/sqlTabelle"
		)

var Mutex sync.Mutex					// erstellt Mutex
	
var Knoepfe []buttons.Button			// Slice für alle erstellten Knöpfe
var BuZurueck buttons.Button
var Akt bool = true						// True gdw. Raum gewechselt wurde
var Raumnummer uint8					// Raumnummer des momentanen Raumes
var Knopftexte []string
var conn SQL.Verbindung



func main () {
	Fenster (1200, 700)
	Fenstertitel(" ###  LWB - Datenbank  ###")
	SetzeFont ("./Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",20)
	
	Knopftexte = append(Knopftexte, "zurück", "Dozenten", "Minispiele", "Fachgebiete", "Semester")
	
	BuZurueck 	= buttons.New(10,620,200,70, 230,50,100, true, Knopftexte[0])		// zurück
	
	bu1			:= buttons.New(100,150,300,70, 230,50,100, true, Knopftexte[1])		// Dozenten
	bu2			:= buttons.New(100,250,300,70, 230,50,100, true, Knopftexte[2])		// Minispiele
	bu3			:= buttons.New(100,350,300,70, 230,50,100, true, Knopftexte[3])		// Fachgebiete
	bu4			:= buttons.New(100,450,300,70, 230,50,100, true, Knopftexte[4])		// Semester
	Knoepfe = append(Knoepfe, bu1, bu2, bu3, bu4)
	
	SetzeFont("../Schriftarten/terminus-font/Terminus-Bold.ttf",20)
	// Verbindungsaufbau
	
	conn = SQL.PgSQL ("user=lewein dbname=lewein")
	defer conn.Beenden ()
	fmt.Println("Verbindung hergestellt.\n")
	
	
	// Das Hauptprogramm startet die View-Komponente als nebenläufigen Prozess!
	go view_komponente()
	
	// Nebenläufig wird die Kontroll-Komponente für die Maus gestartet.
	go maussteuerung()


	TastaturLesen1 ()
}

func ZeichneKnoepfe(knoepfe []buttons.Button) {
	for _,bu := range knoepfe {
		bu.ZeichneButton()
	}
}
func AktiviereKnoepfe(knoepfe []buttons.Button) {
	for _,bu := range knoepfe {
		bu.AktiviereButton()
	}
}
func DeaktiviereKnoepfe(knoepfe []buttons.Button) {
	for _,bu := range knoepfe {
		bu.DeaktiviereButton()
	}
}



// Es folgt die VIEW-Komponente
func view_komponente () { 
		
	var t1 int64 = time.Now().UnixNano() 		//Startzeit
	var anz,anzahl int                  		// zur Bestimmung der Frames pro Sekunde
	var verzögerung = 90

	// defer wg.Done()

	for { //Endlos ...
		Mutex.Lock()
		UpdateAus () 										// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
		
		Stiftfarbe(255,255,255)
		Cls()												// Cleart vollständigen Screen
		
		if Akt {
			ZeichneRaum()
			Archivieren()
			Akt = false
		} else {
			Restaurieren(0,0,1200,700)						// Restauriert das alte Hintergrundbild
		}
		

		SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 35 )
		Stiftfarbe(100,10,155)
		Schreibe (2,2,"FPS:"+fmt.Sprint (anzahl))							// Schreibe links oben FPS
		
		
		if time.Now().UnixNano() - t1 < 1000000000 { 		//noch in der Sekunde ...
			anz++
		} else {
			t1 = time.Now().UnixNano() 						// neue Sekunde
			anzahl = anz
			anz=0
			if anzahl < 100 { verzögerung--}				//Selbstregulierung der 
			if anzahl > 100 { verzögerung++}				//Frame-Rate :-)		-- dieser 8-zeilige Abschnitt wurde  von Herrn Schmidt übernommen
		}
		
		UpdateAn () 										// Nun wird der gezeichnete Frame sichtbar gemacht!
		Mutex.Unlock()
	
		time.Sleep(time.Duration(verzögerung * 1e6)) 		// Immer ca. 100 FPS !!
	}
}

func ZeichneRaum() {
	SetzeFont ("../Schriftarten/Ubuntu-B.ttf", 80 )
	Stiftfarbe(100,100,100)
	switch Raumnummer {
		case 0:	
		SchreibeFont(300,50,"LWB - Datenbank")
		ZeichneKnoepfe(Knoepfe)
		case 1:
		SchreibeFont(300,50,Knopftexte[1])
		anfrage := "SELECT * FROM dozent_innen NATURAL JOIN npcs;"
		zeichneAnfrage(conn,anfrage)
		BuZurueck.ZeichneButton()
		
		
		case 2:
		SchreibeFont(300,50,Knopftexte[2])
		anfrage := "SELECT * FROM minigames;"
		zeichneAnfrage(conn,anfrage)
		BuZurueck.ZeichneButton()
		case 3:
		SchreibeFont(300,50,Knopftexte[3])
		anfrage := "SELECT * FROM themengebiete;"
		zeichneAnfrage(conn,anfrage)
		BuZurueck.ZeichneButton()
		case 4:
		anfrage := "SELECT * FROM veranstaltungen;"
		zeichneAnfrage(conn,anfrage)
		SchreibeFont(300,50,Knopftexte[4])
		BuZurueck.ZeichneButton()
	}
}

// Es folgt die Maus-Komponente 1 --- Kein Bestandteil der Welt, also unabhängig -----
func maussteuerung () {

	for {
		_, status, mausX, mausY := MausLesen1()
		// maus.SetzeKoordinaten(mausX,mausY)								// Aktualisiert Maus-Koordinaten
		
		if status==1 { 													// Maustaste gedrückt
			for _,knopf := range Knoepfe { 									// überprüft Knöpfe im Array
				if knopf.TesteXYPosInButton(mausX,mausY) {
					fmt.Println("Button gedrückt: ", knopf.GibBeschriftung() )
					switch knopf.GibBeschriftung() {
						case Knopftexte[0]:	Raumnummer = 0
						case Knopftexte[1]: Raumnummer = 1; fmt.Println(Raumnummer)
						case Knopftexte[2]:	Raumnummer = 2
						case Knopftexte[3]: Raumnummer = 3
						case Knopftexte[4]:	Raumnummer = 4						
					}
				} else if BuZurueck.TesteXYPosInButton(mausX,mausY) {
					Raumnummer = 0
				}
				Akt = true		// etwas wurde geklickt und muss neu gezeichnet werden
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


