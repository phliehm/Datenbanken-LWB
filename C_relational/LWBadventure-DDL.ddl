/compact_listings on
%
% Miniwelt LWBadventure, eingeschraenkte Auswahl
% erstellt im Rahmen des Datenbankpraktikums (LWB Informatik 2021-2023)
% von Annalena Cyriacus, Philipp Liehm, Benjamin Schneider und Martin Seiß
%
/output           off
/multiline        on

% NPCs ($NPCNr, NPCName) 				NPCs - Non-Playing-Character
CREATE OR REPLACE TABLE NPCs (
  NPCNr 		INTEGER,	
  NPCName 		STRING
);

% Dozent_innen (!NPCNr, Lieblingsgetraenk)
CREATE OR REPLACE TABLE Dozent_innen (
  NPCNr 		INTEGER,
  Lieblingsgetraenk	STRING
);

% sonstigeNPCs (!NPCNr, Aufgabe)
CREATE OR REPLACE TABLE sonstigeNPCs (
  NPCNr 		INTEGER,
  Aufgabe		STRING
);

% Raeume ($RaumNr, RaumName, Ort, Funktion)
CREATE OR REPLACE TABLE Raeume (
  RaumNr 		INTEGER,
  RaumName		STRING,
  Ort			STRING,
  Funktion		STRING
  );

% Themengebiete ($GebietNr, GebietName)
CREATE OR REPLACE TABLE Themengebiete (
  GebietNr 		INTEGER,
  GebietName		STRING
);

% Veranstaltungen ($VNr, VName, Kuerzel, SWS, Semester, !GebietNr)
CREATE OR REPLACE TABLE Veranstaltungen (
  VNr 			INTEGER,
  VName			STRING,
  Kuerzel		STRING,
  SWS			INTEGER,
  Semester		INTEGER,
  GebietNr		INTEGER
);

% Spieler_innen ($SpNr, SpName, Schluesselanzahl, !RaumNr)
CREATE OR REPLACE TABLE Spieler_innen (
  SpNr	 		INTEGER,
  SpName 		STRING,
  Schuesselanzahl	INTEGER,
  RaumNr 		INTEGER
);

% Minigames ($GameNr, GameName, !VNr)
CREATE OR REPLACE TABLE Minigames(
  GameNr 		INTEGER,
  GameName		STRING,
  VNr			INTEGER
);

% Spielstaende (!GameNr, !SpNr, Note, Punktzahl)
CREATE OR REPLACE TABLE Spielstaende (
  GameNr 		INTEGER,
  SpNr			INTEGER,
  Note			FLOAT,
  Punkte		INTEGER
);

% Aufenthaltsorte (!NPCNr, !RaumNr)
CREATE OR REPLACE TABLE Aufenthaltsorte (
  NPCNr 		INTEGER,
  RaumNr		INTEGER
);

% Unterricht (!VNr, !NPCNr, !RaumNr)							
CREATE OR REPLACE TABLE Unterricht (
  VNr 			INTEGER,
  NPCNr			INTEGER,
  RaumNr		INTEGER
);

% Assistenz (!VNr,!NPCNr)
CREATE OR REPLACE TABLE Assistenz (
  VNr 			INTEGER,
  NPCNr			INTEGER
);

% NPCs ($NPCNr, NPCName) 
INSERT INTO NPCs VALUES (1, 'Darth Schmidter');
INSERT INTO NPCs VALUES (2, 'Winnie the K');
INSERT INTO NPCs VALUES (3, 'Fab Web');
INSERT INTO NPCs VALUES (4, 'Amoebi');
INSERT INTO NPCs VALUES (5, 'J.EthI');
INSERT INTO NPCs VALUES (6, 'Herk');
INSERT INTO NPCs VALUES (7, 'Heidi');
INSERT INTO NPCs VALUES (8, 'Palim Palim');
INSERT INTO NPCs VALUES (9, 'Hubi-Horde');


