
-- Skript gibt alle Realtionen zur Miniwelt LWBadventure 
-- als einzelne CSV-Dateien aus

--\o LWBadventure_npcs.tex			
-- > hier werden die Ergebniss hingeschrieben


\copy (select npcnr as "NPC-Nr.",npcname as "NPC Name" from npcs) TO '../B_modell/CVS/LWBadventure-Daten-NPCs.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';

\copy (select npcnr as "NPC-Nr.", aufgabe as "Aufgabe" from sonstigenpcs) TO '../B_modell/CVS/LWBadventure-Daten-sonstige-NPCs.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';

\copy (select npcnr as "NPC-Nr.", lieblingsgetraenk as "Lieblingsgetränk" from dozent_Innen) TO '../B_modell/CVS/LWBadventure-Daten-Dozent_Innen.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';

\copy (select vnr as "Veranstaltungsnr.", npcnr as "NPC-Nr.", raumnr as "Raumnr."from unterricht) TO '../B_modell/CVS/LWBadventure-Daten-Unterricht.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';

\copy (select vnr as "Veranstaltungsnr.", npcnr as "NPC-Nr." from assistenz) TO '../B_modell/CVS/LWBadventure-Assistenz.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';

\copy (select vnr as "Veranstaltungsnr.", vname as "Veranstaltungsname", kuerzel as "Kürzel", sws as "SWS", semester as "Semester", gebietnr as "Themengebietsnr."  from veranstaltungen) TO '../B_modell/CVS/LWBadventure-Veranstaltungen.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';

\copy (select gebietnr as "Themengebietsnr.", gebietname as "Themengebietsname" from themengebiete) TO '../B_modell/CVS/LWBadventure-Themengebiete.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';

\copy (select gamenr as "Gamenr.",gamename as "Gamename", vnr as "Veranstaltungsnr." from minigames) TO '../B_modell/CVS/LWBadventure-Minigames.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';

\copy (select spnr as "Spieler_Innen Nr.", spname as "Spieler_Innen Name", schuesselanzahl as "Schüsselanzahl", raumnr as "Raumnr." from spieler_Innen) TO '../B_modell/CVS/LWBadventure-Daten-Spieler_Innen.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';

\copy (select raumnr as "Raumnr.", raumname as "Raumname", ort as "Ort", funktion as "Funktion" from raeume) TO '../B_modell/CVS/LWBadventure-Daten-Räume.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';

\copy (select gamenr as "Gamenr.", spnr as "Spieler_Innen-Nr.", note as "Note", punkte as "Punkte" from spielstaende) TO '../B_modell/CVS/LWBadventure-Daten-Spielstände.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';

\copy (select npcnr as "NPC-Nr.", raumnr as "Raumnr." from aufenthaltsorte) TO '../B_modell/CVS/LWBadventure-Daten-Aufenthaltsorte.csv' DELIMITER ';'  CSV HEADER ENCODING 'UTF8';
