package SQL

//  Treiber für postgreSQL (Pure Go)
//  Quelle: golang.org/s/sqldrivers
import _ "github.com/lib/pq"

//  Treiberspezifische Funktion  zum Herstellen  einer Verbin-
//  dung mit der Datenbank. Details zum Aufbau der Zeichenket-
//  te <datenquelle> siehe unten.

func PgSQL (datenquelle string) Verbindung {
//
  const (
    DBTYPE="postgres"
    OPTION=" sslmode=disable"
  )
  return verbinden (datenquelle, OPTION, DBTYPE)
}

//  Definition der Zeichenkette <datenquelle>:
//  <datenquelle> besteht aus einer Folge von 'keyword = Wert'
//  Paaren. Folgende Optionen stehen zur Auswahl.Wird eine Op-
//  weggelassen, wird der jeweilige Standardwert verwendet.

//  host     = Name/IP-Adresse des Datenbankservers
//             (Standard: localhost)
//  port     = Portnummer der DB-Verbindung
//             (Standard: 5432)
//  dbname   = Name der Datenbank
//             (Standard: wie Benutzer)
//  user     = Benutzername
//  password = Passwort der Datenbankverbindung
//             (Standard: leer)
//  timeout  = Timeout in Sekunden, 0 bedeutet deaktiviert
//             (Standard: 0)

//  **********************************************************
//          Zuordnung von postgreSQL-Typen zu GO-Typen
//  **********************************************************
//  -----------------+----------------------------------------
//  PostgreSQL       | Go
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
