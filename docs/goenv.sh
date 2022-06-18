#!/usr/bin/env bash

# goenv hasn't been updated for a while. To get the latest version of Go, whe
# need to manually install go, do the shim, and then bind it to goenv.
declare GO_VER='1.17.5'
declare GO_FILE="go${GO_VER}.darwin-amd64.tar.gz"
declare TMP_DIR='/tmp'
declare GOENV_VERSIONS_DIR="${HOME}/.goenv/versions"

# Check that goenv is installed.
if ! which goenv > /dev/null 2>&1; then
  echo 'goenv is not installed. Exiting now.'
  exit 1
else
  if [ ! -d "${GOENV_VERSIONS_DIR}" ]; then
    mkdir -p "${GOENV_VERSIONS_DIR}"
  fi
fi

# Download and install go.
if [ "$(uname -m)" == 'arm64' ]; then
  GO_FILE='go1.17.5.darwin-arm64.tar.gz'
fi

wget -c "https://go.dev/dl/${GO_FILE}" -P "${TMP_DIR}"
tar xvfz "${TMP_DIR}/${GO_FILE}" -C "${GOENV_VERSIONS_DIR}"
mv "${GOENV_VERSIONS_DIR}/go" "${GOENV_VERSIONS_DIR}/${GO_VER}"

# Switch to the latest version and verify.
goenv global "${GO_VER}"
echo 'Go has been installed.'
go version
