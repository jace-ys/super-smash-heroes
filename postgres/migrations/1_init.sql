CREATE TABLE superheroes (
	id SERIAL PRIMARY KEY,
	full_name TEXT UNIQUE,
	alter_ego TEXT UNIQUE,
	image_url TEXT
);
