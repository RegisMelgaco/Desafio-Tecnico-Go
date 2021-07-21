CREATE TABLE IF NOT EXISTS account(
   account_id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   cpf VARCHAR (11) NOT NULL UNIQUE,
   secret VARCHAR (255) NOT NULL
      CHECK(secret <> ''),
   balance INTEGER NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS transfer (
   transfer_id SERIAL PRIMARY KEY,
   origin_account SERIAL REFERENCES account (account_id),
   destination_account SERIAL REFERENCES account (account_id),
   amount INTEGER NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
