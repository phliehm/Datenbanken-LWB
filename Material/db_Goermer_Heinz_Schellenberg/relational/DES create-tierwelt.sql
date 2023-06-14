%
% Tierwelt
%
% SQL Formulation

/multiline on
/abolish
/show_compilations on
/sql

CREATE TABLE Habitattypen (
	HT_ID			INT NOT NULL,
	Bezeichnung		VARCHAR	NOT NULL,
	Beschreibung		VARCHAR,	
	Temperatur		INT,	
	Luftfeuchtigkeit	INT,	
	Ausstattung		VARCHAR,
	CONSTRAINT HT_ID PRIMARY KEY (HT_ID)
) ;


CREATE TABLE Terrarien (
	TR_ID		INT NOT NULL,
	Name		VARCHAR NOT NULL,
	Beschreibung	VARCHAR,
	Raumnummer	VARCHAR,
	Breite		INT,
	Breite_vorn	INT,
	Tiefe		INT,	
	Hoehe		INT,
	HT_ID		INT REFERENCES Habitattypen (HT_ID),
	CONSTRAINT TR_ID PRIMARY KEY (TR_ID)
) ;

CREATE TABLE Unterarten (
	UA_ID			INT NOT NULL,
	Bezeichnung		VARCHAR NOT NULL,
	Heimat			VARCHAR,
	Lebensraum		VARCHAR,
	Temperatur		VARCHAR,
	Luftfeuchtigkeit	VARCHAR,
	Groesse			VARCHAR,
	Lebenserwartung		VARCHAR,
	Ernaehrung		VARCHAR,
	Besonderheiten		VARCHAR,
	CONSTRAINT UA_ID PRIMARY KEY (UA_ID)
) ;

CREATE TABLE Tiere (
	TI_ID		INT NOT NULL,
	Name		VARCHAR NOT NULL,
	Geschlecht	CHAR,
	Laenge		INT,
	Gewicht		INT,
	Zugangsdatum	DATE,
	UA_ID		INT REFERENCES Unterarten (UA_ID),
	TR_ID		INT REFERENCES Terrarien (TR_ID),
	CONSTRAINT TI_ID PRIMARY KEY (TI_ID)
) ;

CREATE TABLE Verbrauchsmittel (
	VM_ID		INT NOT NULL,
	Bezeichnung	VARCHAR NOT NULL,
	Beschreibung	VARCHAR,
	Lager_min	INT,
	Lager_akt	INT,
	Kosten		Float,
	CONSTRAINT VM_ID PRIMARY KEY (VM_ID)
);

CREATE TABLE Leuchtmittel (
	VM_ID		INT NOT NULL REFERENCES Verbrauchsmittel (VM_ID),
	Helligkeit	INT,
	Waermeleistung	INT
);

CREATE TABLE Bodenbelag (
	VM_ID		INT NOT NULL REFERENCES Verbrauchsmittel (VM_ID),
	Farbe		VARCHAR,
	Koernung	VARCHAR
);

CREATE TABLE Pflanzen (
	VM_ID		INT NOT NULL REFERENCES Verbrauchsmittel (VM_ID),
	Wasserverbrauch	INT,
	kuenstlich	INT, 
	spruehen	INT
);

CREATE TABLE Wasserbecken (
	VM_ID		INT NOT NULL REFERENCES Verbrauchsmittel (VM_ID),
	Hoehe		INT,
	Breite		INT,
	Tiefe		INT,
	bepflanzt	INT
);

CREATE TABLE Deko (
	VM_ID		INT NOT NULL REFERENCES Verbrauchsmittel (VM_ID),
	Material	VARCHAR,
	Farbe		VARCHAR,
	Zweck		VARCHAR
);

CREATE TABLE Arbeitsdiensttypen (
	AT_ID		INT NOT NULL,
	Bezeichnung	VARCHAR NOT NULL,
	Beschreibung	VARCHAR,
	Frequenz	INT,
	CONSTRAINT AT_ID PRIMARY KEY (AT_ID)
);

CREATE TABLE Arbeitsdienste (
	AD_ID		INT NOT NULL,
	AT_ID		INT REFERENCES Arbeitsdiensttypen (AT_ID),
	Faelligkeit	DATE,
	CONSTRAINT AD_ID PRIMARY KEY (AD_ID)
);

CREATE TABLE Verbrauch (
	TR_ID	INT REFERENCES Terrarien (TR_ID),
	VM_ID	INT REFERENCES Verbrauchsmittel (VM_ID),
	AD_ID	INT REFERENCES Arbeitsdienste (AD_ID),
	Anzahl	INT NOT NULL
);

CREATE TABLE Futter (
	FU_ID		INT NOT NULL ,
	Bezeichnung	VARCHAR NOT NULL,
	Lagerungsform	VARCHAR,
	CONSTRAINT FU_ID PRIMARY KEY (FU_ID)
);

CREATE TABLE Tierpflege (
	TI_ID		INT REFERENCES Tiere (TI_ID),
	AD_ID		INT REFERENCES Arbeitsdienste (AD_ID),
	FU_ID		INT REFERENCES Futter (FU_ID),
	Menge_min	INT,
	Menge_max	INT,
	Mahlkategorie	INT
);

CREATE TABLE Terrariumpflege (
	TR_ID	INT REFERENCES Terrarien (TR_ID),
	AD_ID	INT REFERENCES Arbeitsdienste (AD_ID)
);

CREATE TABLE Lagerpflege (
	VM_ID	INT REFERENCES Verbrauchsmittel (VM_ID),
	AD_ID	INT REFERENCES Arbeitsdienste (AD_ID)
);

/development off
/dbschema
/development on
/dbschema
