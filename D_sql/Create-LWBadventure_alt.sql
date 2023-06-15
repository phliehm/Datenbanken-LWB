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

-- npcs(npcNr$,npcName) 				NPCs - Non-Playing-Character (Mitarbeiter) 
-- dozentInnen(npcNr !$,lieblingsgetraenk)
-- sNPCs(npcsNr $!,funktion)
-- raeum(raumNr $!, raumName, ort)
-- kursraeume(semster,raumNr !)
-- sRaeume(raumNr $!,sFunktion)										-- sinnvoll?!
-- themengebiete(themenNr $,themenname)
-- veranstaltungen(vNr$, vName, kuerzel, sws, themenNr !)
-- spielerInnen(spNr$, spName, schuesselanzahl, raumNr!)
-- minigames(gameNr$, gameName, vNr!)
-- spielstaende(gameNr!, spNr!, Note, Punktzahl)
-- aufenthalt(npcnr!,raumNr!)
-- unterricht(vNr !, npc !, raumNr !)							
-- assistenz(vNr !,npcNr !)



-- npcs(npcNr$,npcName) 		NPCs - Non-Playing-Character (Mitarbeiter) 
CREATE TABLE npcs (
  npcNr 		INTEGER				NOT NULL,	
  npcName 		VARCHAR (50)		NOT NULL,
  CONSTRAINT npcKEY PRIMARY KEY (npcNr)
);
COMMENT ON Table npcs IS 'Miniwelt LWBadventure';


-- dozentInnen(npcNr !$,lieblingsgetraenk)
CREATE TABLE dozentInnen (
  npcNr 				INTEGER 			REFERENCES npcs (npcNr),		-- NPC-Nummer
  lieblingsgetraenk		VARCHAR (100)		NOT NULL,	-- jeder Dozent braucht ein Lieblingsgetränk
  CONSTRAINT dozentInnenKEY PRIMARY KEY (npcNr)
);
COMMENT ON Table dozentInnen IS 'Miniwelt LWBadventure';


-- sNPCs(npcsNr $!,funktion)
CREATE TABLE sNPCs (
  npcNr 		INTEGER 			REFERENCES npcs (npcNr),		-- NPC-Nummer
  funktion		VARCHAR (100)		NOT NULL,	-- jeder Mitarbeiter hat eine Funktion
  CONSTRAINT sNPCKEY PRIMARY KEY (npcNr)
);
COMMENT ON Table sNPCs IS 'Miniwelt LWBadventure';


-- raeum(raumNr $!, raumName, ort)
CREATE TABLE raeume (
  raumNr 		INTEGER				NOT NULL,
  raumName		VARCHAR (50)		NOT NULL,
  ort			VARCHAR (50)		NOT NULL,
  CONSTRAINT raumKEY PRIMARY KEY (raumNr)
);
COMMENT ON Table raeume IS 'Miniwelt LWBadventure';


-- kursraeume(semster,raumNr !)
CREATE TABLE kursraeume (
  semester		INTEGER				NOT NULL DEFAULT 1,
  raumNr 		INTEGER 			REFERENCES raeume (raumNr),		-- Raumnummer
  
  CONSTRAINT RaumNrCHECK CHECK (raumNr BETWEEN 1 AND 4),
  CONSTRAINT semensterCHECK CHECK (semester BETWEEN 1 AND 4),
  CONSTRAINT kursraumKEY PRIMARY KEY (raumNr)
);
COMMENT ON Table kursraeume IS 'Miniwelt LWBadventure';


-- sRaeume(raumNr $!,sFunktion)										-- sinnvoll?!
CREATE TABLE sRaeume (
  raumNr 		INTEGER 			REFERENCES raeume (raumNr),		-- Raumnummer
  sFunktion		VARCHAR (100)		NOT NULL,	-- jeder Raum hat eine Funktion
  CONSTRAINT sRaumNrCHECK CHECK (raumNr = 0 OR raumNr > 4),	-- 1 bis 4 für Kursräume = Semester
  CONSTRAINT sRaumKEY PRIMARY KEY (raumNr)
);
COMMENT ON Table sRaeume IS 'Miniwelt LWBadventure';


