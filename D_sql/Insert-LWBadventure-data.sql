--  Miniwelt: LWBadventure
--  Umsetzung für das Datenbankpraktikum der LWB Informatik
-- 
--  Annalena Cyriacus, Benjamin Schneider, Philipp Liehm, Martin Seiß
--
--  Start: 13.06.2023


-- npcs(npcNr$,npcName)
INSERT INTO npcs VALUES (1, 'Darth Schmidter');
INSERT INTO npcs VALUES (2, 'Winnie the K');
INSERT INTO npcs VALUES (3, 'Fab Web');
INSERT INTO npcs VALUES (4, 'Amoebi');
INSERT INTO npcs VALUES (5, 'J.EthI');
INSERT INTO npcs VALUES (6, 'Herk');
INSERT INTO npcs VALUES (7, 'Heidi');
INSERT INTO npcs VALUES (8, 'Palim Palim');
INSERT INTO npcs VALUES (9, 'Hubi-Horde');


-- dozentInnen(dNPCnr$!,lieblingsgetreank)
-- Folgende Getränkezuteilung wird verwendet:
-- ******************************************
--  Darth Schmidter		-	Extraschwarzer Kaffee
--	Herk				-	Hefeweizen
--  J.EthI				-	Kaffee mit Milch und 2x Zucker
--  Fab Web				-	Cappuccino
--	Amoebi				-	Grüner Tee
--  Winnie the K		-	Bier
INSERT INTO dozentInnen VALUES (1, 'Extraschwarzer Kaffee');
--INSERT INTO dozentInnen VALUES (6, 'Melissentee');
INSERT INTO dozentInnen VALUES (6, 'Hefeweizen');
INSERT INTO dozentInnen VALUES (5, 'Kaffee mit Milch und 2x Zucker');
INSERT INTO dozentInnen VALUES (3, 'Cappuccino');
INSERT INTO dozentInnen VALUES (4, 'Grüner Tee');
INSERT INTO dozentInnen VALUES (2, 'Bier');


-- sNPCs(sNPCnr$!,funktion)
INSERT INTO sNPCs VALUES (7, 'StEPS-Chefin'); 
-- INSERT INTO sNPC VALUES (2, 'Palim-Palim rufen!');
INSERT INTO sNPCs VALUES (8, 'Helferlein'); 
INSERT INTO sNPCs VALUES (9, 'Hubi'); 


-- raeume(raumNr$, raumName, ort)
INSERT INTO raeume VALUES (0, 'Main Floor', 'LWB-World');
INSERT INTO raeume VALUES (1, '1. Semester', 'FU Berlin');
INSERT INTO raeume VALUES (2, '2. Semester', 'Home Office');
INSERT INTO raeume VALUES (3, '3. Semester', 'FU Berlin');
INSERT INTO raeume VALUES (4, '4. Semester', 'StEPS');
INSERT INTO raeume VALUES (5, 'Nichtzeugnis-Verleihung', 'schöner Ort');


-- Herausgenommen: kursraeume(kRaumNr$!) einfacher und konsitenter
--INSERT INTO kursraeume VALUES (1);
--INSERT INTO kursraeume VALUES (2);
--INSERT INTO kursraeume VALUES (3);
--INSERT INTO kursraeume VALUES (4);


-- sraeume(sRaumNr$!,sFunktion)											// Warum?
INSERT INTO sraeume VALUES (0, 'Zugang zu Kursräume');
INSERT INTO sraeume VALUES (5, 'Zeugnisvergabe');


