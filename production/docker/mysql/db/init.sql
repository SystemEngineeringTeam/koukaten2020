-- -- テスト用クエリ
-- drop database if exists sample;
-- -- テスト用クエリ
-- create database sample;
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

create table emails (
    email_id int auto_increment not null primary key,
    email varchar(128) not null unique,
    password varchar(128) not null
);

create table persons (
    person_id int auto_increment not null primary key,
    card_data varchar(10),
    person_name varchar(50) not null,
    email_id int not null unique,
    person_datetime datetime not null,
    constraint fk_email_id foreign key (email_id) references emails(email_id)
);

create table borrowed_logs(
    borrowed_log_id int auto_increment not null primary key,
    rfid_tag varchar(20) not null,
    person_id int not null,
    constraint fk_rfid_tag foreign key (rfid_tag) references books(rfid_tag),
    constraint fk_person_id foreign key (person_id) references persons(person_id)
);

create table pre_persons (
    pre_person_id int auto_increment not null primary key,
    pre_person_email varchar(128) not null unique key,
    pre_person_token varchar(128) not null,
    pre_person_datetime datetime not null
);

-- ここからテスト用データ挿入
insert into
    places (place_id, place_name)
values
    (-1, "持出中");

insert into
    places (place_id, place_name)
values
    (0, "貸出中");

insert into
    places (place_name)
values
    ("シス研の本棚");

insert into
    books (
        rfid_tag,
        book_name,
        isbn,
        place_id,
        book_datetime
    )
values
    (
        "hoge",
        "苦しんで覚えるC言語",
        "9784798030142",
        "1",
        "2020/04/26 17:46:00"
    );

insert into
    emails(email, password)
values
    (
        "e19070ee@aitech.ac.jp",
        "4c716d4cf211c7b7d2f3233c941771ad0507ea5bacf93b492766aa41ae9f720d"
    );

insert into
    persons(
        card_data,
        person_name,
        email_id,
        person_datetime
    )
values
    (
        "E19070",
        "柴原恒佑",
        1,
        "2020-04-26 17:53:00"
    );

insert into
    borrowed_logs(rfid_tag, person_id)
values
    ("hoge", 1);

insert into
    pre_persons(
        pre_person_email,
        pre_person_token,
        pre_person_datetime
    )
values
    (
        "e19070ee@aitech.ac.jp",
        "792dcf7b8e952b3a10bd8d71303931c148df413caaa162d1777a79493b163ead",
        "20200426180100"
    );

-- datetimeのフォーマットは上記のいずれのフォーマットでも入力でき 、 出力は同じフォーマットになる