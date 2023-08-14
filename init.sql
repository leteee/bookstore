# 用户
drop table if exists users;
create table users
(
    id       int auto_increment
        primary key,
    username varchar(100) not null,
    password varchar(100) not null,
    email    varchar(100) null,
    constraint username
        unique (username)
);

insert into users (username, password, email)
values ('admin', 'admin', 'admin@qq.com');

# 图书
drop table if exists books;
create table books
(
    id       int auto_increment
        primary key,
    title    varchar(100) not null,
    author   varchar(100) not null,
    price    double(11, 2),
    sales    int          not null,
    stock    int          not null,
    img_path varchar(100) null
);

INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('解忧杂货店', '东野圭吾', 27.20, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('边城', '沈从文', 23.00, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('中国哲学史', '冯友兰', 44.5, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('忽然七日', ' 劳伦', 19.33, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('苏东坡传', '林语堂', 19.30, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('百年孤独', '马尔克斯', 29.50, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('扶桑', '严歌苓', 19.8, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('给孩子的诗', '北岛', 22.20, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('为奴十二年', '所罗门', 16.5, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('平凡的世界', '路遥', 55.00, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('悟空传', '今何在', 14.00, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('硬派健身', '斌卡', 31.20, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('从晚清到民国', '唐德刚', 39.90, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('三体', '刘慈欣', 56.5, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('看见', '柴静', 19.50, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('活着', '余华', 11.00, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('小王子', '安托万', 19.20, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('我们仨', '杨绛', 11.30, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('生命不息,折腾不止', '罗永浩', 25.20, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('皮囊', '蔡崇达', 23.90, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('恰到好处的幸福', '毕淑敏', 16.40, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('大数据预测', '埃里克', 37.20, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('人月神话', '布鲁克斯', 55.90, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('C语言入门经典', '霍尔顿', 45.00, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('数学之美', '吴军', 29.90, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('Java编程思想', '埃史尔', 70.50, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('设计模式之禅', '秦小波', 20.20, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('图解机器学习', '杉山将', 33.80, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('艾伦图灵传', '安德鲁', 47.20, 100, 100, 'static/img/default.jpg');
INSERT INTO books (title, author, price, sales, stock, img_path)
VALUES ('教父', '马里奥普佐', 29.00, 100, 100, 'static/img/default.jpg');

# session表
drop table if exists sessions;
create table sessions
(
    session_id varchar(100) primary key,
    username   varchar(100) not null,
    user_id    int          not null,
    FOREIGN KEY (user_id) references users (id)
);

# 购物车表
drop table if exists carts;
create table carts
(
    id           varchar(100) primary key,
    total_count  int           not null,
    total_amount DOUBLE(11, 2) not null,
    user_id      int           not null,
    foreign key (user_id) references users (id)
);
# 购物车项目表
drop table if exists cart_items;
create table cart_items
(
    id      int primary key auto_increment,
    count   int           not null,
    amount  double(11, 2) not null,
    book_id int           not null,
    cart_id varchar(100)  not null,
    foreign key (book_id) references books (id),
    foreign key (cart_id) references carts (id)
);

# 订单表
drop table if exists orders;
create table orders
(
    id           varchar(100) primary key,
    create_time  varchar(100)  not null,
    total_count  int           not null,
    total_amount double(11, 2) not null,
    state        int           not null,
    user_id      int
);

# 订单项
drop table if exists order_items;
create table order_items
(
    id       int primary key AUTO_INCREMENT,
    order_id varchar(100)  not null,
    count    int           not null,
    amount   double(11, 2) not null,
    title    varchar(100)  not null,
    author   varchar(100)  not null,
    price    double(11, 2),
    img_path varchar(100)  null,
    foreign key (order_id) references orders (id)
)
