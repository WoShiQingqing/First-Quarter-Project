#!/bin/bash
set -e
echo "[1/3] curl /healthz"
curl -fsS http://localhost/healthz >/dev/null && echo "OK"

echo "[2/3] curl /"
curl -fsS http://localhost/ | head -n1

echo "[3/3] check gunicorn + nginx"
ss -tnlp | egrep ':(80|8081)\s' || (echo "ports not listening"; exit 1)
echo "Smoke test passed."
