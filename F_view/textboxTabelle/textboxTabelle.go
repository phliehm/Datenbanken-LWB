// Author: Philipp Liehm
// Datum: Juni 2023
// Zweck: Aus einer Tabelle [][]string wird ein 2d-slice [][]textboxen.Texbox generiert
//		  Dieser kann zur Darstellung in gfx.Verwendet werden

package textboxTabelle

import (
		"gfx"
		"../textboxen"
		//"fmt"
		//"time"
		)


// Vor.: Ein gfx-Fenster ist geöffne. Ein Font wurde mit gfx.SetzeFont() gesetzt
type data struct {
	stringTabelle [][]string		// hier könnte man auch sqlTabelle als Datentyp nehmen
	tBTabelle [][]textboxen.Textbox
	kopf 	     header
	zeilenAbstand uint16
	spaltenBreite uint16
	font		string
	schriftgröße 	int
	r,g,b	uint8
	x,y uint16		// Position im gfx-Fenster
	
}

// Tabellenkopf
type header struct {
	tbKopf	[]textboxen.Textbox
	kopf	[]string
	r,g,b	uint8
	font	string
	schriftgröße int
}

func (h *header) formatiere(breite,höhe uint16) {
	for _,t := range h.tbKopf {
		t.SetzeBreite(breite)
		t.SetzeHöhe(höhe)
		t.SetzeSchriftgröße(h.schriftgröße)
		t.SetzeFarbe(h.r,h.g,h.b)
		t.SetzeFont(h.font)
	}
}

func (h *header) setzeFont(f string) {
	h.font = f
}

func (h *header) zeichne() {
	for _,t := range h.tbKopf {
		//r,g,b := t.GibSchriftfarbe()
		//fmt.Println(t.GibX(),t.GibY(),t.GibFont(),r,g,b,t.GibSchriftgröße(),t.GibBreite(),t.GibHöhe(),t.GibText())
		t.Zeichne()
		//time.Sleep(1e9)
	}
}

func (h *header) schreibeTbHeader(x,y uint16,b,höhe uint16) {
	//fmt.Println("Header: ",h.kopf)
	for i,zelle := range h.kopf {
		tb := textboxen.New(x+uint16(i)*b,y,b,höhe)
		tb.SchreibeText(zelle)
		h.tbKopf = append(h.tbKopf,tb)
	}
}

func (h *header) SetzeSchriftgröße(g int) {
	h.schriftgröße = g
}

func New(tabelle[][]string,kopf []string,x,y uint16) *data {
	tT := new(data)
	tT.zeilenAbstand = 5
	tT.spaltenBreite = 300
	tT.font = gfx.GibFont()
	tT.kopf.kopf = kopf
	tT.kopf.font = tT.font
	tT.kopf.schriftgröße = 10
	
	tT.x,tT.y = x,y
	tT.schriftgröße = 20
	
	tT.stringTabelle = tabelle 
	
	tT.schreibeTbTabelle()
	
	return tT
}


func (tT *data) SetzeSchriftgrößeTabelle(g int) {
	tT.schriftgröße = g
}

func (tT *data) SetzeFontKopf(f string) {
	tT.kopf.setzeFont(f)
}

func (tT *data) SetzeFontTabelle(f string) {
	tT.font = f
}

func (tT *data) SetzeSchriftgrößeKopf(g int) {
	tT.kopf.schriftgröße = g
}

func (tT *data) SetzeFont(font string) {
	tT.font = font
}

func (tT *data) SetzeFarbeTabelle(r,g,b uint8) {
	tT.r,tT.g,tT.b = r,g,b
}

func (tT *data) SetzeFarbeKopf(r,g,b uint8){
	tT.kopf.r,tT.kopf.g,tT.kopf.b = r,g,b
}

func (tT *data) SetzeZeilenAbstand(a uint16){
	tT.zeilenAbstand = a
}

func (tT *data) SetzeSpaltenBreite(b uint16) {
	tT.spaltenBreite = b
}

// Füllt Tabelle mit leeren Textboxen
func (tT *data) schreibeTbTabelle() {
	// erstelle für jede Zelle eine Textbox
	// Speichere den Zeiger auf die Textbox in ein feld
	var textbTabelle [][]textboxen.Textbox
	for i,zeile := range tT.stringTabelle {
		temp := make([]textboxen.Textbox,0) 	// Slice aus Textboxen
		// y-Position: y der Tabelle insgesamt + Kopf + letzte Zeile
		y := tT.y + uint16(2*tT.kopf.schriftgröße)+(tT.zeilenAbstand+uint16(tT.schriftgröße))*uint16(i)
		for j,_ := range zeile {
			x:=tT.x + tT.spaltenBreite*uint16(j)	
			t:=textboxen.New(x,y,tT.spaltenBreite,uint16(2*tT.schriftgröße))	// Höhe ist 2*Schriftgröße, muss das größer sein? +10?
			temp = append(temp,t)
		}
		textbTabelle = append(textbTabelle,temp)
	}
	tT.tBTabelle =  textbTabelle
}


func (tT *data) formatiere() {
	tT.schreibeTbTabelle()			// Damit die Textboxen an die richtige Stelle gesetzt werden
	for _,zeile:= range tT.tBTabelle {
		for _,zelle := range zeile {
			zelle.SetzeFarbe(tT.r,tT.g,tT.b)
			zelle.SetzeFont(tT.font)
			zelle.SetzeSchriftgröße(tT.schriftgröße)
		}
	}
}

func (tT *data) Zeichne() {
	// Kopf zeichnen
	tT.kopf.schreibeTbHeader(tT.x,tT.y,tT.spaltenBreite,uint16(tT.schriftgröße))
	tT.kopf.formatiere(tT.spaltenBreite,uint16(tT.schriftgröße))
	tT.kopf.zeichne()
	// Tabelle zeichnen
	tT.formatiere()
	for z,zeile:= range tT.tBTabelle {

		for s,zelle := range zeile {
			//fmt.Println(tT.stringTabelle[z][s])
			zelle.SchreibeText(tT.stringTabelle[z][s])
			//r,g,b := zelle.GibSchriftfarbe()
			//fmt.Println(zelle.GibX(),zelle.GibY(),zelle.GibFont(),r,g,b,zelle.GibSchriftgröße())
			zelle.Zeichne()
			//time.Sleep(0.5e9)
		}
	}
}


//////////////////////
// HILFS-FUNKTIONEN //
//////////////////////



