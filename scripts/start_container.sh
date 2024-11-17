#!/bin/bash
echo "Starting Docker container..."
docker run -d -p 80:8084 --name f-bot -v /var/param/.env:/app/.env $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/f-bot:latest