drop schema if exists dorama_set cascade;

create schema if not exists dorama_set;

create table if not exists dorama_set.Subscription (
    id serial primary key,
    name text not null,
    cost int not null ,
    description text not null,
    duration int not null,
    access_lvl int not null 
);

create table if not exists dorama_set.User (
    username text primary key,
    sub_id int,
    password text not null,
    email text not null,
    registration_date date not null,
    last_active date not null,
    last_subscribe date not null,
    points int not null default 0,
    is_admin bool not null default false,
    emoji text default '2b50',
    color text not null default '#000000',
    foreign key (sub_id) references dorama_set.Subscription(id) on delete cascade
);

create table if not exists dorama_set.List (
    id serial primary key,
    name_creator text not null,
    foreign key (name_creator) references dorama_set.User(username) on delete cascade,
    name_list text not null,
    type text not null,
    description text not null
);

create table if not exists dorama_set.Dorama (
    id serial primary key,
    name text not null,
    description text not null,
    release_year int not null,
    status text not null,
    genre text not null
);

create table if not exists dorama_set.Episode (
    id serial primary key,
    id_dorama int not null,
    foreign key (id_dorama) references dorama_set.Dorama(id) on delete cascade,
    num_season int not null,
    num_episode int not null
);

create table if not exists dorama_set.Staff (
    id serial primary key,
    name text not null,
    birthday date,
    gender text,
    type text not null
);

create table if not exists dorama_set.Picture (
    id serial primary key,
    URL text not null
);

-- развязочные таблицы

create table if not exists dorama_set.UserList (
    username text not null,
    id_list int not null,
    primary key (username, id_list),
    foreign key (username) references dorama_set."user"(username) on delete cascade,
    foreign key (id_list) references dorama_set.list(id) on delete cascade
);

create table if not exists dorama_set.UserEpisode (
    username text not null,
    id_episode int not null,
    primary key (username, id_episode),
    foreign key (username) references dorama_set."user"(username) on delete cascade,
    foreign key (id_episode) references dorama_set.episode(id) on delete cascade
);

create table if not exists dorama_set.DoramaStaff (
    id_dorama int not null,
    id_staff int not null,
    primary key (id_dorama, id_staff),
    foreign key (id_dorama) references dorama_set.dorama(id) on delete cascade,
    foreign key (id_staff) references dorama_set.staff(id) on delete cascade
);

create table if not exists dorama_set.DoramaPicture (
    id_dorama int not null,
    id_picture int not null,
    primary key (id_dorama, id_picture),
    foreign key (id_dorama) references dorama_set.dorama(id) on delete cascade,
    foreign key (id_picture) references dorama_set.picture(id) on delete cascade
);

create table if not exists dorama_set.ListDorama (
    id_dorama int not null,
    id_list int not null,
    primary key (id_dorama, id_list),
    foreign key (id_dorama) references dorama_set.dorama(id) on delete cascade,
    foreign key (id_list) references dorama_set.list(id) on delete cascade
);

create table if not exists dorama_set.StaffPicture (
    id_staff int not null,
    id_picture int not null,
    primary key (id_staff, id_picture),
    foreign key (id_staff) references dorama_set.staff(id) on delete cascade,
    foreign key (id_picture) references dorama_set.picture(id) on delete cascade
);