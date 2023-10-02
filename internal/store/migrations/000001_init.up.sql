CREATE TABLE IF NOT EXISTS
  users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

CREATE
OR REPLACE FUNCTION create_user (p_username VARCHAR(50), p_password VARCHAR(50)) RETURNS INT LANGUAGE plpgsql AS $$
DECLARE
    user_id INT;
BEGIN
    INSERT INTO Users (username, password)
    VALUES (p_username, p_password)
    RETURNING id INTO user_id;
    
    RETURN user_id;
    
    EXCEPTION
        WHEN unique_violation THEN
            RAISE EXCEPTION 'The username already exists.';
        WHEN check_violation THEN
            RAISE EXCEPTION 'One or more input parameters violate the constraints.';
END;
$$;

CREATE
OR REPLACE FUNCTION get_user_by_username (v_username VARCHAR(50)) RETURNS setof users LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT *
    FROM users
    WHERE username = v_username;
END;
$$;

CREATE TABLE IF NOT EXISTS
  monitors (
    id SERIAL PRIMARY KEY,
    hostname VARCHAR(50) NOT NULL UNIQUE,
    interval INT NOT NULL,
    protocol VARCHAR(10) NOT NULL,
    addr VARCHAR(50) NOT NULL,
    port INT DEFAULT NULL,
    tags TEXT[] DEFAULT NULL
  );

CREATE OR REPLACE FUNCTION add_monitor (
  v_hostname VARCHAR(50),
  v_interval INT,
  v_protocol VARCHAR(10),
  v_addr VARCHAR(50),
  v_port INT,
  v_tags TEXT[]
) RETURNS INT LANGUAGE plpgsql AS $$
DECLARE
    monitor_id INT;
BEGIN
    INSERT INTO monitors(hostname, interval, protocol, addr, port, tags)
    VALUES(v_hostname, v_interval, v_protocol, v_addr, v_port, v_tags)
    RETURNING id INTO monitor_id;
    
    RETURN monitor_id;
END;
$$;

CREATE
OR REPLACE FUNCTION get_monitor (v_id INT) RETURNS setof monitors LANGUAGE plpgsql AS $$
BEGIN
    RETURN QUERY
    SELECT id, hostname, interval, protocol, addr, port, tags
    FROM monitors
    WHERE id = v_id;
END;
$$;
