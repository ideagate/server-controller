create table if not exists worker (
    id varchar(20) not null,
    application_id varchar(20) not null,
    project_id varchar(20) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    name varchar(50) not null,
    environment_id varchar(20) not null,
    replicas integer not null default 0,
    foreign key (project_id, application_id) references application(project_id, id),
    primary key (project_id, application_id, id)
)