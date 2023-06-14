<?php

  echo "
    <html>
      <head>
        <title>DB Tierwelt (INDEX)</title>
      </head>
      <body>
        <p>Überblick über die html | php | sql - Anbindung zur Datenbank Reptilien-Tierwelt</p>
        <div>
          <lable>Ansicht aller vorhandenen Tabellen</lable>
          <a href='showOverview.php'>hier</a>
        </div>
        <div>
          <lable>Aufgabenliste</lable>
          <a href='showTasks.php'>hier</a>
        </div>
        <div>
          <lable>Eingabeformular</lable>
          <a href='showFormular.php'>hier</a>
        </div>
      </body>
  </html>
  ";

  //echo "<br> SERVER liegt in:".$_SERVER["DOCUMENT_ROOT"];

  //echo "<br> URI in:".$_SERVER["REQUEST_URI"];

  //echo "<br> ini-Files-Liste:".php_ini_scanned_files();
?>