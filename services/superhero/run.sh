#!/bin/sh

echo "==> Starting service.."
exec service \
  --port "$PORT" \
  --gateway-port "$GATEWAY_PORT" \
  --postgres-host "$POSTGRES_HOST" \
  --postgres-user "$POSTGRES_USER" \
  --postgres-password "$POSTGRES_PASSWORD" \
  --postgres-db "$POSTGRES_DB" \
  --superhero-api-token "$SUPERHERO_API_TOKEN"
