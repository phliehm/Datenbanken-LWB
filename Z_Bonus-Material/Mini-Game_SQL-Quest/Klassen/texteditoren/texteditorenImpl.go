//Annalena Cyriacus
//Juni 2023
//Klasse Texteditoren
//basierend auf dem editor08-Quelltext von Stefan Schmidt (LWB Informatik, ALP 3)

package texteditoren

import (
	"gfx"
	"./editor"
	"./editor1line"
)

type data struct {
	x,y,breite,höhe uint16 	// Koordinaten und breite und Höhe
	schriftgr int
	text string
	einzeilig bool
}

// Vor.: Soll der Texteditor vollständig im gfx-Fenster sichtbar sein,
//       muss ein entsprechendes Fenster geöffnet sein oder die Maße
//       in ein Fenster der Größe 1200x700 hineinpassen. 
// Erg.: Ist kein gfx-Fenster offen, ist ein neues geöffnet (1200x700)
//       ein Texteditor mit der übergebenen Breite und Höhe ist geliefert,
//       dessen obere linke Ecke an den Koordinaten (posx,posy) liegt.
//       Die Schriftart des Editors ist TerminusTTF-Bold-4.49.2 und die
//       Schriftgröße entspricht dem übergebenen Wert. Wurde true übergeben,
//       ist der Texteditor einzeilig, wurde false übergeben, sind
//       Zeilenumbrüche mithilfe der Enter-Taste möglich.
func New(posx,posy,breite,höhe uint16, schriftgr int, einzeilig bool) *data {
	var ted *data
	ted = new(data)
	ted.x = posx
	ted.y = posy
	ted.höhe = höhe
	ted.breite = breite
	ted.einzeilig = einzeilig
	if !gfx.FensterOffen () { gfx.Fenster(1200,700) }
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(posx,posy,breite,höhe)
	if einzeilig {
		ted.text = editor1line.Editor(posx,posy,breite,höhe,schriftgr)
	} else {
		ted.text = editor.Editor(posx,posy,breite,höhe,schriftgr)
	}
	return ted
}


//-------------------------- Methoden ---------------------------- 

// Vor.: -
// Erg.: Die Koordinaten der oberen linken Ecke des Texteditors
//       sind geliefert.
func (ted *data) GibPosition() (x,y uint16) {
	return ted.x, ted.y
}

// Vor.: -
// Eff.: Die Koordinaten der oberen linken Ecke des Texteditors
//       sind entsprechend der übergebenen Werte neu gesetzt.
func (ted *data) SetzePosition(xneu,yneu uint16) {
	ted.x = xneu
	ted.y = yneu
}

// Vor.: -
// Erg.: Die Höhe des Texteditors ist geliefert.
func (ted *data) GibHöhe() (höhe uint16) {
	return ted.höhe
}

// Vor.: -
// Eff.: Die Höhe des Texteditors ist entsprechend des übergebenen
//       Werts neu gesetzt.
func (ted *data) SetzeHöhe(höheneu uint16) {
	ted.höhe = höheneu
}


// Vor.: -
// Erg.: Die Breite des Texteditors ist geliefert.
func (ted *data) GibBreite() (breite uint16) {
	return ted.breite
}

// Vor.: -
// Eff.: Die Breite des Texteditors ist entsprechend des übergebenen
//       Werts neu gesetzt.
func (ted *data) SetzeBreite(breiteneu uint16) {
	ted.breite = breiteneu
}

// Vor.: -
// Erg.: Die Schriftgröße des Texteditors ist geliefert.
func (ted *data) GibSchriftgröße() (schriftgr int) {
	return ted.schriftgr
}

// Vor.: -
// Eff.: Die Schriftgröße des Texteditors ist entsprechend des
//       übergebenen Werts neu gesetzt.
func (ted *data) SetzeSchriftgröße(schriftgrneu int) {
	ted.schriftgr = schriftgrneu
}

// Vor.: -
// Erg.: Ist der Texteditor einzeilig, ist true geliefert,
//       andernfalls false.
func (ted *data) IstEinzeilig() bool {
	return ted.einzeilig
}

// Vor.: -
// Erg.: Der in den Texteditor eingegebene Text ist als string geliefert.
//       Wurde nichts eingegeben, ist ein leerer String geliefert.
func (ted *data) GibString() string {
	return ted.text
}
