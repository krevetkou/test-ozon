create table posts(
    id bigserial primary key unique not null,
    title text unique,
    text text,
    created_at timestamp with time zone,
    can_be_commented boolean,

    comment_id bigserial references comments (id) on delete cascade
);