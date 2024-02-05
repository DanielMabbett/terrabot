#!/usr/bin/env bash

VERSION="v0.3.2"

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  wget "https://github.com/DanielMabbett/terrabot/releases/download/${VERSION}/terrabot-${VERSION}-linux-amd64.tar.gz"
  tar -xzf terrabot-${VERSION}-linux-amd64.tar.gz
  rm terrabot-${VERSION}-linux-amd64.tar.gz
else
  echo "[Information] Operating System $OSTYPE not supported."
fi
