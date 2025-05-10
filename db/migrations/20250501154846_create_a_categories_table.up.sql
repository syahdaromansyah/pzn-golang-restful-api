CREATE TABLE categories(
  id VARCHAR(36) NOT NULL,
  name VARCHAR(128) NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT categories__name__chars_min__check CHECK (length(name) >= 3)
);
