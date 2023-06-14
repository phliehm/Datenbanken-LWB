<?php
class overview {
  private $tableName;
  private $columnNames = [];    //Spaltentitel
  private $tableRows = [];      //Tabelleninhalt

  //Konstruktor
  public function __construct(string $setTableName)   //Tabellenname wird Konstruktur mitgegeben
  {
    $this->tableName = $setTableName;   //Variable im Objekt mit übergebenem Wert belegen 
  }

  //interne Hilfsfunktion
  //Verbindung zur DB herstellen
  private function connectionStart() {
    $dbname = 'lewein';
    $dsn = 'pgsql:host=localhost;dbname='.$dbname;
    $password = 'niewel';
    $con = new PDO ($dsn,$dbname,$password);
    return  $con;
  }

  //Methoden der Klasse
  //Spaltennamen (information_schema.COLUMNS) laden
  public function loadColumns() {
    $con = $this->connectionStart();
    $queryCol = $con -> query("SELECT column_name
                              FROM information_schema.COLUMNS
                              WHERE TABLE_NAME = '$this->tableName';
                            ");
    foreach ($queryCol as $rows) {
      $this->columnNames[] = $rows[0];
      //echo $rows[0]."<br>";
    }
    // Reihenfolge der Spalten explizit vorgeben bei Tabellen, die über NATURAL JOIN entstehen
    if ($this->tableName == "tiere") {
      $this->columnNames = array("ti_id", "tr_id", "ua_id", "ti.name", "geschlecht",	"laenge",	"gewicht",	"zugangsdatum",	"artbezeichnung",	"terrarium");
    } else if ($this->tableName == "terrarien") {
      $this->columnNames = array("tr_id", "ht_id", "name", "beschreibung", "raumnummer", "breite",
                                      "breite_vorn", "tiefe", "hoehe");
    } else if  ($this->tableName == "tierpflege") {
      $this->columnNames = array("ad_id", "ti_id", "fu_id", "menge_min", "menge_max", "mahlkategorie");
    } else if  ($this->tableName == "terrariumpflege") {
      $this->columnNames = array("ad_id", "tr_id");
    } else if ($this->tableName == "verbrauch") {
      $this->columnNames = array("tr_id", "vm_id", "anzahl");
    }
    //"lagerpflege", "arbeitsdienste" funktionieren ohne eine Ersetzung hier

    //TODO: Schlüsselkombination für das Bearbeiten einer Zeile verwenden bei: verbrauch (vm_id & tr_id)
    
    //print_r($this->columnNames);
    //echo "<br>";
    $con = null;
  }

  //Tabelleninhalt (*) laden
  public function loadTable() {
    //PDObjekt erstellen
    $con = $this->connectionStart();
    //PDObjekt mit zurückgegebenen Daten füllen
    $firstColumn = $this->columnNames[0];
    
    if ($this->tableName == "tiere") {
      $queryRows = $con->query("SELECT TI_ID, TR_ID, UA_ID, TI.name,geschlecht,	laenge,	gewicht, 
                                      zugangsdatum,	artbezeichnung,	terrarium
                                FROM $this->tableName AS TI
                                NATURAL JOIN
                                (SELECT UA_ID, U.Bezeichnung AS Artbezeichnung
                                FROM Unterarten U) AS a
                                NATURAL JOIN 
                                (SELECT TR_ID, T.Name AS Terrarium
                                FROM Terrarien T) AS b
                                ORDER BY TI_ID;
                                ")->fetchAll(PDO::FETCH_NAMED);
      //var_dump($queryRows);
    } elseif ($this->tableName == "terrarien") {
      $queryRows = $con->query("SELECT tr_id, ht_id, tr.name, beschreibung, raumnummer, breite,
                                      breite_vorn, tiefe, hoehe, habitatbezeichnung
                                FROM $this->tableName AS TR
                                NATURAL JOIN
                                (SELECT HT_ID, H.Bezeichnung AS Habitatbezeichnung
                                FROM Habitattypen H) AS a
                                ORDER BY TR_ID;
                                ")->fetchAll(PDO::FETCH_NAMED);
    } elseif ($this->tableName == "lagerpflege") {
      $queryRows = $con->query("SELECT ad_id, vm_id, verbrauchsmittel
                                FROM $this->tableName AS AD
                                NATURAL JOIN
                                (SELECT VM_ID, VM.Bezeichnung AS verbrauchsmittel
                                FROM verbrauchsmittel VM) AS a
                                ORDER BY AD_ID;
                                ")->fetchAll(PDO::FETCH_NAMED);
    } elseif ($this->tableName == "arbeitsdienste") {
      $queryRows = $con->query("SELECT ad_id, at_id, arbeitsdienst, faelligkeit
                                FROM $this->tableName AS AD
                                NATURAL JOIN
                                (SELECT AT_ID, A.Bezeichnung AS arbeitsdienst
                                FROM arbeitsdiensttypen A) AS a
                                ORDER BY AD_ID;
                                ")->fetchAll(PDO::FETCH_NAMED);
    } elseif ($this->tableName == "verbrauch") {
      $queryRows = $con->query("SELECT TR_ID, VM_ID, ad_id, verbrauch.anzahl, terrarien.name AS terrarium,
                                verbrauchsmittel.bezeichnung AS verbrauchsmittel
                                
                                FROM $this->tableName
                                JOIN verbrauchsmittel USING (VM_ID)
                                JOIN terrarien USING (TR_ID)
                                
                                
                                ORDER BY TR_ID;
                                ")->fetchAll(PDO::FETCH_NAMED);
                                //var_dump($queryRows);
                                //arbeitsdiensttypen.Bezeichnung AS arbeitsdienst,
                                //arbeitsdienste.faelligkeit
                                //JOIN arbeitsdienste USING (ad_id)
                                //JOIN arbeitsdiensttypen USING (at_id)
    } elseif ($this->tableName == "tierpflege") {
      $queryRows = $con->query("SELECT tierpflege.ad_id, ti_id, fu_id, tierpflege.menge_min, tierpflege.menge_max, tierpflege.mahlkategorie, tiere.name AS tier, futter.bezeichnung AS futter
                                FROM $this->tableName
                                JOIN tiere USING (TI_ID)
                                JOIN futter USING (FU_ID)
                                ORDER BY ad_id;
                                ")->fetchAll(PDO::FETCH_NAMED);
                                //var_dump($queryRows);
    } elseif ($this->tableName == "terrariumpflege") {
      $queryRows = $con->query("SELECT ad_id,  tr_id, terrarium
                                FROM $this->tableName AS TP
                                NATURAL JOIN
                                (SELECT TR_ID, T.name AS terrarium
                                FROM terrarien T) AS a
                                ORDER BY TR_ID;
                                ")->fetchAll(PDO::FETCH_NAMED);
    } else {
      $queryRows = $con->query("SELECT * FROM $this->tableName ORDER BY $firstColumn")->fetchAll(PDO::FETCH_NAMED);
      //$queryRows->execute();
      //var_dump($queryRows);
    }
    
    //fetchAll(PDO::FETCH_GROUP|PDO::FETCH_ASSOC)
    //-> funktioniert prinzipiell IN EINFACHEN FÄLLEN, aber löscht id-Spalte & in verschachtelten Fällen Ergebnisse!
    //PDO::FETCH_ASSOC
    //-> gibt jede Spalte doppelt an (1x mit ihrem Namen als Index & 1x mit Zahl)
    //$queryRows->fetchAll(PDO::FETCH_NAMED); -> PDO::FETCH_NUM: teilen die Ergebnisse entsprechend ASSOC
    
    $i = 0;
    foreach ($queryRows as $row) {
      //var_dump($row); echo "<br>";
      $this->tableRows[$i] = $row; //Daten als Array (Ergebnis von FetchAll) in overview-Objekt gespeichert (OHNE 0.Spalte = ID)
      $i++;
    }
    //Array der Keys / der ERSTEN Spalte
    //$this->tableRows = array_map('reset', $this->tableRows);    
    //var_dump($this->tableRows);
    echo "<br>";
    $con = null;
  }

  //Erstellung & Anzeige der Tabelle mit den geladenen Spaltentiteln & Tabelleninhalten
  public function showOverview() {
    $this->loadColumns();         //jetzt sind die Spaltennamen (columnNames) gesetzt
    $this->loadTable();           //jetzt ist der Tabelleninhalt/-zeilen (tableRows) gesetzt
    //HTML-Ausgabe:

    //print_r(array_keys(current($this->tableRows)));

    if (count($this->tableRows) > 0) {
      echo "Inhalt der Tabelle ".$this->tableName."<br><br>";
      //var_dump($this->tableRows); echo "<br><br>";
      //var_dump($this->columnNames);
      echo "<br>";
      echo "
      <table>
        <thead>
          <tr>
            <th>
      ";
      echo implode('</th><th>', array_keys(current($this->tableRows)));
      echo "
            </th>
          </tr>
        </thead>
        <tbody>
      ";
      foreach ($this->tableRows as $key => $value) {
        array_map('htmlentities', $value);
        //var_dump($this->tableRows[$key]); echo $this->columnNames[0];
        if ($this->tableName == "verbrauch") {
          $id = $this->tableRows[$key][$this->columnNames[0]]."_".$this->tableRows[$key][$this->columnNames[1]];
        } else {
          $id = $this->tableRows[$key][$this->columnNames[0]];
        }
        //echo $key; echo "key-test<br><br>";
        echo "
          <tr>
            <td>
        ";
        echo implode('</td><td>', $value);
        echo "
            </td>
            <td>
              <a href='showChange.php?table=$this->tableName&id=$id'>bearbeiten</a>
            </td>
          </tr>
        ";
      } //link (href) wird mit GET-Variablen gefüllt und gesendet an showChange.php
      //id=$key kommt vom Array & beginnt daher bei 0 (Null)
      echo "
        </tbody>
      </table>
      ";
    }
  }
}
?>