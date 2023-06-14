CREATE TABLE Habitattypen (
	HT_ID				INT NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	Bezeichnung			VARCHAR	NOT NULL,
	Beschreibung		VARCHAR,	
	Temperatur			INT,		/* in °C */
	Luftfeuchtigkeit	INT,		/* in % */
	Ausstattung			VARCHAR
) ;

COMMENT ON TABLE  Habitattypen					IS 'Informationen zu den jeweiligen Terrariumtypen';
COMMENT ON COLUMN Habitattypen.HT_ID			IS 'automatische Identifikationsnummer';
COMMENT ON COLUMN Habitattypen.Bezeichnung		IS 'Bezeichnung des Terrariumtypen';
COMMENT ON COLUMN Habitattypen.Beschreibung		IS 'Evtl. Beschreibung des Terrariumtypen';
COMMENT ON COLUMN Habitattypen.Temperatur		IS 'Temperatur des Terrariumtypen (in °C)';
COMMENT ON COLUMN Habitattypen.Luftfeuchtigkeit	IS 'Luftfeuchtigkeit des Terrariumtypen (in %)';
COMMENT ON COLUMN Habitattypen.Ausstattung		IS 'Evtl. Ausstattung des Terrariumtypen';



CREATE TABLE Terrarien (
	TR_ID			INT NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	Name			VARCHAR NOT NULL,
	Beschreibung	VARCHAR,
	Raumnummer		VARCHAR,
	Breite			INT,		/* in cm */
	Breite_vorn		INT,
	Tiefe			INT,		/* in cm */
	Hoehe			INT,		/* in cm */
	HT_ID			INT REFERENCES Habitattypen (HT_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE
) ;

COMMENT ON TABLE  Terrarien					IS 'Informationen zu den jeweiligen Terrarien';
COMMENT ON COLUMN Terrarien.TR_ID			IS 'automatische Identifikationsnummer';
COMMENT ON COLUMN Terrarien.Name			IS 'Bezeichnung des Terriums';
COMMENT ON COLUMN Terrarien.Beschreibung	IS 'Evtl. Beschreibung des Terrariums';
COMMENT ON COLUMN Terrarien.Raumnummer		IS 'Die Nummer des Raumes, in dem das Terrarium steht';
COMMENT ON COLUMN Terrarien.Breite			IS 'Breite des Terrariums (in cm)';
COMMENT ON COLUMN Terrarien.Breite_vorn		IS 'vordere Breite des Terrariums (in cm)';
COMMENT ON COLUMN Terrarien.Tiefe			IS 'Tiefe des Terrariums (in cm)';
COMMENT ON COLUMN Terrarien.Hoehe			IS 'Höhe des Terrariums (in cm)';
COMMENT ON COLUMN Terrarien.HT_ID			IS 'Verweis auf den Typen des Terraiums';



CREATE TABLE Unterarten (
	UA_ID				INT NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	Bezeichnung			VARCHAR NOT NULL,
	Heimat				VARCHAR,
	Lebensraum			VARCHAR,
	Temperatur			VARCHAR,
	Luftfeuchtigkeit	VARCHAR,
	Groesse				VARCHAR,
	Lebenserwartung		VARCHAR,
	Ernaehrung			VARCHAR,
	Besonderheiten		VARCHAR
) ;

COMMENT ON TABLE  Unterarten					IS 'Informationen zu den jeweiligen Tierarten';
COMMENT ON COLUMN Unterarten.UA_ID				IS 'automatische Identifikationsnummer';
COMMENT ON COLUMN Unterarten.Bezeichnung		IS 'Bezeichnung der Tierart';
COMMENT ON COLUMN Unterarten.Heimat				IS 'Wo lebt die Tierart?';
COMMENT ON COLUMN Unterarten.Lebensraum			IS 'Beschreibung des Lebensraums';
COMMENT ON COLUMN Unterarten.Temperatur			IS 'Bevorzugte Temperaturen (in °C)';
COMMENT ON COLUMN Unterarten.Luftfeuchtigkeit	IS 'Bevorzugte Luftfeuchtigkeit (in %)';
COMMENT ON COLUMN Unterarten.Groesse			IS 'Groesse im ausgewachsenen Zustand';
COMMENT ON COLUMN Unterarten.Lebenserwartung	IS 'Mittlere Lebenserwartung (in Jahren)';
COMMENT ON COLUMN Unterarten.Ernaehrung			IS 'Bevorzugte Ernaehrung';
COMMENT ON COLUMN Unterarten.Besonderheiten		IS 'Evtl. Besonderheiten der Tierart';




CREATE TABLE Tiere (
	TI_ID			INT NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	Name			VARCHAR NOT NULL,
	Geschlecht		CHAR,			/* m/w/d */
	Laenge			INT,			/* in cm */
	Gewicht			INT,			/* in g */
	Zugangsdatum	DATE,
	UA_ID			INT REFERENCES Unterarten (UA_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	TR_ID			INT REFERENCES Terrarien (TR_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE
) ;

COMMENT ON TABLE  Tiere					IS 'Informationen zu den jeweiligen Tieren';
COMMENT ON COLUMN Tiere.TI_ID			IS 'automatische Identifikationsnummer';
COMMENT ON COLUMN Tiere.Name			IS 'Name des Tieres';
COMMENT ON COLUMN Tiere.Geschlecht		IS 'Geschlecht des Tieres (m/w/d)';
COMMENT ON COLUMN Tiere.Laenge			IS 'Länge des Tieres mit Schwanz (in cm)';
COMMENT ON COLUMN Tiere.Gewicht			IS 'Gewicht des Tieres (in g)';
COMMENT ON COLUMN Tiere.Zugangsdatum	IS 'Seit wann ist das Tier in der Schule?';
COMMENT ON COLUMN Tiere.UA_ID			IS 'Verweis auf die Unterart';
COMMENT ON COLUMN Tiere.TR_ID			IS 'Verweis auf das Terrarium, in dem das Tier lebt';



CREATE TABLE Verbrauchsmittel (
	VM_ID			INT NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	Bezeichnung		VARCHAR NOT NULL,
	Beschreibung	VARCHAR,
	Lager_min		INT,
	Lager_akt		INT,
	Kosten			Float
);

COMMENT ON TABLE  Verbrauchsmittel					IS 'Informationen zu den jeweiligen Verbrauchsmitteln';
COMMENT ON COLUMN Verbrauchsmittel.VM_ID			IS 'automatische Identifikationsnummer';
COMMENT ON COLUMN Verbrauchsmittel.Bezeichnung		IS 'Bezeichnung des Verbrauchsmittels';
COMMENT ON COLUMN Verbrauchsmittel.Beschreibung		IS 'Evtl. Beschreibung des Verbrauchsmittels';
COMMENT ON COLUMN Verbrauchsmittel.Lager_min		IS 'Bestand, der minimal verfügbar sein und nicht unterschritten werden sollte (in Stück/g)';
COMMENT ON COLUMN Verbrauchsmittel.Lager_akt		IS 'aktueller Bestand (in Stück/g)';
COMMENT ON COLUMN Verbrauchsmittel.Kosten			IS 'Kosten zur Neuanschaffung pro Einheit (in €)';



CREATE TABLE Leuchtmittel (
	VM_ID			INT NOT NULL REFERENCES Verbrauchsmittel (VM_ID) PRIMARY KEY,
	Helligkeit		INT,
	Waermeleistung	INT
);

COMMENT ON TABLE  Leuchtmittel					IS 'Informationen zu den jeweiligen Leuchtmitteln';
COMMENT ON COLUMN Leuchtmittel.VM_ID			IS 'Verweis auf die ID der allgemeinen Verbrauchsmittel';
COMMENT ON COLUMN Leuchtmittel.Helligkeit		IS 'Helligkeit des Leuchtmittels (in ?)';
COMMENT ON COLUMN Leuchtmittel.Waermeleistung	IS 'Wärmeleistung des Leuchtmittels in ?)';



CREATE TABLE Bodenbelag (
	VM_ID		INT NOT NULL REFERENCES Verbrauchsmittel (VM_ID) PRIMARY KEY,
	Farbe		VARCHAR,
	Koernung	VARCHAR
);

COMMENT ON TABLE  Bodenbelag			IS 'Informationen zu dem jeweiligen Sand';
COMMENT ON COLUMN Bodenbelag.VM_ID		IS 'Verweis auf die ID der allgemeinen Verbrauchsmittel';
COMMENT ON COLUMN Bodenbelag.Farbe		IS 'Farbe des Sandes';
COMMENT ON COLUMN Bodenbelag.Koernung	IS 'Körnung des Sandes (in ?)';



CREATE TABLE Pflanzen (
	VM_ID				INT NOT NULL REFERENCES Verbrauchsmittel (VM_ID) PRIMARY KEY,
	Wasserverbrauch	INT,
	kuenstlich		BOOL,
	spruehen		BOOL
);

COMMENT ON TABLE  Pflanzen					IS 'Informationen zu den jeweiligen Pflanzen';
COMMENT ON COLUMN Pflanzen.VM_ID			IS 'Verweis auf die ID der allgemeinen Verbrauchsmittel';
COMMENT ON COLUMN Pflanzen.Wasserverbrauch	IS 'Wasserverbrauch der Pflanzen (in ml/Woche)';
COMMENT ON COLUMN Pflanzen.kuenstlich		IS 'Ist die Pflanze kuenstlich?';
COMMENT ON COLUMN Pflanzen.spruehen			IS 'Muss die Pflanze regelmaessig besprueht werden?';



CREATE TABLE Wasserbecken (
	VM_ID			INT NOT NULL REFERENCES Verbrauchsmittel (VM_ID) PRIMARY KEY,
	Hoehe		INT,
	Breite		INT,
	Tiefe		INT,
	bepflanzt	BOOL
);

COMMENT ON TABLE  Wasserbecken				IS 'Informationen zu den jeweiligen Wasserbecken';
COMMENT ON COLUMN Wasserbecken.VM_ID		IS 'Verweis auf die ID der allgemeinen Verbrauchsmittel';
COMMENT ON COLUMN Wasserbecken.Hoehe		IS 'Hoehe des Wasserbecken (in mm)';
COMMENT ON COLUMN Wasserbecken.Breite		IS 'Breite des Wasserbecken (in mm)';
COMMENT ON COLUMN Wasserbecken.Tiefe		IS 'Tiefe des Wasserbecken (in mm)';
COMMENT ON COLUMN Wasserbecken.bepflanzt	IS 'Ist das Becken bepflanzt?';



CREATE TABLE Deko (
	VM_ID		INT NOT NULL REFERENCES Verbrauchsmittel (VM_ID) PRIMARY KEY,
	Material	VARCHAR,
	Farbe		VARCHAR,
	Zweck		VARCHAR
);

COMMENT ON TABLE  Deko			IS 'Informationen zu der jeweiligen Deko';
COMMENT ON COLUMN Deko.VM_ID	IS 'Verweis auf die ID der allgemeinen Verbrauchsmittel';
COMMENT ON COLUMN Deko.Material	IS 'Material der Deko';
COMMENT ON COLUMN Deko.Farbe	IS 'Farbe der Deko';
COMMENT ON COLUMN Deko.Zweck	IS 'Zweckbeschreibung der Deko';



CREATE TABLE Arbeitsdiensttypen (
	AT_ID			INT NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	Bezeichnung		VARCHAR NOT NULL,
	Beschreibung	VARCHAR,
	Frequenz		INT
);

COMMENT ON TABLE  Arbeitsdiensttypen				IS 'Informationen zu den jeweiligen Arbeitseinsatztypen';
COMMENT ON COLUMN Arbeitsdiensttypen.AT_ID			IS 'automatische Identifikationsnummer';
COMMENT ON COLUMN Arbeitsdiensttypen.Bezeichnung	IS 'Bezeichnung des Arbeitseinsatzes';
COMMENT ON COLUMN Arbeitsdiensttypen.Beschreibung	IS 'Beschreibung des Arbeitseinsatzes';
COMMENT ON COLUMN Arbeitsdiensttypen.Frequenz		IS 'Mit welchem Zeitabstand soll der Einsatz wiederholt werden? (in Tagen)';




CREATE TABLE Arbeitsdienste (
	AD_ID				INT NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	AT_ID	INT REFERENCES Arbeitsdiensttypen (AT_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	Faelligkeit			DATE
);

COMMENT ON TABLE  Arbeitsdienste					IS 'Informationen zu den jeweiligen Arbeitseinsätzen';
COMMENT ON COLUMN Arbeitsdienste.AD_ID				IS 'automatische Identifikationsnummer';
COMMENT ON COLUMN Arbeitsdienste.AT_ID				IS 'Verweis auf den Arbeitseinsatztypen';
COMMENT ON COLUMN Arbeitsdienste.Faelligkeit		IS 'Datumsstempel der nächsten Fälligkeit';



CREATE TABLE Verbrauch (
	TR_ID			INT REFERENCES Terrarien (TR_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	VM_ID	INT REFERENCES Verbrauchsmittel (VM_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	AD_ID		INT REFERENCES Arbeitsdienste (AD_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	Anzahl				INT NOT NULL
);

COMMENT ON TABLE  Verbrauch			IS 'Welches Terrarium verfügt über welche Verbrauchsmittel';
COMMENT ON COLUMN Verbrauch.TR_ID	IS 'Verweis auf das Terrarium';
COMMENT ON COLUMN Verbrauch.VM_ID	IS 'Verweis auf das Verbrauchsmittel';
COMMENT ON COLUMN Verbrauch.AD_ID	IS 'Verweis auf den Arbeitseinsatz';
COMMENT ON COLUMN Verbrauch.Anzahl	IS 'Vorhandene Anzahl';



CREATE TABLE Futter (
	FU_ID				INT NOT NULL GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
	Bezeichnung			VARCHAR NOT NULL,
	Lagerungsform	VARCHAR
);

COMMENT ON TABLE  Futter				IS 'Informationen zu den jeweiligen Futtersorten';
COMMENT ON COLUMN Futter.FU_ID			IS 'automatische Identifikationsnummer';
COMMENT ON COLUMN Futter.Bezeichnung	IS 'Bezeichnung des Futters';
COMMENT ON COLUMN Futter.Lagerungsform	IS 'Lagerungsform des Futters';



CREATE TABLE Tierpflege (
	TI_ID			INT REFERENCES Tiere (TI_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	AD_ID	INT REFERENCES Arbeitsdienste (AD_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	FU_ID			INT REFERENCES Futter (FU_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	Menge_min		INT,
	Menge_max		INT,
	Mahlkategorie	INT
);

COMMENT ON TABLE  Tierpflege				IS 'Informationen zu den jeweiligen Tierpflegeeinsätzen';
COMMENT ON COLUMN Tierpflege.TI_ID			IS 'Verweis auf das Tier ';
COMMENT ON COLUMN Tierpflege.AD_ID			IS 'Verweis auf den Arbeitseinsatz';
COMMENT ON COLUMN Tierpflege.FU_ID			IS 'Verweis auf das Futter';
COMMENT ON COLUMN Tierpflege.Menge_min		IS 'Wie viel soll mindestens verwendet werden?';
COMMENT ON COLUMN Tierpflege.Menge_max		IS 'Wie viel soll höchstens verwendet werden?';
COMMENT ON COLUMN Tierpflege.Mahlkategorie	IS 'Typ der Mahlzeit: 1 = Hauptmahl, 2 = alt. Hauptmahl, 3 = Nebenmahl';



CREATE TABLE Terrariumpflege (
	TR_ID			INT REFERENCES Terrarien (TR_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	AD_ID			INT REFERENCES Arbeitsdienste (AD_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE
);

COMMENT ON TABLE  Terrariumpflege		IS 'Informationen zu den jeweiligen Terrariumpflegeeinsätzen';
COMMENT ON COLUMN Terrariumpflege.TR_ID	IS 'Verweis auf das Terrarium';
COMMENT ON COLUMN Terrariumpflege.AD_ID	IS 'Verweis auf den Arbeitseinsatz';



CREATE TABLE Lagerpflege (
	VM_ID	INT REFERENCES Verbrauchsmittel (VM_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	AD_ID		INT REFERENCES Arbeitsdienste (AD_ID)
		ON DELETE CASCADE
		ON UPDATE CASCADE
);

COMMENT ON TABLE  Lagerpflege		IS 'Informationen zu den jeweiligen Lagerpflegeeinsätzen';
COMMENT ON COLUMN Lagerpflege.VM_ID	IS 'Verweis auf das Verbrauchsmittel';
COMMENT ON COLUMN Lagerpflege.AD_ID	IS 'Verweis auf den Arbeitseinsatz';
