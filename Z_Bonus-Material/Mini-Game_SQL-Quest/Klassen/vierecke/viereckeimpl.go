package vierecke

// Annalena Cyriacus
// 24.04.2023

import ( "gfx" ; "fmt" ; "unsafe")

// Datenstruktur eines Vierecks
// ----------------------------

type data struct {
	xA, yA uint16
	xB, yB uint16 
	xC, yC uint16
	xD, yD uint16
	aktiv bool
	r, g, b uint8 //optional
}

// Methoden eines Vierecks
// -----------------------

// Vor.: -
// Erg.: Ein Viereck mit den Koordinaten (xA,yA), (xB,yB), (xC,yC) und
//       (xD,yD) ist geliefert. A ist dabei die obere linke Ecke.
func New (xA,yA,xB,yB,xC,yC,xD,yD uint16) *data { // *data erfüllt das Interface Viereck
	var v *data
	v = new(data)
	v.xA = xA
	v.yA = yA
	v.xB = xB
	v.yB = yB
	v.xC = xC
	v.yC = yC
	v.xD = xD
	v.yD = yD
	return v
}

// Vor.: -
// Erg.: Die Koordinaten aller vier Eckpunkte des Vierecks sind
//       geliefert.
func (v *data) GetKoordinaten () (xA,yA,xB,yB,xC,yC,xD,yD uint16) {
	return v.xA,v.yA,v.xB,v.yB,v.xC,v.yC,v.xD,v.yD
}

// Vor.: Die übergebenen Koordinaten müssen im geöffneten gfx-Fenster liegen.
// Eff.: Die Koordinaten der vier Eckpunkte des Vierecks sind
//       entsprechend der übergebenen Werte neu gesetzt.
func (v *data) SetKoordinaten (xA,yA,xB,yB,xC,yC,xD,yD uint16) {
	v.xA = xA
	v.yA = yA
	v.xB = xB
	v.yB = yB
	v.xC = xC
	v.yC = yC
	v.xD = xD
	v.yD = yD
}

// Vor.: -
// Erg.: Die Koordinaten des ersten Eckpunkts des Vierecks sind
//       geliefert.
func (v *data) GibKoordA() (uint16,uint16) {
	return v.xA, v.yA
}

// Vor.: -
// Erg.: Die Koordinaten des zweiten Eckpunkts des Vierecks sind
//       geliefert.
func (v *data) GibKoordB() (uint16,uint16) {
	return v.xB, v.yB
}

// Vor.: -
// Erg.: Die Koordinaten des dritten Eckpunkts des Vierecks sind
//       geliefert.
func (v *data) GibKoordC() (uint16,uint16) {
	return v.xC, v.yC
}

// Vor.: -
// Erg.: Die Koordinaten des vierten Eckpunkts des Vierecks sind
//       geliefert.
func (v *data) GibKoordD() (uint16,uint16) {
	return v.xD, v.yD
}

// Vor.: Die übergebenen Koordinaten müssen im geöffneten gfx-Fenster liegen.
// Eff.: Die Koordinaten des ersten Eckpunkts des Vierecks sind
//       entsprechend der übergebenen Werte neu gesetzt.
func (v *data) SetzeKoordA(x,y uint16) {
	if x < 1200 && y < 700 {
		v.xA = x
		v.yA = y
	}
}

// Vor.: Die übergebenen Koordinaten müssen im geöffneten gfx-Fenster liegen.
// Eff.: Die Koordinaten des zweiten Eckpunkts des Vierecks sind
//       entsprechend der übergebenen Werte neu gesetzt.		
func (v *data) SetzeKoordB(x,y uint16) {
	if x < 1200 && y < 700 {
		v.xB = x
		v.yB = y
	}
}

// Vor.: Die übergebenen Koordinaten müssen im geöffneten gfx-Fenster liegen.
// Eff.: Die Koordinaten des dritten Eckpunkts des Vierecks sind
//       entsprechend der übergebenen Werte neu gesetzt.
func (v *data) SetzeKoordC(x,y uint16) {
	if x < 1200 && y < 700 {
		v.xC = x
		v.yC = y
	}
}

// Vor.: Die übergebenen Koordinaten müssen im geöffneten gfx-Fenster liegen.
// Eff.: Die Koordinaten des vierten Eckpunkts des Vierecks sind
//       entsprechend der übergebenen Werte neu gesetzt.
func (v *data) SetzeKoordD(x,y uint16) {
	if x < 1200 && y < 700 {
		v.xD = x
		v.yD = y
	}
}

