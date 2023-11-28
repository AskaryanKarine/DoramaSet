insert into dorama_set.dorama (id, name, description, release_year, status, genre)
values (1, 'test', 'test', 2000, 'finish', 'test');

insert into dorama_set.picture (id, url)
VALUES (1, 'test_picture1'), (2, 'test_picture2');

insert into dorama_set.doramapicture (id_dorama, id_picture)
VALUES (1, 1), (1, 2)

