package SQL

import (
  "database/sql"
  "log"
  "strings"
  "time"
)

type connection struct { id *sql.DB }
type resultset  struct { id *sql.Rows }

func verbinden (parameter, option, dbtype string) *connection {
//
  var (
    conn    *connection
    p_conn  *sql.DB
    err     error
    source  string
  )

  source = parameter + option
  p_conn, _ = sql.Open (dbtype, source)
  /* Ping() hinzugefügt (23.01.2017, Fehler behoben am 19.01.2020),
     da sql.Open nur die korrekte  Syntax der  Parameter überprüft,
     nicht aber  die Verbindung aufbaut. */
  if err != nil {
    log.Fatal (err)
    return nil
  } else {
    err = p_conn.Ping ()
    if err != nil {
      log.Fatal (err)
      return nil
    }
    conn = new (connection)
    (*conn).id = p_conn
    return conn
  }
}

//  Methoden auf Verbindungen

func (db *connection) Anfrage (SELECT_Anweisung string) Ergebnismenge {
//
  var (
    rs        *resultset
    resultSet *sql.Rows
    err       error
  )
  resultSet, err = ((*db).id).Query (SELECT_Anweisung)
  if err != nil {
    log.Fatal (err)
    return nil
  } else {
    rs = new (resultset)
    (*rs).id = resultSet
    return rs
  }
}

func (db *connection) Ausfuehren (SQL_Anweisung string) (rowsAffected int64) {
//
  var (
    res sql.Result
    err error
  )

  res, err = ((*db).id).Exec (SQL_Anweisung)
  if err != nil {
    log.Fatal (err)
  }
  rowsAffected, _ = res.RowsAffected ()
  return
}

func (db *connection) Beenden () {
//
  var err error
  err = ((*db).id).Close ()
  if err != nil {
    log.Fatal (err)
  }
}

//  Methoden auf Ergebnismengen

func (rs *resultset) GibtTupel () bool {
//
  return ((*rs).id).Next ()
}

func (rs *resultset) Schliessen () () {
//
  ((*rs).id).Close ()
  return
}

func (rs *resultset) LeseTupel (args ...interface{}) () {
//
  var err error
  err = ((*rs).id).Scan (args...)
  if err != nil {
    log.Fatal (err)
  }
  return
}

func (rs *resultset) AnzahlAttribute () int {
//
  var (
    atts  []string
    err   error
  )
  atts, err = ((*rs).id).Columns ()
  if err != nil {
    log.Fatal (err)
  }
  return len (atts)
}

func (rs *resultset) Attribute () []string {
//
  var (
    atts  []string
    err   error
  )
  atts, err = ((*rs).id).Columns ()
  if err != nil {
    log.Fatal (err)
  }
  return atts
}

func (rs *resultset) AttributName (n int) (ret string) {
//
  var (
    atts  []string
    err   error
  )
  atts, err = ((*rs).id).Columns ()
  if err != nil {
    log.Fatal (err)
  }
  if n >= len (atts) {
    return ""
  } else {
    return atts[n]
  }
}

func (rs *resultset) AttributPos (s string) (ret int) {
//
  var (
    atts  []string
    err   error
  )
  atts, err = ((*rs).id).Columns ()
  if err != nil {
    log.Fatal (err)
  }
  for i,attname := range atts {
    if strings.ToLower (s) == strings.ToLower (attname) {
      return i
    }
  }
  return -1
}

//  Funktionen zur Typ-Konvertierung von time.Time-Werten
func Datum   (t time.Time) string { return t.Format ("02.01.2006") }
func Uhrzeit (t time.Time) string { return t.Format ("15:04") }
