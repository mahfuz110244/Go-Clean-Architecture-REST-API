DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS news CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;
-- CREATE EXTENSION IF NOT EXISTS postgis;
-- CREATE EXTENSION IF NOT EXISTS postgis_topology;


CREATE TABLE users
(
    user_id      UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    first_name   VARCHAR(32)                 NOT NULL CHECK ( first_name <> '' ),
    last_name    VARCHAR(32)                 NOT NULL CHECK ( last_name <> '' ),
    email        VARCHAR(64) UNIQUE          NOT NULL CHECK ( email <> '' ),
    password     VARCHAR(250)                NOT NULL CHECK ( octet_length(password) <> 0 ),
    role         VARCHAR(10)                 NOT NULL DEFAULT 'user',
    about        VARCHAR(1024)                        DEFAULT '',
    avatar       VARCHAR(512),
    phone_number VARCHAR(20),
    address      VARCHAR(250),
    city         VARCHAR(30),
    country      VARCHAR(30),
    gender       VARCHAR(20)                 NOT NULL DEFAULT 'male',
    postcode     INTEGER,
    birthday     DATE                                 DEFAULT NULL,
    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,
    login_date   TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE news
(
    news_id    UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    author_id  UUID                     NOT NULL REFERENCES users (user_id),
    title      VARCHAR(250)             NOT NULL CHECK ( title <> '' ),
    content    TEXT                     NOT NULL CHECK ( content <> '' ),
    image_url  VARCHAR(1024) check ( image_url <> '' ),
    category   VARCHAR(250),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments
(
    comment_id UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    author_id  UUID                                               NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
    news_id    UUID                                               NOT NULL REFERENCES news (news_id) ON DELETE CASCADE,
    message    VARCHAR(1024)                                      NOT NULL CHECK ( message <> '' ),
    likes      BIGINT                   DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS news_title_id_idx ON news (title);

-- Create sales order status setting Table
CREATE TYPE status_enum AS ENUM ('estimate', 'issued', 'in_progress', 'fulfilled', 'closed_short', 'void', 'expired');
CREATE TABLE IF NOT EXISTS status (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	deleted_at TIMESTAMP,
	created_by VARCHAR (36),
	updated_by VARCHAR (36),
	name status_enum,
	description VARCHAR (255),
	active BOOLEAN NOT NULL DEFAULT TRUE,
	order_no smallint DEFAULT 0,
	UNIQUE(name)
);


-- Step 1: Create trigger Function
CREATE OR REPLACE FUNCTION trigger_set_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- Step 2: Then create a trigger for each table that has the column updated_at
DO $$
DECLARE
    t text;
BEGIN
    FOR t IN
        SELECT table_name FROM information_schema.columns WHERE column_name = 'updated_at'
    LOOP
        EXECUTE format('CREATE TRIGGER trigger_set_update_timestamp
                    BEFORE UPDATE ON %I
                    FOR EACH ROW EXECUTE PROCEDURE trigger_set_update_timestamp()', t,t);
    END loop;
END;
$$ language 'plpgsql';