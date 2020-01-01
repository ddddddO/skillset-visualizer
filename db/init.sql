-- +migrate Up
CREATE TABLE skills (
	id serial PRIMARY KEY,
	categories json NOT NULL
);

INSERT INTO skills(categories) VALUES ('{"frontend":5,"backend":5,"database":5,"infrastructure":"5","network":"5","facilitation":5}');

-- +migrate Down
DROP TABLE skills;
