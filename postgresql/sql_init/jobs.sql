\c t1
create table jobs
(id serial8 NOT NULL,
error integer,
priority text,
resulturl text,
resultsurl jsonb,
duration double precision,
outobjects jsonb,
primary key (id)
);
