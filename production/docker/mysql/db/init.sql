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
    book_url varchar(1024)
);

create table book_statuses(
    rfid_tag varchar(128) not null primary key,
    book_info_id int not null unique,
    place_id int not null,
    book_datetime datetime not null,
    constraint fk_place_id foreign key (place_id) references places(place_id),
    -- 依存しているbook_info_idが削除されたとき付随して削除される
    constraint fk_book_info_id foreign key (book_info_id) references book_info (book_info_id) on
    delete
        cascade
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
    delete
        cascade
);

create table borrowed_logs(
    borrowed_log_id int auto_increment not null primary key,
    rfid_tag varchar(20) not null,
    person_id int,
    -- 依存しているrfid_tagが削除されたとき付随してレコードは削除される
    constraint fk_rfid_tag foreign key (rfid_tag) references book_statuses(rfid_tag) on
    delete
        cascade,
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

-- ここからテスト用データ挿入
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
        "入門Goプログラミング",
        "zreZDwAAQBAJ",
        "Nathan Youngman他",
        "翔泳社",
        "2019-05-13",
        "【本書の内容】 比較的小さいけれど、どのようなスキルのプログラマ(含む予備軍)にとっても有用で有益なプログラミング言語がGoです。 本書は2009年に発表され、以来、コンパクトさとシンプルさを残したまま、さまざまなスキルをもったプログラマに愛される言語として成長してきたGo言語を意のままに操れるレベルを目指せる学習書です。 ビギナーにとっても、Webにある実行環境を使用することで掲載されたサンプルプログラムや、練習問題を実行できるよう配慮されていますし、スクリプト言語を使った経験があれば、ウォーミングアップは終わっています。スクラッチやエクセルのフォーミュラを使ったことがあったりHTMLを書いた経験があれば、本格的なプログラミング言語の第一歩を踏み出す格好のきっかけになるはずです。 もちろん、それなりの忍耐とか努力は必要ですが、火星探査機を構築しながら(!?)Go言語を完全習得できる、最初の1冊となるでしょう。 【本書のポイント】 ・特定のテーマごとに学習単元をユニット化 ・ユニットはさらに特定のトピックを扱うレッスンに分割 ・レッスン後はクイックチェック ・さらに練習問題を配置して知識を根付かせ応用力を ・問題に対する模範解答も完備 ・ローカルに開発環境を持っていなくても学習できる ・GitHub上にも最新の解答例を掲載 【読者が得られること】 ・Go言語の構文を理解できる ・Go言語を使ったプログラミング ・特色とメリットを活かした開発技法の習得 ※本電子書籍は同名出版物を底本として作成しました。記載内容は印刷出版当時のものです。 ※印刷出版再現のため電子書籍としては不要な情報を含んでいる場合があります。 ※印刷出版とは異なる表記・表現の場合があります。予めご了承ください。 ※プレビューにてお手持ちの電子端末での表示状態をご確認の上、商品をお買い求めください。 (翔泳社)"
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
        "クトゥルフ神話TRPG",
        "z2B3QgAACAAJ",
        "Sandy Petersen",
        "秀和システム",
        "2004-09",
        "能力値ロール、技能ロール、戦闘ルールといった、RPGに必要なシステムや、クトゥルフ神話の資料、シナリオなどプレイに必要なものはすべて、この1冊に含まれている。システムは“ベーシック・ロールプレイング”をベースに、正気度、クトゥルフ神話技能、銃器戦闘などを追加し、ホラーTRPGを楽しむために発展させたものを採用。これまでにない緊張と破滅を体験することができる。クトゥルフ神話に登場する数多くのクリーチャー、神格、魔道書、アーティファクトなどを豊富にカバー。クトゥルフ神話事典としても見逃せない。すぐにプレイ可能な、4本のシナリオを収録。また、20世紀に実際に起こったさまざまなクトゥルフ神話的な事件を紹介。シナリオのアイディアに活用できる。"
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