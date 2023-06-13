package SQL

/*  Einfacher Wrapper zum Umgang mit Datenbasen auf SQL-Basis.
    Der Wrapper baut auf dem Paket database/sql sowie spezifi-
    sche Datenbanktreibern auf, deren Spezifika in den Dateien
    SQL-<Treibername>.go festgelegt werden.
    Der Wrapper stellt eine schmale Schnittstelle für  den Um-
    gang mit  Datenbanksystemen  für den  Schulunterricht  zur
    Verfügung.

    (c) Oliver Schäfer                               2015-2017
        Version 0.40.2017-03-05
*/

type Verbindung interface {
  //  Die treiberspezifische Funktion <Treibername> stellt ei-
  //  ne  Verbindung zur Datenbank  her und liefert ein Daten-
  //  bankhandle vom Typ 'Verbindung'. Die  Funktion  erwartet
  //  eine treiberspezifische Zeichenkette, die die Datenquel-
  //  le beschreibt. Genaueres dazu findet man in der jeweili-
  //  gen Datei SQL-<Treibername>.go.

  Anfrage (SELECT_Anweisung string) Ergebnismenge
  // Vor.: Die Datenbank-Verbindung  ist  geöffnet, die Ergeb-
  //       nismenge ist nicht initialisiert.
  // Eff.: Initialisiert und  liefert  die Ergebnismenge einer
  //       SELECT-Anweisung, sofern  die  SELECT-Anfrage   vom
  //       Datenbank-Backend   fehlerfrei   ausgeführt  werden
  //       konnte. Andernfalls ist 'nil' geliefert.

  Ausfuehren (SQL_Anweisung string) int64
  // Vor.: Die Datenbank-Verbindung ist geöffnet.
  // Eff.: Führt die angegebene SQL-Anweisung auf der geöffne-
  //       ten Datenbankverbindung  aus. Die  Funktion liefert
  //       die Anzahl betroffener Tupel, die bei  einer ausge-
  //       führten INSERT-, UPDATE- oder DELETE-Anweisung  be-
  //       troffen sind.

  Beenden ()
  // Vor.: Die Datenbank-Verbindung ist geöffnet.
  // Eff.: Die geöffnete Datenbank-Verbindung ist geschlossen,
  //       weitere SQL-Kommandos können  nicht mehr  abgesetzt
  //       werden.
}

type Ergebnismenge interface {
  GibtTupel () bool
  //  Vor.: Die Ergebnismenge ist durch eine Anfrage definiert
  //        und zugewiesen.
  //  Eff.: Wenn es noch ein Tupel der Anfrage  hinter der ak-
  //        tuellen Cursor-Position gibt, ist 'true' geliefert
  //        die Cursorposition ist definiert und zeigt auf das
  //        aktuelle Tupel. Andernfalls ist 'false' geliefert.
  //        Befand sich der Cursor bereits auf der letzten Tu-
  //        pel,  bleibt die Cursor-Position  unverändert (auf
  //        dem letzten Tupel der Ergebnismenge).

  LeseTupel (...interface{}) ()
  //  Vor.: Die Cursorposition  ist  definiert, die Anzahl der
  //        übergebenen  Argumente  stimmt mit der Anzahl  der
  //        Felder der aktuellen Ergebnismenge überein.
  //  Eff.: Die Attributwerte des aktuellen Tupels sind an die
  //        übergebenen Variablen(-adressen) kopiert worden.
  //        Dabei findet - abhängig vom DB-Treiber - ein Type-
  //        Casting statt.

  Attribute () []string
  //  Vor.: Die Ergebnismenge ist durch eine Anfrage definiert
  //        und zugewiesen.
  //  Eff.: Eine Liste der Attributnamen der Ergebnismenge ist
  //        geliefert.

  AnzahlAttribute () int
  //  Vor.: Die Ergebnismenge ist durch eine Anfrage definiert
  //        und zugewiesen.
  //  Eff.: Die Anzahl der Attribute der Ergebnismenge ist ge-
  //        liefert.

  AttributName (int) string
  //  Vor.: Die Ergebnismenge ist durch eine Anfrage definiert
  //        und zugewiesen. Die Ergebnismenge ist nicht leer.
  //  Eff.: Der Name des Attributes an  der angegebenen Index-
  //        Position (Zählbeginn: 0) ist geliefert.

  AttributPos (string) int
  //  Vor.: Die Ergebnismenge ist durch eine Anfrage definiert
  //        und zugewiesen. Die Ergebnismenge ist nicht leer.
  //  Eff.: Der Index des Attributes zum Angegebenen Namen ist
  //        geliefert. Existiert  der angegebene  Attributname
  //        nicht, ist -1 geliefert.

  Schliessen ()
  //  Vor.: Die Ergebnismenge ist definiert.
  //  Eff.: Die Ergebnismenge ist nicht definiert.
}

//  **********************************************************
//        Konvertierungsfunktionen für Datum und Zeit
//  **********************************************************
//  Wegen treiberspezifischer  Unterschiede in  der Behandlung
//  der für  Datenbankanwendungen  typischen  Datentypen  DATE
//  und TIME stellt das Modul zwei Funktionen  zur  Verfügung,
//  die eine Umwandlung in eine Zeichenkette erlauben.
//
//  Datum   (time.Time) string
//          liefert das Datum im Format 'DD.MM.YYYY'
//  Uhrzeit (time.Time) string
//          liefert die Uhrzeit im Format 'HH:MM'
