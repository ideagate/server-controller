create table if not exists application (
    id varchar(20) not null,
    project_id varchar(20) not null references project(id),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    name varchar(50) not null,
    description text not null,
    primary key (project_id, id)
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
    foreign key (project_id, application_id) references application(project_id, id),
    primary key (project_id, application_id, id)
);

create table if not exists workflow (
    version int not null,
    entrypoint_id varchar(20) not null,
    application_id varchar(20) not null,
    project_id varchar(20) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    data_bytes bytea,
    foreign key (project_id, application_id, entrypoint_id) references entrypoint(project_id, application_id, id),
    primary key (project_id, application_id, entrypoint_id, version)
);