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
--
-- npc(npcNr$,npcName)
-- dozentIn(dozNr$!,lieblingsgetreank)
-- sNPC(sNPCnr$!,funktion)
-- raum(raumNr$, raumName, ort)
--     --- wird Herausgenommen: kursraum(kRaumNr$!) einfacher und konsitenter
-- sRaum(kRaumNr$!,sFunktion)
-- veranstaltung(vNr$, vName, sws, thema, kuerzel, vDozNr!, semester!)
-- spielerIn(spielerNr$, sName, sRaumNr)
-- minigame(gameNr$, gameName, gameVnr!)
-- sRaum()
-- spielstand(sGameNr!, sSpielerNr!, Note, Punktzahl)
-- aufenthalt(asNPCnr!,aRaumNr!)


-- NPC - Non-Playing-Character (Mitarbeiter):  npc(npcNr$,npcName)
CREATE TABLE npc (
  npcNr 		INTEGER				NOT NULL,	
  npcName 		VARCHAR (50)		NOT NULL,
  CONSTRAINT npcKEY PRIMARY KEY (npcNr)
);
COMMENT ON Table npc IS 'Miniwelt LWBadventure';


-- dozentIn(dozNr$!,lieblingsgetreank)
CREATE TABLE dozentIn (
  dozNr 				INTEGER 			REFERENCES npc (npcNr),		-- NPC-Nummer
  lieblingsgetreank		VARCHAR (100)		NOT NULL,	-- jeder Dozent braucht ein Lieblingsgetränk
  CONSTRAINT dozentInKEY PRIMARY KEY (dozNr)
);
COMMENT ON Table dozentIn IS 'Miniwelt LWBadventure';


-- sNPC(snpcNr$!,funktion)
CREATE TABLE sNPC (
  sNPCnr 		INTEGER 			REFERENCES npc (npcNr),		-- NPC-Nummer
  funktion		VARCHAR (100)		NOT NULL,	-- jeder Mitarbeiter hat eine Funktion
  CONSTRAINT sNPCKEY PRIMARY KEY (sNPCnr)
);
COMMENT ON Table sNPC IS 'Miniwelt LWBadventure';


-- raum(raumNr$, raumName, ort)
CREATE TABLE raum (
  raumNr 		INTEGER				NOT NULL,
  raumName		VARCHAR (50)		NOT NULL,
  ort			VARCHAR (50)		NOT NULL,
  CONSTRAINT raumKEY PRIMARY KEY (raumNr)
);
COMMENT ON Table raum IS 'Miniwelt LWBadventure';


-- Herausgenommen: kursraum(kRaumNr$!) einfacher und konsitenter
-- nicht kursraum(kRaumNr$!,semester) da kRaumNr$! == semester
--CREATE TABLE kursraum (
--  kRaumNr 		INTEGER 			REFERENCES raum (raumNr),		-- Raumnummer
--  CONSTRAINT kRaumNrCHECK CHECK (kRaumNr BETWEEN 1 AND 4),
--  CONSTRAINT kursraumKEY PRIMARY KEY (kRaumNr)
--);
--COMMENT ON Table kursraum IS 'Miniwelt LWBadventure';


-- sRaum(sRaumNr$!,sFunktion)
CREATE TABLE sRaum (
  sRaumNr 		INTEGER 			REFERENCES raum (raumNr),		-- Raumnummer
  sFunktion		VARCHAR (100)		NOT NULL,	-- jeder Raum hat eine Funktion
  CONSTRAINT sRaumNrCHECK CHECK (sRaumNr = 0 OR sRaumNr > 4),	-- 1 bis 4 für Kursräume = Semester
  CONSTRAINT sRaumKEY PRIMARY KEY (sRaumNr)
);
COMMENT ON Table sRaum IS 'Miniwelt LWBadventure';


-- veranstaltung(vNr$, vName, sws, thema, kuerzel, vDozNr!, vRaumNr!)
CREATE TABLE veranstaltung (
  vNr 			INTEGER				NOT NULL,	-- Veranstaltungsnummer
  vName			VARCHAR (50)		NOT NULL,
  sws			INTEGER				CHECK (sws > 0),
  thema			VARCHAR (50)		NOT NULL,
  kuerzel		VARCHAR (5)			NOT NULL,
  vDozNr		INTEGER				REFERENCES dozentIn (dozNr),
  semester 		INTEGER 			REFERENCES raum (raumNr),		-- Raumnummer: 1,2,3,4
  CONSTRAINT vRaumNrCHECK CHECK (semester BETWEEN 1 AND 4),
--  gelesenVon INTEGER REFERENCES Professoren (PersNr)
--    ON DELETE SET NULL
--    ON UPDATE CASCADE,
  CONSTRAINT veranstaltungKEY PRIMARY KEY (vNr)
);
COMMENT ON Table veranstaltung IS 'Miniwelt LWBadventure';


-- spielerIn(spielerNr$, sName, sRaumNr)
CREATE TABLE spielerIn (
  spielerNr 	INTEGER				NOT NULL, 	--CHECK (MatrNr BETWEEN 10000 AND 99999),
  sName 		VARCHAR (50)		NOT NULL,
  sRaumNr 		INTEGER 			REFERENCES raum (raumNr) DEFAULT 1,
   -- NOT NULL CHECK (Semester > 0 AND Semester < 20) DEFAULT 1,
  CONSTRAINT spielerInKEY PRIMARY KEY (spielerNr)
--  CONSTRAINT DEFAULTraum DEFAULT 1
);
COMMENT ON Table SpielerIn IS 'Miniwelt LWBadventure';


-- minigame (gameNr$, gameName, gameVnr!)
CREATE TABLE minigame (
  gameNr 		INTEGER				NOT NULL,	-- oder CHECK (vNr > 0),
  gameName		VARCHAR (50)		NOT NULL,
  gameVnr		INTEGER				REFERENCES veranstaltung (vNr),	-- vNr = Veranstaltungsnummer
  CONSTRAINT minigameKEY PRIMARY KEY (gameNr)
);
COMMENT ON Table minigame IS 'Miniwelt LWBadventure';


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


-- spielstand(sGameNr!, sSpielerNr!, Note, Punktzahl)
CREATE TABLE spielstand (
  sGameNr 		INTEGER			REFERENCES minigame (gameNr),
  sSpielerNr	INTEGER			REFERENCES spielerIn (spielerNr),
  note			NOTEN			,-- nur spezielle Noten
  punkte		INTEGER			NOT NULL
);
COMMENT ON Table spielstand IS 'Miniwelt LWBadventure';


-- aufenthalt(asNPCnr!,aRaumNr!)
CREATE TABLE aufenthalt (
  asNPCnr 		INTEGER			REFERENCES sNPC (sNPCnr),	-- Nummer des sonstigen NPCs im Raum
  aRaumNr		INTEGER			REFERENCES raum (raumNr)
);
COMMENT ON Table aufenthalt IS 'Miniwelt LWBadventure';
