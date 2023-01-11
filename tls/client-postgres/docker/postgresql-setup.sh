#!/usr/bin/env bash

PASSWD=password
EXPIRY=365
POSTGRESQL_DIR=postgresql
KEYPASS=MyBigSecret
CLIENT_SUBJ='/CN=postgres'

### CA key and cert ###

# Create a key for the root ca.
openssl genrsa -des3 \
  -passout pass:${KEYPASS} \
  -out "${POSTGRESQL_DIR}/ca-key-pass" 2048
openssl rsa \
  -passin pass:${KEYPASS} \
  -in "${POSTGRESQL_DIR}/ca-key-pass" \
  -out "${POSTGRESQL_DIR}/ca-key.pem"

# Create self-signed root ca.
openssl req -new -x509 -nodes \
  -key "${POSTGRESQL_DIR}/ca-key.pem" \
  -sha256 \
  -days "${EXPIRY}" \
  -subj '/CN=root-ca' \
  -out "${POSTGRESQL_DIR}/ca-cert.pem"

### Server key and cert ###

# Create server key.
openssl genrsa -des3 \
  -passout pass:${KEYPASS} \
  -out "${POSTGRESQL_DIR}/server-key-pass" 2048
openssl rsa \
  -passin pass:${KEYPASS} \
  -in "${POSTGRESQL_DIR}/server-key-pass" \
  -out "${POSTGRESQL_DIR}/server-key.pem"

# Create a server cert request.
openssl req -new -nodes \
  -subj '/CN=localhost' \
  -key "${POSTGRESQL_DIR}/server-key.pem" \
  -out "${POSTGRESQL_DIR}/server.csr" \

# Create server cert and sign it with root ca.
openssl x509 -req \
  -sha256 \
  -days "${EXPIRY}" \
  -in "${POSTGRESQL_DIR}/server.csr" \
  -CAcreateserial \
  -CA "${POSTGRESQL_DIR}/ca-cert.pem" \
  -CAkey "${POSTGRESQL_DIR}/ca-key.pem" \
  -out "${POSTGRESQL_DIR}/server-cert.pem"

### Client key and cert ###

# Create server key.
openssl genrsa -des3 \
  -passout pass:${KEYPASS} \
  -out "${POSTGRESQL_DIR}/client-key-pass" 2048
openssl rsa \
  -passin pass:${KEYPASS} \
  -in "${POSTGRESQL_DIR}/client-key-pass" \
  -out "${POSTGRESQL_DIR}/client-key.pem"

# Create a server cert request.
openssl req -new -nodes \
  -subj "${CLIENT_SUBJ}" \
  -key "${POSTGRESQL_DIR}/client-key.pem" \
  -out "${POSTGRESQL_DIR}/client.csr" \

# Create server cert and sign it with root ca.
openssl x509 -req \
  -sha256 \
  -days "${EXPIRY}" \
  -in "${POSTGRESQL_DIR}/client.csr" \
  -CAcreateserial \
  -CA "${POSTGRESQL_DIR}/ca-cert.pem" \
  -CAkey "${POSTGRESQL_DIR}/ca-key.pem" \
  -out "${POSTGRESQL_DIR}/client-cert.pem"

# Clean up,
rm "${POSTGRESQL_DIR}"/*-pass
rm "${POSTGRESQL_DIR}"/*.csr
rm "${POSTGRESQL_DIR}"/*.srl

# Change permissions
chmod 600 postgresql/*.pem

# Review the certs.
echo "CA cert:"
openssl x509 -in "${POSTGRESQL_DIR}/ca-cert.pem" -text -noout
echo "Server cert:"
openssl x509 -in "${POSTGRESQL_DIR}/server-cert.pem" -text -noout
echo "Client cert:"
openssl x509 -in "${POSTGRESQL_DIR}/client-cert.pem" -text -noout
