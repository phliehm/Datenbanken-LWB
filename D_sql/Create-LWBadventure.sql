--  Miniwelt: LWBadventure
--  Umsetzung für das Datenbankpraktikum der LWB Informatik
-- 
--  Annalena Cyriacus, Benjamin Schneider, Philipp Liehm, Martin Seiß
--
--  Start: 13.06.2023


-- psql-Befehle:
--		\i Create-LWBadventure.sql	-- zum Erzeugen der Tabellen
--		\i Drop-LWBadventure.sql	-- zum Löschen der Tabellen
-- 		\i Insert-LWBadventure-data.sql -- zum Einfügen der Daten
--		\dt		-	alle Tabelle anzeigen
--		drop table dozentin;	- Löscht Tabelle dozentIn


-- Folgende Tabellen werden implementiert:
-- $ = Schlüssel; ! = Fremdschlüssel

-- npcs ($npcNr, npcName) 						NPCs - Non-Playing-Character
-- dozent_innen (!$npcNr, lieblingsgetraenk)
-- sonstigeNPCs (!$npcNr, aufgabe)
-- veranstaltungen ($vNr, vName, kuerzel, sws, semester, !gebietNr)
-- themengebiete ($gebietNr, gebietName)
-- minigames ($gameNr, gameName, !vNr)
-- spieler_innen ($spNr, spName, schluesselanzahl, !raumNr)
-- raeume ($raumNr, raumName, ort, funktion)
-- unterricht (!$vNr, !npcNr, !raumNr)							
-- spielstaende (!$gameNr, !$spNr, Note, Punkte)
-- aufenthaltsorte (!$npcNr, !$raumNr)
-- assistenz (!$vNr,!npcNr)



-- npcs ($npcNr, npcName) 						NPCs - Non-Playing-Character
CREATE TABLE npcs (
  npcNr 		INTEGER				NOT NULL,	
  npcName 		VARCHAR (50)		NOT NULL,
  CONSTRAINT npcKEY PRIMARY KEY (npcNr)
);
COMMENT ON Table npcs IS 'Miniwelt LWBadventure';


-- dozent_innen (!$npcNr, lieblingsgetraenk)
CREATE TABLE dozent_innen (
  npcNr 				INTEGER 			REFERENCES npcs (npcNr),	-- NPC-Nummer
  lieblingsgetraenk		VARCHAR (100)		NOT NULL,					-- jede/r Dozent_in braucht ein Lieblingsgetränk
  CONSTRAINT dozent_innenKEY PRIMARY KEY (npcNr)
);
COMMENT ON Table dozent_innen IS 'Miniwelt LWBadventure';


-- sonstigeNPCs (!$npcNr, aufgabe)
CREATE TABLE sonstigeNPCs (
  npcNr 		INTEGER 			REFERENCES npcs (npcNr),			-- NPC-Nummer
  aufgabe		VARCHAR (100)		NOT NULL,							-- jede/r sonstige NPC hat eine Aufgabe
  CONSTRAINT sonstigeNPCsKEY PRIMARY KEY (npcNr)
);
COMMENT ON Table sonstigeNPCs IS 'Miniwelt LWBadventure';


-- themengebiete ($gebietNr, gebietName)
CREATE TABLE themengebiete (
  gebietNr 		INTEGER				NOT NULL,
  gebietName	VARCHAR (100)		NOT NULL,
  CONSTRAINT themengebieteKEY PRIMARY KEY (gebietNr)
);
COMMENT ON Table themengebiete IS 'Miniwelt LWBadventure';


-- veranstaltungen ($vNr, vName, kuerzel, sws, semester, !gebietNr)
CREATE TABLE veranstaltungen (
  vNr 			INTEGER				NOT NULL,		-- Veranstaltungsnummer
  vName			VARCHAR (50)		NOT NULL,
  kuerzel		VARCHAR (5)			NOT NULL,
  sws			INTEGER				CHECK (sws > 0),
  semester		INTEGER				CHECK (semester > 0),
  gebietNr		INTEGER				REFERENCES themengebiete (gebietNr),
  CONSTRAINT veranstaltungenKEY PRIMARY KEY (vNr)
);
COMMENT ON Table veranstaltungen IS 'Miniwelt LWBadventure';


