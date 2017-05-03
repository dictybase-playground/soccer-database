create table players (

  Pid serial primary key,
  Pname varchar(64) not null UNIQUE,
  Clubid int not null,
  Countryid int not null

);

create table clubs(
  Clubid serial primary key,
  Clubname varchar(64) not null UNIQUE

);

create table countries(
  Countryid serial primary key,
  Countryname varchar(64) not null UNIQUE

);


drop table players;
drop table countries;
drop table clubs;
