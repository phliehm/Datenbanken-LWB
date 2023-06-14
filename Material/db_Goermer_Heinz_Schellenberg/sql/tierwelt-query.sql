SELECT * FROM Unterarten;

SELECT Name FROM Tiere;

SELECT Name FROM Tiere WHERE Geschlecht = 'W';

SELECT * FROM Tiere NATURAL JOIN Unterarten;

SELECT Name FROM Tiere NATURAL JOIN (SELECT * FROM Unterarten WHERE Bezeichnung = 'Bartagame') AS xyz;

SELECT Name FROM Tiere NATURAL JOIN Terrarien NATURAL JOIN (SELECT * FROM Habitattypen WHERE Bezeichnung = 'Wüste') AS xyz;

SELECT Faelligkeit FROM Tierpflege NATURAL JOIN Arbeitsdienste NATURAL JOIN (SELECT AT_ID FROM Arbeitsdiensttypen WHERE Bezeichnung LIKE 'Fütterung%') AS x NATURAL JOIN (SELECT TI_ID FROM Tiere WHERE Name = 'Exo') AS y;

SELECT Bezeichnung, Beschreibung FROM Arbeitsdiensttypen NATURAL JOIN (SELECT * FROM Arbeitsdienste WHERE DATE_PART('week',Faelligkeit) = DATE_PART('week', CURRENT_DATE)) AS xyz;

SELECT Faelligkeit FROM (SELECT TR_ID FROM Tiere WHERE Name = 'Exo') AS x NATURAL JOIN Terrarien NATURAL JOIN Verbrauch NATURAL JOIN Arbeitsdienste NATURAL JOIN (SELECT AT_ID FROM Arbeitsdiensttypen WHERE Bezeichnung = 'Frischwasser') AS y;

SELECT * FROM Terrarien NATURAL JOIN (SELECT TR_ID, COUNT(*) FROM (SELECT TR_ID, AD_ID FROM Verbrauch UNION SELECT TR_ID,AD_ID FROM Terrariumpflege) AS x GROUP BY TR_ID HAVING COUNT(*) > 1) as y;
