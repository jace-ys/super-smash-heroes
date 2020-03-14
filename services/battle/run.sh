#!/bin/sh

echo "==> Starting service.."
exec service \
  --port "$PORT" \
  --postgres-host "$POSTGRES_HOST" \
  --postgres-user "$POSTGRES_USER" \
  --postgres-password "$POSTGRES_PASSWORD" \
  --postgres-db "$POSTGRES_DB"