-- minigames ($gameNr, gameName, !vNr)
CREATE TABLE minigames(
  gameNr 		INTEGER				NOT NULL,	-- oder CHECK (vNr > 0),
  gameName		VARCHAR (50)		NOT NULL,
  vNr			INTEGER				REFERENCES veranstaltungen (vNr),	-- vNr = Veranstaltungsnummer
  CONSTRAINT minigameKEY PRIMARY KEY (gameNr)
);
COMMENT ON Table minigames IS 'Miniwelt LWBadventure';


-- raeume ($raumNr, raumName, ort, funktion)
CREATE TABLE raeume (
  raumNr 		INTEGER				NOT NULL,
  raumName		VARCHAR (50)		NOT NULL,
  ort			VARCHAR (50)		NOT NULL,
  funktion		VARCHAR (50)		NOT NULL,							-- jeder Raum hat eine Funktion
  CONSTRAINT raumKEY PRIMARY KEY (raumNr)
);
COMMENT ON Table raeume IS 'Miniwelt LWBadventure';


-- spieler_innen ($spNr, spName, schluesselanzahl, !raumNr)
CREATE TABLE spieler_innen (
  spNr	 		  	INTEGER				NOT NULL,
  spName 		  	VARCHAR (50)		NOT NULL,
  schuesselanzahl	INTEGER				NOT NULL,
  raumNr 		 	INTEGER 			REFERENCES raeume (raumNr) DEFAULT 1,
  CONSTRAINT spieler_innenKEY PRIMARY KEY (spNr)
);
COMMENT ON Table spieler_innen IS 'Miniwelt LWBadventure';


-- unterricht (!$vNr, !npcNr, !raumNr)							
CREATE TABLE unterricht (
  vNr 			INTEGER			REFERENCES veranstaltungen (vNr),
  npcNr			INTEGER			REFERENCES dozent_innen (npcNr),
  raumNr		INTEGER			REFERENCES raeume (raumNr),
  CONSTRAINT unterrichtRaumCHECK CHECK (raumNr BETWEEN 1 AND 4)
);
COMMENT ON Table unterricht IS 'Miniwelt LWBadventure';


-- Nur spezielle Notenformate
CREATE DOMAIN NOTEN 
	AS NUMERIC (2,1) 
	DEFAULT 6.0 
	CHECK (VALUE IN (1.0,1.3,1.7,
			2.0,2.3,2.7,
			3.0,3.3,3.7,
			4.0,4.3,4.7,
			5.0,5.3,5.7,
			6.0));


-- spielstaende (!$gameNr, !$spNr, Note, Punkte)
CREATE TABLE spielstaende (
  gameNr 		INTEGER			REFERENCES minigames (gameNr),
  spNr			INTEGER			REFERENCES spieler_innen (spNr),
  note			NOTEN			,-- nur spezielle Noten
  punkte		INTEGER			NOT NULL
);
COMMENT ON Table spielstaende IS 'Miniwelt LWBadventure';


-- aufenthaltsorte (!$npcNr, !$raumNr)
CREATE TABLE aufenthaltsorte (
  npcnr 		INTEGER			REFERENCES sonstigeNPCs (npcnr),				-- Nummer des sonstigen NPCs im Raum
  raumNr		INTEGER			REFERENCES raeume (raumNr)
);
COMMENT ON Table aufenthaltsorte IS 'Miniwelt LWBadventure';


-- assistenz (!$vNr,!npcNr)
CREATE TABLE assistenz (
  vNr 			INTEGER			REFERENCES veranstaltungen (vNr),
  npcNr			INTEGER			REFERENCES dozent_innen (npcNr)
);
COMMENT ON Table assistenz IS 'Miniwelt LWBadventure';


