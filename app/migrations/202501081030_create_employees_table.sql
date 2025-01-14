create table employess (
  id int auto_increment primary key,
  department_id int not null,
  identity_number varchar(33) null,
  name varchar(33) null,
  employee_image_uri varchar(255) null,
  gender enum ('male', 'female') null,
  created_at datetime null,
  updated_at datetime null,
  deleted_at datetime null,
  constraint department_id_fk foreign key (department_id) references departments (id)
);