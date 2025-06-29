-- +goose Up
CREATE TABLE countries (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  shortcut TEXT NOT NULL CHECK (length(shortcut) <= 3),
  federation TEXT NOT NULL CHECK (federation IN ('E', 'AM', 'AZ', 'AF'))
);

CREATE TABLE teams (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  country_id INTEGER NOT NULL,
  is_national_team BOOLEAN NOT NULL DEFAULT 0,
  FOREIGN KEY (country_id) REFERENCES countries(id)
);

-- +goose Down
DROP TABLE IF EXISTS teams;
DROP TABLE IF EXISTS countries;