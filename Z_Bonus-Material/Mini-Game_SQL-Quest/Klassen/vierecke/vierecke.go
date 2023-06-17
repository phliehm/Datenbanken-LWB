package vierecke

// Annalena Cyriacus
// 24.04.2023
// Klasse Vierecke

// Der ADT Viereck dient der Verwaltung von Vierecken. Jedes Viereck
// hat vier Koordinaten, nämlich (xA,yA), (xB,yB), (xC,yC) und (xD,yD)
// durch die es festgelegt wird. A ist dabei die obere linke Ecke.

// Vor.: -
// Erg.: Ein Viereck mit den Koordinaten (xA,yA), (xB,yB), (xC,yC) und
//       (xD,yD) ist geliefert. A ist dabei die obere linke Ecke.
// New (xA,yA,xB,yB,xC,yC,xD,yD uint16) *data // *data erfüllt das Interface Viereck

type Viereck interface {
	
	// Vor.: -
	// Erg.: Die Koordinaten aller vier Eckpunkte des Vierecks sind
	//       geliefert.
    GetKoordinaten () (xA,yA,xB,yB,xC,yC,xD,yD uint16)
    
    // Vor.: Die übergebenen Koordinaten müssen im geöffneten gfx-Fenster liegen.
	// Eff.: Die Koordinaten der vier Eckpunkte des Vierecks sind
	//       entsprechend der übergebenen Werte neu gesetzt.
    SetKoordinaten(xA,yA,xB,yB,xC,yC,xD,yD uint16)
    
    // Vor.: -
	// Erg.: Die Koordinaten des ersten Eckpunkts des Vierecks sind
	//       geliefert.
    GibKoordA() (uint16,uint16) 
    
    // Vor.: -
	// Erg.: Die Koordinaten des zweiten Eckpunkts des Vierecks sind
	//       geliefert.
    GibKoordB() (uint16,uint16)
    
    // Vor.: -
	// Erg.: Die Koordinaten des dritten Eckpunkts des Vierecks sind
	//       geliefert.
    GibKoordC() (uint16,uint16)
    
    // Vor.: -
	// Erg.: Die Koordinaten des vierten Eckpunkts des Vierecks sind
	//       geliefert.
    GibKoordD() (uint16,uint16)
    
    // Vor.: Die übergebenen Koordinaten müssen im geöffneten gfx-Fenster liegen.
	// Eff.: Die Koordinaten des ersten Eckpunkts des Vierecks sind
	//       entsprechend der übergebenen Werte neu gesetzt.
    SetzeKoordA(x,y uint16)
    
    // Vor.: Die übergebenen Koordinaten müssen im geöffneten gfx-Fenster liegen.
	// Eff.: Die Koordinaten des zweiten Eckpunkts des Vierecks sind
	//       entsprechend der übergebenen Werte neu gesetzt.
    SetzeKoordB(x,y uint16)
    
    // Vor.: Die übergebenen Koordinaten müssen im geöffneten gfx-Fenster liegen.
	// Eff.: Die Koordinaten des dritten Eckpunkts des Vierecks sind
	//       entsprechend der übergebenen Werte neu gesetzt.
    SetzeKoordC(x,y uint16)
    
    // Vor.: Die übergebenen Koordinaten müssen im geöffneten gfx-Fenster liegen.
	// Eff.: Die Koordinaten des vierten Eckpunkts des Vierecks sind
	//       entsprechend der übergebenen Werte neu gesetzt.
    SetzeKoordD(x,y uint16)
    
    // Vor.: -
    // Erg.: Die rgb-Farbwerte des Vierecks sind geliefert.
    GibFarbe() (r,g,b uint8)
    
    // Vor.: -
	// Eff.: Die Farbe des Vierecks ist entsprechend der übergebenen
	//       rgb-Werte neu gesetzt.
    SetzeFarbe(r,g,b uint8) 
	
	// Vor.: -
	// Erg.: Eine Repräsentation des Vierecks in Form eines Strings ist geliefert.
	String () string
	
	// Vor.: -
	// Erg.: Ein stets gleichlanger Bytestrom, der die serialisierte Form
	//       des Vierecks darstellt, ist geliefert.
	Kodieren () []byte
	
	// Vor.: -
	// Eff.: Das Viereck hat jetzt die Eigenschaften, die im Bytestrom
	//       enthalten waren. Seine vorher vorhandenen Eigenschaften
	//       gibt es nicht mehr.
	Dekodieren ([]byte)
	
	// Vor.: -
	// Erg.: Eine tiefe Kopie des Vierecks v in Form einer Variablen des
	//       Typs interface {} ist geliefert. Die Objekte v und 
	//       v.Kopie().(Viereck) sind unterschiedliche Objekte, haben aber
	//       komplett identische Eigenschaften.
	Kopie () interface {}
	
	// Vor.: -
	// Erg.: Das Viereck ist in der aktuellen Stiftfarbe gezeichnet.
	Zeichnen()

	// Vor.: -
	// Eff.: Das Viereck ist nun anklickbar und die Methode Angeklickt
	//       kann dafür aufgerufen werden und liefert einen bool-Wert.
	AktiviereKlickbar()
	
	// Vor.: -
	// Eff.: Das Viereck ist nun nicht mehr anklickbar und die Methode
	//       Angeklickt kann dafür nicht mehr aufgerufen werden.
	DeaktiviereKlickbar()
	
	// Vor.: Das Viereck muss anklickbar sein (initial ist es das nicht),
	//       die Methode AkitviereAnklickbar muss also mindestens einmal
	//       aufgerufen worden sein.
	// Erg.: True ist geliefert, wenn mit die übergebenen Koordinaten
	//       innerhalb des Vierecks liegen, ansonsten ist false geliefert.
	Angeklickt(x,y uint16) bool
}

