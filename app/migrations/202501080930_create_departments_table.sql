create table departments (
  id int auto_increment primary key,
  manager_id int not null,
  name varchar(33) not null,
  created_at datetime null,
  updated_at datetime null,
  deleted_at datetime null,
  constraint departments_managers_id_fk foreign key (manager_id) references managers (id)
);