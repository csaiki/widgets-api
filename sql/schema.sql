DROP DATABASE IF EXISTS widgetsapi;
CREATE DATABASE widgetsapi;
USE widgetsapi;

DROP TABLE IF EXISTS users;
CREATE TABLE users
(
  id              INT unsigned NOT NULL, # Unique ID for the record
  name            VARCHAR(150) NOT NULL,                # Name of the user
  gravatar        VARCHAR(150) NOT NULL
);

DROP TABLE IF EXISTS widgets;
CREATE TABLE widgets
(
	id INT NOT NULL,
	name CHAR(32) NOT NULL,
	color CHAR(32) NOT NULL,
  price CHAR(32) NOT NULL,
  inventory CHAR(32) NOT NULL,
  melts BOOLEAN NOT NULL
);
