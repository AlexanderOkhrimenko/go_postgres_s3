\c t1
create table jobs
(id serial8 NOT NULL,
error integer,
errordescription text,
command text,
status text,
complete integer,
task text,
priority text,
resulturl text,
resultsurl jsonb,
duration double precision,
outobjects jsonb,
primary key (id)
);