% Dozent_innen (!NPCNr, Lieblingsgetraenk)
% Folgende Getränkezuteilung wird verwendet:
% ******************************************
%  Darth Schmidter		-	Extraschwarzer Kaffee
%  Herk				-	Hefeweizen
%  J.EthI			-	Kaffee mit Milch und 2x Zucker
%  Fab Web			-	Cappuccino
%  Amoebi			-	Grüner Tee
%  Winnie the K			-	Bier
INSERT INTO Dozent_innen VALUES (1, 'Extraschwarzer Kaffee');
INSERT INTO Dozent_innen VALUES (2, 'Bier');
INSERT INTO Dozent_innen VALUES (3, 'Cappuccino');
INSERT INTO Dozent_innen VALUES (4, 'Grüner Tee');
INSERT INTO Dozent_innen VALUES (5, 'Kaffee mit Milch und 2x Zucker');
INSERT INTO Dozent_innen VALUES (6, 'Hefeweizen');

% sonstigeNPCs (!NPCNr, Aufgabe)
INSERT INTO sonstigeNPCs VALUES (7, 'StEPS-Chefin'); 
INSERT INTO sonstigeNPCs VALUES (8, 'Helferlein'); 
INSERT INTO sonstigeNPCs VALUES (9, 'Kontrolletis'); 


% Raeume ($RaumNr, RaumName, Ort, Funktion)
INSERT INTO Raeume VALUES (0, 'Main Floor', 'LWB-World','Start-Raum');
INSERT INTO Raeume VALUES (1, '1. Semester', 'FU Berlin','Kursraum 1');
INSERT INTO Raeume VALUES (2, '2. Semester', 'Home Office','Kursraum 2');
INSERT INTO Raeume VALUES (3, '3. Semester', 'FU Berlin','Kursraum 3');
INSERT INTO Raeume VALUES (4, '4. Semester', 'StEPS','Kursraum 4');
INSERT INTO Raeume VALUES (5, 'Nichtzeugnis-Verleihung', 'schöner Ort','Schluss-Raum');


% Themengebiete ($GebietNr, GebietName)
INSERT INTO Themengebiete VALUES (1, 'Rechnerarchitektur, Betriebs- und Kommunikationssysteme');
INSERT INTO Themengebiete VALUES (2, 'Programmierung');
INSERT INTO Themengebiete VALUES (3, 'Theoretische und technische Informatik');
INSERT INTO Themengebiete VALUES (4, 'Datenbanken');
INSERT INTO Themengebiete VALUES (5, 'Didaktik');


% Veranstaltungen ($VNr, VName, Kuerzel, SWS, !GebietNr)
INSERT INTO Veranstaltungen VALUES (1, 'Betriebssystemwerkzeuge','BSW',2,1,1);
INSERT INTO Veranstaltungen VALUES (2, 'Funktionale Programmierung','FP',8,1,2);
INSERT INTO Veranstaltungen VALUES (3, 'Grundlagen der Technischen Informatik','RS',6,1,3);
INSERT INTO Veranstaltungen VALUES (4, 'Imperative und objektorientierte Programmierung','ALP2',7,2,2);
INSERT INTO Veranstaltungen VALUES (5, 'Rechnerarchitektur','RO',4,2,1);
INSERT INTO Veranstaltungen VALUES (6, 'Einführung in die Theoretischen Informatik','EthI',5,2,3);
INSERT INTO Veranstaltungen VALUES (7, 'Datenstrukturen und Datenabstraktion','ALP3',6,3,2);
INSERT INTO Veranstaltungen VALUES (8, 'Datenbanksysteme','DBSA',6,3,4);
INSERT INTO Veranstaltungen VALUES (9, 'Fachdidaktik Informatik','DDI',4,3,5);
INSERT INTO Veranstaltungen VALUES (10, 'Nichtsequentielle und verteilte Programmierung','NSP',6,4,2);
INSERT INTO Veranstaltungen VALUES (11, 'Rechnernetze','NET',6,4,1);
INSERT INTO Veranstaltungen VALUES (12, 'Unterrichtsbezogenes Softwarepraktikum','SWP',3,4,2);
INSERT INTO Veranstaltungen VALUES (13, 'Unterrichtsbezogenes Datenbankpraktikum','DBP',3,4,4);
INSERT INTO Veranstaltungen VALUES (14, 'Analyse fachlichen Lernens','AfL',3,4,5);


