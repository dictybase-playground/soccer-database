create table players (

  pid serial primary key,
  pname varchar(64) not null UNIQUE,
  clubid varchar(64) not null,
  countryid varchar(64) not null

);

create table clubs(
  clubid serial primary key,
  clubname varchar(64) not null UNIQUE

);

create table countries(
  countryid serial primary key,
  country varchar(64) not null UNIQUE

);


drop table players;
drop table countries;
drop table clubs;
