--  Miniwelt: LWBadventure
--  Umsetzung für das Datenbankpraktikum der LWB Informatik
-- 
--  Annalena Cyriacus, Benjamin Schneider, Philipp Liehm, Martin Seiß
--
--  Start: 13.06.2023

DROP TABLE betreuung			CASCADE;
DROP TABLE aufenthalt			CASCADE;
DROP TABLE spielstaende      	CASCADE;
DROP TABLE spielerInnen  	  	CASCADE;
DROP TABLE minigames      		CASCADE;
DROP TABLE veranstaltungen    	CASCADE;
DROP TABLE sRaeume  			CASCADE;
DROP TABLE raeume    			CASCADE;
DROP TABLE sNPCs         		CASCADE;
DROP TABLE dozentInnen        	CASCADE;
DROP TABLE npcs		      		CASCADE;
--DROP TABLE kursraum  		CASCADE;

DROP DOMAIN NOTEN;
