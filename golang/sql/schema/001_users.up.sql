CREATE SCHEMA users;

CREATE TABLE users (
  id bigserial NOT NULL PRIMARY KEY,
  email varchar(254) NOT NULL UNIQUE,
  username varchar(60) NOT NULL UNIQUE,
  pass varchar(60) NOT NULL,
  salt varchar(60) NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

