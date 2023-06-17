package textboxen

import (
	"gfx"
	"strings"
	//"fmt"
	"unicode/utf8"
)

type data struct {
	x,y,breite,höhe uint16 	// Koordinaten und breite und Höhe
	font string 
	zeilenAbstand uint16
	text string
	schriftgr int
	r,g,b uint8		// Textfarbe
	linksbündig bool
	zentriert bool
	hintergrund bool
	rahmen bool
	rr,rg,rb uint8 // Rahmenfarbe
	hr,hg,hb uint8 // Hintergrundfarbe
	
}

func New(posx,posy,breite,höhe uint16) *data{
	var tb *data
	tb = new(data)
	tb.x = posx
	tb.y = posy
	tb.höhe = höhe
	tb.breite = breite
	tb.schriftgr = 20
	tb.font = gfx.GibFont() 		
	tb.linksbündig = true
	return tb
}
// Vor.: --
// Eff.: Die Position der Textbox ist gesetzt
func (tb *data)	SetzePosition(x,y uint16) () {
	tb.x = x
	tb.y = y
} 

// Vor.: --
// Erg.: X-Position der Textbox ist geliefert
func (tb *data) GibX() uint16 {
	return tb.x
}

// Vor.: --
// Erg.: Y-Position der Textbox ist geliefert
func (tb *data) GibY() uint16 {
	return tb.y
}

// Vor.: --
// Eff.: Die Breite der Textbox ist gesetzt
func (tb *data)	SetzeBreite(b uint16) {
	tb.breite = b
}
	
	
// Vor.: --
// Eff.: Die Höhe der Textbox ist gesetzt
func (tb *data)	SetzeHöhe(h uint16) {
	tb.höhe = h
}

// Vor.: 
// Erg.: Die Höhe der Textbox ist geliefert
func (tb *data)	GibHöhe() uint16 {
	// Höhe berechnen, dafür muss der Text dargestellt werden
	
	// Zeilenumbrüche generieren
	//var textTemp string = tb.text
	var zeilenAnzahl uint16	
	var zeile string	// Zeileninhalt
	var zeichenbreite int
	var zeilenLänge uint16
	zeichenbreite = tb.schriftgr/2
	if zeichenbreite == 0 {zeichenbreite = 1}	// Vermeide eine Zeichenbreite von 0
	
	var zeileUser []string // Durch den Nutzer festgelegte Zeilen
	
	// Teile den Text in die vorgegbenen Zeilen
	for _,w := range tb.text {
		if w!='\n' {
			zeile+=string(w)
		} else {
			zeileUser = append(zeileUser,zeile)
			zeile = ""
		}
	}
	// Letzte Zeile, denn dann kam eventl. kein \n
	zeileUser = append(zeileUser,zeile)
	zeile = ""
	
	//fmt.Println(zeileUser)
	for i:=0;i<len(zeileUser);i++ {
		// Wörter sollen ganz gelassen werden, Zeilenumbrüche also nur bei Leerzeichen
		worte := strings.Fields(zeileUser[i])	// Slice mit durch Leerzeichen getrennte Worte
		
		for _,w:= range worte {
			//fmt.Println("Zeilenlänge: ", zeilenLänge, "  Zeile: ", zeile, "  ZeilenAnzahl: ",zeilenAnzahl)
			// utf8.RuneCountInString(w) gibt die korrekte Anzahl an Zeichen in einem UTF8-Wort, also die Länge des Wortes
			zeilenLänge = uint16(len(zeile)*zeichenbreite + utf8.RuneCountInString(w)*zeichenbreite)
			if zeilenLänge <= tb.breite {		// Prüfe ob Zeilenlänge überschritten wird
				//l += utf8.RuneCountInString(w)	*zeichenbreite						// Füge Länge des Wortes zur Zeilenlänge hinzu
				zeile += w + " "					// Wort zur Zeile, inklusive Leerzeichen danach

			} else {								// Zeile zu lang
				// Schreibe den Text 
				//gfx.SchreibeFont(tb.x,tb.y+(uint16(tb.schriftgr)+tb.zeilenAbstand)*zeilenAnzahl,zeile)
				
				zeilenLänge=uint16(utf8.RuneCountInString(w)*zeichenbreite)		// setze Zeilenlänge neu (sonst wäre sie jetzt zu lang
				zeile = w + " "					// Beginne eine neue Zeile
				zeilenAnzahl+=1							// Erhöhe die Zeilenanzahl
			}
		}
		//gfx.SchreibeFont(tb.x,tb.y+(uint16(tb.schriftgr)+tb.zeilenAbstand)*zeilenAnzahl,zeile)
		zeile = ""
		zeilenAnzahl+=1
		zeilenLänge = 0
	}
	tb.höhe=(uint16(tb.schriftgr)+tb.zeilenAbstand)*zeilenAnzahl
	return tb.höhe
}
// Vor.: --
// Erg.: Breite der Textbox ist geliefert
func (tb *data) GibBreite() uint16 {
	return tb.breite
}

