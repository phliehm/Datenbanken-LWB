package textboxen

/* Autor: P. Liehm
 * Datum: 28.03.2023
 * 
 * Dieses Paket stellt einen ADT Textbox gemäß der unten angegebenen 
 * Spezifikation zur Verfügung. Eine Textbox hat:
 * Koordinate x,y, Breite, Höhe, Schriftfarbe, Font, Zeilenabstand, Schriftgröße, Inhalt
 * Der Text kann linksbündig oder zentriert ausgegeben werden. Zentriert sieht nicht immer gut aus, das liegt an der
 * unterschiedlichen Zeichenbreite, die wäre vermutlich nur unter erheblichen Performanceeinbußen richtig zu ermitteln.
 * Die Höhe hat derzeit keine Auswirkungen
 * Es können Zeilenumbrüche im eingegeben Text mit "\n" erzwungen werden. 
 * Beispiel: "1. Zeile\n2. Zeile\n3. Zeile\n\nHier war eine Leerzeile."
 *  -------------------------------------------------------------

	
/* Vor.: --
 *  
 *  Erg.:Eine neue Textbox mit den Koordinaten x,y und der Breite und Höhe 
 * 		ist geliefert. Default Werte sind: 
 * 		Schriftfarbe - schwarz
 * 		ZeilenAbstand - 0
 * 		Schriftgröße - 20
 * 		Font - "../Schriftarten/Ubuntu-B.ttf"
 * 		Inhalt - Leer ("")
 *  New (x,y, b, h uint16) *data // *data erfüllt das Interface Textbox
*/

type Textbox interface {
	
// Vor.: --
// Eff.: Die Position der Textbox ist gesetzt
	SetzePosition(x,y uint16) () 
// Vor.: --
// Erg.: Position der Textbox ist geliefert
	GibX()uint16
// Vor.: --
// Erg.: Position der Textbox ist geliefert
	GibY()uint16
// Vor.: --
// Eff.: Die Breite der Textbox ist gesetzt
	SetzeBreite(b uint16)
// Vor.: --
// Eff.: Die Höhe der Textbox ist gesetzt
	SetzeHöhe(h uint16)
// Vor.: --
// Erg.: Breite der Textbox ist geliefert
	GibBreite() uint16
// Vor.: --
// Erg.: Höhe der Textbox ist geliefert
	GibHöhe() uint16
// Vor.: --
// Eff.: Der Text wird der Textbox hinzugefügt 
	SchreibeText(text string) ()
	
// Vor.: --
// Eff.: Der Zeilenabstand wurde geändert
	SetzeZeilenAbstand(za uint16) ()
	
// Vor.: --
// Erg.: Der Zeilenabstand ist geliefert
	GibZeilenAbstand() (font uint16)

// Vor.: Es existiert eine Font-Datei im ttf-Format im angegebenen Ordner
// Eff.: Der Font des Textes in der Textbox ist gesetzt. 
	SetzeFont(string)

// Vor.: --
// Erg.: Der Font der Textbox ist geliefert
	GibFont()string

// Vor.: --
// Erg.: Die Schriftgröße der Textbox ist geändert
	SetzeSchriftgröße (int)

// Vor.: --
// Erg.: Die Schriftgröße der Textbox ist geliefert
	GibSchriftgröße () int


// Vor.: --
// Erg.: Die Farbe des Textes und des Rahmens (gleich) ist gesetzt
	SetzeFarbe(r,g,b uint8) 

// Vor.: --
// Erg.: Text wird beim Aufruf von Zeichne() linksbündig ausgegeben (default)
	SetzeLinksbündig()

// Vor.: --
// Erg.: Text wird beim Aufruf von Zeichne() zentriert ausgeben
	SetzeZentriert()

// Vor.: --
// Erg.: Ist die Eingabe true, wird ein Rahmen in der Rahmenfarbe gezeichnet, sonst nicht
	RahmenAn(r bool)

// Vor.: --
// Erg.: Rahmenfarbe ist gesetzt
	SetzeRahmenFarbe(r,g,b uint8)
	
// Vor.:
// Erg.: Ist die Eingabe true, wird der Hintergrund mit der Hintergrundfarbe gezeichnet, sonst ist der Hintergrund transparent
	HintergrundAn(h bool)
// Vor.: 
// Erg.: Hintergrundfarbe ist gesetzt
	SetzeHintergrundFarbe(r,g,b uint8)

// Vor.: Ein gfx-Fenster ist offen
// Eff.: Zeichnet die Textbox in das gfx-Fenster
	Zeichne()
	
	
}

