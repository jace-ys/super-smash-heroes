#!/bin/sh

echo "==> Running migrations.."
migrate \
  -source "file://$MIGRATIONS_DIR"  \
  -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST/$POSTGRES_DB?sslmode=disable" \
  up
