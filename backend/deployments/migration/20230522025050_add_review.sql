-- +goose Up
-- +goose StatementBegin
create table dorama_set.review (
    id_dorama int not null,
    username text not null,
    mark int not null,
    content text not null default '',
    foreign key (id_dorama) references dorama_set.dorama(id) on delete cascade,
    foreign key (username) references dorama_set."user"(username) on delete cascade,
    primary key (id_dorama, username)
);

alter table dorama_set.review
    add constraint correct_mark check ( mark > 0 and mark <= 5 );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists dorama_set.review;
-- +goose StatementEnd
