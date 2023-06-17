//Annalena Cyriacus
//Juni 2023
//Klasse Texteditoren
//basierend auf dem editor08-Quelltext von Stefan Schmidt (LWB Informatik, ALP 3)

// Der ADT Texteditoren dient der Verwaltung von Vierecken. Jedes Viereck
// hat vier Koordinaten, nämlich (xA,yA), (xB,yB), (xC,yC) und (xD,yD)
// durch die es festgelegt wird. A ist dabei die obere linke Ecke.

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
// New(posx,posy,breite,höhe uint16, schriftgr int, einzeilig bool) *data // *data erfüllt das Interface Texteditor

package texteditoren

type Texteditor interface {
	
	// Vor.: -
	// Erg.: Die Koordinaten der oberen linken Ecke des Texteditors
	//       sind geliefert.
	GibPosition() (x,y uint16) 
	
	// Vor.: -
	// Eff.: Die Koordinaten der oberen linken Ecke des Texteditors
	//       sind entsprechend der übergebenen Werte neu gesetzt.
	SetzePosition(xneu,yneu uint16)
	
	// Vor.: -
	// Erg.: Die Höhe des Texteditors ist geliefert.
	GibHöhe() (höhe uint16)
	
	// Vor.: -
	// Eff.: Die Höhe des Texteditors ist entsprechend des übergebenen
	//       Werts neu gesetzt.
	SetzeHöhe(höheneu uint16)
	
	// Vor.: -
	// Erg.: Die Breite des Texteditors ist geliefert.
	GibBreite() (breite uint16)
	
	// Vor.: -
	// Eff.: Die Breite des Texteditors ist entsprechend des übergebenen
	//       Werts neu gesetzt.
	SetzeBreite(breiteneu uint16)
	
	// Vor.: -
	// Erg.: Die Schriftgröße des Texteditors ist geliefert.
	GibSchriftgröße() (schriftgr int)
	
	// Vor.: -
	// Eff.: Die Schriftgröße des Texteditors ist entsprechend des
	//       übergebenen Werts neu gesetzt.
	SetzeSchriftgröße(schriftgrneu int)
	
	// Vor.: -
	// Erg.: Ist der Texteditor einzeilig, ist true geliefert,
	//       andernfalls false.
	IstEinzeilig() bool
	
	// Vor.: -
	// Erg.: Der in den Texteditor eingegebene Text ist als string geliefert.
	//       Wurde nichts eingegeben, ist ein leerer String geliefert.	
	GibString() string
	
}