-- themengebiete(themenNr $,themenname)
CREATE TABLE themengebiete (
  themenNr 		INTEGER				NOT NULL,
  themenname	VARCHAR (100)		NOT NULL,
  CONSTRAINT themengebieteKEY PRIMARY KEY (themenNr)
);
COMMENT ON Table themengebiete IS 'Miniwelt LWBadventure';


-- veranstaltungen(vNr$, vName, kuerzel, sws, themenNr !)
CREATE TABLE veranstaltungen (
  vNr 			INTEGER				NOT NULL,	-- Veranstaltungsnummer
  vName			VARCHAR (50)		NOT NULL,
  kuerzel		VARCHAR (5)			NOT NULL,
  sws			INTEGER				CHECK (sws > 0),
  themenNr		INTEGER				REFERENCES themengebiete (themenNr),
--  CONSTRAINT vRaumNrCHECK CHECK (semester BETWEEN 1 AND 4),
--  gelesenVon INTEGER REFERENCES Professoren (PersNr)
--    ON DELETE SET NULL
--    ON UPDATE CASCADE,
  CONSTRAINT veranstaltungKEY PRIMARY KEY (vNr)
);
COMMENT ON Table veranstaltungen IS 'Miniwelt LWBadventure';


-- spielerInnen(spNr$, spName, schuesselanzahl, raumNr!)
CREATE TABLE spielerInnen (
  spNr	 		  	INTEGER				NOT NULL, 	--CHECK (MatrNr BETWEEN 10000 AND 99999),
  spName 		  	VARCHAR (50)		NOT NULL,
  schuesselanzahl	INTEGER				NOT NULL,
  raumNr 		 	INTEGER 			REFERENCES raeume (raumNr) DEFAULT 1,
  CONSTRAINT spielerInnenKEY PRIMARY KEY (spNr)
);
COMMENT ON Table spielerInnen IS 'Miniwelt LWBadventure';


-- minigames(gameNr$, gameName, vNr!)
CREATE TABLE minigames(
  gameNr 		INTEGER				NOT NULL,	-- oder CHECK (vNr > 0),
  gameName		VARCHAR (50)		NOT NULL,
  vNr			INTEGER				REFERENCES veranstaltungen (vNr),	-- vNr = Veranstaltungsnummer
  CONSTRAINT minigameKEY PRIMARY KEY (gameNr)
);
COMMENT ON Table minigames IS 'Miniwelt LWBadventure';


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


-- spielstaende(gameNr!, spNr!, Note, Punktzahl)
CREATE TABLE spielstaende (
  gameNr 		INTEGER			REFERENCES minigames (gameNr),
  spNr			INTEGER			REFERENCES spielerInnen (spNr),
  note			NOTEN			,-- nur spezielle Noten
  punkte		INTEGER			NOT NULL
);
COMMENT ON Table spielstaende IS 'Miniwelt LWBadventure';


-- aufenthalt(npcnr!,raumNr!)
CREATE TABLE aufenthalt (
  npcnr 		INTEGER			REFERENCES sNPCs (npcnr),	-- Nummer des sonstigen NPCs im Raum
  raumNr		INTEGER			REFERENCES raeume (raumNr)
);
COMMENT ON Table aufenthalt IS 'Miniwelt LWBadventure';


-- unterricht(vNr !, npc !, raumNr !)							
CREATE TABLE unterricht (
  vNr 			INTEGER			REFERENCES veranstaltungen (vNr),
  npcNr			INTEGER			REFERENCES dozentInnen (npcNr),
  raumNr		INTEGER			REFERENCES kursraeume (raumNr)
);
COMMENT ON Table unterricht IS 'Miniwelt LWBadventure';


-- assistenz(vNr !,npcNr !)
CREATE TABLE assistenz (
  vNr 			INTEGER			REFERENCES veranstaltungen (vNr),
  npcNr			INTEGER			REFERENCES dozentInnen (npcNr)
);
COMMENT ON Table assistenz IS 'Miniwelt LWBadventure';

