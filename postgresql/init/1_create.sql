create table company
(
    company_cd integer not null,
    rr_cd integer not null,
    company_name varchar not null,
    company_name_k varchar,
    company_name_h varchar,
    company_name_r varchar,
    company_url varchar,
    company_type integer,
    e_status integer,
    e_sort integer,
    PRIMARY KEY (company_cd)
);
comment on table company is 'company20200309.csv';

create table line
(
    line_cd integer not null ,
    company_cd integer not null,
    line_name varchar not null,
    line_name_k varchar,
    line_name_h varchar,
    line_color_c varchar,
    line_color_t varchar,
    line_type integer,
    lon float,
    lat float,
    zoom integer,
    e_status integer,
    e_sort integer,
    PRIMARY KEY (line_cd)
);
comment on table line is 'line20200306free.csv';

create table station
(
    station_cd integer not null,
    station_g_cd integer not null,
    station_name varchar not null,
    station_name_k varchar,
    station_name_r varchar,
    line_cd integer,
    pref_cd integer,
    post varchar,
    address varchar,
    lon float,
    lat float,
    open_ymd varchar ,
    close_ymd varchar,
    e_status integer,
    e_sort integer,
    PRIMARY KEY (station_cd)
);
comment on table station is 'station20200316free.csv';

create table station_join
(
    line_cd integer not null,
    station_cd1 integer not null,
    station_cd2 integer not null,
    PRIMARY KEY (line_cd, station_cd1, station_cd2)
);
comment on table station_join is 'join20200306.csv';
