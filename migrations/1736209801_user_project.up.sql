create table if not exists "user" (
    id bigserial primary key,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    name varchar(50) not null,
    email varchar(50) not null,
    password varchar(100) not null
);

create table if not exists project (
    id varchar(20) primary key,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    name varchar(50) not null,
    description text
);

create table if not exists project_user (
    project_id varchar(20) not null references project(id),
    user_id bigint not null references "user"(id),
    created_at timestamptz not null default now(),
    primary key (project_id, user_id)
);
