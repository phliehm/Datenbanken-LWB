<?php
  include('classOverview.php');
  $animalworldTables = array( "Unterarten",         "Habitattypen",   "Terrarien",    "Tiere",      "Verbrauchsmittel", 
                              "Leuchtmittel",       "Pflanzen",       "Bodenbelag",   "Deko",       "Wasserbecken", 
                              "Arbeitsdiensttypen", "Arbeitsdienste", "Lagerpflege",  "Verbrauch",  "Futter", 
                              "Tierpflege",         "Terrariumpflege");
  //var_dump($animalworldTables);
  echo "
    <html>
      <head>
        <link rel='stylesheet' href='db-animalWorld.css' type='text/css'>
        <script src='db-animalWorld.js'></script>
        <title>DB Tierwelt (ALLE Tabellen)</title>
      </head>
      <body>
        <form action='".$_SERVER['PHP_SELF']."' method='POST'>
          <label for='tabels'>Wähle eine Tabelle:</label>
            <select name='tableDrop' id='tableDrop'>
  ";
  for ($table = 0; $table < count($animalworldTables); $table++) {
    echo "
              <option value='".strtolower($animalworldTables[$table])."'>".$animalworldTables[$table]."</option>
    ";
  }
  echo "
            </select>
          <br><br>
          <button type='submit'>Sende</button>
        </form>
  ";

  if ($_POST != null) {
    $tableName = filter_input(INPUT_POST, 'tableDrop');
    //$tableName = 'tiere';
    $overview = new overview($tableName);
    $overview->showOverview();    //$sicht->loadColumns() & loadTable() -> stecken hier drin
  }


  echo "
        <div id='jsAusgabe'></div>
  ";

  echo "
      </body>
    </html>
  ";

?>