use book_management_db;

create table places (
    -- -1:持出中 1:貸出中 2:本棚
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
    description varchar(4096),
    -- 初期値ではnullが入っている
    book_img_url varchar(1024)
);

create table book_statuses(
    rfid_tag varchar(128) not null primary key,
    book_info_id int not null unique,
    place_id int not null,
    book_datetime datetime not null,
    constraint fk_place_id foreign key (place_id) references places(place_id),
    -- 依存しているbook_info_idが削除されたとき付随して削除される
    constraint fk_book_info_id foreign key (book_info_id) references book_info (book_info_id) on
    delete cascade
);

create table emails (
    email_id int auto_increment not null primary key,
    email varchar(128) not null unique,
    password varchar(128) not null
);

create table persons (
    person_id int auto_increment not null primary key,
    card_data varchar(10) unique,
    person_name varchar(50) not null,
    email_id int not null unique,
    person_datetime datetime not null,
    -- 依存しているemail_idが削除されたとき付随して削除される
    constraint fk_email_id foreign key (email_id) references emails(email_id) on
    delete cascade
);

create table borrowed_logs(
    borrowed_log_id int auto_increment not null primary key,
    rfid_tag varchar(20) not null,
    person_id int,
    -- 依存しているrfid_tagが削除されたとき付随してレコードは削除される
    constraint fk_rfid_tag foreign key (rfid_tag) references book_statuses(rfid_tag) on
    delete cascade,
    -- 依存しているperson_idが削除されたときこのテーブルのレコードにnullを入れる
    constraint fk_person_id foreign key (person_id) references persons(person_id) on
    delete
    set
        null
);

create table pre_persons (
    pre_person_id int auto_increment not null primary key,
    pre_person_email varchar(128) not null unique key,
    pre_person_token varchar(128) not null,
    pre_person_datetime datetime not null
);

-- 運用に必要なレコード
insert into
    places (place_id, place_name)
values
    (-1, "持出中");

insert into
    places (place_id, place_name)
values
    (1, "貸出中");

insert into
    places (place_name)
values
    ("シス研本棚");

-- テスト用データ挿入
insert into
    emails(email, password)
values
    (
        "e19070ee@aitech.ac.jp",
        "ecb666d778725ec97307044d642bf4d160aabb76f56c0069c71ea25b1e926825" # hoge
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
        "hoge",
        "guest",
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
        "24031d1fa7ec42af43c24d5f6ba0834a6c4ce7cadf92a8dfd926ca5ffac14b77",
        "20200426180100"
    );

-- datetimeのフォーマットは上記のいずれのフォーマットでも入力でき 、 出力は同じフォーマットになる