-- veranstaltungen(vNr$, vName, sws, thema, kuerzel, vDozNr!, vRaumNr!)
INSERT INTO veranstaltungen VALUES (1, 'Betriebssystemwerkzeuge',2,'Rechnerarchitektur, Betriebs- und Kommunikationssysteme','BSW',2,1);
INSERT INTO veranstaltungen VALUES (2, 'Funktionale Programmierung',8,'Programmierung','FP',3,1);
INSERT INTO veranstaltungen VALUES (3, 'Grundlagen der Technischen Informatik',6,'Theoretische und technische Informatik','RS',2,1);
INSERT INTO veranstaltungen VALUES (4, 'Imperative und objektorientierte Programmierung',7,'Programmierung','ALP2',1,2);
INSERT INTO veranstaltungen VALUES (5, 'Rechnerarchitektur',4,'Rechnerarchitektur, Betriebs- und Kommunikationssysteme','RO',2,2);
INSERT INTO veranstaltungen VALUES (6, 'Grundlagen der Theoretischen Informatik',5,'Theoretische und technische Informatik','EthI',5,1);
INSERT INTO veranstaltungen VALUES (7, 'Datenstrukturen und Datenabstraktion',6,'Programmierung','ALP3',1,3);
INSERT INTO veranstaltungen VALUES (8, 'Datenbanksysteme',6,'Datenbanken','DBSA',6,3);
INSERT INTO veranstaltungen VALUES (9, 'Fachdidaktik Informatik',4,'Didaktik','DDI',2,3);
INSERT INTO veranstaltungen VALUES (10, 'Nichtsequentielle und verteilte Programmierung',6,'Programmierung','NSP',1,4);
INSERT INTO veranstaltungen VALUES (11, 'Rechnernetze',6,'Rechnerarchitektur, Betriebs- und Kommunikationssysteme','NET',2,4);
INSERT INTO veranstaltungen VALUES (12, 'Unterrichtsbezogenes Softwarepraktikum',3,'Programmierung','SWP',4,4);
INSERT INTO veranstaltungen VALUES (13, 'Unterrichtsbezogenes Datenbankpraktikum',3,'Datenbanken','DBP',6,4);
INSERT INTO veranstaltungen VALUES (14, 'Analyse fachlichen Lernens',3,'Didaktik','AfL',3,4);

-- spielerInnen(spNr$, spName, spRaumNr)
INSERT INTO spielerInnen VALUES (1, 'Cyra',0);  -- sRaumNr = 0 default
INSERT INTO spielerInnen VALUES (2, 'Maddi',0);
INSERT INTO spielerInnen VALUES (3, 'Ben',0);
INSERT INTO spielerInnen VALUES (4, 'Phil',0);
INSERT INTO spielerInnen VALUES (5, 'Klocki',0);
INSERT INTO spielerInnen VALUES (6, 'Bob',0);
INSERT INTO spielerInnen VALUES (7, 'LWB-Master',0);
INSERT INTO spielerInnen VALUES (8, 'Nerd42',0);


-- minigames (gameNr$, gameName, gameVnr!)
INSERT INTO minigames VALUES (1, 'Muster-Spiel',2);
INSERT INTO minigames VALUES (2, 'Bauelemente-Spiel',3);
INSERT INTO minigames VALUES (3, 'Vaderobi-Game',4);
INSERT INTO minigames VALUES (4, 'Getränkeautomaten-Spiel',6);
INSERT INTO minigames VALUES (5, 'SQL-Quest',8);
INSERT INTO minigames VALUES (6, 'FachJargon',9);
INSERT INTO minigames VALUES (7, 'Food-Moorhuhn',10);
INSERT INTO minigames VALUES (8, 'theNETgame',11);
INSERT INTO minigames VALUES (9, 'BugAttack',12);


-- aufenthalt(asNPCnr!,aRaumNr!)
INSERT INTO aufenthalt VALUES (7, 0);
INSERT INTO aufenthalt VALUES (7, 4);
INSERT INTO aufenthalt VALUES (7, 5);
INSERT INTO aufenthalt VALUES (8, 0);
INSERT INTO aufenthalt VALUES (9, 0);
INSERT INTO aufenthalt VALUES (9, 4);


-- spielstaende(sGameNr!, sSpNr!, Note, Punktzahl)
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


--INSERT INTO Professoren VALUES (2126, 'Russel', 'C4', 5700.00, 232);
