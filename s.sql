create table user(
	id int not null auto_increment primary key,
	name varchar(20),
	password varchar(20),
	age int(20) not null,
	sex int(1) not null
)