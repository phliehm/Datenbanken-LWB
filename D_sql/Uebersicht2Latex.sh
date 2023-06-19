
## zum Erstellen einer Übersicht der Tabellen

## Martin Seiß		16.6.2023


psql -c '\i Uebersicht2LaTeX.sql' -o ../A_doc/LWBadventure-Daten.tex

# alternative
#psql -c '\i Uebersicht.sql' > ../A/LWBadventure-Daten.txt
