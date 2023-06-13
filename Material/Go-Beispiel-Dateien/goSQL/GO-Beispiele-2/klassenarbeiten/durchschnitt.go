package main

//  Autor:  Oliver Schäfer
//  Datum:  Sa 26. Jan 16:05:12 CET 2019
//  Zweck:  Berechnung der Durchschnitte von Klassenarbeiten
//          Ausgabe der Ergebnisse unter Verwendung des felder-Pakets

import ("SQL"; "gfx"; "felder"; "fmt")

type arbeit struct {
  kname   string
  anr     int
  thema   string
  anz     [6]int
}

func main() {
  var (
    kname,anr,thema,average felder.Feld
    n [6]felder.Feld
    conn SQL.Verbindung
    rs SQL.Ergebnismenge
    tupel arbeit
    query,eingabe string
  )

  // Bitte die IP-Adresse des eigenen Hosts eintragen!
  conn = SQL.PgSQL ("user=lewein dbname=lewein")

  defer conn.Beenden ()

  gfx.Fenster (620, 100)

  // Suchfeld: Erlaubt sind nur Sek-I-Klassen von 7-10 und a-e
  kname = felder.New ( 10, 10,  5, 'l', "Klasse")
  kname.SetzeErlaubteZeichen ("17890abcde%")

  // Ausgabefelder
  anr   = felder.New (  70, 10,  2, 'l', "Nr")
  anr.SetzeErlaubteZeichen ("123456")
  thema = felder.New ( 100, 10, 50, 'l', "Thema")
  thema.SetzeErlaubteZeichen (felder.Ascii)
  for k := uint16(0); k < 6; k++ {
    n[k] = felder.New (100+k*30,50,2,'l',fmt.Sprintf ("%d",k+1))
    (n[k]).SetzeErlaubteZeichen (felder.Digits)
  }
  average = felder.New (290,50,4,'l',"Mittel")
  average.SetzeErlaubteZeichen ("0123456789,.")

  for {
    // Bitte Klasse eingeben
    kname.Leere ()
    eingabe = kname.Edit ()
    if eingabe == "" { break }
    // Hier wird die Anfrage dynamisch erzeugt und an den DB-Server geschickt
    query = fmt.Sprintf (`
      SELECT  k.name,a.anr,a.thema,a.anz1,a.anz2,a.anz3,a.anz4,a.anz5,a.anz6
      FROM    klassen k natural join arbeiten a
      WHERE   k.name LIKE '%s'
      ORDER   BY k.name;`, eingabe)
    rs = conn.Anfrage (query)

    // Ausgabe der gefundenen Tupel
    for nr := uint16(0); rs.GibtTupel (); nr += 90 {
      rs.LeseTupel (&tupel.kname, &tupel.anr, &tupel.thema,
                    &tupel.anz[0], &tupel.anz[1], &tupel.anz[2],
                    &tupel.anz[3], &tupel.anz[4], &tupel.anz[5])
      // Fülle die Felder mit den gelesenen Werten
      kname.Schreibe (tupel.kname)
      anr.Schreibe (fmt.Sprintf ("%d",tupel.anr))
      thema.Schreibe (tupel.thema)
      for k := 0; k < 6; k++ { (n[k]).Schreibe (fmt.Sprintf ("%d",tupel.anz[k])) }
      // Hintergrundfarbe des Durchschnittsfeldes nach Wert steuern
      if arithMittel (tupel.anz) > 3 {
        average.SetzeHintergrundfarbe (0xFF,0xCF,0xCF)
      } else {
        average.SetzeHintergrundfarbe (0xD7,0xF1,0xB4)
      }
      average.Schreibe (fmt.Sprintf ("%3.2f", arithMittel (tupel.anz)))
      gfx.TastaturLesen1 ()
      gfx.TastaturLesen1 ()
    }
    rs.Schliessen ()
  }
}

func arithMittel (noten [6]int) float32 {
  var z,n float32
  for k := 0; k < 6; k++ {
    z += float32(k+1)*float32(noten[k])
    n += float32(noten[k])
  }
  return z/n
}
