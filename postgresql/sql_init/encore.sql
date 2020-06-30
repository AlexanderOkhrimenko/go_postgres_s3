\c t1
create table encore_tab
(id serial8 NOT NULL,
url text NOT NULL,
error integer,
errordescription text,
mimetype text,
container text,
streams jsonb,
statusdownload integer,
filename text,
primary key (id)
);
