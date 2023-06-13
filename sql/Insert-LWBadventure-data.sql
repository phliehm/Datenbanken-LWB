--  Miniwelt: LWBadventure
--  Umsetzung für das Datenbankpraktikum der LWB Informatik
-- 
--  Annalena Cyriacus, Benjamin Schneider, Philipp Liehm, Martin Seiß
--
--  Start: 13.06.2023


-- npc(npcNr$,npcName)
INSERT INTO npc VALUES (1, 'Darth Schmidter');
INSERT INTO npc VALUES (2, 'Winnie the K');
INSERT INTO npc VALUES (3, 'Fab Web');
INSERT INTO npc VALUES (4, 'Amoebi');
INSERT INTO npc VALUES (5, 'J.EthI');
INSERT INTO npc VALUES (6, 'Herk');
INSERT INTO npc VALUES (7, 'Heidi');
INSERT INTO npc VALUES (8, 'Palim Palim');
INSERT INTO npc VALUES (9, 'Hubi-Horde');


-- dozentInnen(dNPCnr$!,lieblingsgetreank)
-- Folgende Getränkezuteilung wird verwendet:
-- ******************************************
--  Darth Schmidter		-	Extraschwarzer Kaffee
--	Herk				-	Melissentee									//lieber Hefeweizen?
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


-- sNPC(sNPCnr$!,funktion)
INSERT INTO sNPC VALUES (1, 'StEPS-Chefin'); 
-- INSERT INTO sNPC VALUES (2, 'Palim-Palim rufen!');
INSERT INTO sNPC VALUES (2, 'Helferlein'); 
INSERT INTO sNPC VALUES (3, 'Kontrolletis'); 


-- raum(raumNr$, raumName, ort)
INSERT INTO raum VALUES (0, 'Main Floor', 'LWB-World');
INSERT INTO raum VALUES (1, '1. Semester', 'FU Berlin');
INSERT INTO raum VALUES (2, '2. Semester', 'Home Office');
INSERT INTO raum VALUES (3, '3. Semester', 'FU Berlin');
INSERT INTO raum VALUES (4, '4. Semester', 'StEPS');
INSERT INTO raum VALUES (5, 'Nichtzeugnis-Verleihung', 'schöner Ort');


-- Herausgenommen: kursraum(kRaumNr$!) einfacher und konsitenter
--INSERT INTO kursraum VALUES (1);
--INSERT INTO kursraum VALUES (2);
--INSERT INTO kursraum VALUES (3);
--INSERT INTO kursraum VALUES (4);


-- sRaum(sRaumNr$!,sFunktion)											// Warum?
INSERT INTO sRaum VALUES (0, 'Zugang zu Kursräume');
INSERT INTO sRaum VALUES (5, 'Zeugnisvergabe');


-- veranstaltung(vNr$, vName, sws, thema, kuerzel, vDozNr!, vRaumNr!)
INSERT INTO veranstaltung VALUES (1, 'Betriebssystemwerkzeuge',2,'Rechnerarchitektur, Betriebs- und Kommunikationssysteme','BSW',2,1);
INSERT INTO veranstaltung VALUES (2, 'Funktionale Programmierung',8,'Programmierung','FP',3,1);
INSERT INTO veranstaltung VALUES (3, 'Grundlagen der Technischen Informatik',6,'Theoretische und technische Informatik','RS',2,1);
INSERT INTO veranstaltung VALUES (4, 'Imperative und objektorientierte Programmierung',7,'Programmierung','ALP2',1,2);
INSERT INTO veranstaltung VALUES (5, 'Rechnerarchitektur',4,'Rechnerarchitektur, Betriebs- und Kommunikationssysteme','RO',2,2);
INSERT INTO veranstaltung VALUES (6, 'Grundlagen der Theoretischen Informatik',5,'Theoretische und technische Informatik','EthI',5,1);
INSERT INTO veranstaltung VALUES (7, 'Datenstrukturen und Datenabstraktion',6,'Programmierung','ALP3',1,3);
INSERT INTO veranstaltung VALUES (8, 'Datenbanksysteme',6,'Datenbanken','DBSA',6,3);
INSERT INTO veranstaltung VALUES (9, 'Fachdidaktik Informatik',4,'Didaktik','DDI',2,3);
INSERT INTO veranstaltung VALUES (10, 'Nichtsequentielle und verteilte Programmierung',6,'Programmierung','NSP',1,4);
INSERT INTO veranstaltung VALUES (11, 'Rechnernetze',6,'Rechnerarchitektur, Betriebs- und Kommunikationssysteme','NET',2,4);
INSERT INTO veranstaltung VALUES (12, 'Unterrichtsbezogenes Softwarepraktikum',3,'Programmierung','SWP',4,4);
INSERT INTO veranstaltung VALUES (13, 'Unterrichtsbezogenes Datenbankpraktikum',3,'Datenbanken','DBP',6,4);
INSERT INTO veranstaltung VALUES (14, 'Analyse fachlichen Lernens',3,'Didaktik','AfL',3,4);

-- spielerInnen(spNr$, spName, spRaumNr)
INSERT INTO spielerInnen VALUES (1, 'Cyra',0);  -- sRaumNr = 0 default
INSERT INTO spielerInnen VALUES (2, 'Maddi',0);
INSERT INTO spielerInnen VALUES (3, 'Ben',0);
INSERT INTO spielerInnen VALUES (4, 'Phil',0);
INSERT INTO spielerInnen VALUES (5, 'Klocki',0);
INSERT INTO spielerInnen VALUES (6, 'Bob',0);
INSERT INTO spielerInnen VALUES (7, 'LWB-Master',0);
INSERT INTO spielerInnen VALUES (8, 'Nerd42',0);


