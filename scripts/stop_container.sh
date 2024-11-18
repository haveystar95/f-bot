#!/bin/bash
echo "Stopping any running Docker containers..."
docker stop f-bot || true
docker rm f-bot || true

docker stop f-bot-telegram || true
docker rm f-bot-telegram || true