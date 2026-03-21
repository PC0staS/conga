#!/usr/bin/env bash
set -euo pipefail

# Usage: ./tools/generate_certs.sh DOMAIN=conga.local
# Defaults:
DOMAIN=${DOMAIN:-conga.local}
OUT_DIR="$(pwd)/certs"
HOSTS=("${DOMAIN}" "localhost" "127.0.0.1")

mkdir -p "$OUT_DIR"

if ! command -v mkcert >/dev/null 2>&1; then
  echo "mkcert no está instalado. Instala mkcert y vuelve a intentarlo: https://github.com/FiloSottile/mkcert"
  exit 2
fi

echo "Instalando/asegurando CA local (si hace falta)..."
mkcert -install

CERT_FILE="$OUT_DIR/${DOMAIN}.crt"
KEY_FILE="$OUT_DIR/${DOMAIN}.key"

echo "Generando certificado para: ${HOSTS[*]} -> $CERT_FILE, $KEY_FILE"
mkcert -cert-file "$CERT_FILE" -key-file "$KEY_FILE" "${HOSTS[@]}"

echo "Certificados generados en: $OUT_DIR"

echo "Recuerda añadir '127.0.0.1 ${DOMAIN}' a /etc/hosts si aún no lo has hecho."
