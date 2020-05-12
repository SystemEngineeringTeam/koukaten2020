use book_management_db;

create table places (
    place_id int auto_increment not null primary key,
    place_name varchar(128) not null
);

create table book_info(
    book_info_id int auto_increment not null primary key,
    api_id varchar(20) not null,
    book_name varchar(128),
    author varchar(128),
    publisher varchar(128),
    published_date varchar(20),
    description varchar(1024)
);

create table book_statuses(
    rfid_tag varchar(128) not null primary key,
    book_info_id int not null unique,
    place_id int not null,
    book_datetime datetime not null,
    constraint fk_place_id foreign key (place_id) references places(place_id),
    constraint fk_book_info_id foreign key (book_info_id) references book_info (book_info_id)
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
    constraint fk_rfid_tag foreign key (rfid_tag) references book_statuses(rfid_tag),
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
    ("シス研本棚");

insert into
    book_info(
        book_name,
        api_id,
        author,
        publisher,
        published_date,
        description
    )
values
    (
        "苦しんで覚えるC言語",
        "VL6_u1Nyic0C",
        "MMGames",
        "秀和システム",
        "2011-06-00",
        "いつの時代も基本はC言語。ネットで人気のC講座を書籍化。"
    );

insert into
    book_info(
        book_name,
        api_id,
        author,
        publisher,
        published_date,
        description
    )
values
    (
        "苦しんで覚えるC言語2",
        "VL6_u1Nyic0B",
        "MMGames",
        "秀和システム",
        "2011-06-00",
        "いつの時代も基本はC言語。ネットで人気のC講座を書籍化。"
    );

insert into
    book_info(
        book_name,
        api_id,
        author,
        publisher,
        published_date,
        description
    )
values
    (
        "苦しんで覚えるC言語3",
        "VL6_u1Nyic0A",
        "MMGames",
        "秀和システム",
        "2011-06-00",
        "いつの時代も基本はC言語。ネットで人気のC講座を書籍化。"
    );

insert into
    book_statuses (
        rfid_tag,
        book_info_id,
        place_id,
        book_datetime
    )
values
    (
        "hoge",
        1,
        1,
        "2020/04/30 04:45:50"
    );

insert into
    book_statuses (
        rfid_tag,
        book_info_id,
        place_id,
        book_datetime
    )
values
    (
        "hogehoge",
        2,
        1,
        "2020/04/30 04:45:50"
    );

insert into
    book_statuses (
        rfid_tag,
        book_info_id,
        place_id,
        book_datetime
    )
values
    (
        "fuga",
        3,
        1,
        "2020/04/30 04:45:50"
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