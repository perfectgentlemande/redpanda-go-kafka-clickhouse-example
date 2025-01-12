CREATE DATABASE IF NOT EXISTS hotels;

CREATE TABLE IF NOT EXISTS hotels.content
(
    `id` String,
    `geo_id` UInt64,
    `emails` Array(String),
    `type` UInt8,
    `content_ru_address` String,
    `content_ru_name` String,
    `content_ru_description` String,
    `created_at` DateTime
)
    ENGINE = ReplacingMergeTree()
    ORDER BY (geo_id);