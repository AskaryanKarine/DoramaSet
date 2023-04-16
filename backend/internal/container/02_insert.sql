insert into dorama_set.Subscription (id, cost, description, duration)
values (1, 0, 'free', 2592000), (2, 50, 'not free', 2592000);

insert into dorama_set.User (username, sub_id, password, email, registration_date, last_active, last_subscribe, points)
values ('test', 1, 'qwerty', 'qwerty@gmail.com', '2023-01-01', '01-01-2023', '2023-01-01', 100),
       ('test1', 1, 'qwerty', 'qwerty@gmail.com', '2023-01-01', '01-01-2023', '2023-01-01', 100);

insert into dorama_set.list (name_creator, name_list, type, description)
values ('test', 'test', 'private', 'test');

insert into dorama_set.userlist (username, id_list)
values ('test', 1);

insert into dorama_set.dorama (id, name, description, release_year, status, genre)
values (1, 'test', 'test', 2000, 'finish', 'test');

insert into dorama_set.episode (id_dorama, num_season, num_episode)
values (1, 1, 1), (1, 1, 2);