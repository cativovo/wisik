-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE image ( 
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  src TEXT NOT NULL,
  label TEXT NOT NULL 
);
