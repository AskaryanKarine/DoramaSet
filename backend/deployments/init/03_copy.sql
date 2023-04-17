copy dorama_set.subscription from '/data/subscription.csv' delimiters ';';
select setval('dorama_set.subscription_id_seq', (select max(id) from dorama_set.subscription));

copy dorama_set.dorama from '/data/dorama.csv' delimiter ';';
select setval('dorama_set.dorama_id_seq', (select max(id) from dorama_set.dorama));

copy dorama_set.episode from '/data/episode.csv' delimiter ';';
select setval('dorama_set.episode_id_seq', (select max(id) from dorama_set.episode));

copy dorama_set.picture from '/data/picture.csv' delimiter ',';
select setval('dorama_set.picture_id_seq', (select max(id) from dorama_set.picture));

copy dorama_set.picture from '/data/staff.csv' delimiter ',';
select setval('dorama_set.staff_id_seq', (select max(id) from dorama_set.staff));

copy dorama_set.doramapicture from '/data/dorama-picture.csv' delimiter ',';
copy dorama_set.doramastaff from '/data/dorama-staff.csv' delimiter ',';
copy dorama_set.staffpicture from '/data/staff-picture.csv' delimiter ',';
















