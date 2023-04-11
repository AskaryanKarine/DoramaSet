drop schema if exists dorama_set cascade;

create schema dorama_set;

create table dorama_set.Subscription (
    id serial primary key,
    cost int not null ,
    description text not null,
    duration int not null
);

create table dorama_set.User (
    username text primary key,
    sub_id int,
    password text not null,
    email text not null,
    registration_date date not null,
    last_active date not null,
    points int not null default 0,
    is_admin bool not null default false,
    foreign key (sub_id) references dorama_set.Subscription(id) on delete cascade
);

create table dorama_set.List (
    id serial primary key,
    name_creator text not null,
    foreign key (name_creator) references dorama_set.User(username) on delete cascade,
    name_list text not null,
    type text not null,
    description text not null
);

create table dorama_set.Dorama (
    id serial primary key,
    name text not null,
    description text not null,
    release_year int not null,
    status text not null,
    genre text not null
);

create table dorama_set.Episode (
    id serial primary key,
    id_dorama int not null,
    foreign key (id_dorama) references dorama_set.Dorama(id) on delete cascade,
    num_season int not null,
    num_episode int not null
);

create table dorama_set.Staff (
    id serial primary key,
    birthday date not null,
    gender text not null,
    bio text not null
);

create table dorama_set.Picture (
    id serial primary key,
    description text not null,
    URL text not null
);

-- развязочные таблицы

create table dorama_set.UserList (
    username text not null,
    id_list int not null,
    primary key (username, id_list),
    foreign key (username) references dorama_set."user"(username),
    foreign key (id_list) references dorama_set.list(id)
);

create table dorama_set.UserEpisode (
    username text not null,
    id_episode int not null,
    primary key (username, id_episode),
    foreign key (username) references dorama_set."user"(username) on delete cascade,
    foreign key (id_episode) references dorama_set.episode(id) on delete cascade
);

create table dorama_set.DoramaStaff (
    id_dorama int not null,
    id_staff int not null,
    primary key (id_dorama, id_staff),
    foreign key (id_dorama) references dorama_set.dorama(id),
    foreign key (id_staff) references dorama_set.staff(id)
);

create table dorama_set.DoramaPicture (
    id_dorama int not null,
    id_picture int not null,
    primary key (id_dorama, id_picture),
    foreign key (id_dorama) references dorama_set.dorama(id),
    foreign key (id_picture) references dorama_set.picture(id)
);

create table dorama_set.StaffPicture (
    id_staff int not null,
    id_picture int not null,
    primary key (id_staff, id_picture),
    foreign key (id_staff) references dorama_set.staff(id),
    foreign key (id_picture) references dorama_set.picture(id)
);