// Vor.: -
// Erg.: Die rgb-Farbwerte des Vierecks sind geliefert.		
func (v *data) GibFarbe() (uint8,uint8,uint8) {
	return v.r, v.g, v.b
}
// Vor.: -
// Eff.: Die Farbe des Vierecks ist entsprechend der übergebenen
//       rgb-Werte neu gesetzt.
func (v *data) SetzeFarbe (r,g,b uint8) {
	v.r = r
	v.g = g
	v.b = b
}


// Vor.: -
// Erg.: Eine Repräsentation des Vierecks in Form eines Strings ist geliefert.
func (v *data) String () string {
	return "Viereck mit den Koordinaten A("+fmt.Sprint(v.xA)+";"+fmt.Sprint(v.yA)+"), B("+fmt.Sprint(v.xB)+";"+fmt.Sprint(v.yB)+"), C("+fmt.Sprint(v.xC)+";"+fmt.Sprint(v.yC)+") und D("+fmt.Sprint(v.xD)+";"+fmt.Sprint(v.yD)+")"
}
	
	
// Vor.: -
// Erg.: Ein stets gleichlanger Bytestrom, der die serialisierte Form
//       des Vierecks darstellt, ist geliefert.
func (v *data) Kodieren () (b []byte) {
	b = (*(*[unsafe.Sizeof(*v)]byte)(unsafe.Pointer(&(*v))))[:] // Es wird rüberkopiert!
	return
}
	

// Vor.: -
// Eff.: Das Viereck hat jetzt die Eigenschaften, die im Bytestrom
//       enthalten waren. Seine vorher vorhandenen Eigenschaften
//       gibt es nicht mehr.
func (v *data) Dekodieren (b []byte) {
	(*v) = *(*data)(unsafe.Pointer(&b[0])) 
}
	
// Vor.: -
// Erg.: Eine tiefe Kopie des Vierecks v in Form einer Variablen des
//       Typs interface {} ist geliefert. Die Objekte v und 
//       v.Kopie().(Viereck) sind unterschiedliche Objekte, haben aber
//       komplett identische Eigenschaften.
func (v * data) Kopie () interface {} {
	var kopie *data = new(data) // 1. Erzeuge eine Kopie-Instanz!
	*kopie = *v // Der Inhalt des Structs von kopie sei bitte mit den Werten des Inhalts des Structs von v belegt.
	return kopie
}

// Vor.: -
// Erg.: Das Viereck ist in der aktuellen Stiftfarbe gezeichnet.
func (v *data) Zeichnen() {
	gfx.Stiftfarbe(v.r, v.g, v.b)
	gfx.Linie(v.xA, v.yA, v.xB, v.yB)
	gfx.Linie(v.xB, v.yB, v.xC, v.yC)
	gfx.Linie(v.xC, v.yC, v.xD, v.yD)
	gfx.Linie(v.xD, v.yD, v.xA, v.yA)
}

// Vor.: -
// Eff.: Das Viereck ist nun anklickbar und die Methode Angeklickt
//       kann dafür aufgerufen werden und liefert einen bool-Wert.
func (v *data) AktiviereKlickbar() {
	v.aktiv = true
}

// Vor.: -
// Eff.: Das Viereck ist nun nicht mehr anklickbar und die Methode
//       Angeklickt kann dafür nicht mehr aufgerufen werden.
func (v *data) DeaktiviereKlickbar() {
	v.aktiv = false
}

// Vor.: Das Viereck muss anklickbar sein (initial ist es das nicht),
//       die Methode AkitviereAnklickbar muss also mindestens einmal
//       aufgerufen worden sein.
// Erg.: True ist geliefert, wenn mit die übergebenen Koordinaten
//       innerhalb des Vierecks liegen, ansonsten ist false geliefert.
func (v *data) Angeklickt(x,y uint16) bool {						// Checkt, ob angeklickt
	
	if v.aktiv {
		var min, max func(a,b uint16) uint16
		min = func (a,b uint16) uint16 {
			if a < b {return a} else {return b}
		}
		max = func (a,b uint16) uint16 {
			if a > b {return a} else {return b}
		}
		
		return min(v.xA,v.xB) <= x && x <= max(v.xC,v.xD) && min(v.yA,v.yD) <= y && y <= max(v.yB,v.yC)
	}
	
	return false
}
