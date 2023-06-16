--  Miniwelt: LWBadventure
--  Umsetzung für das Datenbankpraktikum der LWB Informatik
-- 
--  Annalena Cyriacus, Benjamin Schneider, Philipp Liehm, Martin Seiß
--
--  Start: 13.06.2023


-- npcs ($npcNr, npcName) 
INSERT INTO npcs VALUES (1, 'Darth Schmidter');
INSERT INTO npcs VALUES (2, 'Winnie the K');
INSERT INTO npcs VALUES (3, 'Fab Web');
INSERT INTO npcs VALUES (4, 'Amoebi');
INSERT INTO npcs VALUES (5, 'J.EthI');
INSERT INTO npcs VALUES (6, 'Herk');
INSERT INTO npcs VALUES (7, 'Heidi');
INSERT INTO npcs VALUES (8, 'Palim Palim');
INSERT INTO npcs VALUES (9, 'Hubi-Horde');


-- dozent_innen (!npcNr, lieblingsgetraenk)
-- Folgende Getränkezuteilung wird verwendet:
-- ******************************************
--  Darth Schmidter		-	Extraschwarzer Kaffee
--	Herk				-	Hefeweizen
--  J.EthI				-	Kaffee mit Milch und 2x Zucker
--  Fab Web				-	Cappuccino
--	Amoebi				-	Grüner Tee
--  Winnie the K		-	Bier
INSERT INTO dozent_innen VALUES (1, 'Extraschwarzer Kaffee');
INSERT INTO dozent_innen VALUES (2, 'Bier');
INSERT INTO dozent_innen VALUES (3, 'Cappuccino');
INSERT INTO dozent_innen VALUES (4, 'Grüner Tee');
INSERT INTO dozent_innen VALUES (5, 'Kaffee mit Milch und 2x Zucker');
INSERT INTO dozent_innen VALUES (6, 'Hefeweizen');


-- sonstigeNPCs (!npcNr, aufgabe)
INSERT INTO sonstigeNPCs VALUES (7, 'StEPS-Chefin'); 
INSERT INTO sonstigeNPCs VALUES (8, 'Helferlein'); 
INSERT INTO sonstigeNPCs VALUES (9, 'Kontrolletis'); 


-- raeume ($raumNr, raumName, ort, funktion)
INSERT INTO raeume VALUES (0, 'Main Floor', 'LWB-World','Start-Raum');
INSERT INTO raeume VALUES (1, '1. Semester', 'FU Berlin','Kursraum 1');
INSERT INTO raeume VALUES (2, '2. Semester', 'Home Office','Kursraum 2');
INSERT INTO raeume VALUES (3, '3. Semester', 'FU Berlin','Kursraum 3');
INSERT INTO raeume VALUES (4, '4. Semester', 'StEPS','Kursraum 4');
INSERT INTO raeume VALUES (5, 'Nichtzeugnis-Verleihung', 'schöner Ort','Schluss-Raum');

/*
-- kursraeume(semster,RaumNr !)
INSERT INTO kursraeume VALUES (1,1);
INSERT INTO kursraeume VALUES (2,2);
INSERT INTO kursraeume VALUES (3,3);
INSERT INTO kursraeume VALUES (4,4);


-- sraeume(sRaumNr$!,sFunktion)
INSERT INTO sraeume VALUES (0, 'Zugang zu Kursräume');
INSERT INTO sraeume VALUES (5, 'Zeugnisvergabe');
*/

-- themengebiete ($gebietNr, gebietName)
INSERT INTO themengebiete VALUES (1, 'Rechnerarchitektur, Betriebs- und Kommunikationssysteme');
INSERT INTO themengebiete VALUES (2, 'Programmierung');
INSERT INTO themengebiete VALUES (3, 'Theoretische und technische Informatik');
INSERT INTO themengebiete VALUES (4, 'Datenbanken');
INSERT INTO themengebiete VALUES (5, 'Didaktik');


