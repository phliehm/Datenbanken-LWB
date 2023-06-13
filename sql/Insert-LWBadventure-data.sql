--  Miniwelt: LWBadventure
--  Umsetzung für das Datenbankpraktikum der LWB Informatik
-- 
--  Annalena Cyriacus, Benjamin Schneider, Philipp Liehm, Martin Seiß
--
--  Start: 13.06.2023


-- npc(npcNr$,npcName)
INSERT INTO npc VALUES (1, 'Heidi');
INSERT INTO npc VALUES (2, 'Palim Palim');
INSERT INTO npc VALUES (10, 'Dark Schmidther');
INSERT INTO npc VALUES (11, 'Herk');
INSERT INTO npc VALUES (12, 'J.EthI');
INSERT INTO npc VALUES (13, 'FabFour');
INSERT INTO npc VALUES (14, 'Amoebi');
INSERT INTO npc VALUES (15, 'WtheK');


-- dozentIn(dNPCnr$!,lieblingsgetreank)
-- Folgende Getränkezuteilung wird verwendet:
-- ******************************************
--  Dark Schmidther	-	Extraschwarzer Kaffee
--	Herk			-	Melissentee
--  J.EthI			-	Kaffee mit Milch und 2x Zucker
--  FabFour			-	Cappuccino
--	Amoebi			-	Grüner Tee
--  WtheK			-	Bier
INSERT INTO dozentIn VALUES (10, 'Extraschwarzer Kaffee');
INSERT INTO dozentIn VALUES (11, 'Melissentee');
INSERT INTO dozentIn VALUES (12, 'Kaffee mit Milch und 2x Zucker');
INSERT INTO dozentIn VALUES (13, 'Cappuccino');
INSERT INTO dozentIn VALUES (14, 'Grüner Tee');
INSERT INTO dozentIn VALUES (15, 'Bier');


-- sNPC(sNPCnr$!,funktion)
INSERT INTO sNPC VALUES (1, 'Leitung'); 
INSERT INTO sNPC VALUES (2, 'Palim-Palim rufen!'); 


-- raum(raumNr$, raumName, ort)
INSERT INTO raum VALUES (0, 'Main Floor', 'irgendwo');
INSERT INTO raum VALUES (1, 'Raum 1', 'FU Berlin');
INSERT INTO raum VALUES (2, 'Raum 2', 'Home Office');
INSERT INTO raum VALUES (3, 'Raum 3', 'FU berlin');
INSERT INTO raum VALUES (4, 'Raum 4', 'STEPS');
INSERT INTO raum VALUES (5, 'Zeugnisraum', 'N.N.');


-- Herausgenommen: kursraum(kRaumNr$!) einfacher und konsitenter
--INSERT INTO kursraum VALUES (1);
--INSERT INTO kursraum VALUES (2);
--INSERT INTO kursraum VALUES (3);
--INSERT INTO kursraum VALUES (4);


-- sRaum(sRaumNr$!,sFunktion)
INSERT INTO sRaum VALUES (0, 'Zugang zu Kursräume');
INSERT INTO sRaum VALUES (5, 'Zeugnisvergabe');


-- veranstaltung(vNr$, vName, sws, thema, kuerzel, vDozNr!, vRaumNr!)
-- SWS FP, RS ?????????? check mal Annalena
INSERT INTO veranstaltung VALUES (1, 'Funktionale Programmierung',100,'Programmierung','FP',13,1);
INSERT INTO veranstaltung VALUES (2, 'Technische Informatik',100,'Theoretische und technische Informatik','RS',15,1);


-- spielerIn(spielerNr$, sName, sRaumNr)
INSERT INTO spielerIn VALUES (1, 'Annalena',1);  -- sRaumNr = 1 default
INSERT INTO spielerIn VALUES (2, 'Maddin',1);


-- minigame (gameNr$, gameName, gameVnr!)
INSERT INTO minigame VALUES (1, 'Muster-Spiel',1);
INSERT INTO minigame VALUES (2, 'Bauelemente-Spiel',2);


-- spielstand(sGameNr!, sSpielerNr!, Note, Punktzahl)
INSERT INTO spielstand VALUES (1, 1, 1.3, 325);
INSERT INTO spielstand VALUES (1, 2, 1.7, 225);
INSERT INTO spielstand VALUES (2, 1, 1.0, 475);
INSERT INTO spielstand VALUES (2, 2, 2.7, 125);


-- aufenthalt(asNPCnr!,aRaumNr!)
INSERT INTO aufenthalt VALUES (1, 1);
INSERT INTO aufenthalt VALUES (1, 0);
INSERT INTO aufenthalt VALUES (1, 4);
INSERT INTO aufenthalt VALUES (2, 0);





--INSERT INTO Professoren VALUES (2126, 'Russel', 'C4', 5700.00, 232);
