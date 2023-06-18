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
		"unicode/utf8"
		"math"
		"../sqlTabelle"
		"SQL"
		)


// Vor.: Ein gfx-Fenster ist geöffne. Ein Font wurde mit gfx.SetzeFont() gesetzt
type data struct {
	stringTabelle [][]string		// hier könnte man auch sqlTabelle als Datentyp nehmen
	tBTabelle [][]textboxen.Textbox
	kopf 	     header
	zeilenAbstand uint16
	spaltenBreite uint16
	spaltenAbstand uint16
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
		t.SetzeFarbe(h.r,h.g,h.b)
		t.SetzeFont(h.font)
		t.SetzeSchriftgröße(h.schriftgröße)
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

func (h *header) schreibeTbHeader(x,y uint16,b,höhe uint16,tT *data) {
	//fmt.Println("Header: ",h.kopf)
	breiten := gibVariableBreiten(tT.stringTabelle,h.kopf)
	if len(breiten) != len(h.kopf) {return}	// Wenn es keine Antwort gibt
	//fmt.Println(len(breiten),len(h.kopf))
	verschiebung := float64(x)
	for i,zelle := range h.kopf {
		tb := textboxen.New(uint16(math.Round(verschiebung)),y,b,höhe)
		tb.SchreibeText(zelle)
		h.tbKopf = append(h.tbKopf,tb)
		verschiebung += float64(uint16(tT.schriftgröße)*breiten[i])/2 + float64(tT.spaltenAbstand)
	}
}



func New(tabelle[][]string,kopf []string,x,y uint16) *data {
	tT := new(data)
	tT.zeilenAbstand = 5
	tT.spaltenBreite = 300
	tT.font = gfx.GibFont()
	tT.kopf.kopf = kopf
	tT.kopf.font = tT.font
	
	
	tT.x,tT.y = x,y
	tT.spaltenAbstand = 50
	tT.schriftgröße = 20
	tT.kopf.schriftgröße = tT.schriftgröße
	
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
func (tT *data) SetzeSpaltenAbstand(a uint16){
	tT.spaltenAbstand = a
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
		y := tT.y + uint16(2*tT.schriftgröße)+(tT.zeilenAbstand+uint16(tT.schriftgröße))*uint16(i)
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
	tT.kopf.schriftgröße = tT.schriftgröße
	tT.kopf.schreibeTbHeader(tT.x,tT.y,tT.spaltenBreite,uint16(tT.schriftgröße),tT)
	tT.kopf.formatiere(tT.spaltenBreite,uint16(tT.schriftgröße))
	tT.kopf.zeichne()
	// Tabelle zeichnen
	tT.formatiere()
	tT.VariableBreite()
	for z,zeile:= range tT.tBTabelle {

		for s,zelle := range zeile {
			//fmt.Println(tT.stringTabelle[z][s])
			zelle.SchreibeText(tT.stringTabelle[z][s])
			//r,g,b := zelle.GibSchriftfarbe()
			//fmt.Println(zelle.GibX(),zelle.GibY())
			zelle.Zeichne()
			//time.Sleep(0.5e9)
		}
	}
}

// Verändert die Tabelle so, dass die Spaltenbreiten variable sind
func (tT *data) VariableBreite() {
	breiten := gibVariableBreiten(tT.stringTabelle,tT.kopf.kopf)
	// Ändere Breite
	for i,zeile := range tT.tBTabelle {
		y:= tT.y + uint16(2*tT.schriftgröße)+(tT.zeilenAbstand+uint16(tT.schriftgröße))*uint16(i)
		x := float64(tT.x)
		for j,zelle := range zeile {
			//fmt.Println("Spaltenbreite: ",breiten[j])
			zelle.SetzeBreite(uint16(tT.schriftgröße)*breiten[j])
			zelle.SetzePosition(uint16(math.Round(x)),y)
			// Verschiebe um die Spaltenbreiten links von der aktuellen Position
			// plus einem konstanten Abstand
			x+=math.Round(float64(uint16(tT.schriftgröße)*breiten[j])/2) + float64(tT.spaltenAbstand)	
		}
	} 
}

//////////////////////
// HILFS-FUNKTIONEN //
//////////////////////

// transponiert einen 2d-Slice aus Strings
func transponiere(tabelle [][]string) [][]string{
	var spalten int
	spalten = len(tabelle[0])
	
	transponiert := make([][]string,spalten) 
	
	for i:=0;i<len(tabelle);i++ {
		for j:=0;j<len(tabelle[0]);j++ {
			transponiert[j] = append(transponiert[j], tabelle[i][j]) 
		}
	}
	return transponiert
}


func findeLängstenString(s []string, h string) uint16 {
	var maxL uint16
	maxL = uint16(utf8.RuneCountInString(h))
	for _,w := range s {
		l := uint16(utf8.RuneCountInString(w))
		if l >maxL {maxL =l }
	}
	
	return maxL
	 
}

// Gibt die maximalen Breiten aller Spalten
func gibVariableBreiten(tabelle [][]string,h []string) []uint16 {
	var breiten []uint16
	// transponiere Slice damit man leichter an die Spalten kommt
	transponiert := transponiere(tabelle)
	// finde längste zelle in jeder Spalte
	for i,spalte := range transponiert {
		breiten = append(breiten,findeLängstenString(spalte,h[i]))
	}
	return breiten
}


// 

func ZeichneAnfrage(conn SQL.Verbindung,anfrage string,x,y uint16,zeigeAnfrage bool, 
					rT,gT,bT,rK,gK,bK uint8, schriftgröße int, font string) {
	//Stiftfarbe(255,255,255)
	//Cls()
	sT := sqlTabelle.New(conn,anfrage)
	//fmt.Println(sT.GibTabelle())
	
	// SQL Anfrage anzeigen
	if zeigeAnfrage == true {
		// Nur zum Testen auch SQL Anfrage anzeigen
		gfx.Stiftfarbe(0,0,0)
		tbAnfrage := textboxen.New(10,670,1100,100)
		tbAnfrage.SetzeFont("../Schriftarten/terminus-font/Terminus-Bold.ttf")
		tbAnfrage.SetzeSchriftgröße(12)
		tbAnfrage.SchreibeText(anfrage)
		tbAnfrage.Zeichne()
	}
	
	// Textbox Tabelle
	tbT := New(sT.GibTabelle(),sT.GibKopf(),x,y)
	tbT.SetzeFarbeTabelle(rT,gT,bT)
	tbT.SetzeZeilenAbstand(1)
	tbT.SetzeSchriftgrößeTabelle(schriftgröße)
	tbT.SetzeSpaltenAbstand(20)
	tbT.SetzeFarbeKopf(rK,gK,bK)
	tbT.SetzeFontKopf(font)
	tbT.SetzeFontTabelle(font)
	tbT.Zeichne()
}
