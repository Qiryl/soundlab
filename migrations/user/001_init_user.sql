-- users table
CREATE TABLE users(
    user_id text PRIMARY KEY CHECK (user_id != '') NOT NULL,
    email text unique NOT NULL,
    username text unique NOT NULL,
    full_name text NOT NULL,
    password text NOT NULL,
    bio text,
    avatar_url text,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
	updated_at timestamp without time zone NOT NULL DEFAULT now()
);

CREATE INDEX user_name ON users(username text_pattern_ops);

---- create above / drop below ----

DROP TABLE users;
