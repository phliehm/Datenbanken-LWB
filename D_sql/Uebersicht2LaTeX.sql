
-- Skript gibt alle Realtionen zur Miniwelt LWBadventure aus:

\pset format latex

--\o LWBadventure_npcs.tex			
-- > hier werden die Ergebniss hingeschrieben

\C 'NPCs'
select npcnr as "NPC-Nr.",npcname as "NPC Name" from npcs;

--\o							

\C 'sonstige NPCs'
select npcnr as "NPC-Nr.", aufgabe as "Aufgabe" from sonstigenpcs;

\C 'DozentInnen'
select npcnr as "NPC-Nr.", lieblingsgetraenk as "Lieblingsgetränk" from dozent_Innen;

\C 'Unterricht'
select vnr as "Veranstaltungssnr..", npcnr as "NPC-Nr.", raumnr as "Raumnr."from unterricht;

\C 'Assistenz'
select vnr as "Veranstaltungssnr..", npcnr as "NPC-Nr." from assistenz;

\C 'Veranstaltungen'
select vnr as "Veranstaltungssnr..", vname as "Veranstaltungsname", kuerzel as "Kürzel", sws as "SWS", semester as "Semester", gebietnr as "Themengebietsnr."  from veranstaltungen;

\C 'Themengebiete'
select gebietnr as "Themengebietsnr.", gebietname as "Themengebietsname" from themengebiete;

\C 'Minigames'
select gamenr as "Gamenr.",gamename as "Gamename", vnr as "Veranstaltungsnr." from minigames;

\C 'SpielerInnen'
select spnr as "Spieler_Innen Nr.", spname as "Spieler_Innen Name", schuesselanzahl as "Schüsselanzahl", raumnr as "Raumnr." from spieler_Innen;

\C 'Räume'
select raumnr as "Raumnr.", raumname as "Raumname", ort as "Ort", funktion as "Funktion" from raeume;

\C 'Spielstände'
 select gamenr as "Gamenr.", spnr as "Spieler_Innen-Nr.", note as "Note", punkte as "Punkte" from spielstaende;

\C 'Aufenthaltsorte'
select npcnr as "NPC-Nr.", raumnr as "Raumnr." from aufenthaltsorte;
