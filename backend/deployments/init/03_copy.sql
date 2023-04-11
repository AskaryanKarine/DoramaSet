-- тестовые данные
-- insert into dorama_set.dorama (id, name, description, release_year, status, genre)
-- values (1, '_test_qwerty', 'qwerty', 2020, 'finish', 'comedian'),
--        (2, '_test_qwerty zxcvb', 'asdfg', 2022, 'in progress', 'thriller');
--
-- insert into dorama_set.episode (id, id_dorama, num_season, num_episode)
-- values (1, 1, 1, 1),
--        (2, 1, 1, 2),
--        (3, 2, 1, 1);


-- реальные данные
copy dorama_set.subscription from '/data/subscription.csv' delimiter ',';
