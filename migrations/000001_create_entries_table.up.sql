--Filename migrations/000001_create_entries_table.up.sql

CREATE TABLE IF NOT EXISTS entries (
    english text NOT NULL,
    kriol text NOT NULL
);