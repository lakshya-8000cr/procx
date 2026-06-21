#!/bin/bash

set -e

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo " PROCX Installer"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""


# url for reaching the realese
URL="https://github.com/lakshya-8000cr/procx/releases/download/v1.0.0/procx"

echo "Downloading Procx..."
sudo curl -fsSL "$URL" -o /usr/local/bin/procx  #download the file from url  then sve it to location /usr/local/bin/... okkk

echo "Setting permissions..."
sudo chmod +x /usr/local/bin/procx

echo ""
echo "Procx installed successfully."
echo ""
echo "Try:"
echo "  procx list"
echo "  procx diagnose <pid>"
echo "  procx port <port>"
echo ""