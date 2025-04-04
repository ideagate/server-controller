create table if not exists application (
    id varchar(20) not null,
    project_id varchar(20) not null references project(id),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    name varchar(50) not null,
    description text not null,
    primary key (id, project_id)
);

drop type if exists entrypoint_type;
create type entrypoint_type as enum('rest','cron');

create table if not exists entrypoint (
    id varchar(20) not null,
    application_id varchar(20) not null,
    project_id varchar(20) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    type entrypoint_type not null,
    name varchar(255),
    description text,
    settings jsonb not null default '{}',
    foreign key (application_id, project_id) references application(id, project_id),
    primary key (id, application_id, project_id)
);

create table if not exists workflow (
    version int not null,
    entrypoint_id varchar(20) not null,
    application_id varchar(20) not null,
    project_id varchar(20) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    data jsonb not null default '{}',
    foreign key (entrypoint_id, application_id, project_id) references entrypoint(id, application_id, project_id),
    primary key (version, entrypoint_id, application_id, project_id)
);