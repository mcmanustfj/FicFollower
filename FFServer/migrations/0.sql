create table User (
    id text primary key,
    username text unique
);

create table Story(
    id text primary key,
    title text,
    "site" text,
    link text,
    updated text,
    updatedlink text,
    author text,
    authorlink text,
    checked text
);

create table UserFollows(
    username text,
    storyid text,
    readDate text,
    readlink text,
    primary key(username, storyid),
    foreign key(username) references User(username),
    foreign key(storyid) references Story(id)
);

insert into User(username) values ('McRibbed');

PRAGMA foreign_keys = ON;
PRAGMA user_version = 1;