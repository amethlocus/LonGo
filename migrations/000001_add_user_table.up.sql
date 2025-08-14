Create Table "user"(
    username varchar(30) unique primary key not null,
    email text unique not null,
    password text not null,
    name text,
    lastname text,
    created_at timestamptz not null default(now() at time zone 'UTC'),
    last_update timestamptz not null default(now() at time zone 'UTC'));
