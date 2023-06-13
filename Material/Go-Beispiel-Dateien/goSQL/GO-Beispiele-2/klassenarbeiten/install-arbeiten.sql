CREATE Table klassen (
  KID       integer PRIMARY KEY,
  Name      varchar (10)
);

CREATE Table arbeiten (
  KID       integer REFERENCES klassen (KID),
  ANr       integer,
  Thema     varchar (30),
  Anz1      integer,
  Anz2      integer,
  Anz3      integer,
  Anz4      integer,
  Anz5      integer,
  Anz6      integer,
  PRIMARY KEY (KID,ANr)
);

INSERT INTO klassen VALUES (1, '8a');
INSERT INTO klassen VALUES (2, '7c');
INSERT INTO klassen VALUES (3, '9e');
INSERT INTO klassen VALUES (4, '10b');

INSERT INTO arbeiten VALUES (1, 1, 'Terme und Gleichungen', 3, 6, 9, 7, 4, 0);
INSERT INTO arbeiten VALUES (1, 2, 'Prozentrechnung', 8, 8, 3, 7, 2, 0);
INSERT INTO arbeiten VALUES (2, 1, 'Ebene Figuren', 0, 4, 11, 9, 5, 0);
INSERT INTO arbeiten VALUES (2, 2, 'Proportionalität', 1, 3, 7, 9, 2, 0);
INSERT INTO arbeiten VALUES (3, 1, 'Satzgruppe des Pythagoras', 3, 8, 3, 12, 3, 1);
INSERT INTO arbeiten VALUES (4, 1, 'Trigonometrie', 7, 5, 9, 2, 4, 0);
INSERT INTO arbeiten VALUES (4, 2, 'Stereometrie', 4, 4, 12, 6, 3, 0);
INSERT INTO arbeiten VALUES (4, 3, 'Exponentielles Wachstum', 2, 5, 10, 8, 3, 1);
