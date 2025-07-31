-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS btx (
    id          SERIAL PRIMARY KEY,
    radio       TEXT NOT NULL,
    mcc         INTEGER NOT NULL,
    mnc         INTEGER NOT NULL,
    lac         INTEGER NOT NULL,
    cid         INTEGER NOT NULL,
    lat         DOUBLE PRECISION,
    lon         DOUBLE PRECISION,
    range       INTEGER,
    samples     INTEGER,
    changeable  BOOLEAN DEFAULT false,
    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS btx;
-- +goose StatementEnd
