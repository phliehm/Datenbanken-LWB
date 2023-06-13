create table kunden (
  kid       integer         primary key,
  nname     varchar (100),
  vname     varchar (100),
  strasse   varchar (100),
  plz       char (5)        check (plz between '00000' and '99999'),
  ort       varchar (100)
);

insert into kunden values (1,'Schmidt','Harald','Mercatorweg 3','12345','Pudelsdorf');
insert into kunden values (2,'Schmitt','Hans','Alstergasse 5','98761','Herzogshagen');
insert into kunden values (3,'Schmitt','Klaus','Kirchweg 3','87621','Glücksburg');
insert into kunden values (4,'Schmitt Hirschfelder','Klaus','Wernerweg 1a','03333','Zitzelsberg');
