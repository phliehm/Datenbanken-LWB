<?php
//include('formularClass.php');
$inputName = strval($_GET['id']);
$inputVal = strval($_GET['val']);
$inputTable = strval($_GET['table']);

try {
  $doppelt = false;
  $con = new PDO ('pgsql:host=localhost;dbname=lewein','lewein','niewel');
  $queryString = "SELECT * FROM ".$inputTable." WHERE ".$inputName."='".$inputVal."'";
  //echo $queryString;
  $erg = $con->query($queryString);
  //var_dump($erg);
  //$erg = $con->query("SELECT * FROM buecher WHERE titel = 'Taschenbuch Datenbanken'");
  foreach ($erg as $row) {
    $doppelt = true;
    //print_r ($row);
    //echo "<br/>";
  }
  if ($doppelt) {
    echo 
      $inputTable."-".$inputName." enthält Wert (".$inputVal."): JA"
    ;
  } else {
    echo 
      $inputTable."-".$inputName." enthält Wert (".$inputVal."): NEIN"
    ;
  }

  $con = null;
  //echo $inputTable."-".$inputName." enthält Wert (".$inputVal."): NEIN";
  
} catch (PDOException $err) {
  echo "Fehler: " . htmlspecialchars ($err->getMessage ());
}

//var_dump($inputId);

//$buchFormular1 = new formular('buecher');
//$buchFormular1->loadTable();
?>