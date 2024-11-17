#!/bin/bash
echo "Stopping any running Docker containers..."
docker stop f-bot || true
docker rm f-bot || true