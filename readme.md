# Commands

## Run all DBs

`docker-compose up`

Don't forget to run sql migrations on your ClickHouse

## Run redpanda connect

`rpk connect run ./streams/hotel.yaml`

## Run go service

`go run .`

# UIs for everything

open these URLs in your browser:

`localhost:8089`  
`localhost:18123/play`

# Query to check data

`SELECT * FROM hotels.content`