CREATE table contacts (
	id serial PRIMARY KEY,
	nome VARCHAR ( 100 ) NOT NULL,
	celular VARCHAR ( 13 ) NOT NULL,
	client_id INT,
	CONSTRAINT fk_client
      FOREIGN KEY(client_id) 
	  REFERENCES clients(client_id)
);