-- veranstaltungen ($vNr, vName, kuerzel, sws, !gebietNr)
INSERT INTO veranstaltungen VALUES (1, 'Betriebssystemwerkzeuge','BSW',2,1,1);
INSERT INTO veranstaltungen VALUES (2, 'Funktionale Programmierung','FP',8,1,2);
INSERT INTO veranstaltungen VALUES (3, 'Grundlagen der Technischen Informatik','RS',6,1,3);
INSERT INTO veranstaltungen VALUES (4, 'Imperative und objektorientierte Programmierung','ALP2',7,2,2);
INSERT INTO veranstaltungen VALUES (5, 'Rechnerarchitektur','RO',4,2,1);
INSERT INTO veranstaltungen VALUES (6, 'Einführung in die Theoretischen Informatik','EthI',5,2,3);
INSERT INTO veranstaltungen VALUES (7, 'Datenstrukturen und Datenabstraktion','ALP3',6,3,2);
INSERT INTO veranstaltungen VALUES (8, 'Datenbanksysteme','DBSA',6,3,4);
INSERT INTO veranstaltungen VALUES (9, 'Fachdidaktik Informatik','DDI',4,3,5);
INSERT INTO veranstaltungen VALUES (10, 'Nichtsequentielle und verteilte Programmierung','NSP',6,4,2);
INSERT INTO veranstaltungen VALUES (11, 'Rechnernetze','NET',6,4,1);
INSERT INTO veranstaltungen VALUES (12, 'Unterrichtsbezogenes Softwarepraktikum','SWP',3,4,2);
INSERT INTO veranstaltungen VALUES (13, 'Unterrichtsbezogenes Datenbankpraktikum','DBP',3,4,4);
INSERT INTO veranstaltungen VALUES (14, 'Analyse fachlichen Lernens','AfL',3,4,5);


-- spieler_innen ($spNr, spName, schluesselanzahl, !raumNr)
INSERT INTO spieler_innen VALUES (1, 'Cyra',1,0);  						-- sRaumNr = 0 default
INSERT INTO spieler_innen VALUES (2, 'Maddi',1,0);
INSERT INTO spieler_innen VALUES (3, 'Ben',1,0);
INSERT INTO spieler_innen VALUES (4, 'Phil',1,0);
INSERT INTO spieler_innen VALUES (5, 'Klocki',1,0);
INSERT INTO spieler_innen VALUES (6, 'Bob',1,0);
INSERT INTO spieler_innen VALUES (7, 'LWB-Master',1,0);
INSERT INTO spieler_innen VALUES (8, 'Nerd42',1,0);


-- minigames ($gameNr, gameName, !vNr)
INSERT INTO minigames VALUES (1, 'Muster-Spiel',2);
INSERT INTO minigames VALUES (2, 'Bauelemente-Spiel',3);
INSERT INTO minigames VALUES (3, 'Vaderobi-Game',4);
INSERT INTO minigames VALUES (4, 'Getränkeautomaten-Spiel',6);
INSERT INTO minigames VALUES (5, 'SQL-Quest',8);
INSERT INTO minigames VALUES (6, 'FachJargon',9);
INSERT INTO minigames VALUES (7, 'Food-Moorhuhn',10);
INSERT INTO minigames VALUES (8, 'theNETgame',11);
INSERT INTO minigames VALUES (9, 'BugAttack',12);


-- aufenthaltsorte (!npcNr, !raumNr)
INSERT INTO aufenthaltsorte VALUES (7, 0);
INSERT INTO aufenthaltsorte VALUES (7, 4);
INSERT INTO aufenthaltsorte VALUES (7, 5);
INSERT INTO aufenthaltsorte VALUES (8, 0);
INSERT INTO aufenthaltsorte VALUES (9, 0);
INSERT INTO aufenthaltsorte VALUES (9, 4);


-- unterricht (!vNr, !npcNr, !raumNr)	
-- Achtung: Nur Raumnummer von 1 bis 4 verwenden!						
INSERT INTO unterricht VALUES (1,2,1);
INSERT INTO unterricht VALUES (2,3,1);
INSERT INTO unterricht VALUES (3,2,1);
INSERT INTO unterricht VALUES (4,1,2);
INSERT INTO unterricht VALUES (5,2,2);
INSERT INTO unterricht VALUES (6,5,2);
INSERT INTO unterricht VALUES (7,1,3);
INSERT INTO unterricht VALUES (8,6,3);
INSERT INTO unterricht VALUES (9,2,3);
INSERT INTO unterricht VALUES (10,1,4);
INSERT INTO unterricht VALUES (11,2,4);
INSERT INTO unterricht VALUES (12,4,4);
INSERT INTO unterricht VALUES (13,6,4);
INSERT INTO unterricht VALUES (14,3,4);


-- assistenz (!vNr,!npcNr)
INSERT INTO assistenz VALUES (1,1);
INSERT INTO assistenz VALUES (2,5);
INSERT INTO assistenz VALUES (3,1);
INSERT INTO assistenz VALUES (4,4);
INSERT INTO assistenz VALUES (10,2);
INSERT INTO assistenz VALUES (12,3);
INSERT INTO assistenz VALUES (14,4);


