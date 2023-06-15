package main
// Autor: A. Cyriacus, M. Seiss, P. Liehm, B. Schneider
// Datum: 15.06.2023
// Zweck: DBP - LWB - Datenbank
//--------------------------------------------------------------------

import ( 	. "gfx"
			"fmt"
			"sync"
			"time"
			//"felder"
			"./klassen/buttons"
		)

var Mutex sync.Mutex					// erstellt Mutex
	
var Knoepfe []buttons.Button			// Slice für alle erstellten Knöpfe
var BuZurueck buttons.Button
var Akt bool = true						// True gdw. Raum gewechselt wurde
var Raumnummer uint8					// Raumnummer des momentanen Raumes
var Knopftexte []string

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
	
	
	
	
	// Das Hauptprogramm startet die View-Komponente als nebenläufigen Prozess!
	go view_komponente()
	
	// Nebenläufig wird die Kontroll-Komponente für die Maus gestartet.
	go maussteuerung()

	/*
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


/*
	// Das Hauptprogramm startet die View-Komponente als nebenläufigen Prozess!
	go view_komponente(&obj, maus, okayObjekt, &stop, &akt, &ende, &punkte, &diff, &mutex, &eingabe, &wg)

	// Objekte werden nach und nach in der Welt platziert
	go spielablauf(&obj, maus, random, &mutex, &akt, &tastatur, &stop, &signal, &ende, &zweiter, &eingabe, &wert, &punkte, &punkteArr, kanal, &wg)

	// Nebenläufig wird die Kontroll-Komponente für die Maus gestartet.
	go maussteuerung(&obj, maus, okayObjekt, &signal, &stop, &akt, &ende, &punkte, &diff, &wert, kanal, &wg)

	go musikhintergrund(&ende, &wg)

	// Die Kontroll-Komponente 2 ist die 'Mainloop' im Hauptprogramm	
	// Wir fragen hier nur die Tastatur ab.

	wg.Add(4)						// Wait-Group erhält Counter 4 zum Warten auf das Ende der nebenläufigen Routinen

	SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 28 )


	A:	for {
		taste, gedrueckt, tiefe = TastaturLesen1()
		
		if tastatur {
			if gedrueckt == 1  { 						// Beim Drücken der Taste, nicht beim Loslassen!
				switch {
					case taste == 27:  									// ESC-Taste
					break A
					case taste==13 || taste==271:  						// Enter-Taste(n)
					signal = true
					case taste == 32:  									// Leer-Taste
					eingabe += " "
					case taste ==  8:  									// Backspace-Taste
					if eingabe != "" {
						eingabe = eingabe [:len(eingabe)-1]
					}
					case taste ==  276:  								// LINKS-Taste
					if eingabe != "" {
						eingabe = eingabe [:len(eingabe)-1]
					}
					case taste >= 48 && taste < 58 && tiefe == 0:  		// Zahlen
					eingabe += string(taste)
					case taste == 44:
					eingabe += ","
					case taste == 46:
					eingabe += "."
					case taste == 55 && tiefe > 0:  					// 7
					eingabe += "["
					case taste == 56 && tiefe > 0:  					// 8
					eingabe += "("
					case taste == 57 && tiefe > 0:    					// 9		
					eingabe += ")"
					case taste == 48 && tiefe > 0:  		  			// 0
					eingabe += "]"
					case taste == 46 && tiefe > 0:  		
					eingabe += ":"
					case taste == 49 && tiefe > 0:  		
					eingabe += ":"
					case taste == 50 && tiefe > 0:  		
					eingabe += "\""
					case taste == 51 && tiefe > 0:  		
					eingabe += "'"
					case taste == 92 && tiefe > 0:  		
					eingabe += "'"
					case taste >= 97 && taste < 123 && tiefe == 0:  	// Kleinbuchstaben
					eingabe += string(taste)
					case taste >= 97 && taste < 123 && tiefe > 0:		// Großbuchstaben
					eingabe += string(taste-32)
					default:
				}
			}
		}
*/



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
	
		time.Sleep(time.Duration(verzögerung * 1e5)) 		// Immer ca. 100 FPS !!
	}
}

func ZeichneRaum() {
	SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 80 )
	Stiftfarbe(100,100,100)
	switch Raumnummer {
		case 0:	
		SchreibeFont(300,50,"LWB - Datenbank")
		ZeichneKnoepfe(Knoepfe)
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






