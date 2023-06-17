
-- SQL-ANFRAGEN
-- ------------

-- 1a. Welche Räume gibt es in der LWB-Adventure-World?
SELECT * FROM raeume;

-- 1b. Welche Aufgaben haben die sonstigen NPCs im LWB-Adventure?
SELECT aufgabe FROM sonstigenpcs;


-- 2a. Welche Lehrveranstaltungen haben 6 SWS?
SELECT * FROM veranstaltungen WHERE sws = 6;

-- 2b. Welche Lehrveranstaltungen gibt es im 4. Semester?
SELECT * FROM veranstaltungen WHERE semester = 4;

-- 2c. Welche Minigames gibt es im 4. Semester?
SELECT * FROM minigames NATURAL JOIN veranstaltungen WHERE semester = 4;


-- 3a. Wie heißen die Spieler_innen, die bisher das LWB-Adventure gespielt haben?
SELECT spname FROM spieler_innen;

-- 3b. Wie heißen die Dozenten im LWB-Adventure?
SELECT npcname FROM dozent_innen NATURAL JOIN npcs;

-- 3c. Welche Aufgabe hat NPC 'Heidi'?
SELECT aufgabe FROM sonstigenpcs NATURAL JOIN npcs WHERE npcname = 'Heidi';


-- 4a. Welche Lehrveranstaltungen gehören zum Themengebiet 'Programmierung'?
SELECT * FROM veranstaltungen NATURAL JOIN themengebiete WHERE gebietname = 'Programmierung';

-- 4b. Welche Lehrveranstaltungen haben etwas mit 'Daten' oder 'Programmierung' zu tun?
SELECT * FROM veranstaltungen WHERE vname LIKE '%Daten%' OR vname LIKE '%Programmierung%';

-- oder
SELECT * FROM veranstaltungen NATURAL JOIN themengebiete WHERE gebietname LIKE '%Daten%' OR gebietname LIKE '%Programmierung%';


-- 5. Was ist das Lieblingsgetränk von Darth Schmidter?
SELECT lieblingsgetraenk FROM dozent_innen NATURAL JOIN npcs WHERE npcname = 'Darth Schmidter';


-- 6. Welche Lehrveranstaltungen finden nicht in der 'FU Berlin' statt?
SELECT vname, semester, ort FROM raeume NATURAL JOIN unterricht NATURAL JOIN veranstaltungen WHERE ort != 'FU Berlin';

-- alternativ (falls man nicht sicher ist, ob 'FU Berlin' die vollständige Orts-Bezeichnung ist:
SELECT vname, semester, ort FROM raeume NATURAL JOIN unterricht NATURAL JOIN veranstaltungen WHERE ort NOT LIKE '%FU%Berlin%';


-- 7. Welche Dozenten sind in der LWB nur leitend tätig und machen keine Assistenz?
SELECT npcname FROM npcs NATURAL JOIN (SELECT npcnr FROM dozent_innen EXCEPT SELECT npcnr FROM assistenz) AS xyz;
-- Kommentar: 	Hier braucht es einen Alias, damit der NATURAL JOIN mit der Unterabfrage funktioniert.
--				Die Bezeichnung ist jedoch egal, da nur auf den NPCNamen projiziert wird.


-- Anfragen, die nur mit erweiterter relationaler Algebra beschrieben werden können:
-- ---------------------------------------------------------------------------------

-- 8. Wieviele Mini-Games gibt es in der LWB-Adventure-World? (Ausgaben-Titel: AnzahlMinigames)
SELECT COUNT(*) AS AnzahlMinigames FROM minigames;


-- 9. Wieviele SWS müssen in der LWB ingesamt absolviert werden? (Ausgaben-Titel: GesamtanzahlSWS)
SELECT SUM(sws) AS GesamtanzahlSWS FROM veranstaltungen;


-- 10. Wie heißt die Veranstaltung mit den meisten SWS?
SELECT vname FROM veranstaltungen WHERE sws = (SELECT MAX(sws) FROM veranstaltungen);


-- 11. Gesucht sind die Namen, Semester und SWS aller Veranstaltungen von Winnie the K absteigend sortiert nach SWS-Anzahl!
SELECT vname, sws, semester FROM veranstaltungen NATURAL JOIN unterricht NATURAL JOIN npcs WHERE npcname = 'Winnie the K' ORDER BY sws DESC;


-- 12. Wieviele Veranstaltungen gibt es pro Standort?
SELECT ort, COUNT(*) AS AnzahlVeranstaltungen FROM raeume NATURAL JOIN unterricht GROUP BY ort ORDER BY COUNT(*);


--13. Welche Spieler_innen haben einen Gesamt-Notendurchschnitt, der nicht zwischen 2.0 und 4.0 liegt? (Sortierung nach Gesamt-Notendurchschnitt aufsteigend, also bester Schnitt zuerst)
SELECT SpName FROM spieler_innen NATURAL JOIN spielstaende GROUP BY spname HAVING AVG(note) NOT BETWEEN 2.0 AND 4.0 ORDER BY AVG(note),spname;

-- oder mit Ausgabe des jeweiligen Notendurchschnitts:
SELECT SpName, ROUND(AVG(note),2) AS Notendurchschnitt FROM spieler_innen NATURAL JOIN spielstaende GROUP BY spname HAVING AVG(note) NOT BETWEEN 2.0 AND 4.0 ORDER BY AVG(note),spname;

