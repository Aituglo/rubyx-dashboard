CREATE SCHEMA program;

CREATE TYPE program_type AS ENUM (
  'public',
  'private'
);

CREATE TABLE program (
  id bigserial NOT NULL PRIMARY KEY,
  platform_id bigserial REFERENCES platform(id) ON DELETE CASCADE,
  name varchar(200) NOT NULL,
  slug varchar(200) NOT NULL UNIQUE,
  vdp boolean NOT NULL,
  favourite boolean NOT NULL DEFAULT false,
  tag varchar NOT NULL,
  url varchar NOT NULL,
  type program_type NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);