% Spieler_innen ($SpNr, SpName, Schluesselanzahl, !RaumNr)
INSERT INTO Spieler_innen VALUES (1, 'Cyra',1,0);
INSERT INTO Spieler_innen VALUES (2, 'Maddi',1,0);
INSERT INTO Spieler_innen VALUES (3, 'Ben',1,0);
INSERT INTO Spieler_innen VALUES (4, 'Phil',1,0);
INSERT INTO Spieler_innen VALUES (5, 'Klocki',1,0);
INSERT INTO Spieler_innen VALUES (6, 'Bob',1,0);
INSERT INTO Spieler_innen VALUES (7, 'LWB-Master',1,0);
INSERT INTO Spieler_innen VALUES (8, 'Nerd42',1,0);


% Minigames ($GameNr, GameName, !VNr)
INSERT INTO Minigames VALUES (1, 'Muster-Spiel',2);
INSERT INTO Minigames VALUES (2, 'Bauelemente-Spiel',3);
INSERT INTO Minigames VALUES (3, 'Vaderobi-Game',4);
INSERT INTO Minigames VALUES (4, 'Getränkeautomaten-Spiel',6);
INSERT INTO Minigames VALUES (5, 'SQL-Quest',8);
INSERT INTO Minigames VALUES (6, 'FachJargon',9);
INSERT INTO Minigames VALUES (7, 'Food-Moorhuhn',10);
INSERT INTO Minigames VALUES (8, 'theNETgame',11);
INSERT INTO Minigames VALUES (9, 'BugAttack',12);


% Aufenthaltsorte (!NPCNr, !RaumNr)
INSERT INTO Aufenthaltsorte VALUES (7, 0);
INSERT INTO Aufenthaltsorte VALUES (7, 4);
INSERT INTO Aufenthaltsorte VALUES (7, 5);
INSERT INTO Aufenthaltsorte VALUES (8, 0);
INSERT INTO Aufenthaltsorte VALUES (9, 0);
INSERT INTO Aufenthaltsorte VALUES (9, 4);


% Unterricht (!VNr, !NPCNr, !RaumNr)							
INSERT INTO Unterricht VALUES (1,2,1);
INSERT INTO Unterricht VALUES (2,3,1);
INSERT INTO Unterricht VALUES (3,2,1);
INSERT INTO Unterricht VALUES (4,1,2);
INSERT INTO Unterricht VALUES (5,2,2);
INSERT INTO Unterricht VALUES (6,5,2);
INSERT INTO Unterricht VALUES (7,1,3);
INSERT INTO Unterricht VALUES (8,6,3);
INSERT INTO Unterricht VALUES (9,2,3);
INSERT INTO Unterricht VALUES (10,1,4);
INSERT INTO Unterricht VALUES (11,2,4);
INSERT INTO Unterricht VALUES (12,4,4);
INSERT INTO Unterricht VALUES (13,6,4);
INSERT INTO Unterricht VALUES (14,3,4);


% Assistenz (!VNr,!NPCNr)
INSERT INTO Assistenz VALUES (1,1);
INSERT INTO Assistenz VALUES (2,5);
INSERT INTO Assistenz VALUES (3,1);
INSERT INTO Assistenz VALUES (4,4);
INSERT INTO Assistenz VALUES (10,2);
INSERT INTO Assistenz VALUES (12,3);
INSERT INTO Assistenz VALUES (14,4);

% spielstaende (!GameNr, !SpNr, Note, Punktzahl)
INSERT INTO spielstaende VALUES (1, 1, 1.7, 1333);
INSERT INTO spielstaende VALUES (1, 2, 1.3, 1456);
INSERT INTO spielstaende VALUES (1, 3, 1.0, 1800);
INSERT INTO spielstaende VALUES (1, 4, 2.3, 999);
INSERT INTO spielstaende VALUES (1, 5, 2.0, 1210);
INSERT INTO spielstaende VALUES (1, 6, 4.0, 400);
INSERT INTO spielstaende VALUES (1, 7, 1.0, 2000);
INSERT INTO spielstaende VALUES (1, 8, 3.0, 789);

