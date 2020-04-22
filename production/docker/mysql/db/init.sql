use sample;

create table places (
    place_id int auto_increment not null primary key,
    place_name varchar(1024) not null
);

create table books(
    rfid_tag varchar(20) not null primary key,
    book_name varchar(1024),
    isbn varchar(20),
    place_id int not null,
    book_datetime datetime not null,
    constraint fk_place_id foreign key (place_id) references places(place_id)
);

create table persons (
    person_id int auto_increment not null primary key,
    card_data varchar(10),
    person_name varchar(50) not null,
    person_email varchar(128) not null,
    password varchar(128) not null,
    person_datetime datetime not null
);

create table borrowed_logs(
    borrowed_log_id int auto_increment not null primary key,
    rfid_tag varchar(20) not null,
    person_id int not null,
    constraint fk_rfid_tag foreign key (rfid_tag) references books(rfid_tag),
    constraint fk_person_id foreign key (person_id) references persons(person_id)
);

create table pre_persons (
    pre_preson_id int auto_increment not null primary key,
    pre_person_email varchar(128) not null,
    pre_person_token varchar(128) not null,
    pre_person_datetime datetime not null
);