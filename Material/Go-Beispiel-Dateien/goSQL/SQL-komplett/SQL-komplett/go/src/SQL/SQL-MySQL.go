package SQL

//  Treiber für mySQL (erfordert Go 1.2 oder größer)
//  Quelle: golang.org/s/sqldrivers
import _ "github.com/go-sql-driver/mysql"

//  Treiberspezifische Funktion  zum Herstellen  einer Verbin-
//  dung mit der Datenbank. Details zum Aufbau der Zeichenket-
//  te <datenquelle> siehe unten.

func MySQL (datenquelle string) Verbindung {
//
  const (
    DBTYPE = "mysql"
    OPTION = "?parseTime=true&charset=utf8"
  )
  return verbinden (datenquelle, OPTION, DBTYPE)
}

//  Definition der Zeichenkette  <datenquelle> anhand  einiger
//  Beispiele. Weitere Informationen findet man  im Netz unter
//  github.com/go-sql-driver/mysql/blob/master/README.md.

//  Localhost, Standardport 3306
//  username:password@/dbname

//  Ferner Host
//  username:password@tcp(host:port)/dbname

//  Minimal
//  username@/dbname

//  Maximal
//  username:password@protocol(address)/dbname?param=value


//  **********************************************************
//             Zuordnung von mySQL-Typen zu GO-Typen
//  **********************************************************
//  -----------------+----------------------------------------
//  mySQL            | Go
//  -----------------+----------------------------------------
//  Smallint         | int64
//  Integer          | int64
//  Bigint           | int64
//  Boolean          | bool
//  Char             | string
//  Text             | string
//  Varchar          | string
//  Double           | float64
//  Numeric          | float64
//  Real             | float64
//  Date             | time.Time
//  Time             | time.Time
//  TimeTZ           | time.Time
//  Timestamp        | time.Time
//  TimestampTZ      | time.Time
//  -----------------+----------------------------------------
