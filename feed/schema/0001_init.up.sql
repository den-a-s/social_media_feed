CREATE TABLE IF NOT EXISTS post
(
    "id"            serial       primary key,
    "name"          varchar(255) not null,
    "image_path"    varchar(255) null,
    "content"       varchar(255) null
);
