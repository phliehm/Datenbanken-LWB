// ADT button - Spezifikation und Implementation

// Martin Seiß    29.3.2023

package buttons

import "gfx"
// import "fmt"



type data struct {
	x,y 		uint16		// bezeichet die obere linke Ecke.
	br,h 		uint16		// die Breite und Höhe des Button
	r,g,b 		uint8		// Hintergrundfarbwerte des Buttons
	aktiv		bool		// true = Button aktiv
	beschriftung 	string	// String mit Beschriftung des Buttons
	font			string	// Name mit Pfad von Font
	sound			string  // Name mit Pfad von Sound
}


func New(x,y,bb,hh uint16, r,g,b uint8, aktiv bool, beschriftung string) *data {
	var but *data = new(data)
	but.x = x
	but.y = y
	but.h = hh
	but.br = bb
	but.r = r
	but.g = g
	but.b = b
	but.aktiv = aktiv
	but.beschriftung = beschriftung
	but.font = gfx.GibFont() 			//--> ACHTUNG: muss vorher einmal gesetzt worden sein!!! (Spez.)
	but.sound = ""
	return but
}


func (but *data) SetzePosition(x,y uint16) {
	but.x = x
	but.y = y
}

func (but *data) TesteXYPosInButton(x,y uint16) bool {
	if but.aktiv {
		if but.x <= x &&  x <= but.x+but.br {
			if but.y <= y &&  y <= but.y+but.h{
				if but.sound != "" {gfx.SpieleSound(but.sound)}
				return true
			}
		} 
	}
	return false
}



func (but *data) AktiviereButton() {
	but.aktiv = true
}

	
	
func (but *data) DeaktiviereButton() {
	but.aktiv = false
}
	
	
func (but *data) GibAktivitaetButton() bool {
	return but.aktiv
}

	
func (but *data) AendereBeschriftung(text string) {
	but.beschriftung = text
}	
	
	
	
func (but *data) SetzeFont(font string) {
	but.font = font
}


func (but *data) SetzeSound(sound string) {
	but.sound = sound
}
	
	
func (but *data) ZeichneButton() {
	var diff uint16 = 2
	var fsize int = int(but.h/2)
	
	gfx.SetzeFont(but.font,fsize)

	if but.aktiv {
		gfx.Stiftfarbe(but.r,but.g,but.b)
		gfx.Vollrechteck(but.x,but.y,but.br,but.h)
		gfx.Stiftfarbe(0,0,0)
		gfx.Rechteck(but.x,but.y,but.br,but.h)
		gfx.Rechteck(but.x+diff,but.y+diff,but.br-2*diff,but.h-2*diff)
		gfx.SchreibeFont(but.x+10,but.y+but.h/2-uint16(fsize)/2,but.beschriftung)
	} else {
		gfx.Stiftfarbe(but.r/2,but.g/2,but.b/2)
		gfx.Vollrechteck(but.x,but.y,but.br,but.h)
		gfx.Stiftfarbe(100,100,100)
		gfx.Rechteck(but.x,but.y,but.br,but.h)
		gfx.Rechteck(but.x+diff,but.y+diff,but.br-2*diff,but.h-2*diff)
		gfx.SchreibeFont(but.x+10,but.y+but.h/2-uint16(fsize)/2,but.beschriftung)		
		gfx.Stiftfarbe(0,0,0)
	}
}
