<?php
  include('classFormular.php');
  echo "
    <html>
      <head>
        <link rel='stylesheet' href='db-animalWorld.css' type='text/css'>
        <script src='db-animalWorld.js'></script>
        <title>DB Tierwelt (ÄNDERUNG einer Tabellenzeile)</title>
      </head>
      <body>
  ";

  if (!empty($_GET['table'])) {
    //erster Aufruf der Seite lädt ein Formular mit eingetragenen Werten einer Zeile (Tabellenname und id kommen per GET) -> Aufruf erfolgt durch senden von Variablen mit einem Link -> Link in showOverview
    //echo $_GET['table'] . $_GET['id'];
    $tableName = $_GET['table'];
    $changeFormular = new formular($tableName);
    $id = $_GET['id'];
    if ($tableName == "verbrauch") {  //Tabelle Verbrauch hat eine Schlüsselkombinbation aus 2 Schlüsseln -> Schlüsselteile werden hier in Array aufgespalten und an showChangeRowArr (kann mit Array als Input umgehen) übergeben
      $id = explode('_', $id);
      //var_dump($id);
      for ($n=0; $n < count($id); $n++) { 
        $id[$n] = intval($id[$n]);  // Strings in Integer umwandeln
      }
      $changeFormular->showChangeRowArr($id);
      //$id = $this->tableRows[$key][$this->columnNames[0]]."_".$this->tableRows[$key][$this->columnNames[1]];
    } else {
      $changeFormular->showChangeRow($id);
    }
  } else if (isset($_POST['sendUpdate'])) { //check if form was submitted
    //nach dem Absenden-Button im ersten Aufruf werden nun POST-Variablen verwendet, um das Update zu schicken
    $id = $_POST['id'];
    //echo "showChange.php - POST-Teil";
    echo "Für Tabelle ".$_POST['table'].", Zeile (ID) ".$id." wurde das Update versendet.<br>";
    $tableName = $_POST['table'];   //über hidden-Felder Werte von table & id wieder aufgegriffen
    //$id = $_POST['id'];
    $sendFormular = new formular($tableName);
    $sendFormular->loadColumns();
    $changePosts = array();
    for ($n = 0; $n < count($sendFormular->columnNames); $n++) {
      $changePosts[$n] = $_POST[$sendFormular->columnNames[$n]]; //get input text
      //echo "showChange.php: POST-Feld mit Namen '".$sendFormular->columnNames[$n]."' enthält den POST ".$changePosts[$n]."<br>";
    }
    $sendFormular->updateRow($changePosts);
    echo "
      <br><br>
      <div>
        <lable>zurück zur Ansicht aller vorhandenen Tabellen</lable>
          <a href='showOverview.php'>hier</a>
      </div>
    ";

    //$sendFormular->updateRow($id);
  } else {
    echo "Weder GET noch POST ist angekommen...";
  }
  
  echo "
      </body>
    </html>
  ";
?>