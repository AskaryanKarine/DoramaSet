insert into dorama_set.Subscription (id, cost, description, duration, access_lvl, name)
values (1, 0, 'free', 2592000, 1, 'basic'), (2, 50, 'not free', 2592000, 2, 'premium');

insert into dorama_set.dorama (id, name, description, release_year, status, genre)
values (1, 'test', 'test', 2000, 'finish', 'test'),
       (2, 'test2', 'test2', 2002, 'finish', 'test2');

insert into dorama_set.picture (id, url)
VALUES (1, 'test_picture1'), (2, 'test_picture2');

insert into dorama_set.doramapicture (id_dorama, id_picture)
VALUES (1, 1), (1, 2)

