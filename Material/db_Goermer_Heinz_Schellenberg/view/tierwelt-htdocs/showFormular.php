<?php
  include('classFormular.php');
  $animalworldTables = array( "Unterarten", "Habitattypen", "Terrarien", "Tiere", "Verbrauchsmittel", 
                                "Leuchtmittel", "Pflanzen", "Bodenbelag", "Deko", "Wasserbecken", 
                                "Arbeitsdiensttypen", "Arbeitsdienste", "Verbrauch", "Futter", "Lagerpflege", "Tierpflege", "Terrariumpflege"
                                );
  //TODO "Verbrauch", "Lagerpflege", "Tierpflege" & "Terrariumpflege" müssen anders behandelt werden, da hier nicht nur in Arbeitsdiensten geguckt werden muss, ob die neu eingetragenen Werte nicht bereits existiren
  //var_dump($animalworldTables);
  echo "
    <html>
      <head>
        <link rel='stylesheet' href='db-animalWorld.css' type='text/css'>
        <script src='db-animalWorld.js'></script>
        <title>DB Tierwelt (EINGABE)</title>
      </head>
      <body>
        <form action='".$_SERVER['PHP_SELF']."' method='POST'>
          <label for='tabels'>Wähle eine Tabelle für das Eingabe-Formular:</label>
            <select name='tableDrop' id='tableDrop'>
  ";
  for ($table = 0; $table < count($animalworldTables); $table++) {
    //Dropdown wieder auf zuvor gesetzten Wert setzen:
    if ($_POST != null AND $_POST['tableDrop']==strtolower($animalworldTables[$table])) { 
      echo "
              <option value='".strtolower($animalworldTables[$table])."' selected>".$animalworldTables[$table]."</option>
    ";
    } else {  //keine Vorauswahl im Dropdown
      echo "
              <option value='".strtolower($animalworldTables[$table])."'>".$animalworldTables[$table]."</option>
    ";
    }
  }
  echo "
            </select>
          <br><br>
  ";
  if (isset($_POST['insertSubmit'])) {
    if ($_POST['insertSubmit'] == 'table') {
      echo "
            <button type='submit' name='insertSubmit' value='insert'>Eingabe senden</button>
    ";
    } else if ($_POST['insertSubmit'] == 'insert') {
      echo "
              <button type='submit' name='insertSubmit' value='table'>Tabelle wählen</button>
      ";
    }
  }
  if ($_POST == null) { //als OR in der obigen Abfrage gibt es eine Fehlermeldung
    echo "
            <button type='submit' name='insertSubmit' value='table'>Tabelle wählen</button>
    ";
  }
  if ($_POST != null) {
    $tableName = filter_input(INPUT_POST, 'tableDrop');
    $formular = new formular($tableName);
    if ($_POST['insertSubmit'] == 'table') {
      //wird über insertSubmit 'table' übertragen -> Eingabeformular aufbauen
      if (isset($_POST['tableDrop'])) {
        //$tableName = 'tiere';
        $formular->showFormular();    //$formular->loadTable(); -> steckt hier drin
      }
    }
    if ($_POST['insertSubmit'] == 'insert') {
      //wird über insertSubmit 'insert' übertragen -> Bestätigung oder Fehler des Inserts ausgeben, kein Eingabeformular aufbauen
      echo "<br><br>";
      $insertArr = array();
      foreach($_POST as $key => $value) {
        if (!(($key == 'tableDrop') OR ($key == 'insertSubmit'))) {
          $insertArr[$key] = $value;
          echo $key." = ".$value."<br>";
        } else {
          //echo $key." = ".$value."<br>";
        }
      }
      $formular->insertRowArr($insertArr);
    }
  }
  echo "
        </form>
  ";

  
  //$tiereFormular = new formular('tierpflege');
  //$tiereFormular->showFormular();    //$tiereFormular->loadTable(); -> steckt hier drin

echo "
      <div id='jsAusgabe'></div>
";

echo "
    </body>
  </html>
";
?>