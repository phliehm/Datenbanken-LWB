package editor1line

// Autor: St. Schmidt, adaptiert von Annalena Cyriacus im Juni 2023
// Datum: 13.09.2022
// Zweck: "Anwendung" des ADT Folge (polymorph) - Der Mehrzeilen-Editor!!
// ---> Mit Verbesserungen: 1.) Es gehen keine Tastendrücke mehr verloren!
// --->                     2.) Kein Flackern des Editorfensters mehr!
//-----------------------------------------------------------------------------------------------

import ("fmt";"gfx"; . "../folgen")

//const spaltenanzahl, zeilenanzahl = 64, 21
var spaltenanzahl, zeilenanzahl uint
	
var f Folge = New (rune(0)) //neue Folge mit Elementen vom Typ rune
var zeilen Folge = New (f)  //neue Folge mit Elementen vom Typ 'Folge von rune-Elementen'
	
var gedrueckt uint8
var taste,tiefe uint16
	
var startspalte,startzeile, index uint = 0,0,0

func String (f Folge, startposition uint) (erg string) {
	var index uint = f.AktuellerIndex ()
	f.Positionieren (startposition)
	for i:=uint(0); i < f.Laenge (); i++ { 
		e,_:=f.AktuellesElement()
		erg = erg + string(e.(rune)) //Folgenelemente sind vom Typ Element !!
		f.Vor ()
	}
	f.Positionieren (index)
	return
}

func updateFenster (xpos,ypos,breite,höhe uint16, schriftgr int) {
		
	zeilenanzahl = uint((höhe)/(uint16(schriftgr+8)))+1
	spaltenanzahl = uint(((2*breite)/(uint16(schriftgr)))-2)
	fmt.Println(spaltenanzahl, zeilenanzahl)
	//spaltenanzahl = 81
	gfx.UpdateAus()
	
	//schwarze Schrift auf weißem Hintergrund
	//gfx.Stiftfarbe(255,255,255)
	//gfx.Vollrechteck (xpos+1,ypos+1,breite-2,höhe-2)
	//gfx.Stiftfarbe(0,0,0)
	//weiße Schrift auf schwarzem Hintergrund
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck (xpos+1,ypos+1,breite-2,höhe-2)
	gfx.Stiftfarbe(255,255,255)
	
	z,_:= zeilen.AktuellesElement ()
	f = z.(Folge)
	for f.AktuellerIndex () >= startspalte + spaltenanzahl {startspalte++}
	for f.AktuellerIndex () < startspalte {startspalte--}
	index = zeilen.AktuellerIndex () //merke aktuelle Zeile
	for index > startzeile + zeilenanzahl - 2 {startzeile++}
	for index < startzeile {startzeile--}
	//------------------------
	//Versuch:
	//if zeilen.Laenge() > uint(höhe) {
	//	zeilen.Positionieren(0)
	//	zeilen.Loeschen()
	//}
	//------------------------
	for i:= startzeile; i < startzeile + zeilenanzahl - 1; i++ {
	//for i:= startzeile; i < uint(höhe) - 1; i++ {
		zeilen.Positionieren (i)
		if a,ok:= zeilen.AktuellesElement (); ok {
			if a.(Folge).Laenge () > startspalte {
				gfx.SchreibeFont (xpos+8,ypos+8 + uint16((i-startzeile)*uint(schriftgr)),String(a.(Folge),startspalte) + "  ")
			}
		}
	}
	/*
	//Design für schwarze Schrift auf weißem Hintergrund
	gfx.Stiftfarbe(255,255,255)
	//transparentes weißes Rechteck für "alte Eingaben"
	gfx.Transparenz(100)
	gfx.Vollrechteck(702,398,296,276)
	//weißes Rechteck unter Editor
	gfx.Transparenz(0)
	gfx.Vollrechteck(700,675,300,25)
	gfx.Vollrechteck(702,662,296,12)
	*/
	
	//Design für weiße Schrift auf schwarzem Hintergrund
	gfx.Stiftfarbe(0,0,0)
	//transparentes weißes Rechteck für "alte Eingaben"
	gfx.Transparenz(150)
	//gfx.Vollrechteck(702,398,296,276)
	gfx.Vollrechteck(xpos+2,ypos+uint16(schriftgr)+8,breite-4,höhe-uint16(schriftgr)-9)
	//schwarzes und weißes Rechteck am Ende bzw. unter Editor
	gfx.Transparenz(0)
	//gfx.Vollrechteck(702,662,296,12)
	//gfx.Vollrechteck(xpos+2,ypos+höhe/25*uint16(schriftgr)+10,breite-4,höhe-höhe/25*uint16(schriftgr)-10)
	gfx.Vollrechteck(xpos+2,ypos+höhe/uint16(schriftgr+5)*uint16(schriftgr)+10,breite-4,höhe-höhe/uint16(schriftgr+5)*uint16(schriftgr)-10)
	gfx.Stiftfarbe(255,255,255)
	gfx.Vollrechteck(xpos,ypos+höhe,breite,700-ypos+höhe)
	
	zeilen.Positionieren (index) //aktiviere wieder aktuelle Zeile		
	//gfx.SchreibeFont (0,572,"Sp:       ") 
	//gfx.SchreibeFont (70,572,fmt.Sprint(f.AktuellerIndex ()+1))
	//gfx.SchreibeFont (120,572,"Zeile:        ")
	//gfx.SchreibeFont  (230,572,fmt.Sprint(zeilen.AktuellerIndex()+1))
	gfx.Stiftfarbe(255,255,255)
	gfx.SchreibeFont (xpos+8 + uint16((f.AktuellerIndex()-startspalte)*(uint(schriftgr)/2)),ypos+8 + (uint16(schriftgr)*uint16(index-startzeile)),"_") //Kursor
	
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(xpos+breite-12,ypos,17,höhe)
	
	gfx.UpdateAn()
}
	
