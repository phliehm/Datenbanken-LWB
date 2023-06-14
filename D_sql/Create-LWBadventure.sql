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
-- spielerInnen(spNr$, spName, spRaumNr)
-- minigame(gameNr$, gameName, gameVnr!)
-- sRaum()
-- spielstand(sGameNr!, sSpNr!, Note, Punktzahl)
-- aufenthalt(asNPCnr!,aRaumNr!)


-- NPCs - Non-Playing-Character (Mitarbeiter):  npcs(npcNr$,npcName)
CREATE TABLE npcs (
  npcNr 		INTEGER				NOT NULL,	
  npcName 		VARCHAR (50)		NOT NULL,
  CONSTRAINT npcKEY PRIMARY KEY (npcNr)
);
COMMENT ON Table npcs IS 'Miniwelt LWBadventure';


-- dozentInnen(lieblingsgetreank)
CREATE TABLE dozentInnen (
  npcNr 				INTEGER 			REFERENCES npcs (npcNr),		-- NPC-Nummer
  lieblingsgetreank		VARCHAR (100)		NOT NULL,	-- jeder Dozent braucht ein Lieblingsgetränk
  CONSTRAINT dozentInnenKEY PRIMARY KEY (npcNr)
);
COMMENT ON Table dozentInnen IS 'Miniwelt LWBadventure';


-- sNPCs(snpcsNr$!,funktion)
CREATE TABLE sNPCs (
  npcNr 		INTEGER 			REFERENCES npcs (npcNr),		-- NPC-Nummer
  funktion		VARCHAR (100)		NOT NULL,	-- jeder Mitarbeiter hat eine Funktion
  CONSTRAINT sNPCKEY PRIMARY KEY (npcNr)
);
COMMENT ON Table sNPCs IS 'Miniwelt LWBadventure';


-- raeum(raeumeNr$, raumName, ort)
CREATE TABLE raeume (
  raumNr 		INTEGER				NOT NULL,
  raumName		VARCHAR (50)		NOT NULL,
  ort			VARCHAR (50)		NOT NULL,
  CONSTRAINT raumKEY PRIMARY KEY (raumNr)
);
COMMENT ON Table raeume IS 'Miniwelt LWBadventure';


-- Herausgenommen: kursraum(kRaumNr$!) einfacher und konsitenter
-- nicht kursraum(kRaumNr$!,semester) da kRaumNr$! == semester
--CREATE TABLE kursraum (
--  kRaumNr 		INTEGER 			REFERENCES raum (raumNr),		-- Raumnummer
--  CONSTRAINT kRaumNrCHECK CHECK (kRaumNr BETWEEN 1 AND 4),
--  CONSTRAINT kursraumKEY PRIMARY KEY (kRaumNr)
--);
--COMMENT ON Table kursraum IS 'Miniwelt LWBadventure';


-- sRaeume(sRaumNr$!,sFunktion)										-- sinnvoll?!
CREATE TABLE sRaeume (
  raumNr 		INTEGER 			REFERENCES raeume (raumNr),		-- Raumnummer
  sFunktion		VARCHAR (100)		NOT NULL,	-- jeder Raum hat eine Funktion
  CONSTRAINT sRaumNrCHECK CHECK (raumNr = 0 OR raumNr > 4),	-- 1 bis 4 für Kursräume = Semester
  CONSTRAINT sRaumKEY PRIMARY KEY (raumNr)
);
COMMENT ON Table sRaeume IS 'Miniwelt LWBadventure';


-- veranstaltungen(vNr$, vName, sws, thema, kuerzel, vDozNr!, vRaumNr!)
CREATE TABLE veranstaltungen (
  vNr 			INTEGER				NOT NULL,	-- Veranstaltungsnummer
  vName			VARCHAR (50)		NOT NULL,
  sws			INTEGER				CHECK (sws > 0),
  thema			VARCHAR (100)		NOT NULL,
  kuerzel		VARCHAR (5)			NOT NULL,
  npcNr			INTEGER				REFERENCES dozentInnen (npcNr),
  semester 		INTEGER 			REFERENCES raeume (raumNr),		-- Raumnummer: 1,2,3,4
  CONSTRAINT vRaumNrCHECK CHECK (semester BETWEEN 1 AND 4),
--  gelesenVon INTEGER REFERENCES Professoren (PersNr)
--    ON DELETE SET NULL
--    ON UPDATE CASCADE,
  CONSTRAINT veranstaltungKEY PRIMARY KEY (vNr)
);
COMMENT ON Table veranstaltungen IS 'Miniwelt LWBadventure';


-- betreuungen(vNr!,dozNr!)											-- statt sRaum?!
CREATE TABLE betreuung (
  vNr 			INTEGER			REFERENCES veranstaltungen (vNr),
  npcNr			INTEGER			REFERENCES dozentInnen (npcNr)
);
COMMENT ON Table betreuung IS 'Miniwelt LWBadventure';


-- spielerInnen(spNr$, spName, spRaumNr!)
CREATE TABLE spielerInnen (
  spNr	 		INTEGER				NOT NULL, 	--CHECK (MatrNr BETWEEN 10000 AND 99999),
  spName 		VARCHAR (50)		NOT NULL,
  raumNr 		INTEGER 			REFERENCES raeume (raumNr) DEFAULT 1,
   -- NOT NULL CHECK (Semester > 0 AND Semester < 20) DEFAULT 1,
  CONSTRAINT spielerInnenKEY PRIMARY KEY (spNr)
--  CONSTRAINT DEFAULTraum DEFAULT 0
);
COMMENT ON Table spielerInnen IS 'Miniwelt LWBadventure';


-- minigames(gameNr$, gameName, gameVnr!)
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


-- spielstaende(sGameNr!, sSpNr!, Note, Punktzahl)
CREATE TABLE spielstaende (
  gameNr 		INTEGER			REFERENCES minigames (gameNr),
  spNr			INTEGER			REFERENCES spielerInnen (spNr),
  note			NOTEN			,-- nur spezielle Noten
  punkte		INTEGER			NOT NULL
);
COMMENT ON Table spielstaende IS 'Miniwelt LWBadventure';


-- aufenthalt(asNPCnr!,aRaumNr!)
CREATE TABLE aufenthalt (
  npcnr 		INTEGER			REFERENCES sNPCs (npcnr),	-- Nummer des sonstigen NPCs im Raum
  raumNr		INTEGER			REFERENCES raeume (raumNr)
);
COMMENT ON Table aufenthalt IS 'Miniwelt LWBadventure';