-- minigame (gameNr$, gameName, gameVnr!)
INSERT INTO minigame VALUES (1, 'Muster-Spiel',2);
INSERT INTO minigame VALUES (2, 'Bauelemente-Spiel',3);
INSERT INTO minigame VALUES (3, 'Vaderobi-Game',4);
INSERT INTO minigame VALUES (4, 'Getränkeautomaten-Spiel',6);
INSERT INTO minigame VALUES (5, 'SQL-Quest',8);
INSERT INTO minigame VALUES (6, 'FachJargon',9);
INSERT INTO minigame VALUES (7, 'Food-Moorhuhn',10);
INSERT INTO minigame VALUES (8, 'theNETgame',11);
INSERT INTO minigame VALUES (9, 'BugAttack',12);


-- aufenthalt(asNPCnr!,aRaumNr!)
INSERT INTO aufenthalt VALUES (7, 0);
INSERT INTO aufenthalt VALUES (7, 4);
INSERT INTO aufenthalt VALUES (7, 5);
INSERT INTO aufenthalt VALUES (8, 0);
INSERT INTO aufenthalt VALUES (9, 0);
INSERT INTO aufenthalt VALUES (9, 4);


-- spielstand(sGameNr!, sSpNr!, Note, Punktzahl)
INSERT INTO spielstand VALUES (1, 1, 1.7, 325);
INSERT INTO spielstand VALUES (1, 2, 1.3, 325);
INSERT INTO spielstand VALUES (1, 3, 1.0, 325);
INSERT INTO spielstand VALUES (1, 4, 2.3, 325);
INSERT INTO spielstand VALUES (1, 5, 2.0, 325);
INSERT INTO spielstand VALUES (1, 6, 4.0, 325);
INSERT INTO spielstand VALUES (1, 7, 1.0, 325);
INSERT INTO spielstand VALUES (1, 8, 3.0, 325);

INSERT INTO spielstand VALUES (2, 1, 1.3, 225);
INSERT INTO spielstand VALUES (2, 2, 1.0, 225);
INSERT INTO spielstand VALUES (2, 3, 1.3, 225);
INSERT INTO spielstand VALUES (2, 4, 1.7, 225);
INSERT INTO spielstand VALUES (2, 5, 2.0, 225);
INSERT INTO spielstand VALUES (2, 6, 3.7, 225);
INSERT INTO spielstand VALUES (2, 7, 1.0, 225);
INSERT INTO spielstand VALUES (2, 8, 2.7, 225);

INSERT INTO spielstand VALUES (3, 1, 1.0, 630);
INSERT INTO spielstand VALUES (3, 2, 2.3, 432);
INSERT INTO spielstand VALUES (3, 3, 1.3, 555);
INSERT INTO spielstand VALUES (3, 4, 1.7, 512);
INSERT INTO spielstand VALUES (3, 5, 2.0, 487);
INSERT INTO spielstand VALUES (3, 6, 3.3, 333);
INSERT INTO spielstand VALUES (3, 7, 1.0, 650);
INSERT INTO spielstand VALUES (3, 8, 1.0, 600);

INSERT INTO spielstand VALUES (4, 1, 1.3, 325);
INSERT INTO spielstand VALUES (4, 2, 1.0, 325);
INSERT INTO spielstand VALUES (4, 3, 1.7, 325);
INSERT INTO spielstand VALUES (4, 4, 1.3, 325);
INSERT INTO spielstand VALUES (4, 5, 2.0, 325);
INSERT INTO spielstand VALUES (4, 6, 3.0, 325);
INSERT INTO spielstand VALUES (4, 7, 2.3, 325);

INSERT INTO spielstand VALUES (5, 1, 1.0, 100);
INSERT INTO spielstand VALUES (5, 2, 1.7, 88);
INSERT INTO spielstand VALUES (5, 3, 1.3, 96);
INSERT INTO spielstand VALUES (5, 4, 2.0, 79);
INSERT INTO spielstand VALUES (5, 6, 6.0, 44);
INSERT INTO spielstand VALUES (5, 7, 1.0, 100);
INSERT INTO spielstand VALUES (5, 8, 2.3, 73);

INSERT INTO spielstand VALUES (6, 1, 2.3, 325);
INSERT INTO spielstand VALUES (6, 2, 2.0, 325);
INSERT INTO spielstand VALUES (6, 3, 1.7, 325);
INSERT INTO spielstand VALUES (6, 4, 1.0, 325);
INSERT INTO spielstand VALUES (6, 7, 1.0, 325);

INSERT INTO spielstand VALUES (7, 1, 1.7, 325);
INSERT INTO spielstand VALUES (7, 2, 2.3, 325);
INSERT INTO spielstand VALUES (7, 3, 1.0, 325);
INSERT INTO spielstand VALUES (7, 4, 1.3, 325);
INSERT INTO spielstand VALUES (7, 7, 1.0, 325);

INSERT INTO spielstand VALUES (8, 1, 1.3, 325);
INSERT INTO spielstand VALUES (8, 2, 1.0, 325);
INSERT INTO spielstand VALUES (8, 3, 1.7, 325);
INSERT INTO spielstand VALUES (8, 4, 1.3, 325);
INSERT INTO spielstand VALUES (8, 7, 1.0, 325);

INSERT INTO spielstand VALUES (9, 1, 3.0, 325);
INSERT INTO spielstand VALUES (9, 2, 2.0, 325);
INSERT INTO spielstand VALUES (9, 3, 1.3, 325);
INSERT INTO spielstand VALUES (9, 4, 1.0, 325);
INSERT INTO spielstand VALUES (9, 7, 1.0, 325);


--INSERT INTO Professoren VALUES (2126, 'Russel', 'C4', 5700.00, 232);
