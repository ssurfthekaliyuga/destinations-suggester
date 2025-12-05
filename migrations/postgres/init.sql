create table rides
(
    id            uuid        not null,
    user_id       uuid        not null,
    from_place_id uuid        not null,
    to_place_id   uuid        not null,
    time          timestamptz not null,
    created_at    timestamptz not null default now()
);

create table rides_template
(
    like rides including all
);

create index ride_time_idx on rides_template (time);
create index user_id_idx on rides_template (user_id);

select partman.apply_template('public.rides');

select partman.create_parent(
               p_parent_table := 'public.rides',
               p_control := 'time',
               p_type := 'native',
               p_interval := '1 month',
               p_premake := 3
       );
