#!/bin/sh

cd /app && ./grata &
echo "grata start all server.."
tail -f /dev/null