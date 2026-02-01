CREATE TABLE writers (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);

CREATE TABLE books (
  id   BIGSERIAL PRIMARY KEY,
  writer_id bigint NOT NULL,
  name text      NOT NULL,
  bio  text,
  CONSTRAINT fk_books_writer
    FOREIGN KEY (writer_id) 
    REFERENCES writers(id)
    ON DELETE CASCADE
);