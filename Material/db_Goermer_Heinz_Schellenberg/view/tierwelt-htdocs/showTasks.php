<?php
  include('classTasks.php');
  $taskNames = array("Lagerpflege", "Tierpflege", "Terrariumpflege");
  $dateRange = array("week", "day", "over", "future");
  $datumsbereich = array("fällig in der laufenden Woche", "heute fällig", "überfällig", "zukünftige");
  echo "
    <html>
      <head>
        <link rel='stylesheet' href='db-animalWorld.css' type='text/css'>
        <script src='db-animalWorld.js'></script>
        <title>DB Tierwelt (Pflege-DIENSTE)</title>
      </head>
      <body>
        <form action='".$_SERVER['PHP_SELF']."' method='POST'>
            <label for='tabels'>Wähle einen Arbeitsdienstbereich:</label>
              <select name='taskDrop' id='taskDrop'>
  ";
  //Dropdown für Task-/Dienst-Auswahl
  for ($task = 0; $task < count($taskNames); $task++) {
    //Dropdown wieder auf zuvor gesetzten Wert setzen:
    if ($_POST != null AND $_POST['taskDrop']==strtolower($taskNames[$task])) { 
      echo "
                  <option value='".strtolower($taskNames[$task])."' selected>".$taskNames[$task]."</option>
      ";
    } else {  //keine Vorauswahl im Dropdown
      echo "
                  <option value='".strtolower($taskNames[$task])."'>".$taskNames[$task]."</option>
      ";
    }
  }
  echo "
              </select>
            <br><br>
            <label for='tabels'>Wähle einen Datumsbereich:</label>
              <select name='dateDrop' id='dateDrop'>
  ";
  //Dropdown für Due-/Fälligkeits-Auswahl
  for ($due = 0; $due < count($dateRange); $due++) {
    //Dropdown wieder auf zuvor gesetzten Wert setzen:
    if ($_POST != null AND $_POST['dateDrop']==strtolower($dateRange[$due])) { 
      echo "
                <option value='".strtolower($dateRange[$due])."' selected>".$datumsbereich[$due]."</option>
    ";
    } else {  //keine Vorauswahl im Dropdown
      echo "
                <option value='".strtolower($dateRange[$due])."'>".$datumsbereich[$due]."</option>
    ";
    }
  }
  echo "
              </select>
            <br><br>
  ";
  if (isset($_POST['taskSubmit'])) {
    if ($_POST['taskSubmit'] == 'tasks') {
      echo "
            <button type='submit' name='taskSubmit' value='taskNames.dateRange'>Dienst wählen</button>
    ";
    } else if ($_POST['taskSubmit'] == 'taskNames.dateRange') {
      echo "
              <button type='submit' name='taskSubmit' value='tasks'>Häkchen senden</button>
      ";
    }
  }
  if ($_POST == null) { //als OR in der obigen Abfrage gibt es eine Fehlermeldung
    echo "
            <button type='submit' name='taskSubmit' value='taskNames.dateRange'>Dienst wählen</button>
    ";
  } 
  if ($_POST != null) {
    if ($_POST['taskSubmit'] == 'taskNames.dateRange') {
      //alle POSTs:
      /*foreach($_POST as $key => $value) {
        echo $key." = ".$value."<br>";
      }*/
      $taskName = filter_input(INPUT_POST, 'taskDrop');
      //$tableName = 'tiere';
      $due = filter_input(INPUT_POST, 'dateDrop');
      //echo "Ausgewählt wurde der Arbeitsdienstbereich: ".$taskName." und der Datumsbereich: ".$due;
      $tasks = new tasks($taskName);
      $tasks->showOverview($due);    //$sicht->loadColumns() & loadTable() -> stecken hier drin
    } else if ($_POST['taskSubmit'] == 'tasks') {
      $taskName = filter_input(INPUT_POST, 'taskSubmit');
      $tasks = new tasks($taskName);
      echo "<br>";
      //alle POSTs / Schleife über alle POSTs (auch zu den Radiobuttons)
      foreach($_POST as $key => $value) {
        //$_POST[$i] auseinander nehmen und in Zeilennr. sowie ad_id auftrennen
        $keyId = explode('_', $key);
        //var_dump($keyId);
        if (count($keyId)>1) {  //ersten 3 POSTs mit name 'taskDrop', 'dateDrop', 'taskSubmit' werden übersprungen, da hier der Array genau 1 Element hat
          echo "<br><br>Der ".$_POST['taskDrop']."-Arbeitsdienst in Zeile ".$keyId[0]." (mit 'ad_id' = ".$keyId[1].") wurde mit '".$_POST[$key]."' markiert.";
          
          if ($_POST[$key] == 'done') { //bei done auf das nächste Fälligkeitsdatum vorsetzen
            //UPDATE Datum: Fälligkeit in arbeitsdienste bei ad_id mit aktuellem Datum ERHÖHT um Frequenz setzen & prüfen, ob neues Datum auf Wochentag fällt, sonst auf nächsten Mo setzen
            //UPDATE arbeitsdienste bei ad_id -> Fälligkeit updaten
            $tasks->updateTask($keyId[1], "next");
          } else if ($_POST[$key] == 'shift') { //bei shift Fälligkeitsdatum unabhängig von der Frequenz um einen Tag erhöhen
            //so ist es auch heute und aktuelle Woche 
            //UPDATE Datum: Fälligkeit in arbeitsdienste bei ad_id mit aktuellem Datum ERHÖHT um einen Tag setzen & prüfen, ob neues Datum auf Wochentag fällt, sonst auf nächsten Mo setzen
            $tasks->updateTask($keyId[1], "one");
          }
        }
        //echo $key." = ".$value."<br>";    //$key = $keyId[0]."_".$keyId[1]
      }
    }
  }
  echo "
          </form>
  ";
  echo "
          <div>
  ";
  //$tasksDueWeek = new tasks();
  //$tasksDueWeek->loadDueWeek();    //$tasksDueWeek->loadTable(); -> steckt hier drin
  echo "
        </div>
        <div>
  ";
  //$tasksDueToday = new tasks();
  //$tasksDueToday->loadDueToday();    //$tasksDueToday->loadTable(); -> steckt hier drin
  echo "
        </div>
        <div>
  ";
  //$tasksOverdue = new tasks();
  //$tasksOverdue->loadOverdue();    //$tasksOverdue->loadTable(); -> steckt hier drin
  echo "
        </div>
        <div>
  ";
  //$tasksFuture = new tasks();
  //$tasksFuture->loadFuture();    //$tasksFuture->loadTable(); -> steckt hier drin
  echo "
        </div>
        <div>
  ";
  //$tasksDueTodayTerrarium = new tasks();
  //$tasksDueTodayTerrarium->loadDueTodayTerrarium();    //$tasksFuture->loadTable(); -> steckt hier drin
  echo "
              </div>
              <div>
  ";
  //$tasksDueTodayAnimals = new tasks();
  //$tasksDueTodayAnimals->loadDueTodayAnimals();    //$tasksFuture->loadTable(); -> steckt hier drin
  echo "
              </div>
        <div id='jsAusgabe'></div>
      </body>
    </html>
  ";
?>