// Vor.: --
// Eff.: Der Text wird der Textbox hinzugefügt 
func (tb *data)	SchreibeText(text string) () {
	tb.text = text
}
	
// Vor.: --
// Eff.: Der Zeilenabstand wurde geändert
func (tb *data)	SetzeZeilenAbstand(za uint16) () {
	tb.zeilenAbstand = za
}
	
// Vor.: --
// Erg.: Der Zeilenabstand ist geliefert
func (tb *data)	GibZeilenAbstand() (uint16) {
	return tb.zeilenAbstand
}

// Vor.: Es existiert eine Font-Datei im ttf-Format im angegebenen Ordner
// Eff.: Der Font des Textes in der Textbox ist gesetzt. 
func (tb *data)	SetzeFont(font string) {
	tb.font = font
}

// Vor.: --
// Erg.: Der Font der Textbox ist geliefert
func (tb *data)	GibFont()string {
	return tb.font
}

// Vor.: --
// Erg.: Die Schriftgröße der Textbox ist geändert
func (tb *data)	SetzeSchriftgröße (gr int) {
	tb.schriftgr = gr
}

// Vor.: --
// Erg.: Die Schriftgröße der Textbox ist geliefert
func (tb *data)	GibSchriftgröße () int {
	return tb.schriftgr
}


// Vor.: --
// Erg.: Die Farbe des Textes und des Rahmens (gleich) ist gesetzt
func (tb *data) SetzeFarbe(r,g,b uint8) {
	tb.r,tb.g,tb.b  = r,g,b
}

// Vor.: --
// Erg.: Text wird beim Aufruf von Zeichne() linksbündig ausgegeben (default)
func (tb *data) SetzeLinksbündig() {
	tb.linksbündig = true
	tb.zentriert = false
}

// Vor.: --
// Erg.: Text wird beim Aufruf von Zeichne() zentriert ausgeben
func (tb *data) SetzeZentriert() {
	tb.linksbündig = false
	tb.zentriert = true
}

// Vor.: --
// Erg.: Ist die Eingabe true, wird ein Rahmen in der Rahmenfarbe gezeichnet, sonst nicht
func (tb *data)	RahmenAn(r bool) {
	tb.rahmen = r
}

// Vor.: --
// Erg.: Rahmenfarbe ist gesetzt
func (tb* data)	SetzeRahmenFarbe(r,g,b uint8) {
	tb.rr,tb.rg,tb.rb = r,g,b
}
	
// Vor.:
// Erg.: Ist die Eingabe true, wird der Hintergrund mit der Hintergrundfarbe gezeichnet, sonst ist der Hintergrund transparent
func (tb *data) HintergrundAn(h bool) {
	tb.hintergrund = h
}

// Vor.: 
// Erg.: Hintergrundfarbe ist gesetzt
func (tb *data)	SetzeHintergrundFarbe(r,g,b uint8) {
	tb.hr,tb.hg,tb.hb = r,g,b
}

