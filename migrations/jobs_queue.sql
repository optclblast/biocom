create table if not exists jobs (
        id varchar(36) PRIMARY KEY,
        job bytea not null,
        status int8 default 0,
        ttr int,
        delayed_to timestamp default null,
        priority int default 3
);

create index index_jobs_id_status on jobs  (id, status); 
create index index_jobs_id_status_priority on jobs  (id, status, priority);
create index index_jobs_id_status_delayed on jobs  (id, status, delayed_to); 
create index index_jobs_id_status_priority_delayed on jobs  (id, status, priority, delayed_to);

