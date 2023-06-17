
## zum Erstellen einer Übersicht der Tabellen

## Martin Seiß		16.6.2023


psql -c '\i Uebersicht.sql' -o ../B_modell/LWBadventure-Daten.txt

# alternative
#psql -c '\i Uebersicht.sql' > ../B_modell/LWBadventure-Daten.txt
