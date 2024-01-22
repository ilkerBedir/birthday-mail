\c email-db

create table users (
    id serial primary key,
	name_surname varchar(255) not null,
	birth_date date not null,
	email varchar(255) not null
);

insert into users (name_surname,birth_date,email) values ('İlker Bedir','1998-01-21','test@gmail.com')