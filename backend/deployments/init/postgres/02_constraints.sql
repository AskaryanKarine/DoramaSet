alter table dorama_set.list
    add constraint correct_type check (type = 'private' or type = 'public');

alter table dorama_set.dorama
    add constraint correct_status check (status = 'in progress' or status = 'finish' or status = 'will released'),
    add constraint correct_year check (release_year >= 1923);
