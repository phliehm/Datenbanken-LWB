package main
// Autor: A. Cyriacus, M. Seiss, P. Liehm, B. Schneider
// Datum: 15.06.2023
// Zweck: DBP - LWB - Datenbank
//--------------------------------------------------------------------

import ( 	. "gfx"
			"fmt"
			"sync"
			"time"
			"felder"
			"./Klassen/buttons"
		)

var Mutex sync.Mutex					// erstellt Mutex
	
var Knoepfe, Suchknoepfe []buttons.Button		// Slices für alle erstellten Knöpfe / die Suchfelder
var Suchfelder []felder.Feld
var BuZurueck,BuEintrag,BuEnde buttons.Button
var Akt bool = true						// True gdw. Raum gewechselt wurde
var Raumnummer uint8					// Raumnummer des momentanen Raumes
var Knopftexte, Suchknopftexte []string

func main () {
	Fenster (1200, 700)
	Fenstertitel(" ###  LWB - Datenbank  ###")
	SetzeFont ("./Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",20)
	
	Knopftexte = append(Knopftexte, "", "Dozenten", "Minispiele", "Fachgebiete", "Semester")
	Suchknopftexte = append(Knopftexte, "", "vname", "nname", "str", "plz")
	
	BuZurueck 	= buttons.New(10,620,200,70, 230,50,100, true, " zurück")				// zurück
	BuEintrag 	= buttons.New(400,620,350,70, 230,50,100, true, " Erstelle Eintrag")	// Erstelle Eintrag
	BuEnde 	= buttons.New(1000,620,200,70, 230,50,100, true, " Beenden")				// Beenden
	
	ErstelleKnoepfe()
	ErstelleSuchknoepfe()
	ErstelleSuchFelder()
	
	ZeichneRaum()

	
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
	vname		:= buttons.New(100,100,300,50, 230,50,100, true, Suchknopftexte[1])		// Dozenten
	nname		:= buttons.New(300,100,300,50, 230,50,100, true, Suchknopftexte[2])		// Minispiele
	str			:= buttons.New(500,100,300,50, 230,50,100, true, Suchknopftexte[3])		// Fachgebiete
	plz			:= buttons.New(700,100,300,50, 230,50,100, true, Suchknopftexte[4])		// Semester
	Suchknoepfe = append(Suchknoepfe, vname, nname, str, plz)
}
func ErstelleSuchFelder() {
	vname := felder.New (100, 100, 30, 'l', "Vorname")		// Position 10/10; Länge von 30 Zeichen; linksbündig; Name des Feldes
	nname := felder.New (300, 100, 30, 'l', "Nachname")
	str   := felder.New (500, 100, 30, 'l', "Straße")
	plz   := felder.New (700, 100,  5, 'l', "PLZ")
	Suchfelder = append(Suchfelder, vname, nname, str, plz)
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

/*

	// Die Kontroll-Komponente 2 ist die 'Mainloop' im Hauptprogramm	
	// Wir fragen hier nur die Tastatur ab.

	wg.Add(4)						// Wait-Group erhält Counter 4 zum Warten auf das Ende der nebenläufigen Routinen

	SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 28 )

			}
		}
*/


// Es folgt die VIEW-Komponente		--- wird wahrscheinlich so nicht benötigt
func view_komponente () { 
		
	var t1 int64 = time.Now().UnixNano() 		//Startzeit
	var anz,anzahl int                  		// zur Bestimmung der Frames pro Sekunde
	var verzögerung = 90

	// defer wg.Done()

	for { 
		if Akt {
			//Endlos ...
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
		
			time.Sleep(time.Duration(verzögerung * 1e5)) 		// Immer ca. 100 FPS !!
		}
	}
}



func ZeichneRaum() {
	UpdateAus () 										// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
	Stiftfarbe(255,255,255)
	Cls()												// Cleart vollständigen Screen
	SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 80 )
	Stiftfarbe(100,100,100)
	switch Raumnummer {
		case 0:	
		SchreibeFont(300,50,"LWB - Datenbank")
		ZeichneKnoepfe()
		BuEintrag.ZeichneButton()
		BuEnde.ZeichneButton()
		case 1:
		SchreibeFont(300,50,Knopftexte[1])
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
				}
			}
			for _,suchknopf := range Suchknoepfe { 									// überprüft Knöpfe im Array
				if suchknopf.TesteXYPosInButton(mausX,mausY) {
					fmt.Println("Suchknopf gedrückt: ", suchknopf.GibBeschriftung() )
					switch suchknopf.GibBeschriftung() {
						case Suchknopftexte[1]: 
						Suchfelder[0].Edit()
						case Suchknopftexte[2]:	
						Suchfelder[1].Edit()
						case Suchknopftexte[3]: 
						Suchfelder[2].Edit()
						case Suchknopftexte[4]:	
						Suchfelder[3].Edit()					
					}
				}
			}
			if BuZurueck.TesteXYPosInButton(mausX,mausY) {
				AktiviereKnoepfe()
				Raumnummer = 0
			} else if BuEintrag.TesteXYPosInButton(mausX,mausY) {
				Raumnummer = 10
				//ZeichneSuchknoepfe()
			} 
				
			ZeichneRaum()			// etwas wurde geklickt und muss neu gezeichnet werden
		}
	}
}






