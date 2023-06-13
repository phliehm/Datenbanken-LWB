// Autor: Sebastian Herker
// Datum: 16.01.2022, zuletzt geändert: 16.01.2022
// Zweck: Abfragen von Daten mit GO

package main
import ("SQL")

func main () {
  var (
    conn SQL.Verbindung
    rs SQL.Ergebnismenge
    persnr,raum int
    name,rang string
    gehalt float32
  )

  conn = SQL.PgSQL ("user=lewein dbname=lewein")
  defer conn.Beenden ()

  println ("Verbindungstest erfolgreich")

	rs = conn.Anfrage ("SELECT * FROM Professoren;")
	for rs.GibtTupel () {
	  rs.LeseTupel (&persnr, &name, &rang, &gehalt, &raum)
	  println (persnr, name, rang, gehalt, raum)
	}
	rs.Schliessen ()
}
