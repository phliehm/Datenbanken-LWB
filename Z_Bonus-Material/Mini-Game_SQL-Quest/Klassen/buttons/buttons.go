// ADT button - Spezifikation

// Martin Seiß    29.3.2023 (Start)

// Buttons haben recheckige Form.
// werden.
// x,y - bezeichet die obere linke Ecke.
// b,h - die Breite und Höhe
// r,g,b - Hintergrundfarbwerte des Buttons
// aktiv - Button ist aktiv und bereit zur Positionsbestimmung
// beschriftung	- String mit Beschriftung des Buttons
// font - Font mit dem der Button beschriftet wird
// sound - wav-Dateianame mit Pfad als String für Klang bei Tastendruck

// --> ACHTUNG: Vor dem ersten Aufrufen muss der Font einmal
//				gesetzt worden sein!
//				z.B: gfx.SetzeFont ("Schriftarten/Ubuntu-B.ttf",20)
// New(x,y,bb,hh,r,g,b,aktiv,beschriftung) 

package buttons

// ------------  importierte Pakete und Klassen  -----------------//
// import "gfx"


type Button interface {

	// Vor: -
	// Eff: Die x-y-Position ist entsprechend der Eingabe neu gesetzt.
	// Erg: -
	SetzePosition(x,y uint16)

	// Vor: Button ist aktiv.
	// Eff: -
	// Erg: True ist geliefert, wenn die x-y-Position mit dem 
	//		Buttonbereich übereinstimmt, andernfalls False.
	TesteXYPosInButton(x,y uint16) bool

	// Vor: -
	// Eff: Button ist auf aktiv gestellt. x,y-Position kann gecheckt
	//		werden und Button ist schwarz-farbig gezeichnet.
	// Erg: -
	AktiviereButton()
	
	// Vor: -
	// Eff: Button ist auf inaktiv gestellt. x,y-Position kann gecheckt
	//		werden, liefert aber immer false, und Button ist 
	//		grau-blassfarbig gezeichnet.
	// Erg: -
	DeaktiviereButton()

	// Vor: -
	// Eff: -
	// Erg: Der Aktivwert des Buttons ist geliefert.
	GibAktivitaetButton() bool
	
	// Vor: -
	// Eff: Die Beschriftung des Button ist geändert.
	// Erg: -
	AendereBeschriftung(text string) 

	// Vor: -
	// Eff: Der angegebene Font wird beim Schreiben verwendet.
	// Erg: -
	SetzeFont(font string)

	// Vor: -
	// Eff: Der angegebene Sound-String wird verwendet, und erklingt,
	//		wenn Button aktiv und getroffen.
	// Erg: -
	SetzeSound(sound string)

	// Vor: gfx Fenster ist geöffnet und Button kann an gegebener Stelle
	//		gezeichnet werden.
	// Eff: Der Button ist gezeichnet.
	// Erg: -
	ZeichneButton()

}