-- spielstaende (!gameNr, !spNr, Note, Punktzahl)
INSERT INTO spielstaende VALUES (1, 1, 1.7, 325);
INSERT INTO spielstaende VALUES (1, 2, 1.3, 325);
INSERT INTO spielstaende VALUES (1, 3, 1.0, 325);
INSERT INTO spielstaende VALUES (1, 4, 2.3, 325);
INSERT INTO spielstaende VALUES (1, 5, 2.0, 325);
INSERT INTO spielstaende VALUES (1, 6, 4.0, 325);
INSERT INTO spielstaende VALUES (1, 7, 1.0, 325);
INSERT INTO spielstaende VALUES (1, 8, 3.0, 325);

INSERT INTO spielstaende VALUES (2, 1, 1.3, 225);
INSERT INTO spielstaende VALUES (2, 2, 1.0, 225);
INSERT INTO spielstaende VALUES (2, 3, 1.3, 225);
INSERT INTO spielstaende VALUES (2, 4, 1.7, 225);
INSERT INTO spielstaende VALUES (2, 5, 2.0, 225);
INSERT INTO spielstaende VALUES (2, 6, 3.7, 225);
INSERT INTO spielstaende VALUES (2, 7, 1.0, 225);
INSERT INTO spielstaende VALUES (2, 8, 2.7, 225);

INSERT INTO spielstaende VALUES (3, 1, 1.0, 630);
INSERT INTO spielstaende VALUES (3, 2, 2.3, 432);
INSERT INTO spielstaende VALUES (3, 3, 1.3, 555);
INSERT INTO spielstaende VALUES (3, 4, 1.7, 512);
INSERT INTO spielstaende VALUES (3, 5, 2.0, 487);
INSERT INTO spielstaende VALUES (3, 6, 3.3, 333);
INSERT INTO spielstaende VALUES (3, 7, 1.0, 650);
INSERT INTO spielstaende VALUES (3, 8, 1.0, 600);

INSERT INTO spielstaende VALUES (4, 1, 1.3, 325);
INSERT INTO spielstaende VALUES (4, 2, 1.0, 325);
INSERT INTO spielstaende VALUES (4, 3, 1.7, 325);
INSERT INTO spielstaende VALUES (4, 4, 1.3, 325);
INSERT INTO spielstaende VALUES (4, 5, 2.0, 325);
INSERT INTO spielstaende VALUES (4, 6, 3.0, 325);
INSERT INTO spielstaende VALUES (4, 7, 2.3, 325);

INSERT INTO spielstaende VALUES (5, 1, 1.0, 100);
INSERT INTO spielstaende VALUES (5, 2, 1.7, 88);
INSERT INTO spielstaende VALUES (5, 3, 1.3, 96);
INSERT INTO spielstaende VALUES (5, 4, 2.0, 79);
INSERT INTO spielstaende VALUES (5, 6, 6.0, 44);
INSERT INTO spielstaende VALUES (5, 7, 1.0, 100);
INSERT INTO spielstaende VALUES (5, 8, 2.3, 73);

INSERT INTO spielstaende VALUES (6, 1, 2.3, 325);
INSERT INTO spielstaende VALUES (6, 2, 2.0, 325);
INSERT INTO spielstaende VALUES (6, 3, 1.7, 325);
INSERT INTO spielstaende VALUES (6, 4, 1.0, 325);
INSERT INTO spielstaende VALUES (6, 7, 1.0, 325);

INSERT INTO spielstaende VALUES (7, 1, 1.7, 325);
INSERT INTO spielstaende VALUES (7, 2, 2.3, 325);
INSERT INTO spielstaende VALUES (7, 3, 1.0, 325);
INSERT INTO spielstaende VALUES (7, 4, 1.3, 325);
INSERT INTO spielstaende VALUES (7, 7, 1.0, 325);

INSERT INTO spielstaende VALUES (8, 1, 1.3, 325);
INSERT INTO spielstaende VALUES (8, 2, 1.0, 325);
INSERT INTO spielstaende VALUES (8, 3, 1.7, 325);
INSERT INTO spielstaende VALUES (8, 4, 1.3, 325);
INSERT INTO spielstaende VALUES (8, 7, 1.0, 325);

INSERT INTO spielstaende VALUES (9, 1, 3.0, 325);
INSERT INTO spielstaende VALUES (9, 2, 2.0, 325);
INSERT INTO spielstaende VALUES (9, 3, 1.3, 325);
INSERT INTO spielstaende VALUES (9, 4, 1.0, 325);
INSERT INTO spielstaende VALUES (9, 7, 1.0, 325);