// Vor.: Ein gfx-Fenster ist offen
// Eff.: Zeichnet die Textbox in das gfx-Fenster
func (tb *data)	Zeichne() {
	var tempFont string
	tempFont = gfx.GibFont()	// Speichere aktuellen Font um diesen später wieder darzustellen
	gfx.SetzeFont(tb.font,tb.schriftgr)
	// Zeichne Hintergrund
	if tb.hintergrund {
		gfx.Stiftfarbe(tb.hr,tb.hg,tb.hb)
		gfx.Vollrechteck(tb.x-10,tb.y-10,tb.breite+20,tb.höhe+20)
	}
	if tb.rahmen {
		gfx.Stiftfarbe(tb.rr,tb.rg,tb.rb)
		gfx.Rechteck(tb.x-10,tb.y-10,tb.breite+20,tb.höhe+20)
	}
	gfx.Stiftfarbe(tb.r, tb.g, tb.b)
	// Zeilenumbrüche generieren
	//var textTemp string = tb.text
	var zeilenAnzahl uint16	
	var zeile string	// Zeileninhalt
	var zeichenbreite int
	var zeilenLänge, zeilenLänge_temp uint16
	//var linkerAbstand uint16 // Abstand zur linken Boxgrenze zur Darstellung von zentriertem Text
	zeichenbreite = tb.schriftgr/2
	if zeichenbreite == 0 {zeichenbreite = 1}	// Vermeide eine Zeichenbreite von 0
	
	var zeileUser []string // Durch den Nutzer festgelegte Zeilen
	
	// Teile den Text in die vorgegbenen Zeilen
	for _,w := range tb.text {
		if w!='\n' {
			zeile+=string(w)
		} else {
			zeileUser = append(zeileUser,zeile)
			zeile = ""
		}
	}
	// Letzte Zeile, denn dann kam eventl. kein \n
	zeileUser = append(zeileUser,zeile)
	zeile = ""
	
	//fmt.Println(zeileUser)
	for i:=0;i<len(zeileUser);i++ {
		//fmt.Println(i,zeileUser[i])
		// Wörter sollen ganz gelassen werden, Zeilenumbrüche also nur bei Leerzeichen
		worte := strings.Fields(zeileUser[i])	// Slice mit durch Leerzeichen getrennte Worte
		
		for _,w:= range worte {
			//fmt.Println("Zeilenlänge: ", zeilenLänge, "  Zeile: ", zeile, "  ZeilenAnzahl: ",zeilenAnzahl)
			// utf8.RuneCountInString(w) gibt die korrekte Anzahl an Zeichen in einem UTF8-Wort, also die Länge des Wortes
			zeilenLänge = uint16(len(zeile)*zeichenbreite + utf8.RuneCountInString(w)*zeichenbreite)
			if zeilenLänge <= tb.breite {		// Prüfe ob Zeilenlänge überschritten wird
				//l += utf8.RuneCountInString(w)	*zeichenbreite						// Füge Länge des Wortes zur Zeilenlänge hinzu
				zeile += w + " "					// Wort zur Zeile, inklusive Leerzeichen danach
				//fmt.Println(w, len(w),utf8.RuneCountInString(w))
				zeilenLänge_temp = zeilenLänge			// temp wird gebraucht um später den Text zentriert darstellen zu können
				

			} else {								// Zeile zu lang
				// Schreibe den Text entweder linksbündig oder zentriert
				if tb.linksbündig {
					gfx.SchreibeFont(tb.x,tb.y+(uint16(tb.schriftgr)+tb.zeilenAbstand)*zeilenAnzahl,zeile)
				} else {
					gfx.SchreibeFont(tb.x+(tb.breite-zeilenLänge_temp)/2,tb.y+(uint16(tb.schriftgr)+tb.zeilenAbstand)*zeilenAnzahl,zeile)	// Schreibe Zeile an richtige Stelle
				}
				
				
				zeilenLänge=uint16(utf8.RuneCountInString(w)*zeichenbreite)		// setze Zeilenlänge neu (sonst wäre sie jetzt zu lang
				zeilenLänge_temp = zeilenLänge									// Setze auch temp neu, sonst wird die falsche länge bei nur einem Wort pro Zeile genommen
				zeile = w + " "					// Beginne eine neue Zeile
				zeilenAnzahl+=1							// Erhöhe die Zeilenanzahl
				//fmt.Println(w, len(w),utf8.RuneCountInString(w))
			}
		}
		if tb.linksbündig {
			gfx.SchreibeFont(tb.x,tb.y+(uint16(tb.schriftgr)+tb.zeilenAbstand)*zeilenAnzahl,zeile)
		} else {
			gfx.SchreibeFont(tb.x+(tb.breite-zeilenLänge_temp)/2,tb.y+(uint16(tb.schriftgr)+tb.zeilenAbstand)*zeilenAnzahl,zeile)	// Schreibe Zeile an richtige Stelle
		}
		zeile = ""
		zeilenAnzahl+=1
		zeilenLänge = 0
	}
	tb.höhe=(uint16(tb.schriftgr)+tb.zeilenAbstand)*zeilenAnzahl-tb.zeilenAbstand
	//gfx.Rechteck(tb.x,tb.y,tb.breite,tb.höhe)
	gfx.SetzeFont(tempFont,20)		// vorheriger Font wird wieder hergestellt
}


