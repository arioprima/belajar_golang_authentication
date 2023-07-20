create table admin(
                      id uuid primary key,
                      email varchar(255) not null,
                      password varchar(255) not null,
                      created_at timestamp not null default now(),
                      updated_at timestamp not null default now()
);

create table customers(
                          id uuid primary key,
                          firstname varchar(255) not null,
                          lastname varchar(255) not null,
                          username varchar(255) not null,
                          email varchar(255) not null,
                          password varchar(255) not null,
                          phone varchar(255),
                          address varchar(255),
                          created_at timestamp not null default now(),
                          updated_at timestamp not null default now()
);
create table tokens(
                       id uuid primary key,
                       user_id uuid not null,
                       token varchar(255) not null,
                       refresh_token varchar(255) not null
);