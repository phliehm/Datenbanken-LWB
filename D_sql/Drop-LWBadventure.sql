--  Miniwelt: LWBadventure
--  Umsetzung für das Datenbankpraktikum der LWB Informatik
-- 
--  Annalena Cyriacus, Benjamin Schneider, Philipp Liehm, Martin Seiß
--
--  Start: 13.06.2023

DROP TABLE standorte			CASCADE;
DROP TABLE spielstaende      	CASCADE;
DROP TABLE unterricht			CASCADE;
DROP TABLE assistenz			CASCADE;
DROP TABLE spieler_innen  	  	CASCADE;
DROP TABLE minigames      		CASCADE;
DROP TABLE veranstaltungen    	CASCADE;
DROP TABLE themengebiete		CASCADE;
--DROP TABLE kursraeume  			CASCADE;
--DROP TABLE sRaeume  			CASCADE;
DROP TABLE raeume    			CASCADE;
DROP TABLE sNPCs         		CASCADE;
DROP TABLE dozent_innen        	CASCADE;
DROP TABLE npcs		      		CASCADE;
--DROP TABLE kursraum  		CASCADE;

DROP DOMAIN NOTEN;

