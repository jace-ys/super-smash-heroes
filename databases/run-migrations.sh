#!/bin/sh

echo "==> Running migrations.."
migrate \
  -source "${SOURCE:-file://migrations}" \
  -database "$DATABASE_URL" \
  up
