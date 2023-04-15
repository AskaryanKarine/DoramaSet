insert into dorama_set.Subscription (id, cost, description, duration)
values (1, 0, 'free', 2592000), (2, 50, 'not free', 2592000);

insert into dorama_set.User (username, sub_id, password, email, registration_date, last_active, last_subscribe, points)
values ('test', 1, 'qwerty', 'qwerty@gmail.com', '2023-01-01', '01-01-2023', '2023-01-01', 100);

insert into dorama_set.list (id, name_creator, name_list, type, description)
values (1, 'test', 'test', 'private', 'test');

insert into dorama_set.userlist (username, id_list)
values ('test', 1);