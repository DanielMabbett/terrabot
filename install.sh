#!/usr/bin/env bash

# Script to download terrabot for Linux

VERSION="v0.3.2"
FILE_NAME="terrabot-${VERSION}-linux-amd64"
DOWNLOAD_URL="https://github.com/DanielMabbett/terrabot/releases/download/${VERSION}/${FILE_NAME}"

# Function to download file using wget or curl
download_file() {
    if command -v wget > /dev/null; then
        wget "$1" -O "$2"
    elif command -v curl > /dev/null; then
        curl -L "$1" -o "$2"
    else
        echo "[Error] Neither wget nor curl is installed. Please install one and retry."
        exit 1
    fi
}

# Main installation process
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    echo "[Information] Starting download of terrabot ${VERSION} for Linux..."
    download_file "${DOWNLOAD_URL}" "${FILE_NAME}"

    if [ -f "${FILE_NAME}" ]; then
        chmod +x "${FILE_NAME}"
        echo "[Success] Downloaded and prepared ${FILE_NAME}."
        echo "Run './${FILE_NAME}' to start using terrabot."
    else
        echo "[Error] Failed to download ${FILE_NAME}. Please check your internet connection and try again."
    fi
else
    echo "[Error] Operating System $OSTYPE not supported. terrabot only supports Linux."
fi
