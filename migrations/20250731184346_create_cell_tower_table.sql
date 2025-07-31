-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS btx (
    id          BIGSERIAL PRIMARY KEY,
    radio       TEXT,
    mcc         BIGINT,
    net         BIGINT,
    area         BIGINT,
    cell         BIGINT,
    unit         BIGINT,
    lat         NUMERIC,
    lon         NUMERIC,
    range       BIGINT,
    samples     BIGINT,
    changeable  BIGINT,
    averageSignal BIGINT ,
    created  BIGINT,
    updated  BIGINT 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS btx;
-- +goose StatementEnd
