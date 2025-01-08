CREATE TABLE managers (
  id int auto_increment primary key,
  email varchar(255) not null,
  password varchar(255) null,
  name varchar(52) null,
  user_image_uri varchar(255) null,
  company_name varchar(52) null,
  company_image_uri varchar(255) null,
  created_at datetime not null,
  updated_at datetime null,
  deleted_at datetime null
);