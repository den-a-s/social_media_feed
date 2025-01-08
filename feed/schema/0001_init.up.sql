CREATE TABLE IF NOT EXISTS post
(
    "id"            serial       primary key,
    "name"          varchar(255) not null,
    "image_path"    varchar(255) null,
    "content"       varchar(255) null
);

CREATE TABLE IF NOT EXISTS "like"
(
    "id"         serial       primary key,
    "post_id"    integer not null,
    "user_id"    integer not null
);