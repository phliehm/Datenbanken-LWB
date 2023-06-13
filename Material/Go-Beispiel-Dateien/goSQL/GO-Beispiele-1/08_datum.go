package main

//  Autor:  Oliver Schäfer
//  Datum:  Sa 26. Jan 15:51:07 CET 2019
//  Zweck:  Datumsverarbeitung in SQL

import (
  "SQL"
  "fmt"
  "time"
)

type freund struct {
  name    string
  groesse float32
  gebdat  time.Time
}

func main() {
  var (
    conn  SQL.Verbindung
    rs    SQL.Ergebnismenge
    tupel freund
  )

  conn = SQL.PgSQL ("user=lewein dbname=lewein")
  defer conn.Beenden ()

  conn.Ausfuehren (`
    CREATE TABLE freunde (
      name    VARCHAR,
      groesse NUMERIC (3,2),
      gebdat  TIMESTAMP
    );
    INSERT INTO freunde
      VALUES ('Hans', 1.92, '1956-01-13 10:00:00');
    INSERT INTO freunde
      VALUES ('Lisa', 1.63, '1979-02-11 13:30:56');
    INSERT INTO freunde
      VALUES ('Theo', 1.18, '2010-06-09 23:47:44');
  `)

  rs = conn.Anfrage (`SELECT name, groesse, gebdat
                      FROM   freunde;`)
  for rs.GibtTupel () {
    rs.LeseTupel (&tupel.name, &tupel.groesse, &tupel.gebdat)
    fmt.Println (tupel.name, tupel.groesse, tupel.gebdat,
                 SQL.Datum (tupel.gebdat),
                 SQL.Uhrzeit (tupel.gebdat))
  }
  rs.Schliessen ()

  conn.Ausfuehren ("DROP TABLE Freunde;")
}
