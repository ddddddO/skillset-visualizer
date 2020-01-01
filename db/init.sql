-- +migrate Up
CREATE TABLE skills (
	id serial PRIMARY KEY,
	categories json NOT NULL
);

INSERT INTO skills(categories) VALUES ('{"frontend": 1,"backend": 2,"database": 3,"infrastructure": 3,"network": 2,"facilitation": 4}');
