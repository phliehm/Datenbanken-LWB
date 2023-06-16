
-- Beispiel für Abfrageausgabe

\o AbfrageBsp.txt			-- > hier werden die Ergebniss hingeschrieben

\C 'select * from npcs natural join dozent_innen;'
select * from npcs natural join dozent_innen;

\C 'select spname,note,punkte  from minigames natural join spielstaende natural join spieler_innen where gamename like 'theNETgame' order by note,punkte;'
select spname,note,punkte  from minigames natural join spielstaende natural join spieler_innen where gamename like 'theNETgame' order by note,punkte;

\o							-- > Ende des Schreibens in das File
