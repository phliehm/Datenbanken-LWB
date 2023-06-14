<?php

use formular as GlobalFormular;

class formular {
  private $tableName;
  public $columnNames = [];
  private $inserts = [];

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
      $this->columnNames = array("ti_id", "tr_id", "ua_id", "geschlecht",	"laenge",	"gewicht",	"zugangsdatum");
    } else if ($this->tableName == "terrarien") {
      $this->columnNames = array("tr_id", "ht_id", "name", "beschreibung", "raumnummer", "breite",
                                  "breite_vorn", "tiefe", "hoehe");
    } else if  ($this->tableName == "tierpflege") {
      $this->columnNames = array("ad_id", "ti_id", "fu_id", "menge_min", "menge_max", "mahlkategorie");
    } else if  ($this->tableName == "terrariumpflege") {
      $this->columnNames = array("ad_id", "tr_id");
    } else if ($this->tableName == "verbrauch") {
      $this->columnNames = array("tr_id", "vm_id", "anzahl", "ad_id");
    }
    //print_r($this->columnNames);
    //echo "<br>";
  }

  //Spalteninhalt einer angegebenen Spalte laden
  public function loadFirstColumns() {
    $con = $this->connectionStart();
    if ($this->tableName == 'habitattypen') {
      $columns = array("ht_id", "bezeichnung");
    } else if ($this->tableName == 'terrarien') {
      $columns = array("tr_id", "name");
    } else if ($this->tableName == 'unterarten') {
      $columns = array("ua_id", "bezeichnung");
    } else if ($this->tableName == 'verbrauchsmittel') {
      $columns = array("vm_id", "bezeichnung");
    } else if ($this->tableName == 'arbeitsdiensttypen') {
      $columns = array("at_id", "bezeichnung");
    } else if ($this->tableName == 'arbeitsdienste') {
      $columns = array("ad_id", "arbeitsdienst");
    } else if ($this->tableName == 'tiere') {
      $columns = array("ti_id", "name");
    } else if ($this->tableName == 'futter') {
      $columns = array("fu_id", "bezeichnung");
    }
    $columsNames = $columns[0].", ".$columns[1];
    $order = $columns[0];

    if ($this->tableName == 'arbeitsdienste') {
      $queryCol = $con -> query("SELECT $columsNames            
                              FROM arbeitsdienste AS AD
                              NATURAL JOIN
                              (SELECT AT_ID, A.Bezeichnung AS arbeitsdienst
                              FROM arbeitsdiensttypen A) AS a
                              ORDER BY AD_ID;
                            ")->fetchAll(PDO::FETCH_KEY_PAIR);
                            /*
                            SELECT ad_id, at_id, arbeitsdienst, faelligkeit
                                FROM $this->tableName AS AD
                                NATURAL JOIN
                                (SELECT AT_ID, A.Bezeichnung AS arbeitsdienst
                                FROM arbeitsdiensttypen A) AS a
                                ORDER BY AD_ID;
                             */
    } else {
      $queryCol = $con -> query("SELECT $columsNames
                              FROM $this->tableName
                              ORDER BY $order;
                            ")->fetchAll(PDO::FETCH_KEY_PAIR);
    }
    
    $i = 0;
    foreach ($queryCol as $rows) {
      //$this->tableRows[$i] = $row;
      $this->inserts[$i] = $rows;
      $i++;
    }
    //echo "loadFirstColumns: <br>";
    //var_dump($this->inserts);
    //echo "<br>";
  }

  //Inhalt einer Zeile laden
  public function loadRow($id) {
    $this->loadColumns();
    //echo $id;
    $con = $this->connectionStart();
    $idName = $this->columnNames[0];
    //echo $idName;
    $queryCol = $con -> query("SELECT *
                              FROM $this->tableName
                              WHERE $idName = $id;
                            ")->fetchAll(PDO::FETCH_NAMED);
                            /*WITH cte AS (SELECT *,
                              ROW_NUMBER() OVER () AS row_number 
                              FROM 
                              $this->tableName)
                              SELECT *
                              FROM cte
                              WHERE row_number = $id;
                              dasselbe Ergebnis erzeugt folgende Abfrage:
                              SELECT *
                              FROM
                                (SELECT *,
                                ROW_NUMBER() OVER () AS row_number 
                                FROM 
                                $this->tableName) AS sub
                              WHERE 
                              row_number = $id;
 */
    $i = 0;
    foreach ($queryCol as $rows) {
      //$this->tableRows[$i] = $row;
      $this->inserts[$i] = $rows;
      //echo "loadRow: <br>";
      //var_dump($this->inserts[$i]);
    }
    //echo "loadRow: ".$this->inserts[0][$this->columnNames[0]]."<br>";
    //var_dump($this->inserts[0][$this->columnNames[0]]);
  }

  //Inhalt einer Zeile laden, die über eine Schlüsselkombination gegeben ist
  public function loadRowArr(array $id) {
    $this->loadColumns();
    //echo $id;
    $con = $this->connectionStart();

    //WHERE-String für query zusammenbauen
    $idName = $this->columnNames[0];
    $idKey = $id[0];
    $whereString = $idName."=".$idKey;
    for ($n=1; $n < count($id); $n++) { 
      $idName = $this->columnNames[$n];
      $idKey = $id[$n];
      $whereString = $whereString." AND ".$idName."=".$idKey;
    }
    
    //echo $whereString;
    $queryCol = $con -> query("SELECT *
                              FROM $this->tableName
                              WHERE $whereString;
                            ")->fetchAll(PDO::FETCH_NAMED);
                            /*WITH cte AS (SELECT *,
                              ROW_NUMBER() OVER () AS row_number 
                              FROM 
                              $this->tableName)
                              SELECT *
                              FROM cte
                              WHERE row_number = $id;
                              dasselbe Ergebnis erzeugt folgende Abfrage:
                              SELECT *
                              FROM
                                (SELECT *,
                                ROW_NUMBER() OVER () AS row_number 
                                FROM 
                                $this->tableName) AS sub
                              WHERE 
                              row_number = $id;
 */
    $i = 0;
    foreach ($queryCol as $rows) {
      //$this->tableRows[$i] = $row;
      $this->inserts[$i] = $rows;
      //echo "loadRow: <br>";
      //var_dump($this->inserts[$i]);
    }
    //echo "loadRow: ".$this->inserts[0][$this->columnNames[0]]."<br>";
    //var_dump($this->inserts);
  }

  //sendet ein Update an die DB mit this->tableName und der überegebenen ID sowie den als Array übergebenen Werten/POSTs
  public function updateRow(array $updates) {
    $this->loadColumns();
    $con = $this->connectionStart();
    //var_dump($updates);
    $updateString = "";
    $notStringAtribut = array('temperatur', 'luftfeuchtigkeit', 'helligkeit', 'waermeleistung', 'lager_min', 'lager_akt', 'kosten', 'wasserverbrauch', 'hoehe', 'breite', 'tiefe', 'laenge', 'gewicht', 'breite_vorn', 'menge_min', 'menge_max', 'mahlkategorie', 'anzahl', );
    $boolAtribut = array('kuenstlich', 'spruehen', 'bepflanzt');
    $dateAtribut = array('faelligkeit', 'zugangsdatum');
    for ($n = 0; $n < count($updates); $n++) {
      //echo "<br>classFormular.php/updateRow: POST-Feld mit Namen '".$this->columnNames[$n]."' enthält den POST-Wert ".$updates[$n];
      //WENN $updateString nicht leer ist, wird Komma angehängt
      if ($updateString != "") {
        $updateString = $updateString.", ";
      }
      //Anhängen der einzelnen Einträge an $updateString
      //Muss der Eintrag ein String oder ein Int/Date/...-Wert sein? ('' oder null eintragen?)
      //Alternative zu strpos: gettype($updates[$n]) -> müsste man aber mit den Eintragungen in DB & nicht aus dem Formular machen, aber die können auch fehlen -> keine Schlussfolgerung möglich
      if (strpos($this->columnNames[$n], "_id") !== false
          OR in_array($this->columnNames[$n], $notStringAtribut) 
          OR in_array($this->columnNames[$n], $boolAtribut)) {
        //für bool-Atribute die Integer umwandeln
        if (in_array($this->columnNames[$n], $boolAtribut) AND $updates[$n] == 0) {
          $updates[$n] = "false";
        } else if (in_array($this->columnNames[$n], $boolAtribut) AND $updates[$n] == 1) {
          $updates[$n] = "true";
        }
        //echo "<br>ID-Feld bzw. notStringAtriburt ohne '' eintragen für ".$this->columnNames[$n];
        //für leere Einträge null-Werte ohne '' anhängen
        if ($updates[$n] == "") {
          //echo "<br>NULL-Eintrag für ".$this->columnNames[$n];
          $updates[$n] = "null";
        }
        $updateString = $updateString.$this->columnNames[$n]."=".$updates[$n];
      } else {
        //date im leergelassenen Feld als NULL-Wert OHNE '' eintragen
        if (in_array($this->columnNames[$n], $dateAtribut) AND ($updates[$n] == "")) {
          //echo "<br>NULL-Eintrag für ".$this->columnNames[$n];
          $updates[$n] = "null";
          //echo "<br>String-Feld OHNE '' eintragen für ".$this->columnNames[$n];
          $updateString = $updateString.$this->columnNames[$n]."=".$updates[$n];
        } else {
          //echo "<br>String-Feld mit '' eintragen für ".$this->columnNames[$n];
          $updateString = $updateString.$this->columnNames[$n]."='".$updates[$n]."'";
        }
        
      }
      //echo "<br>";
    }
    //echo "<br>updateString:";
    //print($updateString);
    if ($this->tableName == "verbrauch") {  //hier Schlüsselkombination benötigt
      $idName0 = $this->columnNames[0];
      $idName1 = $this->columnNames[1];
      $whereString = $idName0 ."=". $updates[0] ." AND ". $idName1 ."=".$updates[1];
      $idKey = $updates[0] ."_". $updates[1];
    } else {  //einfacher Schlüssel
      $idName = $this->columnNames[0];
      $whereString = $idName ."=".$updates[0];
      $idKey = $updates[0];
    }
    
    $updateQueryString = "UPDATE ".$this->tableName." SET ".$updateString ." WHERE ".$whereString.";";
    //TODO: folgendes wieder auskommentieren:
    echo "<br>";
    print($updateQueryString);
    //echo "<br>";
    
    $queryUps = $con -> query($updateQueryString);
    //var_dump($queryUps);

    //UPDATE arbeitsdienste SET faelligkeit = '2022-05-23' WHERE ad_id=14;
    if ($queryUps == false) {
      echo "<br>Fehler in updateQueryString";
    } else {
      echo "<br>In Tabelle ".$this->tableName." wurde die Zeile mit der ID ".$idKey." erfolgreich mit folgenden Werten gespeichert:<br>".$updateString.".
      ";

      /*auffangen der Daten nicht nötig, da es kein Ergebnis im eigentlichen Sinne gibt
      foreach ($queryUps as $row) {var_dump($row);}*/
    }
  }

  //übergebene Insert-Zeile (via Array) in DB einfügen
  public function insertRowArr(array $insertArr) {
    $notStringAtribut = array('temperatur', 'luftfeuchtigkeit', 'helligkeit', 'waermeleistung', 'lager_min', 'lager_akt', 'kosten', 'wasserverbrauch', 'hoehe', 'breite', 'tiefe', 'laenge', 'gewicht', 'breite_vorn', 'menge_min', 'menge_max', 'mahlkategorie', 'anzahl');
    $boolAtribut = array('kuenstlich', 'spruehen', 'bepflanzt');
    $con = $this->connectionStart();
    //var_dump($insertArr);
    $insertColumns = "";
    $insertValues = "";
    foreach ($insertArr as $key => $value) {
      //Muss der Eintrag ein String oder ein Int/Date/...-Wert sein? ('' oder null eintragen?)
      if (strpos($key, "_id") !== false OR in_array($key, $notStringAtribut) OR in_array($key, $boolAtribut)) {
        //für bool-Atribute die Integer umwandeln
        if (in_array($key, $boolAtribut) AND $value == 0) {
          $value = "false";
        } else if (in_array($key, $boolAtribut) AND $value == 1) {
          $value = "true";
        }
        //für leere Einträge und bool-Atribute (null-)Werte ohne '' anhängen
        if ($insertColumns == "") {
          $insertColumns = $key;
          $insertValues = $value;
        } else {
          $insertColumns = $insertColumns.", ".$key;
          $insertValues = $insertValues.", ".$value;
        }
      } else {
        //leere Einträge MIT '' anhängen
        if ($insertColumns == "") {
          $insertColumns = $key;
          $insertValues = "'".$value."'";
        } else {
          $insertColumns = $insertColumns.", ".$key;
          $insertValues = $insertValues.", '".$value."'";
        }
      }
    }

    //INSERT INTO Lagerpflege (AD_ID, VM_ID) VALUES (71, 10);
    $insertQueryString = "INSERT INTO ".$this->tableName." (".$insertColumns.") VALUES (".$insertValues.");";
    print($insertQueryString);
    echo "<br>";

    $queryIns = $con -> query($insertQueryString);

    if ($queryIns == false) {
      echo "<br>Fehler in insertQueryString";
    } else {
      echo "<br>In Tabelle ".$this->tableName." wurde erfolgreich eine neue Zeile eingefügt. Zu den Spalten <br>".$insertColumns." wurden folgenden Werten gespeichert:<br>".$insertValues.".
      ";

      /*auffangen der Daten nicht nötig, da es kein Ergebnis im eigentlichen Sinne gibt
      foreach ($queryUps as $row) {var_dump($row);}*/
    }
  }

  //Erstellung & Anzeige eines Formulars mit den geladenen Spalten
  public function showFormular() {
    $this->loadColumns();         //jetzt sind die Spaltennamen (columnNames) gesetzt
    //HTML-Ausgabe:
    echo "<br><br>Eingabeformular für ".$this->tableName."<br><br>";
    /*echo "
      <form class='cssFormular' action='".$_SERVER['PHP_SELF']."' method='POST' name='".$this->tableName."'>
    ";*/
    for ($i = 0; $i < count($this->columnNames) ;$i++) {
      //Vorbereitung der Variablen für die HTML-Input-Felder
      $inputId = $this->tableName."_".$i;
      $inputName = $this->columnNames[$i];
      //Formular-Zeilen erstellen => <Spaltentitel>: <Input-Feld>
      echo "
        <div class='label'>
          ".$this->columnNames[$i]."".":
        </div>
      ";
      //für folgende Fälle Dropdownlisten erstellen
      if ((($this->tableName == 'terrarien') AND ($this->columnNames[$i] == 'ht_id')) OR 
      (($this->tableName == 'tiere') AND ($this->columnNames[$i] == 'tr_id')) OR 
      (($this->tableName == 'tiere') AND ($this->columnNames[$i] == 'ua_id')) OR 
      (($this->tableName == 'leuchtmittel') AND ($this->columnNames[$i] == 'vm_id')) OR 
      (($this->tableName == 'pflanzen') AND ($this->columnNames[$i] == 'vm_id')) OR 
      (($this->tableName == 'bodenbelag') AND ($this->columnNames[$i] == 'vm_id')) OR 
      (($this->tableName == 'deko') AND ($this->columnNames[$i] == 'vm_id')) OR 
      (($this->tableName == 'wasserbecken') AND ($this->columnNames[$i] == 'vm_id')) OR 
      (($this->tableName == 'arbeitsdienste') AND ($this->columnNames[$i] == 'at_id')) OR 
      (($this->tableName == 'verbrauch') AND ($this->columnNames[$i] == 'tr_id')) OR 
      (($this->tableName == 'verbrauch') AND ($this->columnNames[$i] == 'vm_id')) OR 
      (($this->tableName == 'verbrauch') AND ($this->columnNames[$i] == 'ad_id')) OR 
      (($this->tableName == 'lagerpflege') AND ($this->columnNames[$i] == 'vm_id')) OR 
      (($this->tableName == 'lagerpflege') AND ($this->columnNames[$i] == 'ad_id')) OR 
      (($this->tableName == 'tierpflege') AND ($this->columnNames[$i] == 'ad_id')) OR 
      (($this->tableName == 'tierpflege') AND ($this->columnNames[$i] == 'ti_id')) OR 
      (($this->tableName == 'tierpflege') AND ($this->columnNames[$i] == 'fu_id')) OR 
      (($this->tableName == 'terrariumpflege') AND ($this->columnNames[$i] == 'ad_id')) OR 
      (($this->tableName == 'terrariumpflege') AND ($this->columnNames[$i] == 'tr_id'))) {

        if ($this->columnNames[$i] == 'ht_id') {
        //if ($this->tableName == 'terrarien') {
          $helpTable = 'habitattypen';
        } else if ($this->columnNames[$i] == 'tr_id') {
          $helpTable = 'terrarien';
        } else if ($this->columnNames[$i] == 'ua_id') {
          $helpTable = 'unterarten';
        } else if ($this->columnNames[$i] == 'vm_id') {
          $helpTable = 'verbrauchsmittel';
        } else if ($this->columnNames[$i] == 'at_id') {
          $helpTable = 'arbeitsdiensttypen';
        } else if ($this->columnNames[$i] == 'ad_id') {
          $helpTable = 'arbeitsdienste';
        } else if ($this->columnNames[$i] == 'ti_id') {
          $helpTable = 'tiere';
        } else if ($this->columnNames[$i] == 'fu_id') {
          $helpTable = 'futter';
        }
        $helpFormular = new GlobalFormular($helpTable);
        //var_dump($helpFormular->tableName);
        //$helpFormular->inserts = array();
        $helpFormular->loadFirstColumns();  //ID und Bezeichnung geladen für Dropdown
        //var_dump($helpFormular->inserts);
        echo "
        <select name='".$inputName."' id='".$inputId."'>
        ";
        foreach ($helpFormular->inserts as $key => $value) {
          echo "
          <option value='".$key."' >".$key.": ".$value."</option>
          ";
        }
        echo "
        </select><br>
          ";
      } else {
        echo "
        <input type='text' onblur='Pruefen(\"$inputName\",\"$this->tableName\")' id='".$inputId."' name='".$inputName."'>
        ";
        if ($i == 0) {
          echo "* automatische Vergabe der ID bei leerer Eingabe";
        }
        echo "<br>";
      }
    } //onclick='jsTest()'

    //Prüf- & Absendebutton erstellen
    /*echo "<input type='submit' name='sendNewRow' value='Absenden'>";
    echo "<button type='button' onclick='jsPruef(\"$this->tableName\")'>Prüfen</button>";
    echo "</form>";*/
  }

  //Erstellung & Anzeige eines Formulars mit den geladenen Spalten UND eingetragenen Werten einer Zeile der DB
  public function showChangeRow(int $id) {
    //echo $id;
    $this->loadRow($id);         //jetzt sind die Spaltennamen (columnNames) gesetzt
    //HTML-Ausgabe:
    echo "Bearbeitungsformular für Tabelle ".$this->tableName.", Zeile ".$id." in der Tierwelt-DB<br><br>";
    echo "
      <form class='cssFormular' action='".$_SERVER['PHP_SELF']."' method='POST' name='".$this->tableName."'>
    ";
    //Tabellennamen per POST übertragen
    echo "<input type='hidden' name='table' value='".$this->tableName."'>";
    //ID per POST übertragen
    echo "<input type='hidden' name='id' value='".$id."'>";

    //var_dump($this->columnNames);

    for ($i = 0; $i < count($this->columnNames) ;$i++) {
      //Vorbereitung der Variablen für die HTML-Input-Felder
      $inputId = $this->tableName."_".$i;
      $inputName = $this->columnNames[$i];
      //Formular-Zeilen erstellen => <Spaltentitel>: <Input-Feld>
      echo "
        <div class='label'>
          ".$this->columnNames[$i]."".":
        </div>
        <input type='text' id='".$inputId."' name='".$inputName."' 
          value='".$this->inserts[0][$this->columnNames[$i]]."'>
        <br>
      ";
    }

    //Absendebutton erstellen
    echo "
        <input type='submit' name='sendUpdate' value='Absenden'>
      </form>
    ";
  }

  //Erstellung & Anzeige eines Formulars mit den geladenen Spalten UND eingetragenen Werten einer Zeile der DB, diese Zeile ist über eine Schlüsselkombination, welche über ein Array mitgegeben wird, identifizierbar
  public function showChangeRowArr(array $id) {
    //echo $id;
    $this->loadRowArr($id);         //jetzt sind die Spaltennamen (columnNames) gesetzt
    //HTML-Ausgabe:
    $idKombi = $id[0];
    for ($n=1; $n < count($id); $n++) { 
      $idKombi = $idKombi."_".$id[$n];
    }
    echo "Bearbeitungsformular für Tabelle ".$this->tableName." mit Schlüsselkombination ".$idKombi." in der Tierwelt-DB<br><br>";
    echo "
      <form class='cssFormular' action='".$_SERVER['PHP_SELF']."' method='POST' name='".$this->tableName."'>
    ";
    //Tabellennamen per POST übertragen
    echo "<input type='hidden' name='table' value='".$this->tableName."'>";
    //ID per POST übertragen
    echo "<input type='hidden' name='id' value='".$idKombi."'>";

    for ($i = 0; $i < count($this->columnNames) ;$i++) {
      //Vorbereitung der Variablen für die HTML-Input-Felder
      $inputId = $this->tableName."_".$i;
      $inputName = $this->columnNames[$i];
      //Formular-Zeilen erstellen => <Spaltentitel>: <Input-Feld>
      echo "
        <div class='label'>
          ".$this->columnNames[$i]."".":
        </div>
        <input type='text' id='".$inputId."' name='".$inputName."' 
          value='".$this->inserts[0][$this->columnNames[$i]]."'>
        <br>
      ";
    }

    //Absendebutton erstellen
    echo "
        <input type='submit' name='sendUpdate' value='Absenden'>
      </form>
    ";
  }
  
}
?>