create database db_book
default char set utf8
default collate utf8_unicode_ci;

create table tb_categorie(
id int PRIMARY KEY auto_increment,
name varchar(45) not null,
status bool default true,
createdat datetime default current_timestamp on update now(),
UpdatedAt timestamp default current_timestamp on update current_timestamp
)ENGINE = InnoDB;

create table tb_product(
id int PRIMARY KEY auto_increment,
name varchar(45) not null,
edition varchar(10),
releasedata date not null,
idcategorie int not null,
status bool default true,
createdat datetime default current_timestamp on update now(),
UpdatedAt timestamp default current_timestamp on update current_timestamp,

constraint `fk_produto_categoria` foreign key (idcategorie) references tb_categorie(id)
);

create table tb_user(
id int PRIMARY KEY auto_increment,
name varchar(45) not null,
bi varchar(15) not null,
birthdate varchar(15) not null,
adress varchar(45) not null,
idproduct int not null,
status bool default true,
createdat datetime default current_timestamp on update now(),
UpdatedAt timestamp default current_timestamp on update current_timestamp,
constraint `fk_user_produto` foreign key (idproduct) references tb_product(id)
);