#!/bin/bash

set -e

//
echo "Installing Procx..."

# url for reaching the realese
URL="https://github.com/lakshya-8000cr/procx/releases/download/v1.0.0/procx"

sudo curl -L $URL -o /usr/local/bin/procx #download the file from url  then sve it to location /usr/local/bin/... okkk


# permissions 
sudo chmod +x /usr/local/bin/procx  # make the binary executable 

echo ""
echo "Procx installed successfully!" #this will show on terminal
echo ""
echo "Try:"
echo ""
echo "procx list"
echo "procx diagnose <pid>"
echo ""