CREATE table contacts (
	id serial PRIMARY KEY,
	nome VARCHAR ( 200 ) NOT NULL,
	celular VARCHAR ( 20 ) NOT NULL,
	clientId INT,
	FOREIGN KEY (clientId) REFERENCES clients(Id)
);


