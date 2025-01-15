create table activity_types (
  id char(36) PRIMARY KEY,
  activity_type varchar(33) not null,
  calories_per_minute int not null
);