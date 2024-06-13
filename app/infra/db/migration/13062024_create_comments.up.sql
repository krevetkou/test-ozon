create table comments(
    id bigserial primary key unique not null,
    post_id text,
    created_at timestamp with time zone,
    comment varchar(2000) not null,

    comment_id bigserial references comments (id) on delete cascade
);