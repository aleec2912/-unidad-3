create database if not exists biblioteca;
use biblioteca;
create table if not exists biblioteca(
    id bigint unsigned not null auto_increment,
    nombre varchar(100) not null,
    descripcion varchar(450) not null,
    autor varchar(200) not null,
    editorial varchar(200) not null,
    fechapublicacion date,
    primary key(id)
);