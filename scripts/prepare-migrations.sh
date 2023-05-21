#!/bin/bash

## Prepare migrations: clean old files and copy all migrations from services to migrations folder.
mkdir -p ./migrations
rm -Rvf ./migrations/*.sql !("00000000000000-db_preparing.sql")
cp -Rvf ./internal/**/repository/sql/migrations/*.sql ./migrations/