package folgeninvert

/* Eine Folge ist eine Auflistung von fortlaufend nummerierten Objekten 
 * (des gleichen Datentyps). Dasselbe Objekt kann in einer Folge mehr-
 * fach auftreten. Die Indizierung der Objekte beginnt mit 0!
 *
 * Als Objekte der Folge betrachten wir beliebige Instanzen eines Datentyps,
 * ggf. auch eines abstrakten Datentyps.*/
 
type Element interface {}   // <-- Das ist unser heutiges Ziel!

// Vor.: -
// Erg.: eine neue, leere Folge von Objekten des 'echten' Datentyps von e,
//       genannt: 'Elementtyp der Folge' 
// New (e Element) Folge

//BEMERKUNG: In Go muss man sich nicht um die Speicherbereinigung kümmern!
//           Daher kein DISPOSE bzw. FREE.

type Folgeinvert interface {
// Vor.: -
// Eff.: War e nicht vom Elementtyp der Folge, so ist nichts passiert,
//       andernfalls: e ist in die Folge NACH dem aktuellen Element
//       eingefügt. Gab es kein aktuelles Element, so ist e am Ende
//       der Folge eingefügt.
	Einfuegen (e Element)
	
// Vor.: -
// Eff.: Gab es ein Element mit Index n in der Folge, so ist es nun
//       aktuelles Element. Andernfalls ist kein Element aktuell.
	Positionieren (n uint)
	
// Vor.: -
// Erg.: War die Folge leer oder gab es kein aktuelles Element, so 
//       ist die Länge der Folge geliefert, andernfalls ist der
//       (Positions-)Index des aktuellen Elements geliefert. 
	AktuellerIndex () uint
	
// Erg.: Falls es kein aktuelles Element gab, ist ok false und e
//       ist eine Elementinstanz mit dem Typs der Folge, die selbst jedoch 
//       kein Bestandteil der Objektfolge ist. Ansonsten ist ok true und
//       das aktuelle Element e ist geliefert.
	AktuellesElement () (e Element, ok bool)
	
// Vor.: -
// Eff.: War kein Element aktuell so ist nichts passiert.
//       Ansonsten ist das dem ehemals aktuellen Element folgende Element
//       nun aktuell. Gibt es ein solches Element nicht, so ist kein 
//       Element aktuell.
	Vor ()

// Vor.: -
// Eff.: War das Element mit Index 0 aktuell, so ist nichts passiert.
//       War kein Element aktuell, so ist nun das letzte Element aktuell.
//       Ansonsten ist das dem ehemals aktuellen Element vorhergehende
//	     Element aktuell. 
	Zurueck ()

// Vor.: -
// Eff.: Gab es ein aktuelles Element, so ist es aus der Liste
//       entfernt, alle anderen Elemente bleiben in gleicher Reihenfolge.
//       Das vorherige folgende Element ist jetzt aktuell. Gibt es ein 
//       solches Element nicht, so ist kein Element aktuell.
//       Gab es kein aktuelles Element, so ist nichts passiert.
	Loeschen ()
	
// Vor.: -
// Erg.: Die Laenge der Folge, d. h. die Anzahl der Elemente
//       in der Folge, ist geliefert.
	Laenge () uint

}
