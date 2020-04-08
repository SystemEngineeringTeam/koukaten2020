USE sample;

CREATE TABLE tasks(
    id INT auto_increment not null,
    date date,
    contents varchar(4096),
    index(id)
);

create table persons(
    id int auto_increment,
    number varchar(128),
    name varchar(128),
    index(id)
);

insert into persons(number,name) values ("e19070","Shibahara"),("k19092","Fukuda"),("x19012","Inagaki"),("x19053","Suzuki")

