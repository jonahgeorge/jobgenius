create table Articles_Categories (
	id int not null comment 'Linked to id column in Articles',
	value int not null comment 'Linked to id column in Articles_Categories_Lookup'
) comment='Contains the article ids and associated article category ids, used as a bridge between Articles <-> Articles_Categories_Lookup';

create table Articles_Categories_Lookup (
	id int not null auto_increment comment 'Linked to value column in Articles_Categories',
	value text not null,
	primary key (id)
) comment='Contains string values for the different types of article categories';
