Table users as U{
id bigserial [pk]
first_name varchar [not null]
last_name varchar [not null]
email varchar [unique, not null]
password varchar [not null]
created_at timestamptz [not null, default: `now()`]
updated_at timestamptz [not null, default: `now()`]
access_level integer
}

Table reservations as Res{
id bigserial [pk]
first_name varchar [not null]
last_name varchar [not null]
email varchar [unique, not null]
phone varchar [not null]
start_date date [not null]
end_date date [not null]
room_id integer [ref: > Rom.id, not null]
created_at timestamptz [not null, default: `now()`]
updated_at timestamptz [not null, default: `now()`]
}

Table rooms as Rom{
id bigserial [pk]
room_name varchar [not null]
created_at timestamptz [not null, default: `now()`]
updated_at timestamptz [not null, default: `now()`]
}

Table room_restrictions as Rres{
id bigserial [pk]
start_date date [not null]
end_date date [not null]
room_id integer [ref: > Rom.id, not null]
restriction_id integer [ref: > restrict.id, not null]
reservations_id integer [ref: > Res.id, not null]
created_at timestamptz [not null, default: `now()`]
updated_at timestamptz [not null, default: `now()`]
}

Table restrictions as restrict{
id bigserial [pk]
restriction_name varchar [not null]
created_at timestamptz [not null, default: `now()`]
updated_at timestamptz [not null, default: `now()`]
}
