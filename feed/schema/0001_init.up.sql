CREATE TABLE IF NOT EXISTS post
(
    "id"            serial       primary key,
    "image_path"    varchar(255) not null,
    "content"       varchar(255) null
);

CREATE TABLE IF NOT EXISTS "like"
(
    "id"         serial      primary key,
    "post_id"    integer not null,
    "user_id"    integer not null
);