func Editor(xpos, ypos, breite, höhe uint16, schriftgr int) (eingabe string) {
		
	zeilen.EinfuegenVor (New (rune(0)))  //Die erste Zeile ist immer da.
	zeilen.Positionieren (0) //Die erste Zeile ist am Anfang aktuell.
	if !gfx.FensterOffen () { gfx.Fenster(1200,700) }
	//gfx.Fenster (900,600)
	//gfx.Fenstertitel ("Unser erster eigener Editor!")
	gfx.SetzeFont ("./Mini-Game_SQL-Quest/Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",schriftgr)
	//gfx.SetzeFont ("./terminus-font/TerminusTTF-Bold-4.49.2.ttf",schriftgr)
	gfx.TastaturpufferAn()
	updateFenster (xpos,ypos,breite,höhe,schriftgr) //Damit man der Cursor in der oberen Ecke sieht.
A:	for {
		taste, gedrueckt, tiefe = gfx.TastaturpufferLesen1()
		//fmt.Println (taste,gedrueckt,tiefe, gfx.Tastaturzeichen(taste, tiefe), string (gfx.Tastaturzeichen(taste,tiefe)))
		if gedrueckt == 1  { // Beim Drücken der Taste, nicht beim Loslassen!
			switch {
				case taste == 27:  // ESC-Taste
				break A
				case taste == 276:  // Pfeil nach links
				f.Zurueck ()		
				case taste == 275:  // Pfeil nach rechts
				f.Vor ()
				case taste == 278:  // POS1-Taste
				f.Positionieren (0)
				case taste == 279: // Ende - Taste
				f.Positionieren (f.Laenge ())
				case taste == 127: // Entf-Taste
				f.Loeschen ()
				case taste ==  8:  //Backspace-Taste
				if f.AktuellerIndex () > 0 {
					f.Zurueck ()
					f.Loeschen ()
				}
				case taste == 13:  //ENTER-Taste
				//zeilen.EinfuegenVor (New (rune(0)))
				//zeilen.Vor ()
				//zeilen.Zurueck ()
				return
				case taste == 273:  // Pfeil nach oben
				zeilen.Zurueck ()
				case taste == 274:  // Pfeil nach unten
				if zeilen.AktuellerIndex() < zeilen.Laenge () - 1 {
					zeilen.Vor ()
				}
				case taste >= 32 && taste <= 270:  
				z:= gfx.Tastaturzeichen(taste, tiefe)  
				if taste == 93 {
					z = gfx.Tastaturzeichen(43,1)
				} else if taste == 92 {
					z = gfx.Tastaturzeichen(35,1)
				} else if taste == 47 {
					if tiefe == 1 {
						z = gfx.Tastaturzeichen(45,1)
					} else {
						z = gfx.Tastaturzeichen(45,0)
					}
				} else {
					z = gfx.Tastaturzeichen(taste, tiefe)
				}
				if z != 0 { f.EinfuegenVor (z) }
			}
			updateFenster(xpos,ypos,breite,höhe,schriftgr)
			eingabe = String(f,0)
		}
	}
	return
}
