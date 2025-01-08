CREATE TABLE IF NOT EXISTS Users
(
    id        INTEGER PRIMARY KEY,
    email     TEXT NOT NULL UNIQUE,
    pass_hash BLOB NOT NULL,
    is_admin  BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO Users(email, pass_hash, is_admin)
VALUES(
    "matfak@uniyar.ac.ru",
    "$2a$10$mN21yv6wSRu91xfl0/s6kuwBT77LnI0Wv9H18Ka.3qaXCMNld9c.e",
    True
)