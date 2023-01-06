create table url (
	id varchar(5) not null primary key,
	criacao timestamp not null default current_timestamp,
	destino varchar(200) not null
);

create table clique (
	url_id varchar(5) not null,
	contagem integer,
	constraint fk_clique_url foreign key (url_id)
	    references url(id) on delete no action
);
