USE sample;

CREATE TABLE  tasks(
    id INT auto_increment not null,
    datetime datetime,
    person_id int,
    contents varchar(4096),
    primary key(id)
);

create table persons(
    id int auto_increment not null,
    number varchar(128),
    name varchar(128),
    primary key(id)
);

insert into persons(number,name) values ("e19070","Shibahara"),("k19092","Fukuda"),("x19012","Inagaki"),("x19053","Suzuki");

