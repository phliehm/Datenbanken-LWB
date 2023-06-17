package folgen

// Autor: St. Schmidt
// Datum: 13.09.2022
// Zweck: Implementierung des ADT Folge mit einer doppelt-verketteten
//        mit Verwaltungsblock und mit Ringliste inklusive Dummy-Element

//NEU: Polymorph - Die Elemente können von EINEM beliebigen Typ sein!!

import "reflect"

type data struct {
	laenge  uint
	aktIndex uint
	aktElement *knoten
	anker     *knoten
}

type knoten struct {
	inhalt     Element
	naechster  *knoten
	vorheriger *knoten
}

// Vor.: -
// Erg.: eine neue, leere Folge von Objekten des 'echten' Datentyps von e,
//       genannt: 'Elementtyp der Folge' 
func New (e Element) *data {
	var f *data
	var k *knoten
	f = new(data)
	k = new(knoten)
	f.aktElement = k
	f.anker      = k	
	k.inhalt = e  // ACHTUNG: UNSAUBER FÜR NICHT-KONKRETE DTen!!---------
	k.naechster = k
	k.vorheriger= k
	return f
}

// Vor.: -
// Eff.: War e nicht vom Elementtyp der Folge, so ist nichts passiert,
//       andernfalls: e ist in die Folge VOR dem aktuellen Element
//       eingefügt. Gab es kein aktuelles Element, so ist e am Ende
//       der Folge eingefügt.
func (f *data)	EinfuegenVor (e Element) {
	if reflect.TypeOf (e) == reflect.TypeOf (f.anker.inhalt) {
		var k *knoten
		k = new (knoten)
		//neuen Knoten mit seinen Werten setzen ...
		k.inhalt = e  // ACHTUNG: UNSAUBER FÜR NICHT-KONKRETE DTen!!---------
		k.naechster  = f.aktElement
		k.vorheriger = f.aktElement.vorheriger
		//neuen Knoten 'einklinken'
		f.aktElement.vorheriger.naechster= k
		f.aktElement.vorheriger = k
		//weitere Invarianten ...
		f.laenge++
		f.aktIndex++
	}
}

// Vor.: -
// Eff.: War e nicht vom Elementtyp der Folge, so ist nichts passiert,
//       andernfalls: e ist in die Folge NACH dem aktuellen Element
//       eingefügt. Gab es kein aktuelles Element, so ist e am Ende
//       der Folge eingefügt.
func (f *data)	EinfuegenNach (e Element) {
	if reflect.TypeOf (e) == reflect.TypeOf (f.anker.inhalt) {
		var k *knoten
		k = new (knoten)
		//neuen Knoten mit seinen Werten setzen ...
		k.inhalt = e  // ACHTUNG: UNSAUBER FÜR NICHT-KONKRETE DTen!!---------
		k.naechster  = f.aktElement.naechster
		k.vorheriger = f.aktElement
		//neuen Knoten 'einklinken'
		f.aktElement.naechster.vorheriger = k
		f.aktElement.naechster = k
		//weitere Invarianten ...
		f.laenge++
		f.aktIndex++
	}
}
	
// Vor.: -
// Eff.: Gab es ein Element mit Index n in der Folge, so ist es nun
//       aktuelles Element. Andernfalls ist kein Element aktuell.
func (f *data) Positionieren (n uint) {
	if n < f.laenge {
		f.aktIndex = n
	} else {
		f.aktIndex = f.laenge
	}
	//nun noch den Zeiger 'aktuelles 'setzen ..
	f.aktElement = f.anker
	for i:= uint(0);i <= f.aktIndex;i++ {
		f.aktElement = f.aktElement.naechster
	}
}
	
// Vor.: -
// Erg.: War die Folge leer oder gab es kein aktuelles Element, so 
//       ist die Länge der Folge geliefert, andernfalls ist der
//       (Positions-)Index des aktuellen Elements geliefert. 
func (f *data) AktuellerIndex () uint {
	if f.aktIndex == f.laenge {
		return f.laenge
	} else {
		return f.aktIndex
	}
}
	
// Vor.: - 
// Erg.: Falls es kein aktuelles Element gab, ist ok false und e
//       ist eine Elementinstanz mit dem Typs der Folge, die selbst jedoch 
//       kein Bestandteil der Objektfolge ist. Ansonsten ist ok true und
//       das aktuelle Element e ist geliefert.
func (f *data) AktuellesElement () (e Element, ok bool) {
	if f.aktIndex < f.laenge {
		ok = true
	} else {
		ok = false
	}
	e = f.aktElement.inhalt  // ACHTUNG: UNSAUBER FÜR NICHT-KONKRETE DTen!!---------
	return
}
	
// Vor.: -
// Eff.: War kein Element aktuell so ist nichts passiert.
//       Ansonsten ist das dem ehemals aktuellen Element folgende Element
//       nun aktuell. Gibt es ein solches Element nicht, so ist kein 
//       Element aktuell.
func (f *data) Vor () {
	if f.aktIndex == f.laenge {
		return
	} else {
		f.aktIndex++
		f.aktElement = f.aktElement.naechster
	}
}	
		
// Vor.: -
// Eff.: War das Element mit Index 0 aktuell, so ist nichts passiert.
//       War kein Element aktuell, so ist nun das letzte Element aktuell.
//       Ansonsten ist das dem ehemals aktuellen Element vorhergehende
//	     Element aktuell. 
func (f *data) Zurueck () {
	if f.aktIndex == 0 {
		return
	} else {
		f.aktIndex--
		f.aktElement = f.aktElement.vorheriger
	}
}

// Vor.: -
// Eff.: Gab es ein aktuelles Element, so ist es aus der Liste
//       entfernt, alle anderen Elemente bleiben in gleicher Reihenfolge.
//       Das vorherige folgende Element ist jetzt aktuell. Gibt es ein 
//       solches Element nicht, so ist kein Element aktuell.
//       Gab es kein aktuelles Element, so ist nichts passiert.
func (f *data) Loeschen () {
	if f.laenge == f.aktIndex { //kein akt. Element: Nichts zu tun
		return
	} else { // Es gibt ein aktuelles Element ...
		f.aktElement.vorheriger.naechster = f.aktElement.naechster
		f.aktElement.naechster.vorheriger = f.aktElement.vorheriger
		f.aktElement = f.aktElement.naechster
		f.laenge--
	}
}
	
// Vor.: -
// Erg.: Die Laenge der Folge, d. h. die Anzahl der Elemente
//       in der Folge, ist geliefert.
func (f *data) Laenge () uint {
	return f.laenge
}


	
