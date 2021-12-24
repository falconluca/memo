create database if not exists `memo`;

use `memo`;

create table if not exists `items` (
    `id` bigint unsigned not null auto_increment primary key,
    `title` varchar(255) default null,
    `description` varchar(255) default null,
    `remind_at` timestamp null default null
);