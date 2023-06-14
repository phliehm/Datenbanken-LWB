<?php
  class tasks {       //Klasse Arbeitsdienste (tasks; services; chores = Hausarbeiten)
    private $taskName;
    private $tasks = [];   //Sammlung der Arbeitsdienste des entsprechenden taskName und Datumsbereichs in Array
    //private $taskColumns = [];

    public function __construct($setTaskName)
    {
      $this->taskName = $setTaskName;    
    }

    private function connectionStart() {
      //Verbindung zur DB
      $dbname = 'lewein';
      $dsn = 'pgsql:host=localhost;dbname='.$dbname;
      $password = 'niewel';
      $con = new PDO ($dsn,$dbname,$password);
      return  $con;
    }

    // lade die Fälligen, d.h. diejenigen in Arbeitsdiensten die in den angegebenen Datumsbereich fallen (mit übergebenem string = week / day / over / future)
    public function loadDue($due) {
      //Datumsbereich in WHERE-String umarbeiten
      if ($due == "week") {
        $whereString = "WHERE DATE_PART('week',faelligkeit) = DATE_PART('week', CURRENT_DATE)";
      } else if ($due == "day") {
        $whereString = "WHERE DATE_PART('day',Faelligkeit) = DATE_PART('day', CURRENT_DATE)
        AND DATE_PART('week',Faelligkeit) = DATE_PART('week', CURRENT_DATE)";
      } else if ($due == "over") {
        $whereString = "WHERE DATE_PART('day',Faelligkeit) < DATE_PART('day', CURRENT_DATE) 
        AND DATE_PART('week',Faelligkeit) = DATE_PART('week', CURRENT_DATE) 
        OR DATE_PART('week',Faelligkeit) < DATE_PART('week', CURRENT_DATE)";
      } else if ($due == "future") {
        $whereString = "WHERE DATE_PART('day',Faelligkeit) > DATE_PART('day', CURRENT_DATE) 
        AND DATE_PART('week',Faelligkeit) = DATE_PART('week', CURRENT_DATE) 
        OR DATE_PART('week',Faelligkeit) > DATE_PART('week', CURRENT_DATE)";
      }
      $con = $this->connectionStart();
      if ($this->taskName == "lagerpflege") {
        $queryTasks = $con->query("SELECT ad_id, vm_id, at_id, arbeitsdienst, verbrauchsmittel, faelligkeit
                                  FROM lagerpflege AS LP
                                  NATURAL JOIN
                                  (SELECT VM_ID, VM.Bezeichnung AS verbrauchsmittel
                                  FROM verbrauchsmittel VM) AS a
                                  NATURAL JOIN
                                  (SELECT ad_id, at_id, faelligkeit
                                  FROM arbeitsdienste AS AD) AS b
                                  NATURAL JOIN
                                  (SELECT AT_ID, A.Bezeichnung AS arbeitsdienst
                                  FROM arbeitsdiensttypen A) AS c
                                  $whereString;
                                  ")->fetchAll(PDO::FETCH_NAMED);   // PostgreSQL uses the yyyy-mm-dd
      } else if ($this->taskName == "tierpflege") {
        $queryTasks = $con->query("SELECT  ad_ID, tiere.name AS tier, terrarien.name AS terrarium, 
                                  arbeitsdiensttypen.bezeichnung AS arbeitsdienst,
                                  arbeitsdiensttypen.beschreibung AS dienstbeschreibung,
                                  futter.bezeichnung AS futter, tierpflege.mahlkategorie,
                                  tierpflege.menge_min, tierpflege.menge_max,
                                  arbeitsdienste.faelligkeit
                                  FROM tierpflege
                                  JOIN tiere USING (ti_id)
                                  JOIN futter USING (fu_ID)
                                  JOIN terrarien USING (tr_ID)
                                  JOIN arbeitsdienste USING (ad_ID)
                                  JOIN arbeitsdiensttypen USING (at_ID)
                                  $whereString;
                                  ")->fetchAll(PDO::FETCH_NAMED);
      } else if ($this->taskName == "terrariumpflege") {
        $queryTasks = $con->query("SELECT  terrariumpflege.ad_id, arbeitsdiensttypen.bezeichnung AS arbeitsdienst,
                                  tr_id, terrarien.name AS terrarium,
                                  arbeitsdienste.faelligkeit
                                  FROM terrariumpflege
                                  JOIN terrarien USING (tr_ID)
                                  JOIN arbeitsdienste USING (ad_ID)
                                  JOIN arbeitsdiensttypen USING (at_ID)
                                  $whereString;
                                  ")->fetchAll(PDO::FETCH_NAMED);
      }
      $i = 0;
      $this->tasks[] = array();   //eventuell vorhandenen Werte im tasks-Array löschen
      foreach ($queryTasks as $row) {
        //Daten als Array (Ergebnis von FetchAll) in tasks-Objekt gespeichert (OHNE 0.Spalte = ID)
        $this->tasks[$i] = $row; 
        $i++;
        //var_dump($row);
      }
      echo "<br><br>Fällige aus ".$this->taskName." (Datumsbereich: ".$due."):<br>";
      //var_dump($this->tasks);
    }

    //bei übergbener AD_ID wird die dazugehörige AT_ID geholt 
    private function loadAT_ID($ad_id){
      $con = $this->connectionStart();
      //at_id	aus arbeitsdienste holen mit ad_id
      $queryTasks = $con->query("SELECT at_id
                                  FROM arbeitsdienste
                                  WHERE ad_id=$ad_id;
                                  ")->fetchAll(PDO::FETCH_NAMED);
      foreach ($queryTasks as $row) {
        //var_dump($row);
        $at_id = $row;
      }
      //var_dump($at_id);
      //echo "<br>at_id lautet: ".$at_id['at_id'];
      return $at_id['at_id'];
    }

    //bei übergbener AT_ID wird die dazugehörige Frequenz geholt 
    private function loadFrequenz($at_id){
      $con = $this->connectionStart();
      //frequenz aus arbeitsdiensttypen holen mit at_id
      $queryTasks = $con->query("SELECT frequenz
                                  FROM arbeitsdiensttypen
                                  WHERE at_id=$at_id;
                                  ")->fetchAll(PDO::FETCH_NAMED);
      foreach ($queryTasks as $row) {
        //var_dump($row);
        $frequenz = $row;
      }
      //var_dump($at_id);
      //echo "<br>Frequenz lautet: ".$frequenz['frequenz'];
      return $frequenz['frequenz'];
    }

    //bei übergbener AD_ID wird das dazugehörige Datum geholt 
    private function loadDate($ad_id){
      $con = $this->connectionStart();
      //frequenz aus arbeitsdiensttypen holen mit at_id
      $queryTasks = $con->query("SELECT faelligkeit
                                  FROM arbeitsdienste
                                  WHERE ad_id=$ad_id;
                                  ")->fetchAll(PDO::FETCH_NAMED);
      foreach ($queryTasks as $row) {
        //var_dump($row);
        $date = $row;
      }
      //var_dump($at_id);
      //echo "<br>Datum lautet: ".$date['faelligkeit'];
      return $date['faelligkeit'];
    }

    //die für jede Ergebniszeile in Arbietsdiensten ($ad_id) gesetzten Häkchen (erledigt / verschieben) in neues Datum "übersetzen" und in DB abspeichern -> $duration (Dauer): Erhöhung des Datums entsprechend der Frequenz oder nur um einen Tag
    public function updateTask($ad_id, $duration) {
      //at_id	aus arbeitsdienste holen mit ad_id
      $at_id = $this->loadAT_ID($ad_id);
      //frequenz aus arbeitsdiensttypen holen mit at_id
      $frequenz = $this->loadFrequenz($at_id);
      //vom heutigen Datum wird hochgezählt -> aktuell eingetragenes Datum unwichtig
      //neues Datum berechnen:
      date_default_timezone_set("Europe/Berlin");
      $timestamp = time();
      $currentDate = date("Y-m-d",$timestamp);
      //echo "<br>".$currentDate;

      //je nachdem, ob der Arbeitsdienst als erledigt oder verschoben markiert wird, wird entweder die eingetragene Frequenz oder nur ein einzelner Tag dazu addiert
      //NEXT = DONE: eingetragene Frequenz addieren bei Häkchen = 'done'
      if ($duration == "next") {
        //bei Frequenzen kleiner/gleich 7 ist der ausführende Wochentag wichtig -> wurde durch vorheriges verschieben die Fälligkeit auf einen falschen Wochentag gesetzt, so muss dies beim nächsten Aufruf von "erledigt" für diesen Arbeitsdienst korrigiert werden -> zurücksetzten auf letzten tonusmäßigen Wochentag für diesen Arbeitsdienst und dann erst mit Frequenz erhöhen erfolgt hier durch Reduktion der Variable Frequenz
        //gespeichertes Fälligkeitsdatum hier holen nicht nötig -> Vergleich muss mit $currentDate geschehen
        switch ($frequenz) {
          case 1:
            //Mo verschoben auf Di -> Di erledigt -> verschoben auf Mi -> keine Korrektur nötig
            break;    //damit die anderen cases nicht mehr geprüft werden müssen
          case 2:
            //liegt aktuell gespeichertes Datum auf Mo, Mi, Fr -> alles OK
            //Mo verschoben auf Di -> Di erledigt -> verschieben auf Do falsch, sondern auf Mi
            //Mi verschoben auf Do -> Do erledigt -> verschieben auf Sa->Mo falsch, sondern auf Fr
            //Do von Fr vorgezogen oder weil Mi schon auf Di vorgezogen wurde, am Do aufgeführt -> trotzdem Fr nochmal, denn Do bis Mo zu lang bei Intervall 2
            //Di von Mi vorgezogen -> ohne Korrektur -> Do aufgeführt, danach anderer Fall
            if (date("w", strtotime($currentDate)) == 2 OR date("w", strtotime($currentDate)) == 4 ) { 
              //1:Mo, 2:Di, 3:Mi, 4:Do, 5:Fr
              $frequenz = $frequenz - 1;
            }
            //echo "Frequenzkorrektur vor Speicherung der Fälligkeit vorgenommen (für Frequenz: ".$frequenz.")";
            break;
          case 4:
            //liegt aktuelles Datum auf Mo -> Mo Haken erledigt gesetzt -> Verschiebung auf Fr -> alles OK
            //Mo verschoben auf Di -> Di erledigt -> verschieben auf Sa->Mo falsch, sondern auf Fr (+3)
            //Di verschoben auf Mi -> Mi erledigt -> verschieben auf So->Mo falsch, sondern auf Fr (+2)
            //liegt aktuelles Datum auf Do -> Haken erledigt -> Verschiebung auf Mo -> alles OK (4 Tage-Intervall passt)
            //liegt aktuelles Datum auf Fr -> Fr Haken erledigt gesetzt -> Verschiebung auf Di (+4) falsch, sondern auf Mo (+3)
            if ((date("w", strtotime($currentDate)) == 2) OR (date("w", strtotime($currentDate)) == 5)) {  //Di, Fr
              $frequenz = $frequenz - 1;
            } else if (date("w", strtotime($currentDate)) == 3) { //Mi
              $frequenz = $frequenz - 2;
            } 
            //echo "Frequenzkorrektur für Frequenz: ".$frequenz." vor Speicherung der Fälligkeit vorgenommen.";
            break;
          case 7:
            //liegt aktuelles Datum auf Mo (1) -> Mo Haken erledigt gesetzt -> Verschiebung auf Mo -> +9 nötig
            //aktuelles Datum auf Di (2) -> Di erledigt -> Verschiebung auf Di (+7) -> +8 nötig, dann wieder Mi
            //liegt aktuelles Datum auf Mi (3) -> Mi Haken erledigt gesetzt -> Verschiebung auf Mi -> alles OK
            //aktuelles Datum auf Do (4) -> Do erledigt -> Verschiebung auf Do (+7) -> +6 nötig, dann wieder Mi
            //aktuelles Datum auf Fr (5) -> Fr erledigt -> Verschiebung auf Fr (+7) -> +5 nötig, dann wieder Mi
            if ((date("w", strtotime($currentDate)) == 1)) {  //Mo
              $frequenz = $frequenz + 2;
            } else if (date("w", strtotime($currentDate)) == 2) { //Di
              $frequenz = $frequenz + 1;
            } else if (date("w", strtotime($currentDate)) == 4) { //Do
              $frequenz = $frequenz - 1;
            } else if (date("w", strtotime($currentDate)) == 2) { //Fr
              $frequenz = $frequenz - 2;
            }
            //echo "Frequenzkorrektur für Frequenz: ".$frequenz." vor Speicherung der Fälligkeit vorgenommen.";
        }
        $newDate = new DateTime($currentDate);
        //$newDate->add(new DateInterval('P1D')); // P1D means a period of 1 day
        $period = "P".$frequenz."D";
        $newDate->add(new DateInterval($period));
        $fomattedDate = $newDate->format('Y-m-d');

        //ONE = SHIFT: einzelnen Tag addieren bei Häkchen = 'shift'
      } else if ($duration == "one") {
        //bei größeren Frequenzen soll ein Verschieben nicht dazu führen, dass etwas was in 1 Monat erst erledigt werden müsste, schon morgen auf der Liste steht -> für fehlerhafte Usereingaben (Häkchen setzen) VOR dem eigentlichen Datum -> ergo prüfen, ob neues Datum vor dem fälligen in DB liegt
        $dueDate = $this->loadDate($ad_id);   //gespeichertes Fälligkeitsdatum holen
        if ($dueDate > $currentDate) {
          //eingetragenes Datum liegt NACH heutigem Datum -> Fälligkeitsdatum ist Referenztag für ALLE Frequenzen von 1, 2, 4, 7, 30, 61, 91, 182, 365 Tagen
          $dueTimestamp = strtotime($dueDate);
          $newTimestamp = strtotime('+1 day', $dueTimestamp);
          $fomattedDate = date("Y-m-d",$newTimestamp);
        } else if ($dueDate == $currentDate) {
          //eingetragenes Datum gleich heute -> heute ist Referenztag -> Verschiebung auf morgen
          $newTimestamp = time() + (60*60*24);   //60*60*24 = 1 Tag
          //negative Berechnung, also Tage abziehen von einem Datum:
          //$newTimestamp = time() - (60*60*24) * $frequenz;   //60*60*24 = 1 Tag
          $fomattedDate = date("Y-m-d",$newTimestamp);
        } else if ($dueDate < $currentDate) {
          //eingetragenes Datum liegt VOR heute (also in Vergangenheit) -> heute ist Referenztag -> Verschiebung AUF heute
          $newTimestamp = time();   //ohne weitere Korrektur auf heute verschieben
          $fomattedDate = date("Y-m-d",$newTimestamp);
        }
      }
      //prüfen, ob neues Datum auf Wochentag fällt, sonst auf nächsten Mo setzen
      if (date("w", strtotime($fomattedDate)) == 0) { //Sonntag + 1 Tag
        $newDate = new DateTime($fomattedDate);
        $newDate->add(new DateInterval('P1D')); // P1D means a period of 1 day
        $fomattedDate = $newDate->format('Y-m-d');
      } else if (date("w", strtotime($fomattedDate)) == 6) {  //Samstag + 2 Tage
        $newDate = new DateTime($fomattedDate);
        $newDate->add(new DateInterval('P2D')); // P1D means a period of 2 days
        $fomattedDate = $newDate->format('Y-m-d');
      }
      
      //echo "<br>".$fomattedDate;
      
      //UPDATE Datum: Fälligkeit in arbeitsdienste bei ad_id mit neuem $fomattedDate setzen
      $con = $this->connectionStart();
      $updateString = "faelligkeit='".$fomattedDate."'";
      $whereString = "ad_id=".$ad_id;
      $updateQueryString = "UPDATE arbeitsdienste SET ".$updateString ." WHERE ".$whereString.";";
      //UPDATE arbeitsdienste SET faelligkeit = '2022-05-23' WHERE ad_id=14;
      //print($updateQueryString); echo "<br>";
    
      $queryUps = $con -> query($updateQueryString);
      //var_dump($queryUps);

      if ($queryUps == false) {
        echo "<br>Fehler in updateTask->updateQueryString";
      } else {
        echo "<br>In Tabelle arbeitsdienste wurde die Zeile mit der ID ".$ad_id." erfolgreich mit folgenden Werten gespeichert: ".$updateString.".
        ";

        /*auffangen der Daten nicht nötig, da es kein Ergebnis im eigentlichen Sinne gibt
        foreach ($queryUps as $row) {var_dump($row);}*/
      }
      
    }

    //Erstellung & Anzeige der Tabelle mit den geladenen Spaltentiteln & Tabelleninhalten
    public function showOverview($due) {
      //Spaltennamen (taskColumns) setzten
      /*if ($this->taskName == "lagerpflege") {
        $this->taskColumns = array("ad_id", "vm_id", "at_id", "arbeitsdienst", "verbrauchsmittel", "faelligkeit");
      } else if ($this->taskName == "ierpflege") {
        $this->taskColumns = array("menge_min", "menge_max", "mahlkategorie",
                                  "tiername", "futter",
                                  "terrariumname", "faelligkeit", "arbeitsdienst");
      }*/
      //Tabelleninhalt/-zeilen (tasks) setzten (mit übergebenem string = week / day / over / future)
      $this->loadDue($due);
      //todo: mahlkategorie erstzen in Wortlaut - 1=Hauptmahl, 2=alternatives Hauptmahl, 3=Nebenmahl

      //var_dump($this->tasks);
      //HTML-Ausgabe:
      if (count($this->tasks[0]) > 0) {
        //echo "Inhalt des Arbeitsdienstbereiches ".$this->taskName."<br><br>";
        echo "<br>";
        echo "
        <table>
          <thead>
            <tr>
              <th>
        ";
        echo implode('</th><th>', array_keys(current($this->tasks)));
        echo "
              </th>
            </tr>
          </thead>
          <tbody>
        ";
        //var_dump($this->tasks);
        foreach ($this->tasks as $key => $value) {
          array_map('htmlentities', $value);
          //var_dump($this->tableRows[$key]); echo $this->columnNames[0];
          /*if ($this->tableName == "verbrauch") {
            $id = $this->tableRows[$key][$this->columnNames[0]]."_".$this->tableRows[$key][$this->columnNames[1]];
          } else {
            $id = $this->tableRows[$key][$this->columnNames[0]];
          }*/
          //echo $key; echo "key-test<br><br>";
          echo "
            <tr>
              <td>
          ";
          echo implode('</td><td>', $value);
          echo "
              </td>
              <td>
                <div>
                  <input type='radio' id='done' value='done' name='".$key."_".$this->tasks[$key]["ad_id"]."'>
                  <label for='done'>erledigt</label>
                  <input type='radio' id='shift' value='shift' name='".$key."_".$this->tasks[$key]["ad_id"]."'>
                  <label for='shift'>verschieben</label>
                </div>
              </td>
            </tr>
          ";  //key ist die Zeilennummer der Ergebnisliste, $this->tasks[$key]["ad_id"] ergibt die ad_id
        }
        echo "
          </tbody>
        </table>
        ";
      }
    }
  }
?>