INSERT INTO spielstaende VALUES (2, 1, 1.3, 29);
INSERT INTO spielstaende VALUES (2, 2, 1.0, 33);
INSERT INTO spielstaende VALUES (2, 3, 1.3, 29);
INSERT INTO spielstaende VALUES (2, 4, 1.7, 27);
INSERT INTO spielstaende VALUES (2, 5, 2.0, 25);
INSERT INTO spielstaende VALUES (2, 6, 3.7, 17);
INSERT INTO spielstaende VALUES (2, 7, 1.0, 33);
INSERT INTO spielstaende VALUES (2, 8, 2.7, 22);

INSERT INTO spielstaende VALUES (3, 1, 1.0, 630);
INSERT INTO spielstaende VALUES (3, 2, 2.3, 432);
INSERT INTO spielstaende VALUES (3, 3, 1.3, 555);
INSERT INTO spielstaende VALUES (3, 4, 1.7, 512);
INSERT INTO spielstaende VALUES (3, 5, 2.0, 487);
INSERT INTO spielstaende VALUES (3, 6, 3.3, 333);
INSERT INTO spielstaende VALUES (3, 7, 1.0, 650);
INSERT INTO spielstaende VALUES (3, 8, 1.0, 600);

INSERT INTO spielstaende VALUES (4, 1, 1.3, 5);
INSERT INTO spielstaende VALUES (4, 2, 1.0, 6);
INSERT INTO spielstaende VALUES (4, 3, 1.7, 4);
INSERT INTO spielstaende VALUES (4, 4, 1.3, 5);
INSERT INTO spielstaende VALUES (4, 5, 2.0, 3);
INSERT INTO spielstaende VALUES (4, 6, 3.0, 2);
INSERT INTO spielstaende VALUES (4, 7, 1.0, 6);

INSERT INTO spielstaende VALUES (5, 1, 1.0, 100);
INSERT INTO spielstaende VALUES (5, 2, 1.7, 88);
INSERT INTO spielstaende VALUES (5, 3, 1.3, 96);
INSERT INTO spielstaende VALUES (5, 4, 2.0, 79);
INSERT INTO spielstaende VALUES (5, 6, 6.0, 44);
INSERT INTO spielstaende VALUES (5, 7, 1.0, 100);
INSERT INTO spielstaende VALUES (5, 8, 2.3, 73);

INSERT INTO spielstaende VALUES (6, 1, 2.3, 71);
INSERT INTO spielstaende VALUES (6, 2, 2.0, 76);
INSERT INTO spielstaende VALUES (6, 3, 1.7, 83);
INSERT INTO spielstaende VALUES (6, 4, 1.0, 99);
INSERT INTO spielstaende VALUES (6, 7, 1.0, 100);

INSERT INTO spielstaende VALUES (7, 1, 1.7, 456);
INSERT INTO spielstaende VALUES (7, 2, 2.3, 369);
INSERT INTO spielstaende VALUES (7, 3, 1.0, 555);
INSERT INTO spielstaende VALUES (7, 4, 1.3, 512);
INSERT INTO spielstaende VALUES (7, 7, 1.0, 600);

INSERT INTO spielstaende VALUES (8, 1, 1.3, 140);
INSERT INTO spielstaende VALUES (8, 2, 1.0, 150);
INSERT INTO spielstaende VALUES (8, 3, 1.7, 128);
INSERT INTO spielstaende VALUES (8, 4, 1.3, 138);
INSERT INTO spielstaende VALUES (8, 7, 1.0, 160);

INSERT INTO spielstaende VALUES (9, 1, 3.0, 3210);
INSERT INTO spielstaende VALUES (9, 2, 2.0, 3775);
INSERT INTO spielstaende VALUES (9, 3, 1.3, 4252);
INSERT INTO spielstaende VALUES (9, 4, 1.0, 4999);
INSERT INTO spielstaende VALUES (9, 7, 1.0, 5000);

/output on
/dbschema
