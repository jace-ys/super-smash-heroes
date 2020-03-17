CREATE TABLE superheroes (
	id SERIAL PRIMARY KEY,
	full_name TEXT UNIQUE NOT NULL,
	alter_ego TEXT UNIQUE NOT NULL,
	image_url TEXT NOT NULL,
	intelligence INTEGER DEFAULT 0,
	strength INTEGER DEFAULT 0,
	speed INTEGER DEFAULT 0,
	durability INTEGER DEFAULT 0,
	power INTEGER DEFAULT 0,
	combat INTEGER DEFAULT 0
);
